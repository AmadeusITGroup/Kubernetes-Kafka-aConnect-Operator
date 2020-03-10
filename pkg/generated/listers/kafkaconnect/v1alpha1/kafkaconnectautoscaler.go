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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/AmadeusITGroup/Kubernetes-Kafka-Connect-Operator/pkg/apis/kafkaconnect/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// KafkaConnectAutoScalerLister helps list KafkaConnectAutoScalers.
type KafkaConnectAutoScalerLister interface {
	// List lists all KafkaConnectAutoScalers in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.KafkaConnectAutoScaler, err error)
	// KafkaConnectAutoScalers returns an object that can list and get KafkaConnectAutoScalers.
	KafkaConnectAutoScalers(namespace string) KafkaConnectAutoScalerNamespaceLister
	KafkaConnectAutoScalerListerExpansion
}

// kafkaConnectAutoScalerLister implements the KafkaConnectAutoScalerLister interface.
type kafkaConnectAutoScalerLister struct {
	indexer cache.Indexer
}

// NewKafkaConnectAutoScalerLister returns a new KafkaConnectAutoScalerLister.
func NewKafkaConnectAutoScalerLister(indexer cache.Indexer) KafkaConnectAutoScalerLister {
	return &kafkaConnectAutoScalerLister{indexer: indexer}
}

// List lists all KafkaConnectAutoScalers in the indexer.
func (s *kafkaConnectAutoScalerLister) List(selector labels.Selector) (ret []*v1alpha1.KafkaConnectAutoScaler, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.KafkaConnectAutoScaler))
	})
	return ret, err
}

// KafkaConnectAutoScalers returns an object that can list and get KafkaConnectAutoScalers.
func (s *kafkaConnectAutoScalerLister) KafkaConnectAutoScalers(namespace string) KafkaConnectAutoScalerNamespaceLister {
	return kafkaConnectAutoScalerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// KafkaConnectAutoScalerNamespaceLister helps list and get KafkaConnectAutoScalers.
type KafkaConnectAutoScalerNamespaceLister interface {
	// List lists all KafkaConnectAutoScalers in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.KafkaConnectAutoScaler, err error)
	// Get retrieves the KafkaConnectAutoScaler from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.KafkaConnectAutoScaler, error)
	KafkaConnectAutoScalerNamespaceListerExpansion
}

// kafkaConnectAutoScalerNamespaceLister implements the KafkaConnectAutoScalerNamespaceLister
// interface.
type kafkaConnectAutoScalerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all KafkaConnectAutoScalers in the indexer for a given namespace.
func (s kafkaConnectAutoScalerNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.KafkaConnectAutoScaler, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.KafkaConnectAutoScaler))
	})
	return ret, err
}

// Get retrieves the KafkaConnectAutoScaler from the indexer for a given namespace and name.
func (s kafkaConnectAutoScalerNamespaceLister) Get(name string) (*v1alpha1.KafkaConnectAutoScaler, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("kafkaconnectautoscaler"), name)
	}
	return obj.(*v1alpha1.KafkaConnectAutoScaler), nil
}
