/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	v1alpha1 "github.com/grafana/crossplane-provider-grafana/apis/cloud/v1alpha1"
	grafana "github.com/grafana/crossplane-provider-grafana/config/grafana"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this APIKey.
func (mg *APIKey) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CloudStackSlug),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.CloudStackRef,
		Selector:     mg.Spec.ForProvider.CloudStackSelector,
		To: reference.To{
			List:    &v1alpha1.StackList{},
			Managed: &v1alpha1.Stack{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CloudStackSlug")
	}
	mg.Spec.ForProvider.CloudStackSlug = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CloudStackRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Dashboard.
func (mg *Dashboard) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Folder),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.FolderRef,
		Selector:     mg.Spec.ForProvider.FolderSelector,
		To: reference.To{
			List:    &FolderList{},
			Managed: &Folder{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Folder")
	}
	mg.Spec.ForProvider.Folder = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.FolderRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Team.
func (mg *Team) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var mrsp reference.MultiResolutionResponse
	var err error

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.Members),
		Extract:       grafana.UserEmailExtractor(),
		References:    mg.Spec.ForProvider.MemberRefs,
		Selector:      mg.Spec.ForProvider.MemberSelector,
		To: reference.To{
			List:    &UserList{},
			Managed: &User{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Members")
	}
	mg.Spec.ForProvider.Members = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.MemberRefs = mrsp.ResolvedReferences

	return nil
}
