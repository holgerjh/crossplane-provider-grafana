package config

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
	grafanaProvider "github.com/grafana/terraform-provider-grafana/v4/pkg/provider"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func TestMetricsEndpointScrapeJobFrameworkRead(t *testing.T) {
	ctx := context.Background()
	const resourceName = "grafana_connections_metrics_endpoint_scrape_job"
	configured := &ujconfig.Resource{Name: resourceName}
	ExternalNameConfigurations()(configured)
	for name, test := range map[string]struct {
		status int
		id     any
	}{
		"existing without id": {status: http.StatusOK},
		"existing empty id":   {status: http.StatusOK, id: ""},
		"existing seeded id":  {status: http.StatusOK, id: "12345:existing-job"},
		"missing job":         {status: http.StatusNotFound},
		"unauthorized":        {status: http.StatusUnauthorized},
		"forbidden":           {status: http.StatusForbidden},
	} {
		t.Run(name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
				if request.Method != http.MethodGet || request.URL.Path != "/api/v1/stacks/12345/metrics-endpoint/jobs/existing-job" {
					t.Errorf("unexpected request: %s %s", request.Method, request.URL.Path)
				}
				writer.Header().Set("Content-Type", "application/json")
				writer.WriteHeader(test.status)
				_, _ = io.WriteString(writer, `{"data":{"enabled":true,"authentication_method":"bearer","url":"https://example.com/metrics","scrape_interval_seconds":60}}`)
			}))
			defer server.Close()
			clients, err := grafanaProvider.CreateClients(grafanaProvider.ProviderConfig{
				ConnectionsAPIAccessToken: types.StringValue("test-token"),
				ConnectionsAPIURL:         types.StringValue(server.URL),
				HTTPHeaders:               types.MapNull(types.StringType),
				RetryStatusCodes:          types.SetNull(types.StringType),
			})
			if err != nil {
				t.Fatal(err)
			}
			var job resource.ResourceWithConfigure
			for _, factory := range grafanaProvider.FrameworkProvider("test").Resources(ctx) {
				candidate := factory()
				metadata := resource.MetadataResponse{}
				candidate.Metadata(ctx, resource.MetadataRequest{}, &metadata)
				if metadata.TypeName == resourceName {
					job = candidate.(resource.ResourceWithConfigure)
					break
				}
			}
			if job == nil {
				t.Fatal("metrics endpoint resource not registered")
			}
			configuration := resource.ConfigureResponse{}
			job.Configure(ctx, resource.ConfigureRequest{ProviderData: clients}, &configuration)
			if configuration.Diagnostics.HasError() {
				t.Fatal(configuration.Diagnostics)
			}
			schema := resource.SchemaResponse{}
			job.Schema(ctx, resource.SchemaRequest{}, &schema)
			stateType := schema.Schema.Type().TerraformType(ctx).(tftypes.Object)
			values := map[string]tftypes.Value{}
			for attribute, attributeType := range stateType.AttributeTypes {
				values[attribute] = tftypes.NewValue(attributeType, nil)
			}
			values["stack_id"] = tftypes.NewValue(tftypes.String, "12345")
			values["name"] = tftypes.NewValue(tftypes.String, "existing-job")
			values["id"] = tftypes.NewValue(tftypes.String, test.id)
			state := tfsdk.State{Schema: schema.Schema, Raw: tftypes.NewValue(stateType, values)}
			response := resource.ReadResponse{State: state}
			job.Read(ctx, resource.ReadRequest{State: state}, &response)
			diagnostics := []*tfprotov6.Diagnostic{}
			for _, diagnostic := range response.Diagnostics {
				diagnostics = append(diagnostics, &tfprotov6.Diagnostic{
					Severity: tfprotov6.DiagnosticSeverityError,
					Summary:  diagnostic.Summary(),
					Detail:   diagnostic.Detail(),
				})
			}
			if got := configured.ExternalName.IsNotFoundDiagnosticFn(diagnostics); got != (test.status == http.StatusNotFound) {
				t.Fatalf("not-found classification = %t, diagnostics = %v", got, response.Diagnostics)
			}
			if test.status != http.StatusOK {
				if !response.Diagnostics.HasError() {
					t.Fatal("failed API read did not return an error")
				}
				return
			}
			if response.Diagnostics.HasError() {
				t.Fatal(response.Diagnostics)
			}
			attributes := map[string]any{}
			for _, attribute := range []string{"id", "stack_id", "name"} {
				var value types.String
				if diags := response.State.GetAttribute(ctx, path.Root(attribute), &value); diags.HasError() {
					t.Fatal(diags)
				}
				attributes[attribute] = value.ValueString()
			}
			externalName, err := configured.ExternalName.GetExternalNameFn(attributes)
			if err != nil || externalName != "12345:existing-job" {
				t.Fatalf("discovered external name = %q, %v", externalName, err)
			}
		})
	}
}
