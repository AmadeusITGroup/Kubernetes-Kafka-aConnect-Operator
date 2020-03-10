// Copyright 2020 The Amadeus Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/AmadeusITGroup/Kubernetes-Kafka-Connect-Operator/pkg/apis/kafkaconnect/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeKafkaConnects implements KafkaConnectInterface
type FakeKafkaConnects struct {
	Fake *FakeKafkaconnectV1alpha1
	ns   string
}

var kafkaconnectsResource = schema.GroupVersionResource{Group: "kafkaconnect.operator.io", Version: "v1alpha1", Resource: "kafkaconnects"}

var kafkaconnectsKind = schema.GroupVersionKind{Group: "kafkaconnect.operator.io", Version: "v1alpha1", Kind: "KafkaConnect"}

// Get takes name of the kafkaConnect, and returns the corresponding kafkaConnect object, and an error if there is any.
func (c *FakeKafkaConnects) Get(name string, options v1.GetOptions) (result *v1alpha1.KafkaConnect, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(kafkaconnectsResource, c.ns, name), &v1alpha1.KafkaConnect{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KafkaConnect), err
}

// List takes label and field selectors, and returns the list of KafkaConnects that match those selectors.
func (c *FakeKafkaConnects) List(opts v1.ListOptions) (result *v1alpha1.KafkaConnectList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(kafkaconnectsResource, kafkaconnectsKind, c.ns, opts), &v1alpha1.KafkaConnectList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.KafkaConnectList{ListMeta: obj.(*v1alpha1.KafkaConnectList).ListMeta}
	for _, item := range obj.(*v1alpha1.KafkaConnectList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested kafkaConnects.
func (c *FakeKafkaConnects) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(kafkaconnectsResource, c.ns, opts))

}

// Create takes the representation of a kafkaConnect and creates it.  Returns the server's representation of the kafkaConnect, and an error, if there is any.
func (c *FakeKafkaConnects) Create(kafkaConnect *v1alpha1.KafkaConnect) (result *v1alpha1.KafkaConnect, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(kafkaconnectsResource, c.ns, kafkaConnect), &v1alpha1.KafkaConnect{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KafkaConnect), err
}

// Update takes the representation of a kafkaConnect and updates it. Returns the server's representation of the kafkaConnect, and an error, if there is any.
func (c *FakeKafkaConnects) Update(kafkaConnect *v1alpha1.KafkaConnect) (result *v1alpha1.KafkaConnect, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(kafkaconnectsResource, c.ns, kafkaConnect), &v1alpha1.KafkaConnect{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KafkaConnect), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeKafkaConnects) UpdateStatus(kafkaConnect *v1alpha1.KafkaConnect) (*v1alpha1.KafkaConnect, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(kafkaconnectsResource, "status", c.ns, kafkaConnect), &v1alpha1.KafkaConnect{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KafkaConnect), err
}

// Delete takes name of the kafkaConnect and deletes it. Returns an error if one occurs.
func (c *FakeKafkaConnects) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(kafkaconnectsResource, c.ns, name), &v1alpha1.KafkaConnect{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKafkaConnects) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(kafkaconnectsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.KafkaConnectList{})
	return err
}

// Patch applies the patch and returns the patched kafkaConnect.
func (c *FakeKafkaConnects) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.KafkaConnect, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(kafkaconnectsResource, c.ns, name, pt, data, subresources...), &v1alpha1.KafkaConnect{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KafkaConnect), err
}
