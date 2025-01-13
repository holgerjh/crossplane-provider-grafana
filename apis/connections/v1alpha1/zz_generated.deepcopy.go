//go:build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricsEndpointScrapeJob) DeepCopyInto(out *MetricsEndpointScrapeJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricsEndpointScrapeJob.
func (in *MetricsEndpointScrapeJob) DeepCopy() *MetricsEndpointScrapeJob {
	if in == nil {
		return nil
	}
	out := new(MetricsEndpointScrapeJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MetricsEndpointScrapeJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricsEndpointScrapeJobInitParameters) DeepCopyInto(out *MetricsEndpointScrapeJobInitParameters) {
	*out = *in
	if in.AuthenticationBasicPasswordSecretRef != nil {
		in, out := &in.AuthenticationBasicPasswordSecretRef, &out.AuthenticationBasicPasswordSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.AuthenticationBasicUsername != nil {
		in, out := &in.AuthenticationBasicUsername, &out.AuthenticationBasicUsername
		*out = new(string)
		**out = **in
	}
	if in.AuthenticationBearerTokenSecretRef != nil {
		in, out := &in.AuthenticationBearerTokenSecretRef, &out.AuthenticationBearerTokenSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.AuthenticationMethod != nil {
		in, out := &in.AuthenticationMethod, &out.AuthenticationMethod
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ScrapeIntervalSeconds != nil {
		in, out := &in.ScrapeIntervalSeconds, &out.ScrapeIntervalSeconds
		*out = new(float64)
		**out = **in
	}
	if in.StackID != nil {
		in, out := &in.StackID, &out.StackID
		*out = new(string)
		**out = **in
	}
	if in.URL != nil {
		in, out := &in.URL, &out.URL
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricsEndpointScrapeJobInitParameters.
func (in *MetricsEndpointScrapeJobInitParameters) DeepCopy() *MetricsEndpointScrapeJobInitParameters {
	if in == nil {
		return nil
	}
	out := new(MetricsEndpointScrapeJobInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricsEndpointScrapeJobList) DeepCopyInto(out *MetricsEndpointScrapeJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MetricsEndpointScrapeJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricsEndpointScrapeJobList.
func (in *MetricsEndpointScrapeJobList) DeepCopy() *MetricsEndpointScrapeJobList {
	if in == nil {
		return nil
	}
	out := new(MetricsEndpointScrapeJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MetricsEndpointScrapeJobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricsEndpointScrapeJobObservation) DeepCopyInto(out *MetricsEndpointScrapeJobObservation) {
	*out = *in
	if in.AuthenticationBasicUsername != nil {
		in, out := &in.AuthenticationBasicUsername, &out.AuthenticationBasicUsername
		*out = new(string)
		**out = **in
	}
	if in.AuthenticationMethod != nil {
		in, out := &in.AuthenticationMethod, &out.AuthenticationMethod
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ScrapeIntervalSeconds != nil {
		in, out := &in.ScrapeIntervalSeconds, &out.ScrapeIntervalSeconds
		*out = new(float64)
		**out = **in
	}
	if in.StackID != nil {
		in, out := &in.StackID, &out.StackID
		*out = new(string)
		**out = **in
	}
	if in.URL != nil {
		in, out := &in.URL, &out.URL
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricsEndpointScrapeJobObservation.
func (in *MetricsEndpointScrapeJobObservation) DeepCopy() *MetricsEndpointScrapeJobObservation {
	if in == nil {
		return nil
	}
	out := new(MetricsEndpointScrapeJobObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricsEndpointScrapeJobParameters) DeepCopyInto(out *MetricsEndpointScrapeJobParameters) {
	*out = *in
	if in.AuthenticationBasicPasswordSecretRef != nil {
		in, out := &in.AuthenticationBasicPasswordSecretRef, &out.AuthenticationBasicPasswordSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.AuthenticationBasicUsername != nil {
		in, out := &in.AuthenticationBasicUsername, &out.AuthenticationBasicUsername
		*out = new(string)
		**out = **in
	}
	if in.AuthenticationBearerTokenSecretRef != nil {
		in, out := &in.AuthenticationBearerTokenSecretRef, &out.AuthenticationBearerTokenSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.AuthenticationMethod != nil {
		in, out := &in.AuthenticationMethod, &out.AuthenticationMethod
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ScrapeIntervalSeconds != nil {
		in, out := &in.ScrapeIntervalSeconds, &out.ScrapeIntervalSeconds
		*out = new(float64)
		**out = **in
	}
	if in.StackID != nil {
		in, out := &in.StackID, &out.StackID
		*out = new(string)
		**out = **in
	}
	if in.URL != nil {
		in, out := &in.URL, &out.URL
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricsEndpointScrapeJobParameters.
func (in *MetricsEndpointScrapeJobParameters) DeepCopy() *MetricsEndpointScrapeJobParameters {
	if in == nil {
		return nil
	}
	out := new(MetricsEndpointScrapeJobParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricsEndpointScrapeJobSpec) DeepCopyInto(out *MetricsEndpointScrapeJobSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricsEndpointScrapeJobSpec.
func (in *MetricsEndpointScrapeJobSpec) DeepCopy() *MetricsEndpointScrapeJobSpec {
	if in == nil {
		return nil
	}
	out := new(MetricsEndpointScrapeJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricsEndpointScrapeJobStatus) DeepCopyInto(out *MetricsEndpointScrapeJobStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricsEndpointScrapeJobStatus.
func (in *MetricsEndpointScrapeJobStatus) DeepCopy() *MetricsEndpointScrapeJobStatus {
	if in == nil {
		return nil
	}
	out := new(MetricsEndpointScrapeJobStatus)
	in.DeepCopyInto(out)
	return out
}
