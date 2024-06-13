/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Escalation.
func (mg *Escalation) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EscalationChainID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.EscalationChainRef,
		Selector:     mg.Spec.ForProvider.EscalationChainSelector,
		To: reference.To{
			List:    &EscalationChainList{},
			Managed: &EscalationChain{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.EscalationChainID")
	}
	mg.Spec.ForProvider.EscalationChainID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.EscalationChainRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.EscalationChainID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.EscalationChainRef,
		Selector:     mg.Spec.InitProvider.EscalationChainSelector,
		To: reference.To{
			List:    &EscalationChainList{},
			Managed: &EscalationChain{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.EscalationChainID")
	}
	mg.Spec.InitProvider.EscalationChainID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.EscalationChainRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Integration.
func (mg *Integration) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.DefaultRoute); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DefaultRoute[i3].EscalationChainID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.DefaultRoute[i3].EscalationChainRef,
			Selector:     mg.Spec.ForProvider.DefaultRoute[i3].EscalationChainSelector,
			To: reference.To{
				List:    &EscalationChainList{},
				Managed: &EscalationChain{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.DefaultRoute[i3].EscalationChainID")
		}
		mg.Spec.ForProvider.DefaultRoute[i3].EscalationChainID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.DefaultRoute[i3].EscalationChainRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.DefaultRoute); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.DefaultRoute[i3].EscalationChainID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.DefaultRoute[i3].EscalationChainRef,
			Selector:     mg.Spec.InitProvider.DefaultRoute[i3].EscalationChainSelector,
			To: reference.To{
				List:    &EscalationChainList{},
				Managed: &EscalationChain{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.DefaultRoute[i3].EscalationChainID")
		}
		mg.Spec.InitProvider.DefaultRoute[i3].EscalationChainID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.DefaultRoute[i3].EscalationChainRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this Route.
func (mg *Route) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EscalationChainID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.EscalationChainRef,
		Selector:     mg.Spec.ForProvider.EscalationChainSelector,
		To: reference.To{
			List:    &EscalationChainList{},
			Managed: &EscalationChain{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.EscalationChainID")
	}
	mg.Spec.ForProvider.EscalationChainID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.EscalationChainRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.IntegrationID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.IntegrationRef,
		Selector:     mg.Spec.ForProvider.IntegrationSelector,
		To: reference.To{
			List:    &IntegrationList{},
			Managed: &Integration{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.IntegrationID")
	}
	mg.Spec.ForProvider.IntegrationID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.IntegrationRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.EscalationChainID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.EscalationChainRef,
		Selector:     mg.Spec.InitProvider.EscalationChainSelector,
		To: reference.To{
			List:    &EscalationChainList{},
			Managed: &EscalationChain{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.EscalationChainID")
	}
	mg.Spec.InitProvider.EscalationChainID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.EscalationChainRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.IntegrationID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.IntegrationRef,
		Selector:     mg.Spec.InitProvider.IntegrationSelector,
		To: reference.To{
			List:    &IntegrationList{},
			Managed: &Integration{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.IntegrationID")
	}
	mg.Spec.InitProvider.IntegrationID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.IntegrationRef = rsp.ResolvedReference

	return nil
}
