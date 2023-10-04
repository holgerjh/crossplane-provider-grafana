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

type DashboardPermissionObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type DashboardPermissionParameters struct {

	// Reference to a Dashboard in oss to populate dashboardUid.
	// +kubebuilder:validation:Optional
	DashboardRef *v1.Reference `json:"dashboardRef,omitempty" tf:"-"`

	// Selector for a Dashboard in oss to populate dashboardUid.
	// +kubebuilder:validation:Optional
	DashboardSelector *v1.Selector `json:"dashboardSelector,omitempty" tf:"-"`

	// UID of the dashboard to apply permissions to.
	// +crossplane:generate:reference:type=github.com/grafana/crossplane-provider-grafana/apis/oss/v1alpha1.Dashboard
	// +crossplane:generate:reference:extractor=github.com/grafana/crossplane-provider-grafana/config/grafana.UIDExtractor()
	// +crossplane:generate:reference:refFieldName=DashboardRef
	// +crossplane:generate:reference:selectorFieldName=DashboardSelector
	// +kubebuilder:validation:Optional
	DashboardUID *string `json:"dashboardUid,omitempty" tf:"dashboard_uid,omitempty"`

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

	// The permission items to add/update. Items that are omitted from the list will be removed.
	// +kubebuilder:validation:Required
	Permissions []PermissionsParameters `json:"permissions" tf:"permissions,omitempty"`
}

type PermissionsObservation struct {
}

type PermissionsParameters struct {

	// Permission to associate with item. Must be one of `View`, `Edit`, or `Admin`.
	// +kubebuilder:validation:Required
	Permission *string `json:"permission" tf:"permission,omitempty"`

	// Manage permissions for `Viewer` or `Editor` roles.
	// +kubebuilder:validation:Optional
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// ID of the team to manage permissions for. Defaults to `0`.
	// +kubebuilder:validation:Optional
	TeamID *string `json:"teamId,omitempty" tf:"team_id,omitempty"`

	// ID of the user or service account to manage permissions for. Defaults to `0`.
	// +kubebuilder:validation:Optional
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`
}

// DashboardPermissionSpec defines the desired state of DashboardPermission
type DashboardPermissionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DashboardPermissionParameters `json:"forProvider"`
}

// DashboardPermissionStatus defines the observed state of DashboardPermission.
type DashboardPermissionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DashboardPermissionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DashboardPermission is the Schema for the DashboardPermissions API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,grafana}
type DashboardPermission struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DashboardPermissionSpec   `json:"spec"`
	Status            DashboardPermissionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DashboardPermissionList contains a list of DashboardPermissions
type DashboardPermissionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DashboardPermission `json:"items"`
}

// Repository type metadata.
var (
	DashboardPermission_Kind             = "DashboardPermission"
	DashboardPermission_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DashboardPermission_Kind}.String()
	DashboardPermission_KindAPIVersion   = DashboardPermission_Kind + "." + CRDGroupVersion.String()
	DashboardPermission_GroupVersionKind = CRDGroupVersion.WithKind(DashboardPermission_Kind)
)

func init() {
	SchemeBuilder.Register(&DashboardPermission{}, &DashboardPermissionList{})
}
