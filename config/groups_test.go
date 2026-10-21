/*
Copyright 2026 Grafana Labs
*/

package config

import (
	"context"
	"testing"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

func TestMetricsEndpointScrapeJobExternalName(t *testing.T) {
	resource := &ujconfig.Resource{Name: "grafana_connections_metrics_endpoint_scrape_job"}
	ExternalNameConfigurations()(resource)
	externalName := resource.ExternalName
	parameters := map[string]any{"stack_id": "12345", "name": "existing:job"}
	want := "12345:existing:job"

	id, err := externalName.GetIDFn(context.Background(), "", parameters, nil)
	if err != nil || id != want {
		t.Fatalf("discovery ID = %q, %v; want %q", id, err, want)
	}
	for _, stateID := range []any{nil, "", want} {
		state := map[string]any{"stack_id": "12345", "name": "existing:job"}
		if stateID != nil {
			state["id"] = stateID
		}
		got, err := externalName.GetExternalNameFn(state)
		if err != nil || got != want {
			t.Fatalf("external name with state ID %v = %q, %v; want %q", stateID, got, err, want)
		}
	}

	importID := "67890:imported:job"
	id, err = externalName.GetIDFn(context.Background(), importID, parameters, nil)
	if err != nil || id != importID {
		t.Fatalf("import ID = %q, %v; want %q", id, err, importID)
	}
	got, err := externalName.GetExternalNameFn(map[string]any{"id": importID})
	if err != nil || got != importID {
		t.Fatalf("import external name = %q, %v; want %q", got, err, importID)
	}
	externalName.SetIdentifierArgumentFn(parameters, importID)
	if parameters["name"] != "existing:job" || parameters["stack_id"] != "12345" {
		t.Fatalf("identifier mapping changed parameters: %v", parameters)
	}
	if !externalName.DisableNameInitializer || len(externalName.OmittedFields) != 0 {
		t.Fatal("name initialization or field omission changed")
	}
}

func TestMetricsEndpointScrapeJobExternalNameMissingIdentity(t *testing.T) {
	resource := &ujconfig.Resource{Name: "grafana_connections_metrics_endpoint_scrape_job"}
	ExternalNameConfigurations()(resource)
	for name, attributes := range map[string]map[string]any{
		"missing stack": {"name": "job"},
		"empty stack":   {"stack_id": "", "name": "job"},
		"numeric stack": {"stack_id": 12345, "name": "job"},
		"missing name":  {"stack_id": "12345"},
		"empty name":    {"stack_id": "12345", "name": ""},
	} {
		t.Run(name, func(t *testing.T) {
			if _, err := resource.ExternalName.GetExternalNameFn(attributes); err == nil {
				t.Fatal("expected error extracting incomplete identity")
			}
			if _, err := resource.ExternalName.GetIDFn(context.Background(), "", attributes, nil); err == nil {
				t.Fatal("expected error constructing incomplete discovery ID")
			}
		})
	}
}

func TestShouldTreatDiagnosticsAsResourceNotFound(t *testing.T) {
	isNotFound := identifierFromProviderTreatingNotFoundDiagnostics().IsNotFoundDiagnosticFn

	cases := map[string]struct {
		diags []*tfprotov6.Diagnostic
		want  bool
	}{
		"empty int id": {
			diags: []*tfprotov6.Diagnostic{{
				Severity: tfprotov6.DiagnosticSeverityError,
				Summary:  "Failed to parse resource ID",
				Detail:   `expected int for field "id", got ""`,
			}},
			want: true,
		},
		"empty composite id": {
			diags: []*tfprotov6.Diagnostic{{
				Severity: tfprotov6.DiagnosticSeverityError,
				Summary:  "Unable to parse resource ID",
				Detail:   `id "" does not match expected format. Should be in the format: folderUID:type (role, team, or user):identifier`,
			}},
			want: true,
		},
		"wrapped empty id": {
			diags: []*tfprotov6.Diagnostic{{
				Severity: tfprotov6.DiagnosticSeverityError,
				Detail:   `Failed to parse resource ID: expected int for field "id", got ""`,
			}},
			want: true,
		},
		"non-empty invalid id": {
			diags: []*tfprotov6.Diagnostic{{
				Severity: tfprotov6.DiagnosticSeverityError,
				Summary:  "Failed to parse resource ID",
				Detail:   `expected int for field "id", got "abc"`,
			}},
			want: false,
		},
		"warning": {
			diags: []*tfprotov6.Diagnostic{{
				Severity: tfprotov6.DiagnosticSeverityWarning,
				Summary:  "Failed to parse resource ID",
				Detail:   `expected int for field "id", got ""`,
			}},
			want: false,
		},
		"nil diagnostic": {
			diags: []*tfprotov6.Diagnostic{nil},
			want:  false,
		},
		"metrics endpoint scrape job not found": {
			diags: []*tfprotov6.Diagnostic{{
				Severity: tfprotov6.DiagnosticSeverityError,
				Summary:  "failed to get metrics endpoint scrape job",
				Detail:   `failed to get metrics endpoint scrape job "staging-smoke": not found`,
			}},
			want: true,
		},
		"metrics endpoint scrape job read failure": {
			diags: []*tfprotov6.Diagnostic{{
				Severity: tfprotov6.DiagnosticSeverityError,
				Summary:  "failed to get metrics endpoint scrape job",
				Detail:   `failed to get metrics endpoint scrape job "staging-smoke": request not authorized for stack`,
			}},
			want: false,
		},
		"other resource not found": {
			diags: []*tfprotov6.Diagnostic{{
				Severity: tfprotov6.DiagnosticSeverityError,
				Summary:  "failed to get dashboard",
				Detail:   `failed to get dashboard "example": not found`,
			}},
			want: false,
		},
		"metrics endpoint scrape job not found warning": {
			diags: []*tfprotov6.Diagnostic{{
				Severity: tfprotov6.DiagnosticSeverityWarning,
				Summary:  "failed to get metrics endpoint scrape job",
				Detail:   `failed to get metrics endpoint scrape job "staging-smoke": not found`,
			}},
			want: false,
		},
		"string id hits list endpoint - folder": {
			diags: []*tfprotov6.Diagnostic{{
				Severity: tfprotov6.DiagnosticSeverityError,
				Summary:  "Failed to get folder",
				Detail:   `json: cannot unmarshal array into Go value of type models.Folder`,
			}},
			want: true,
		},
		"string id hits list endpoint - message template": {
			diags: []*tfprotov6.Diagnostic{{
				Severity: tfprotov6.DiagnosticSeverityError,
				Summary:  "Failed to read message template",
				Detail:   `json: cannot unmarshal array into Go value of type models.NotificationTemplate`,
			}},
			want: true,
		},
		"unmarshal array warning ignored": {
			diags: []*tfprotov6.Diagnostic{{
				Severity: tfprotov6.DiagnosticSeverityWarning,
				Summary:  "Failed to get folder",
				Detail:   `json: cannot unmarshal array into Go value of type models.Folder`,
			}},
			want: false,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			if got := isNotFound(tc.diags); got != tc.want {
				t.Fatalf("got %t, want %t", got, tc.want)
			}
		})
	}
}
