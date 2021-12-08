/*
Copyright 2019-2021 OpenFaaS Authors

Licensed under the MIT license. See LICENSE file in the project root for full license information.
*/

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/openfaas/faas-netes/pkg/apis/openfaas/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// FunctionLister helps list Functions.
type FunctionLister interface {
	// List lists all Functions in the indexer.
	List(selector labels.Selector) (ret []*v1.Function, err error)
	// Functions returns an object that can list and get Functions.
	Functions(namespace string) FunctionNamespaceLister
	FunctionListerExpansion
}

// functionLister implements the FunctionLister interface.
type functionLister struct {
	indexer cache.Indexer
}

// NewFunctionLister returns a new FunctionLister.
func NewFunctionLister(indexer cache.Indexer) FunctionLister {
	return &functionLister{indexer: indexer}
}

// List lists all Functions in the indexer.
func (s *functionLister) List(selector labels.Selector) (ret []*v1.Function, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Function))
	})
	return ret, err
}

// Functions returns an object that can list and get Functions.
func (s *functionLister) Functions(namespace string) FunctionNamespaceLister {
	return functionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FunctionNamespaceLister helps list and get Functions.
type FunctionNamespaceLister interface {
	// List lists all Functions in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.Function, err error)
	// Get retrieves the Function from the indexer for a given namespace and name.
	Get(name string) (*v1.Function, error)
	FunctionNamespaceListerExpansion
}

// functionNamespaceLister implements the FunctionNamespaceLister
// interface.
type functionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Functions in the indexer for a given namespace.
func (s functionNamespaceLister) List(selector labels.Selector) (ret []*v1.Function, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Function))
	})
	return ret, err
}

// Get retrieves the Function from the indexer for a given namespace and name.
func (s functionNamespaceLister) Get(name string) (*v1.Function, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("function"), name)
	}
	return obj.(*v1.Function), nil
}
