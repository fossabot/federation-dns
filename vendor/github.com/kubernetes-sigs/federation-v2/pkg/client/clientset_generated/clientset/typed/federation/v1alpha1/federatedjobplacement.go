/*
Copyright 2018 The Federation v2 Authors.

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
package v1alpha1

import (
	v1alpha1 "github.com/kubernetes-sigs/federation-v2/pkg/apis/federation/v1alpha1"
	scheme "github.com/kubernetes-sigs/federation-v2/pkg/client/clientset_generated/clientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// FederatedJobPlacementsGetter has a method to return a FederatedJobPlacementInterface.
// A group's client should implement this interface.
type FederatedJobPlacementsGetter interface {
	FederatedJobPlacements(namespace string) FederatedJobPlacementInterface
}

// FederatedJobPlacementInterface has methods to work with FederatedJobPlacement resources.
type FederatedJobPlacementInterface interface {
	Create(*v1alpha1.FederatedJobPlacement) (*v1alpha1.FederatedJobPlacement, error)
	Update(*v1alpha1.FederatedJobPlacement) (*v1alpha1.FederatedJobPlacement, error)
	UpdateStatus(*v1alpha1.FederatedJobPlacement) (*v1alpha1.FederatedJobPlacement, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.FederatedJobPlacement, error)
	List(opts v1.ListOptions) (*v1alpha1.FederatedJobPlacementList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.FederatedJobPlacement, err error)
	FederatedJobPlacementExpansion
}

// federatedJobPlacements implements FederatedJobPlacementInterface
type federatedJobPlacements struct {
	client rest.Interface
	ns     string
}

// newFederatedJobPlacements returns a FederatedJobPlacements
func newFederatedJobPlacements(c *FederationV1alpha1Client, namespace string) *federatedJobPlacements {
	return &federatedJobPlacements{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the federatedJobPlacement, and returns the corresponding federatedJobPlacement object, and an error if there is any.
func (c *federatedJobPlacements) Get(name string, options v1.GetOptions) (result *v1alpha1.FederatedJobPlacement, err error) {
	result = &v1alpha1.FederatedJobPlacement{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("federatedjobplacements").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of FederatedJobPlacements that match those selectors.
func (c *federatedJobPlacements) List(opts v1.ListOptions) (result *v1alpha1.FederatedJobPlacementList, err error) {
	result = &v1alpha1.FederatedJobPlacementList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("federatedjobplacements").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested federatedJobPlacements.
func (c *federatedJobPlacements) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("federatedjobplacements").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a federatedJobPlacement and creates it.  Returns the server's representation of the federatedJobPlacement, and an error, if there is any.
func (c *federatedJobPlacements) Create(federatedJobPlacement *v1alpha1.FederatedJobPlacement) (result *v1alpha1.FederatedJobPlacement, err error) {
	result = &v1alpha1.FederatedJobPlacement{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("federatedjobplacements").
		Body(federatedJobPlacement).
		Do().
		Into(result)
	return
}

// Update takes the representation of a federatedJobPlacement and updates it. Returns the server's representation of the federatedJobPlacement, and an error, if there is any.
func (c *federatedJobPlacements) Update(federatedJobPlacement *v1alpha1.FederatedJobPlacement) (result *v1alpha1.FederatedJobPlacement, err error) {
	result = &v1alpha1.FederatedJobPlacement{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("federatedjobplacements").
		Name(federatedJobPlacement.Name).
		Body(federatedJobPlacement).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *federatedJobPlacements) UpdateStatus(federatedJobPlacement *v1alpha1.FederatedJobPlacement) (result *v1alpha1.FederatedJobPlacement, err error) {
	result = &v1alpha1.FederatedJobPlacement{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("federatedjobplacements").
		Name(federatedJobPlacement.Name).
		SubResource("status").
		Body(federatedJobPlacement).
		Do().
		Into(result)
	return
}

// Delete takes name of the federatedJobPlacement and deletes it. Returns an error if one occurs.
func (c *federatedJobPlacements) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("federatedjobplacements").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *federatedJobPlacements) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("federatedjobplacements").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched federatedJobPlacement.
func (c *federatedJobPlacements) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.FederatedJobPlacement, err error) {
	result = &v1alpha1.FederatedJobPlacement{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("federatedjobplacements").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
