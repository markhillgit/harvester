/*
Copyright 2025 Rancher Labs, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by main. DO NOT EDIT.

package fake

import (
	"context"

	v3 "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeGlobalDnsProviders implements GlobalDnsProviderInterface
type FakeGlobalDnsProviders struct {
	Fake *FakeManagementV3
	ns   string
}

var globaldnsprovidersResource = v3.SchemeGroupVersion.WithResource("globaldnsproviders")

var globaldnsprovidersKind = v3.SchemeGroupVersion.WithKind("GlobalDnsProvider")

// Get takes name of the globalDnsProvider, and returns the corresponding globalDnsProvider object, and an error if there is any.
func (c *FakeGlobalDnsProviders) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.GlobalDnsProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(globaldnsprovidersResource, c.ns, name), &v3.GlobalDnsProvider{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.GlobalDnsProvider), err
}

// List takes label and field selectors, and returns the list of GlobalDnsProviders that match those selectors.
func (c *FakeGlobalDnsProviders) List(ctx context.Context, opts v1.ListOptions) (result *v3.GlobalDnsProviderList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(globaldnsprovidersResource, globaldnsprovidersKind, c.ns, opts), &v3.GlobalDnsProviderList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v3.GlobalDnsProviderList{ListMeta: obj.(*v3.GlobalDnsProviderList).ListMeta}
	for _, item := range obj.(*v3.GlobalDnsProviderList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested globalDnsProviders.
func (c *FakeGlobalDnsProviders) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(globaldnsprovidersResource, c.ns, opts))

}

// Create takes the representation of a globalDnsProvider and creates it.  Returns the server's representation of the globalDnsProvider, and an error, if there is any.
func (c *FakeGlobalDnsProviders) Create(ctx context.Context, globalDnsProvider *v3.GlobalDnsProvider, opts v1.CreateOptions) (result *v3.GlobalDnsProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(globaldnsprovidersResource, c.ns, globalDnsProvider), &v3.GlobalDnsProvider{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.GlobalDnsProvider), err
}

// Update takes the representation of a globalDnsProvider and updates it. Returns the server's representation of the globalDnsProvider, and an error, if there is any.
func (c *FakeGlobalDnsProviders) Update(ctx context.Context, globalDnsProvider *v3.GlobalDnsProvider, opts v1.UpdateOptions) (result *v3.GlobalDnsProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(globaldnsprovidersResource, c.ns, globalDnsProvider), &v3.GlobalDnsProvider{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.GlobalDnsProvider), err
}

// Delete takes name of the globalDnsProvider and deletes it. Returns an error if one occurs.
func (c *FakeGlobalDnsProviders) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(globaldnsprovidersResource, c.ns, name, opts), &v3.GlobalDnsProvider{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGlobalDnsProviders) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(globaldnsprovidersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v3.GlobalDnsProviderList{})
	return err
}

// Patch applies the patch and returns the patched globalDnsProvider.
func (c *FakeGlobalDnsProviders) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.GlobalDnsProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(globaldnsprovidersResource, c.ns, name, pt, data, subresources...), &v3.GlobalDnsProvider{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.GlobalDnsProvider), err
}
