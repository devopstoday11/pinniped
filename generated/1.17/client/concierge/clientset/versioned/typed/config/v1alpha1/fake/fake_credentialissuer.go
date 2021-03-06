// Copyright 2020 the Pinniped contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "go.pinniped.dev/generated/1.17/apis/concierge/config/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCredentialIssuers implements CredentialIssuerInterface
type FakeCredentialIssuers struct {
	Fake *FakeConfigV1alpha1
	ns   string
}

var credentialissuersResource = schema.GroupVersionResource{Group: "config.concierge.pinniped.dev", Version: "v1alpha1", Resource: "credentialissuers"}

var credentialissuersKind = schema.GroupVersionKind{Group: "config.concierge.pinniped.dev", Version: "v1alpha1", Kind: "CredentialIssuer"}

// Get takes name of the credentialIssuer, and returns the corresponding credentialIssuer object, and an error if there is any.
func (c *FakeCredentialIssuers) Get(name string, options v1.GetOptions) (result *v1alpha1.CredentialIssuer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(credentialissuersResource, c.ns, name), &v1alpha1.CredentialIssuer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CredentialIssuer), err
}

// List takes label and field selectors, and returns the list of CredentialIssuers that match those selectors.
func (c *FakeCredentialIssuers) List(opts v1.ListOptions) (result *v1alpha1.CredentialIssuerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(credentialissuersResource, credentialissuersKind, c.ns, opts), &v1alpha1.CredentialIssuerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CredentialIssuerList{ListMeta: obj.(*v1alpha1.CredentialIssuerList).ListMeta}
	for _, item := range obj.(*v1alpha1.CredentialIssuerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested credentialIssuers.
func (c *FakeCredentialIssuers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(credentialissuersResource, c.ns, opts))

}

// Create takes the representation of a credentialIssuer and creates it.  Returns the server's representation of the credentialIssuer, and an error, if there is any.
func (c *FakeCredentialIssuers) Create(credentialIssuer *v1alpha1.CredentialIssuer) (result *v1alpha1.CredentialIssuer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(credentialissuersResource, c.ns, credentialIssuer), &v1alpha1.CredentialIssuer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CredentialIssuer), err
}

// Update takes the representation of a credentialIssuer and updates it. Returns the server's representation of the credentialIssuer, and an error, if there is any.
func (c *FakeCredentialIssuers) Update(credentialIssuer *v1alpha1.CredentialIssuer) (result *v1alpha1.CredentialIssuer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(credentialissuersResource, c.ns, credentialIssuer), &v1alpha1.CredentialIssuer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CredentialIssuer), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCredentialIssuers) UpdateStatus(credentialIssuer *v1alpha1.CredentialIssuer) (*v1alpha1.CredentialIssuer, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(credentialissuersResource, "status", c.ns, credentialIssuer), &v1alpha1.CredentialIssuer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CredentialIssuer), err
}

// Delete takes name of the credentialIssuer and deletes it. Returns an error if one occurs.
func (c *FakeCredentialIssuers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(credentialissuersResource, c.ns, name), &v1alpha1.CredentialIssuer{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCredentialIssuers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(credentialissuersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.CredentialIssuerList{})
	return err
}

// Patch applies the patch and returns the patched credentialIssuer.
func (c *FakeCredentialIssuers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CredentialIssuer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(credentialissuersResource, c.ns, name, pt, data, subresources...), &v1alpha1.CredentialIssuer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CredentialIssuer), err
}
