/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type DashboardsInitParameters struct {

	// (Map of String) Add report variables to the dashboard. Values should be separated by commas.
	// Add report variables to the dashboard. Values should be separated by commas.
	// +mapType=granular
	ReportVariables map[string]*string `json:"reportVariables,omitempty" tf:"report_variables,omitempty"`

	// (Block List, Max: 1, Deprecated) Time range of the report. (see below for nested schema)
	// Time range of the report.
	TimeRange []TimeRangeInitParameters `json:"timeRange,omitempty" tf:"time_range,omitempty"`

	// (String) Dashboard uid.
	// Dashboard uid.
	UID *string `json:"uid,omitempty" tf:"uid,omitempty"`
}

type DashboardsObservation struct {

	// (Map of String) Add report variables to the dashboard. Values should be separated by commas.
	// Add report variables to the dashboard. Values should be separated by commas.
	// +mapType=granular
	ReportVariables map[string]*string `json:"reportVariables,omitempty" tf:"report_variables,omitempty"`

	// (Block List, Max: 1, Deprecated) Time range of the report. (see below for nested schema)
	// Time range of the report.
	TimeRange []TimeRangeObservation `json:"timeRange,omitempty" tf:"time_range,omitempty"`

	// (String) Dashboard uid.
	// Dashboard uid.
	UID *string `json:"uid,omitempty" tf:"uid,omitempty"`
}

type DashboardsParameters struct {

	// (Map of String) Add report variables to the dashboard. Values should be separated by commas.
	// Add report variables to the dashboard. Values should be separated by commas.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	ReportVariables map[string]*string `json:"reportVariables,omitempty" tf:"report_variables,omitempty"`

	// (Block List, Max: 1, Deprecated) Time range of the report. (see below for nested schema)
	// Time range of the report.
	// +kubebuilder:validation:Optional
	TimeRange []TimeRangeParameters `json:"timeRange,omitempty" tf:"time_range,omitempty"`

	// (String) Dashboard uid.
	// Dashboard uid.
	// +kubebuilder:validation:Optional
	UID *string `json:"uid" tf:"uid,omitempty"`
}

type ReportInitParameters struct {

	// Reference to a Dashboard in oss to populate dashboardUid.
	// +kubebuilder:validation:Optional
	DashboardRef *v1.Reference `json:"dashboardRef,omitempty" tf:"-"`

	// Selector for a Dashboard in oss to populate dashboardUid.
	// +kubebuilder:validation:Optional
	DashboardSelector *v1.Selector `json:"dashboardSelector,omitempty" tf:"-"`

	// (String, Deprecated) Dashboard to be sent in the report.
	// Dashboard to be sent in the report.
	// +crossplane:generate:reference:type=github.com/grafana/crossplane-provider-grafana/apis/oss/v1alpha1.Dashboard
	// +crossplane:generate:reference:extractor=github.com/grafana/crossplane-provider-grafana/config/grafana.UIDExtractor()
	// +crossplane:generate:reference:refFieldName=DashboardRef
	// +crossplane:generate:reference:selectorFieldName=DashboardSelector
	DashboardUID *string `json:"dashboardUid,omitempty" tf:"dashboard_uid,omitempty"`

	// (Block List) List of dashboards to render into the report (see below for nested schema)
	// List of dashboards to render into the report
	Dashboards []DashboardsInitParameters `json:"dashboards,omitempty" tf:"dashboards,omitempty"`

	// (Set of String) Specifies what kind of attachment to generate for the report. Allowed values: pdf, csv, image.
	// Specifies what kind of attachment to generate for the report. Allowed values: `pdf`, `csv`, `image`.
	// +listType=set
	Formats []*string `json:"formats,omitempty" tf:"formats,omitempty"`

	// (Boolean) Whether to include a link to the dashboard in the report. Defaults to true.
	// Whether to include a link to the dashboard in the report. Defaults to `true`.
	IncludeDashboardLink *bool `json:"includeDashboardLink,omitempty" tf:"include_dashboard_link,omitempty"`

	// (Boolean) Whether to include a CSV file of table panel data. Defaults to false.
	// Whether to include a CSV file of table panel data. Defaults to `false`.
	IncludeTableCsv *bool `json:"includeTableCsv,omitempty" tf:"include_table_csv,omitempty"`

	// (String) Layout of the report. Allowed values: simple, grid. Defaults to grid.
	// Layout of the report. Allowed values: `simple`, `grid`. Defaults to `grid`.
	Layout *string `json:"layout,omitempty" tf:"layout,omitempty"`

	// (String) Message to be sent in the report.
	// Message to be sent in the report.
	Message *string `json:"message,omitempty" tf:"message,omitempty"`

	// (String) Name of the report.
	// Name of the report.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The Organization ID. If not set, the Org ID defined in the provider block will be used.
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	// +crossplane:generate:reference:type=github.com/grafana/crossplane-provider-grafana/apis/oss/v1alpha1.Organization
	// +crossplane:generate:reference:refFieldName=OrganizationRef
	// +crossplane:generate:reference:selectorFieldName=OrganizationSelector
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	// Reference to a Organization in oss to populate orgId.
	// +kubebuilder:validation:Optional
	OrganizationRef *v1.Reference `json:"organizationRef,omitempty" tf:"-"`

	// Selector for a Organization in oss to populate orgId.
	// +kubebuilder:validation:Optional
	OrganizationSelector *v1.Selector `json:"organizationSelector,omitempty" tf:"-"`

	// (String) Orientation of the report. Allowed values: landscape, portrait. Defaults to landscape.
	// Orientation of the report. Allowed values: `landscape`, `portrait`. Defaults to `landscape`.
	Orientation *string `json:"orientation,omitempty" tf:"orientation,omitempty"`

	// (List of String) List of recipients of the report.
	// List of recipients of the report.
	Recipients []*string `json:"recipients,omitempty" tf:"recipients,omitempty"`

	// to email address of the report.
	// Reply-to email address of the report.
	ReplyTo *string `json:"replyTo,omitempty" tf:"reply_to,omitempty"`

	// (Block List, Min: 1, Max: 1) Schedule of the report. (see below for nested schema)
	// Schedule of the report.
	Schedule []ScheduleInitParameters `json:"schedule,omitempty" tf:"schedule,omitempty"`

	// (Block List, Max: 1, Deprecated) Time range of the report. (see below for nested schema)
	// Time range of the report.
	TimeRange []ReportTimeRangeInitParameters `json:"timeRange,omitempty" tf:"time_range,omitempty"`
}

type ReportObservation struct {

	// (String, Deprecated) Dashboard to be sent in the report.
	// Dashboard to be sent in the report.
	DashboardUID *string `json:"dashboardUid,omitempty" tf:"dashboard_uid,omitempty"`

	// (Block List) List of dashboards to render into the report (see below for nested schema)
	// List of dashboards to render into the report
	Dashboards []DashboardsObservation `json:"dashboards,omitempty" tf:"dashboards,omitempty"`

	// (Set of String) Specifies what kind of attachment to generate for the report. Allowed values: pdf, csv, image.
	// Specifies what kind of attachment to generate for the report. Allowed values: `pdf`, `csv`, `image`.
	// +listType=set
	Formats []*string `json:"formats,omitempty" tf:"formats,omitempty"`

	// (String) Generated identifier of the report.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Boolean) Whether to include a link to the dashboard in the report. Defaults to true.
	// Whether to include a link to the dashboard in the report. Defaults to `true`.
	IncludeDashboardLink *bool `json:"includeDashboardLink,omitempty" tf:"include_dashboard_link,omitempty"`

	// (Boolean) Whether to include a CSV file of table panel data. Defaults to false.
	// Whether to include a CSV file of table panel data. Defaults to `false`.
	IncludeTableCsv *bool `json:"includeTableCsv,omitempty" tf:"include_table_csv,omitempty"`

	// (String) Layout of the report. Allowed values: simple, grid. Defaults to grid.
	// Layout of the report. Allowed values: `simple`, `grid`. Defaults to `grid`.
	Layout *string `json:"layout,omitempty" tf:"layout,omitempty"`

	// (String) Message to be sent in the report.
	// Message to be sent in the report.
	Message *string `json:"message,omitempty" tf:"message,omitempty"`

	// (String) Name of the report.
	// Name of the report.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The Organization ID. If not set, the Org ID defined in the provider block will be used.
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	// (String) Orientation of the report. Allowed values: landscape, portrait. Defaults to landscape.
	// Orientation of the report. Allowed values: `landscape`, `portrait`. Defaults to `landscape`.
	Orientation *string `json:"orientation,omitempty" tf:"orientation,omitempty"`

	// (List of String) List of recipients of the report.
	// List of recipients of the report.
	Recipients []*string `json:"recipients,omitempty" tf:"recipients,omitempty"`

	// to email address of the report.
	// Reply-to email address of the report.
	ReplyTo *string `json:"replyTo,omitempty" tf:"reply_to,omitempty"`

	// (Block List, Min: 1, Max: 1) Schedule of the report. (see below for nested schema)
	// Schedule of the report.
	Schedule []ScheduleObservation `json:"schedule,omitempty" tf:"schedule,omitempty"`

	// (Block List, Max: 1, Deprecated) Time range of the report. (see below for nested schema)
	// Time range of the report.
	TimeRange []ReportTimeRangeObservation `json:"timeRange,omitempty" tf:"time_range,omitempty"`
}

type ReportParameters struct {

	// Reference to a Dashboard in oss to populate dashboardUid.
	// +kubebuilder:validation:Optional
	DashboardRef *v1.Reference `json:"dashboardRef,omitempty" tf:"-"`

	// Selector for a Dashboard in oss to populate dashboardUid.
	// +kubebuilder:validation:Optional
	DashboardSelector *v1.Selector `json:"dashboardSelector,omitempty" tf:"-"`

	// (String, Deprecated) Dashboard to be sent in the report.
	// Dashboard to be sent in the report.
	// +crossplane:generate:reference:type=github.com/grafana/crossplane-provider-grafana/apis/oss/v1alpha1.Dashboard
	// +crossplane:generate:reference:extractor=github.com/grafana/crossplane-provider-grafana/config/grafana.UIDExtractor()
	// +crossplane:generate:reference:refFieldName=DashboardRef
	// +crossplane:generate:reference:selectorFieldName=DashboardSelector
	// +kubebuilder:validation:Optional
	DashboardUID *string `json:"dashboardUid,omitempty" tf:"dashboard_uid,omitempty"`

	// (Block List) List of dashboards to render into the report (see below for nested schema)
	// List of dashboards to render into the report
	// +kubebuilder:validation:Optional
	Dashboards []DashboardsParameters `json:"dashboards,omitempty" tf:"dashboards,omitempty"`

	// (Set of String) Specifies what kind of attachment to generate for the report. Allowed values: pdf, csv, image.
	// Specifies what kind of attachment to generate for the report. Allowed values: `pdf`, `csv`, `image`.
	// +kubebuilder:validation:Optional
	// +listType=set
	Formats []*string `json:"formats,omitempty" tf:"formats,omitempty"`

	// (Boolean) Whether to include a link to the dashboard in the report. Defaults to true.
	// Whether to include a link to the dashboard in the report. Defaults to `true`.
	// +kubebuilder:validation:Optional
	IncludeDashboardLink *bool `json:"includeDashboardLink,omitempty" tf:"include_dashboard_link,omitempty"`

	// (Boolean) Whether to include a CSV file of table panel data. Defaults to false.
	// Whether to include a CSV file of table panel data. Defaults to `false`.
	// +kubebuilder:validation:Optional
	IncludeTableCsv *bool `json:"includeTableCsv,omitempty" tf:"include_table_csv,omitempty"`

	// (String) Layout of the report. Allowed values: simple, grid. Defaults to grid.
	// Layout of the report. Allowed values: `simple`, `grid`. Defaults to `grid`.
	// +kubebuilder:validation:Optional
	Layout *string `json:"layout,omitempty" tf:"layout,omitempty"`

	// (String) Message to be sent in the report.
	// Message to be sent in the report.
	// +kubebuilder:validation:Optional
	Message *string `json:"message,omitempty" tf:"message,omitempty"`

	// (String) Name of the report.
	// Name of the report.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The Organization ID. If not set, the Org ID defined in the provider block will be used.
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	// +crossplane:generate:reference:type=github.com/grafana/crossplane-provider-grafana/apis/oss/v1alpha1.Organization
	// +crossplane:generate:reference:refFieldName=OrganizationRef
	// +crossplane:generate:reference:selectorFieldName=OrganizationSelector
	// +kubebuilder:validation:Optional
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	// Reference to a Organization in oss to populate orgId.
	// +kubebuilder:validation:Optional
	OrganizationRef *v1.Reference `json:"organizationRef,omitempty" tf:"-"`

	// Selector for a Organization in oss to populate orgId.
	// +kubebuilder:validation:Optional
	OrganizationSelector *v1.Selector `json:"organizationSelector,omitempty" tf:"-"`

	// (String) Orientation of the report. Allowed values: landscape, portrait. Defaults to landscape.
	// Orientation of the report. Allowed values: `landscape`, `portrait`. Defaults to `landscape`.
	// +kubebuilder:validation:Optional
	Orientation *string `json:"orientation,omitempty" tf:"orientation,omitempty"`

	// (List of String) List of recipients of the report.
	// List of recipients of the report.
	// +kubebuilder:validation:Optional
	Recipients []*string `json:"recipients,omitempty" tf:"recipients,omitempty"`

	// to email address of the report.
	// Reply-to email address of the report.
	// +kubebuilder:validation:Optional
	ReplyTo *string `json:"replyTo,omitempty" tf:"reply_to,omitempty"`

	// (Block List, Min: 1, Max: 1) Schedule of the report. (see below for nested schema)
	// Schedule of the report.
	// +kubebuilder:validation:Optional
	Schedule []ScheduleParameters `json:"schedule,omitempty" tf:"schedule,omitempty"`

	// (Block List, Max: 1, Deprecated) Time range of the report. (see below for nested schema)
	// Time range of the report.
	// +kubebuilder:validation:Optional
	TimeRange []ReportTimeRangeParameters `json:"timeRange,omitempty" tf:"time_range,omitempty"`
}

type ReportTimeRangeInitParameters struct {

	// (String) Start of the time range.
	// Start of the time range.
	From *string `json:"from,omitempty" tf:"from,omitempty"`

	// (String) End of the time range.
	// End of the time range.
	To *string `json:"to,omitempty" tf:"to,omitempty"`
}

type ReportTimeRangeObservation struct {

	// (String) Start of the time range.
	// Start of the time range.
	From *string `json:"from,omitempty" tf:"from,omitempty"`

	// (String) End of the time range.
	// End of the time range.
	To *string `json:"to,omitempty" tf:"to,omitempty"`
}

type ReportTimeRangeParameters struct {

	// (String) Start of the time range.
	// Start of the time range.
	// +kubebuilder:validation:Optional
	From *string `json:"from,omitempty" tf:"from,omitempty"`

	// (String) End of the time range.
	// End of the time range.
	// +kubebuilder:validation:Optional
	To *string `json:"to,omitempty" tf:"to,omitempty"`
}

type ScheduleInitParameters struct {

	// (String) Custom interval of the report.
	// Note: This field is only available when frequency is set to custom.
	// Custom interval of the report.
	// **Note:** This field is only available when frequency is set to `custom`.
	CustomInterval *string `json:"customInterval,omitempty" tf:"custom_interval,omitempty"`

	// 01-02T15:04:05 format if you want to set a custom timezone
	// End time of the report. If empty, the report will be sent indefinitely (according to frequency). Note that times will be saved as UTC in Grafana. Use 2006-01-02T15:04:05 format if you want to set a custom timezone
	EndTime *string `json:"endTime,omitempty" tf:"end_time,omitempty"`

	// (String) Frequency of the report. Allowed values: never, once, hourly, daily, weekly, monthly, custom.
	// Frequency of the report. Allowed values: `never`, `once`, `hourly`, `daily`, `weekly`, `monthly`, `custom`.
	Frequency *string `json:"frequency,omitempty" tf:"frequency,omitempty"`

	// (Boolean) Send the report on the last day of the month Defaults to false.
	// Send the report on the last day of the month Defaults to `false`.
	LastDayOfMonth *bool `json:"lastDayOfMonth,omitempty" tf:"last_day_of_month,omitempty"`

	// 01-02T15:04:05 format if you want to set a custom timezone
	// Start time of the report. If empty, the start date will be set to the creation time. Note that times will be saved as UTC in Grafana. Use 2006-01-02T15:04:05 format if you want to set a custom timezone
	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`

	// (String) Set the report time zone. Defaults to GMT.
	// Set the report time zone. Defaults to `GMT`.
	Timezone *string `json:"timezone,omitempty" tf:"timezone,omitempty"`

	// (Boolean) Whether to send the report only on work days. Defaults to false.
	// Whether to send the report only on work days. Defaults to `false`.
	WorkdaysOnly *bool `json:"workdaysOnly,omitempty" tf:"workdays_only,omitempty"`
}

type ScheduleObservation struct {

	// (String) Custom interval of the report.
	// Note: This field is only available when frequency is set to custom.
	// Custom interval of the report.
	// **Note:** This field is only available when frequency is set to `custom`.
	CustomInterval *string `json:"customInterval,omitempty" tf:"custom_interval,omitempty"`

	// 01-02T15:04:05 format if you want to set a custom timezone
	// End time of the report. If empty, the report will be sent indefinitely (according to frequency). Note that times will be saved as UTC in Grafana. Use 2006-01-02T15:04:05 format if you want to set a custom timezone
	EndTime *string `json:"endTime,omitempty" tf:"end_time,omitempty"`

	// (String) Frequency of the report. Allowed values: never, once, hourly, daily, weekly, monthly, custom.
	// Frequency of the report. Allowed values: `never`, `once`, `hourly`, `daily`, `weekly`, `monthly`, `custom`.
	Frequency *string `json:"frequency,omitempty" tf:"frequency,omitempty"`

	// (Boolean) Send the report on the last day of the month Defaults to false.
	// Send the report on the last day of the month Defaults to `false`.
	LastDayOfMonth *bool `json:"lastDayOfMonth,omitempty" tf:"last_day_of_month,omitempty"`

	// 01-02T15:04:05 format if you want to set a custom timezone
	// Start time of the report. If empty, the start date will be set to the creation time. Note that times will be saved as UTC in Grafana. Use 2006-01-02T15:04:05 format if you want to set a custom timezone
	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`

	// (String) Set the report time zone. Defaults to GMT.
	// Set the report time zone. Defaults to `GMT`.
	Timezone *string `json:"timezone,omitempty" tf:"timezone,omitempty"`

	// (Boolean) Whether to send the report only on work days. Defaults to false.
	// Whether to send the report only on work days. Defaults to `false`.
	WorkdaysOnly *bool `json:"workdaysOnly,omitempty" tf:"workdays_only,omitempty"`
}

type ScheduleParameters struct {

	// (String) Custom interval of the report.
	// Note: This field is only available when frequency is set to custom.
	// Custom interval of the report.
	// **Note:** This field is only available when frequency is set to `custom`.
	// +kubebuilder:validation:Optional
	CustomInterval *string `json:"customInterval,omitempty" tf:"custom_interval,omitempty"`

	// 01-02T15:04:05 format if you want to set a custom timezone
	// End time of the report. If empty, the report will be sent indefinitely (according to frequency). Note that times will be saved as UTC in Grafana. Use 2006-01-02T15:04:05 format if you want to set a custom timezone
	// +kubebuilder:validation:Optional
	EndTime *string `json:"endTime,omitempty" tf:"end_time,omitempty"`

	// (String) Frequency of the report. Allowed values: never, once, hourly, daily, weekly, monthly, custom.
	// Frequency of the report. Allowed values: `never`, `once`, `hourly`, `daily`, `weekly`, `monthly`, `custom`.
	// +kubebuilder:validation:Optional
	Frequency *string `json:"frequency" tf:"frequency,omitempty"`

	// (Boolean) Send the report on the last day of the month Defaults to false.
	// Send the report on the last day of the month Defaults to `false`.
	// +kubebuilder:validation:Optional
	LastDayOfMonth *bool `json:"lastDayOfMonth,omitempty" tf:"last_day_of_month,omitempty"`

	// 01-02T15:04:05 format if you want to set a custom timezone
	// Start time of the report. If empty, the start date will be set to the creation time. Note that times will be saved as UTC in Grafana. Use 2006-01-02T15:04:05 format if you want to set a custom timezone
	// +kubebuilder:validation:Optional
	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`

	// (String) Set the report time zone. Defaults to GMT.
	// Set the report time zone. Defaults to `GMT`.
	// +kubebuilder:validation:Optional
	Timezone *string `json:"timezone,omitempty" tf:"timezone,omitempty"`

	// (Boolean) Whether to send the report only on work days. Defaults to false.
	// Whether to send the report only on work days. Defaults to `false`.
	// +kubebuilder:validation:Optional
	WorkdaysOnly *bool `json:"workdaysOnly,omitempty" tf:"workdays_only,omitempty"`
}

type TimeRangeInitParameters struct {

	// (String) Start of the time range.
	// Start of the time range.
	From *string `json:"from,omitempty" tf:"from,omitempty"`

	// (String) End of the time range.
	// End of the time range.
	To *string `json:"to,omitempty" tf:"to,omitempty"`
}

type TimeRangeObservation struct {

	// (String) Start of the time range.
	// Start of the time range.
	From *string `json:"from,omitempty" tf:"from,omitempty"`

	// (String) End of the time range.
	// End of the time range.
	To *string `json:"to,omitempty" tf:"to,omitempty"`
}

type TimeRangeParameters struct {

	// (String) Start of the time range.
	// Start of the time range.
	// +kubebuilder:validation:Optional
	From *string `json:"from,omitempty" tf:"from,omitempty"`

	// (String) End of the time range.
	// End of the time range.
	// +kubebuilder:validation:Optional
	To *string `json:"to,omitempty" tf:"to,omitempty"`
}

// ReportSpec defines the desired state of Report
type ReportSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ReportParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider ReportInitParameters `json:"initProvider,omitempty"`
}

// ReportStatus defines the observed state of Report.
type ReportStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ReportObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Report is the Schema for the Reports API. Note: This resource is available only with Grafana Enterprise 7.+. Official documentation https://grafana.com/docs/grafana/latest/dashboards/create-reports/HTTP API https://grafana.com/docs/grafana/latest/developers/http_api/reporting/
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,grafana}
type Report struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.recipients) || (has(self.initProvider) && has(self.initProvider.recipients))",message="spec.forProvider.recipients is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.schedule) || (has(self.initProvider) && has(self.initProvider.schedule))",message="spec.forProvider.schedule is a required parameter"
	Spec   ReportSpec   `json:"spec"`
	Status ReportStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ReportList contains a list of Reports
type ReportList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Report `json:"items"`
}

// Repository type metadata.
var (
	Report_Kind             = "Report"
	Report_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Report_Kind}.String()
	Report_KindAPIVersion   = Report_Kind + "." + CRDGroupVersion.String()
	Report_GroupVersionKind = CRDGroupVersion.WithKind(Report_Kind)
)

func init() {
	SchemeBuilder.Register(&Report{}, &ReportList{})
}
