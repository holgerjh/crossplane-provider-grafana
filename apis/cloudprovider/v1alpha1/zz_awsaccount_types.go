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

type AwsAccountInitParameters struct {

	// (Set of String) A set of regions that this AWS Account resource applies to.
	// A set of regions that this AWS Account resource applies to.
	// +listType=set
	Regions []*string `json:"regions,omitempty" tf:"regions,omitempty"`

	// (String) An IAM Role ARN string to represent with this AWS Account resource.
	// An IAM Role ARN string to represent with this AWS Account resource.
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// (String) The StackID of the Grafana Cloud instance.
	// The StackID of the Grafana Cloud instance.
	StackID *string `json:"stackId,omitempty" tf:"stack_id,omitempty"`
}

type AwsAccountObservation struct {

	// This has the format "{{ stack_id }}:{{ resource_id }}".
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Set of String) A set of regions that this AWS Account resource applies to.
	// A set of regions that this AWS Account resource applies to.
	// +listType=set
	Regions []*string `json:"regions,omitempty" tf:"regions,omitempty"`

	// (String) The ID given by the Grafana Cloud Provider API to this AWS Account resource.
	// The ID given by the Grafana Cloud Provider API to this AWS Account resource.
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`

	// (String) An IAM Role ARN string to represent with this AWS Account resource.
	// An IAM Role ARN string to represent with this AWS Account resource.
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// (String) The StackID of the Grafana Cloud instance.
	// The StackID of the Grafana Cloud instance.
	StackID *string `json:"stackId,omitempty" tf:"stack_id,omitempty"`
}

type AwsAccountParameters struct {

	// (Set of String) A set of regions that this AWS Account resource applies to.
	// A set of regions that this AWS Account resource applies to.
	// +kubebuilder:validation:Optional
	// +listType=set
	Regions []*string `json:"regions,omitempty" tf:"regions,omitempty"`

	// (String) An IAM Role ARN string to represent with this AWS Account resource.
	// An IAM Role ARN string to represent with this AWS Account resource.
	// +kubebuilder:validation:Optional
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// (String) The StackID of the Grafana Cloud instance.
	// The StackID of the Grafana Cloud instance.
	// +kubebuilder:validation:Optional
	StackID *string `json:"stackId,omitempty" tf:"stack_id,omitempty"`
}

// AwsAccountSpec defines the desired state of AwsAccount
type AwsAccountSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AwsAccountParameters `json:"forProvider"`
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
	InitProvider AwsAccountInitParameters `json:"initProvider,omitempty"`
}

// AwsAccountStatus defines the observed state of AwsAccount.
type AwsAccountStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AwsAccountObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// AwsAccount is the Schema for the AwsAccounts API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,grafana}
type AwsAccount struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.regions) || (has(self.initProvider) && has(self.initProvider.regions))",message="spec.forProvider.regions is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.roleArn) || (has(self.initProvider) && has(self.initProvider.roleArn))",message="spec.forProvider.roleArn is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.stackId) || (has(self.initProvider) && has(self.initProvider.stackId))",message="spec.forProvider.stackId is a required parameter"
	Spec   AwsAccountSpec   `json:"spec"`
	Status AwsAccountStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AwsAccountList contains a list of AwsAccounts
type AwsAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AwsAccount `json:"items"`
}

// Repository type metadata.
var (
	AwsAccount_Kind             = "AwsAccount"
	AwsAccount_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AwsAccount_Kind}.String()
	AwsAccount_KindAPIVersion   = AwsAccount_Kind + "." + CRDGroupVersion.String()
	AwsAccount_GroupVersionKind = CRDGroupVersion.WithKind(AwsAccount_Kind)
)

func init() {
	SchemeBuilder.Register(&AwsAccount{}, &AwsAccountList{})
}
