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

type ScheduleInitParameters struct {

	// (String) The URL of external iCal calendar which override primary events.
	// The URL of external iCal calendar which override primary events.
	IcalURLOverrides *string `json:"icalUrlOverrides,omitempty" tf:"ical_url_overrides,omitempty"`

	// (String) The URL of the external calendar iCal file.
	// The URL of the external calendar iCal file.
	IcalURLPrimary *string `json:"icalUrlPrimary,omitempty" tf:"ical_url_primary,omitempty"`

	// (String) The schedule's name.
	// The schedule's name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// call shifts.
	// The list of ID's of on-call shifts.
	Shifts []*string `json:"shifts,omitempty" tf:"shifts,omitempty"`

	// specific settings for a schedule. (see below for nested schema)
	// The Slack-specific settings for a schedule.
	Slack []ScheduleSlackInitParameters `json:"slack,omitempty" tf:"slack,omitempty"`

	// (String) The ID of the OnCall team. To get one, create a team in Grafana, and navigate to the OnCall plugin (to sync the team with OnCall). You can then get the ID using the grafana_oncall_team datasource.
	// The ID of the OnCall team. To get one, create a team in Grafana, and navigate to the OnCall plugin (to sync the team with OnCall). You can then get the ID using the `grafana_oncall_team` datasource.
	TeamID *string `json:"teamId,omitempty" tf:"team_id,omitempty"`

	// (String) The schedule's time zone.
	// The schedule's time zone.
	TimeZone *string `json:"timeZone,omitempty" tf:"time_zone,omitempty"`

	// (String) The schedule's type.
	// The schedule's type.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ScheduleObservation struct {

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The URL of external iCal calendar which override primary events.
	// The URL of external iCal calendar which override primary events.
	IcalURLOverrides *string `json:"icalUrlOverrides,omitempty" tf:"ical_url_overrides,omitempty"`

	// (String) The URL of the external calendar iCal file.
	// The URL of the external calendar iCal file.
	IcalURLPrimary *string `json:"icalUrlPrimary,omitempty" tf:"ical_url_primary,omitempty"`

	// (String) The schedule's name.
	// The schedule's name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// call shifts.
	// The list of ID's of on-call shifts.
	Shifts []*string `json:"shifts,omitempty" tf:"shifts,omitempty"`

	// specific settings for a schedule. (see below for nested schema)
	// The Slack-specific settings for a schedule.
	Slack []ScheduleSlackObservation `json:"slack,omitempty" tf:"slack,omitempty"`

	// (String) The ID of the OnCall team. To get one, create a team in Grafana, and navigate to the OnCall plugin (to sync the team with OnCall). You can then get the ID using the grafana_oncall_team datasource.
	// The ID of the OnCall team. To get one, create a team in Grafana, and navigate to the OnCall plugin (to sync the team with OnCall). You can then get the ID using the `grafana_oncall_team` datasource.
	TeamID *string `json:"teamId,omitempty" tf:"team_id,omitempty"`

	// (String) The schedule's time zone.
	// The schedule's time zone.
	TimeZone *string `json:"timeZone,omitempty" tf:"time_zone,omitempty"`

	// (String) The schedule's type.
	// The schedule's type.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ScheduleParameters struct {

	// (String) The URL of external iCal calendar which override primary events.
	// The URL of external iCal calendar which override primary events.
	// +kubebuilder:validation:Optional
	IcalURLOverrides *string `json:"icalUrlOverrides,omitempty" tf:"ical_url_overrides,omitempty"`

	// (String) The URL of the external calendar iCal file.
	// The URL of the external calendar iCal file.
	// +kubebuilder:validation:Optional
	IcalURLPrimary *string `json:"icalUrlPrimary,omitempty" tf:"ical_url_primary,omitempty"`

	// (String) The schedule's name.
	// The schedule's name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// call shifts.
	// The list of ID's of on-call shifts.
	// +kubebuilder:validation:Optional
	Shifts []*string `json:"shifts,omitempty" tf:"shifts,omitempty"`

	// specific settings for a schedule. (see below for nested schema)
	// The Slack-specific settings for a schedule.
	// +kubebuilder:validation:Optional
	Slack []ScheduleSlackParameters `json:"slack,omitempty" tf:"slack,omitempty"`

	// (String) The ID of the OnCall team. To get one, create a team in Grafana, and navigate to the OnCall plugin (to sync the team with OnCall). You can then get the ID using the grafana_oncall_team datasource.
	// The ID of the OnCall team. To get one, create a team in Grafana, and navigate to the OnCall plugin (to sync the team with OnCall). You can then get the ID using the `grafana_oncall_team` datasource.
	// +kubebuilder:validation:Optional
	TeamID *string `json:"teamId,omitempty" tf:"team_id,omitempty"`

	// (String) The schedule's time zone.
	// The schedule's time zone.
	// +kubebuilder:validation:Optional
	TimeZone *string `json:"timeZone,omitempty" tf:"time_zone,omitempty"`

	// (String) The schedule's type.
	// The schedule's type.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ScheduleSlackInitParameters struct {

	// (String) Slack channel id. Reminder about schedule shifts will be directed to this channel in Slack.
	// Slack channel id. Reminder about schedule shifts will be directed to this channel in Slack.
	ChannelID *string `json:"channelId,omitempty" tf:"channel_id,omitempty"`

	// call users change.
	// Slack user group id. Members of user group will be updated when on-call users change.
	UserGroupID *string `json:"userGroupId,omitempty" tf:"user_group_id,omitempty"`
}

type ScheduleSlackObservation struct {

	// (String) Slack channel id. Reminder about schedule shifts will be directed to this channel in Slack.
	// Slack channel id. Reminder about schedule shifts will be directed to this channel in Slack.
	ChannelID *string `json:"channelId,omitempty" tf:"channel_id,omitempty"`

	// call users change.
	// Slack user group id. Members of user group will be updated when on-call users change.
	UserGroupID *string `json:"userGroupId,omitempty" tf:"user_group_id,omitempty"`
}

type ScheduleSlackParameters struct {

	// (String) Slack channel id. Reminder about schedule shifts will be directed to this channel in Slack.
	// Slack channel id. Reminder about schedule shifts will be directed to this channel in Slack.
	// +kubebuilder:validation:Optional
	ChannelID *string `json:"channelId,omitempty" tf:"channel_id,omitempty"`

	// call users change.
	// Slack user group id. Members of user group will be updated when on-call users change.
	// +kubebuilder:validation:Optional
	UserGroupID *string `json:"userGroupId,omitempty" tf:"user_group_id,omitempty"`
}

// ScheduleSpec defines the desired state of Schedule
type ScheduleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ScheduleParameters `json:"forProvider"`
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
	InitProvider ScheduleInitParameters `json:"initProvider,omitempty"`
}

// ScheduleStatus defines the observed state of Schedule.
type ScheduleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ScheduleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Schedule is the Schema for the Schedules API. HTTP API https://grafana.com/docs/oncall/latest/oncall-api-reference/schedules/
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,grafana}
type Schedule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || (has(self.initProvider) && has(self.initProvider.type))",message="spec.forProvider.type is a required parameter"
	Spec   ScheduleSpec   `json:"spec"`
	Status ScheduleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ScheduleList contains a list of Schedules
type ScheduleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Schedule `json:"items"`
}

// Repository type metadata.
var (
	Schedule_Kind             = "Schedule"
	Schedule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Schedule_Kind}.String()
	Schedule_KindAPIVersion   = Schedule_Kind + "." + CRDGroupVersion.String()
	Schedule_GroupVersionKind = CRDGroupVersion.WithKind(Schedule_Kind)
)

func init() {
	SchemeBuilder.Register(&Schedule{}, &ScheduleList{})
}
