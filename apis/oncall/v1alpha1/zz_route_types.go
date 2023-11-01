// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

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

type RouteInitParameters struct {

	// (String) The ID of the escalation chain.
	// The ID of the escalation chain.
	EscalationChainID *string `json:"escalationChainId,omitempty" tf:"escalation_chain_id,omitempty"`

	// (String) The ID of the integration.
	// The ID of the integration.
	IntegrationID *string `json:"integrationId,omitempty" tf:"integration_id,omitempty"`

	// specific settings for a route. (see below for nested schema)
	// MS teams-specific settings for a route.
	Msteams []RouteMsteamsInitParameters `json:"msteams,omitempty" tf:"msteams,omitempty"`

	// (Number) The position of the route (starts from 0).
	// The position of the route (starts from 0).
	Position *float64 `json:"position,omitempty" tf:"position,omitempty"`

	// (String) Python Regex query. Route is chosen for an alert if there is a match inside the alert payload.
	// Python Regex query. Route is chosen for an alert if there is a match inside the alert payload.
	RoutingRegex *string `json:"routingRegex,omitempty" tf:"routing_regex,omitempty"`

	// (String) The type of route. Can be jinja2, regex Defaults to regex.
	// The type of route. Can be jinja2, regex Defaults to `regex`.
	RoutingType *string `json:"routingType,omitempty" tf:"routing_type,omitempty"`

	// specific settings for a route. (see below for nested schema)
	// Slack-specific settings for a route.
	Slack []RouteSlackInitParameters `json:"slack,omitempty" tf:"slack,omitempty"`

	// specific settings for a route. (see below for nested schema)
	// Telegram-specific settings for a route.
	Telegram []RouteTelegramInitParameters `json:"telegram,omitempty" tf:"telegram,omitempty"`
}

type RouteMsteamsInitParameters struct {

	// (Boolean) Enable notification in MS teams. Defaults to true.
	// Enable notification in MS teams. Defaults to `true`.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) The ID of this resource.
	// MS teams channel id. Alerts will be directed to this channel in Microsoft teams.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type RouteMsteamsObservation struct {

	// (Boolean) Enable notification in MS teams. Defaults to true.
	// Enable notification in MS teams. Defaults to `true`.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) The ID of this resource.
	// MS teams channel id. Alerts will be directed to this channel in Microsoft teams.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type RouteMsteamsParameters struct {

	// (Boolean) Enable notification in MS teams. Defaults to true.
	// Enable notification in MS teams. Defaults to `true`.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) The ID of this resource.
	// MS teams channel id. Alerts will be directed to this channel in Microsoft teams.
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type RouteObservation struct {

	// (String) The ID of the escalation chain.
	// The ID of the escalation chain.
	EscalationChainID *string `json:"escalationChainId,omitempty" tf:"escalation_chain_id,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The ID of the integration.
	// The ID of the integration.
	IntegrationID *string `json:"integrationId,omitempty" tf:"integration_id,omitempty"`

	// specific settings for a route. (see below for nested schema)
	// MS teams-specific settings for a route.
	Msteams []RouteMsteamsObservation `json:"msteams,omitempty" tf:"msteams,omitempty"`

	// (Number) The position of the route (starts from 0).
	// The position of the route (starts from 0).
	Position *float64 `json:"position,omitempty" tf:"position,omitempty"`

	// (String) Python Regex query. Route is chosen for an alert if there is a match inside the alert payload.
	// Python Regex query. Route is chosen for an alert if there is a match inside the alert payload.
	RoutingRegex *string `json:"routingRegex,omitempty" tf:"routing_regex,omitempty"`

	// (String) The type of route. Can be jinja2, regex Defaults to regex.
	// The type of route. Can be jinja2, regex Defaults to `regex`.
	RoutingType *string `json:"routingType,omitempty" tf:"routing_type,omitempty"`

	// specific settings for a route. (see below for nested schema)
	// Slack-specific settings for a route.
	Slack []RouteSlackObservation `json:"slack,omitempty" tf:"slack,omitempty"`

	// specific settings for a route. (see below for nested schema)
	// Telegram-specific settings for a route.
	Telegram []RouteTelegramObservation `json:"telegram,omitempty" tf:"telegram,omitempty"`
}

type RouteParameters struct {

	// (String) The ID of the escalation chain.
	// The ID of the escalation chain.
	// +kubebuilder:validation:Optional
	EscalationChainID *string `json:"escalationChainId,omitempty" tf:"escalation_chain_id,omitempty"`

	// (String) The ID of the integration.
	// The ID of the integration.
	// +kubebuilder:validation:Optional
	IntegrationID *string `json:"integrationId,omitempty" tf:"integration_id,omitempty"`

	// specific settings for a route. (see below for nested schema)
	// MS teams-specific settings for a route.
	// +kubebuilder:validation:Optional
	Msteams []RouteMsteamsParameters `json:"msteams,omitempty" tf:"msteams,omitempty"`

	// (Number) The position of the route (starts from 0).
	// The position of the route (starts from 0).
	// +kubebuilder:validation:Optional
	Position *float64 `json:"position,omitempty" tf:"position,omitempty"`

	// (String) Python Regex query. Route is chosen for an alert if there is a match inside the alert payload.
	// Python Regex query. Route is chosen for an alert if there is a match inside the alert payload.
	// +kubebuilder:validation:Optional
	RoutingRegex *string `json:"routingRegex,omitempty" tf:"routing_regex,omitempty"`

	// (String) The type of route. Can be jinja2, regex Defaults to regex.
	// The type of route. Can be jinja2, regex Defaults to `regex`.
	// +kubebuilder:validation:Optional
	RoutingType *string `json:"routingType,omitempty" tf:"routing_type,omitempty"`

	// specific settings for a route. (see below for nested schema)
	// Slack-specific settings for a route.
	// +kubebuilder:validation:Optional
	Slack []RouteSlackParameters `json:"slack,omitempty" tf:"slack,omitempty"`

	// specific settings for a route. (see below for nested schema)
	// Telegram-specific settings for a route.
	// +kubebuilder:validation:Optional
	Telegram []RouteTelegramParameters `json:"telegram,omitempty" tf:"telegram,omitempty"`
}

type RouteSlackInitParameters struct {

	// (String) Slack channel id. Alerts will be directed to this channel in Slack.
	// Slack channel id. Alerts will be directed to this channel in Slack.
	ChannelID *string `json:"channelId,omitempty" tf:"channel_id,omitempty"`

	// (Boolean) Enable notification in MS teams. Defaults to true.
	// Enable notification in Slack. Defaults to `true`.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type RouteSlackObservation struct {

	// (String) Slack channel id. Alerts will be directed to this channel in Slack.
	// Slack channel id. Alerts will be directed to this channel in Slack.
	ChannelID *string `json:"channelId,omitempty" tf:"channel_id,omitempty"`

	// (Boolean) Enable notification in MS teams. Defaults to true.
	// Enable notification in Slack. Defaults to `true`.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type RouteSlackParameters struct {

	// (String) Slack channel id. Alerts will be directed to this channel in Slack.
	// Slack channel id. Alerts will be directed to this channel in Slack.
	// +kubebuilder:validation:Optional
	ChannelID *string `json:"channelId,omitempty" tf:"channel_id,omitempty"`

	// (Boolean) Enable notification in MS teams. Defaults to true.
	// Enable notification in Slack. Defaults to `true`.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type RouteTelegramInitParameters struct {

	// (Boolean) Enable notification in MS teams. Defaults to true.
	// Enable notification in Telegram. Defaults to `true`.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) The ID of this resource.
	// Telegram channel id. Alerts will be directed to this channel in Telegram.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type RouteTelegramObservation struct {

	// (Boolean) Enable notification in MS teams. Defaults to true.
	// Enable notification in Telegram. Defaults to `true`.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) The ID of this resource.
	// Telegram channel id. Alerts will be directed to this channel in Telegram.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type RouteTelegramParameters struct {

	// (Boolean) Enable notification in MS teams. Defaults to true.
	// Enable notification in Telegram. Defaults to `true`.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) The ID of this resource.
	// Telegram channel id. Alerts will be directed to this channel in Telegram.
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

// RouteSpec defines the desired state of Route
type RouteSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RouteParameters `json:"forProvider"`
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
	InitProvider RouteInitParameters `json:"initProvider,omitempty"`
}

// RouteStatus defines the observed state of Route.
type RouteStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RouteObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Route is the Schema for the Routes API. HTTP API https://grafana.com/docs/oncall/latest/oncall-api-reference/routes/
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,grafana}
type Route struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.escalationChainId) || (has(self.initProvider) && has(self.initProvider.escalationChainId))",message="spec.forProvider.escalationChainId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.integrationId) || (has(self.initProvider) && has(self.initProvider.integrationId))",message="spec.forProvider.integrationId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.position) || (has(self.initProvider) && has(self.initProvider.position))",message="spec.forProvider.position is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.routingRegex) || (has(self.initProvider) && has(self.initProvider.routingRegex))",message="spec.forProvider.routingRegex is a required parameter"
	Spec   RouteSpec   `json:"spec"`
	Status RouteStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RouteList contains a list of Routes
type RouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Route `json:"items"`
}

// Repository type metadata.
var (
	Route_Kind             = "Route"
	Route_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Route_Kind}.String()
	Route_KindAPIVersion   = Route_Kind + "." + CRDGroupVersion.String()
	Route_GroupVersionKind = CRDGroupVersion.WithKind(Route_Kind)
)

func init() {
	SchemeBuilder.Register(&Route{}, &RouteList{})
}
