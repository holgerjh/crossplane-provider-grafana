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

type AlgorithmInitParameters struct {

	// For DBSCAN only, specify the configuration map
	Config []ConfigInitParameters `json:"config,omitempty" tf:"config,omitempty"`

	// The name of the algorithm to use ('mad' or 'dbscan').
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specify the sensitivity of the detector (in range [0,1]).
	Sensitivity *float64 `json:"sensitivity,omitempty" tf:"sensitivity,omitempty"`
}

type AlgorithmObservation struct {

	// For DBSCAN only, specify the configuration map
	Config []ConfigObservation `json:"config,omitempty" tf:"config,omitempty"`

	// The name of the algorithm to use ('mad' or 'dbscan').
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specify the sensitivity of the detector (in range [0,1]).
	Sensitivity *float64 `json:"sensitivity,omitempty" tf:"sensitivity,omitempty"`
}

type AlgorithmParameters struct {

	// For DBSCAN only, specify the configuration map
	// +kubebuilder:validation:Optional
	Config []ConfigParameters `json:"config,omitempty" tf:"config,omitempty"`

	// The name of the algorithm to use ('mad' or 'dbscan').
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// Specify the sensitivity of the detector (in range [0,1]).
	// +kubebuilder:validation:Optional
	Sensitivity *float64 `json:"sensitivity" tf:"sensitivity,omitempty"`
}

type ConfigInitParameters struct {

	// Specify the epsilon parameter (positive float)
	Epsilon *float64 `json:"epsilon,omitempty" tf:"epsilon,omitempty"`
}

type ConfigObservation struct {

	// Specify the epsilon parameter (positive float)
	Epsilon *float64 `json:"epsilon,omitempty" tf:"epsilon,omitempty"`
}

type ConfigParameters struct {

	// Specify the epsilon parameter (positive float)
	// +kubebuilder:validation:Optional
	Epsilon *float64 `json:"epsilon" tf:"epsilon,omitempty"`
}

type OutlierDetectorInitParameters struct {

	// The algorithm to use and its configuration. See https://grafana.com/docs/grafana-cloud/machine-learning/outlier-detection/ for details.
	Algorithm []AlgorithmInitParameters `json:"algorithm,omitempty" tf:"algorithm,omitempty"`

	// Reference to a DataSource in oss to populate datasourceUid.
	// +kubebuilder:validation:Optional
	DataSourceRef *v1.Reference `json:"dataSourceRef,omitempty" tf:"-"`

	// Selector for a DataSource in oss to populate datasourceUid.
	// +kubebuilder:validation:Optional
	DataSourceSelector *v1.Selector `json:"dataSourceSelector,omitempty" tf:"-"`

	// The id of the datasource to query.
	DatasourceID *float64 `json:"datasourceId,omitempty" tf:"datasource_id,omitempty"`

	// The type of datasource being queried. Currently allowed values are prometheus, graphite, loki, postgres, and datadog.
	DatasourceType *string `json:"datasourceType,omitempty" tf:"datasource_type,omitempty"`

	// The uid of the datasource to query.
	// +crossplane:generate:reference:type=github.com/grafana/crossplane-provider-grafana/apis/oss/v1alpha1.DataSource
	// +crossplane:generate:reference:extractor=github.com/grafana/crossplane-provider-grafana/config/grafana.UIDExtractor()
	// +crossplane:generate:reference:refFieldName=DataSourceRef
	// +crossplane:generate:reference:selectorFieldName=DataSourceSelector
	DatasourceUID *string `json:"datasourceUid,omitempty" tf:"datasource_uid,omitempty"`

	// A description of the outlier detector.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The data interval in seconds to monitor. Defaults to `300`.
	Interval *float64 `json:"interval,omitempty" tf:"interval,omitempty"`

	// The metric used to query the outlier detector results.
	Metric *string `json:"metric,omitempty" tf:"metric,omitempty"`

	// The name of the outlier detector.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// An object representing the query params to query Grafana with.
	// +mapType=granular
	QueryParams map[string]*string `json:"queryParams,omitempty" tf:"query_params,omitempty"`
}

type OutlierDetectorObservation struct {

	// The algorithm to use and its configuration. See https://grafana.com/docs/grafana-cloud/machine-learning/outlier-detection/ for details.
	Algorithm []AlgorithmObservation `json:"algorithm,omitempty" tf:"algorithm,omitempty"`

	// The id of the datasource to query.
	DatasourceID *float64 `json:"datasourceId,omitempty" tf:"datasource_id,omitempty"`

	// The type of datasource being queried. Currently allowed values are prometheus, graphite, loki, postgres, and datadog.
	DatasourceType *string `json:"datasourceType,omitempty" tf:"datasource_type,omitempty"`

	// The uid of the datasource to query.
	DatasourceUID *string `json:"datasourceUid,omitempty" tf:"datasource_uid,omitempty"`

	// A description of the outlier detector.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The data interval in seconds to monitor. Defaults to `300`.
	Interval *float64 `json:"interval,omitempty" tf:"interval,omitempty"`

	// The metric used to query the outlier detector results.
	Metric *string `json:"metric,omitempty" tf:"metric,omitempty"`

	// The name of the outlier detector.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// An object representing the query params to query Grafana with.
	// +mapType=granular
	QueryParams map[string]*string `json:"queryParams,omitempty" tf:"query_params,omitempty"`
}

type OutlierDetectorParameters struct {

	// The algorithm to use and its configuration. See https://grafana.com/docs/grafana-cloud/machine-learning/outlier-detection/ for details.
	// +kubebuilder:validation:Optional
	Algorithm []AlgorithmParameters `json:"algorithm,omitempty" tf:"algorithm,omitempty"`

	// Reference to a DataSource in oss to populate datasourceUid.
	// +kubebuilder:validation:Optional
	DataSourceRef *v1.Reference `json:"dataSourceRef,omitempty" tf:"-"`

	// Selector for a DataSource in oss to populate datasourceUid.
	// +kubebuilder:validation:Optional
	DataSourceSelector *v1.Selector `json:"dataSourceSelector,omitempty" tf:"-"`

	// The id of the datasource to query.
	// +kubebuilder:validation:Optional
	DatasourceID *float64 `json:"datasourceId,omitempty" tf:"datasource_id,omitempty"`

	// The type of datasource being queried. Currently allowed values are prometheus, graphite, loki, postgres, and datadog.
	// +kubebuilder:validation:Optional
	DatasourceType *string `json:"datasourceType,omitempty" tf:"datasource_type,omitempty"`

	// The uid of the datasource to query.
	// +crossplane:generate:reference:type=github.com/grafana/crossplane-provider-grafana/apis/oss/v1alpha1.DataSource
	// +crossplane:generate:reference:extractor=github.com/grafana/crossplane-provider-grafana/config/grafana.UIDExtractor()
	// +crossplane:generate:reference:refFieldName=DataSourceRef
	// +crossplane:generate:reference:selectorFieldName=DataSourceSelector
	// +kubebuilder:validation:Optional
	DatasourceUID *string `json:"datasourceUid,omitempty" tf:"datasource_uid,omitempty"`

	// A description of the outlier detector.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The data interval in seconds to monitor. Defaults to `300`.
	// +kubebuilder:validation:Optional
	Interval *float64 `json:"interval,omitempty" tf:"interval,omitempty"`

	// The metric used to query the outlier detector results.
	// +kubebuilder:validation:Optional
	Metric *string `json:"metric,omitempty" tf:"metric,omitempty"`

	// The name of the outlier detector.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// An object representing the query params to query Grafana with.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	QueryParams map[string]*string `json:"queryParams,omitempty" tf:"query_params,omitempty"`
}

// OutlierDetectorSpec defines the desired state of OutlierDetector
type OutlierDetectorSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OutlierDetectorParameters `json:"forProvider"`
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
	InitProvider OutlierDetectorInitParameters `json:"initProvider,omitempty"`
}

// OutlierDetectorStatus defines the observed state of OutlierDetector.
type OutlierDetectorStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OutlierDetectorObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// OutlierDetector is the Schema for the OutlierDetectors API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,grafana}
type OutlierDetector struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.algorithm) || (has(self.initProvider) && has(self.initProvider.algorithm))",message="spec.forProvider.algorithm is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.datasourceType) || (has(self.initProvider) && has(self.initProvider.datasourceType))",message="spec.forProvider.datasourceType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.metric) || (has(self.initProvider) && has(self.initProvider.metric))",message="spec.forProvider.metric is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.queryParams) || (has(self.initProvider) && has(self.initProvider.queryParams))",message="spec.forProvider.queryParams is a required parameter"
	Spec   OutlierDetectorSpec   `json:"spec"`
	Status OutlierDetectorStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OutlierDetectorList contains a list of OutlierDetectors
type OutlierDetectorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OutlierDetector `json:"items"`
}

// Repository type metadata.
var (
	OutlierDetector_Kind             = "OutlierDetector"
	OutlierDetector_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: OutlierDetector_Kind}.String()
	OutlierDetector_KindAPIVersion   = OutlierDetector_Kind + "." + CRDGroupVersion.String()
	OutlierDetector_GroupVersionKind = CRDGroupVersion.WithKind(OutlierDetector_Kind)
)

func init() {
	SchemeBuilder.Register(&OutlierDetector{}, &OutlierDetectorList{})
}
