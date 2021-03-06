/*


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
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	multitenancyv1 "multitenancy.metricsadvisor.ai/apiservice/apis/multitenancy/v1"
)

// FakeApiServices implements ApiServiceInterface
type FakeApiServices struct {
	Fake *FakeMultitenancyV1
	ns   string
}

var apiservicesResource = schema.GroupVersionResource{Group: "multitenancy.metricsadvisor.ai", Version: "v1", Resource: "apiservices"}

var apiservicesKind = schema.GroupVersionKind{Group: "multitenancy.metricsadvisor.ai", Version: "v1", Kind: "ApiService"}

// Get takes name of the apiService, and returns the corresponding apiService object, and an error if there is any.
func (c *FakeApiServices) Get(ctx context.Context, name string, options v1.GetOptions) (result *multitenancyv1.ApiService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(apiservicesResource, c.ns, name), &multitenancyv1.ApiService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*multitenancyv1.ApiService), err
}

// List takes label and field selectors, and returns the list of ApiServices that match those selectors.
func (c *FakeApiServices) List(ctx context.Context, opts v1.ListOptions) (result *multitenancyv1.ApiServiceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(apiservicesResource, apiservicesKind, c.ns, opts), &multitenancyv1.ApiServiceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &multitenancyv1.ApiServiceList{ListMeta: obj.(*multitenancyv1.ApiServiceList).ListMeta}
	for _, item := range obj.(*multitenancyv1.ApiServiceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested apiServices.
func (c *FakeApiServices) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(apiservicesResource, c.ns, opts))

}

// Create takes the representation of a apiService and creates it.  Returns the server's representation of the apiService, and an error, if there is any.
func (c *FakeApiServices) Create(ctx context.Context, apiService *multitenancyv1.ApiService, opts v1.CreateOptions) (result *multitenancyv1.ApiService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(apiservicesResource, c.ns, apiService), &multitenancyv1.ApiService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*multitenancyv1.ApiService), err
}

// Update takes the representation of a apiService and updates it. Returns the server's representation of the apiService, and an error, if there is any.
func (c *FakeApiServices) Update(ctx context.Context, apiService *multitenancyv1.ApiService, opts v1.UpdateOptions) (result *multitenancyv1.ApiService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(apiservicesResource, c.ns, apiService), &multitenancyv1.ApiService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*multitenancyv1.ApiService), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeApiServices) UpdateStatus(ctx context.Context, apiService *multitenancyv1.ApiService, opts v1.UpdateOptions) (*multitenancyv1.ApiService, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(apiservicesResource, "status", c.ns, apiService), &multitenancyv1.ApiService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*multitenancyv1.ApiService), err
}

// Delete takes name of the apiService and deletes it. Returns an error if one occurs.
func (c *FakeApiServices) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(apiservicesResource, c.ns, name), &multitenancyv1.ApiService{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeApiServices) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(apiservicesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &multitenancyv1.ApiServiceList{})
	return err
}

// Patch applies the patch and returns the patched apiService.
func (c *FakeApiServices) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *multitenancyv1.ApiService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(apiservicesResource, c.ns, name, pt, data, subresources...), &multitenancyv1.ApiService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*multitenancyv1.ApiService), err
}
