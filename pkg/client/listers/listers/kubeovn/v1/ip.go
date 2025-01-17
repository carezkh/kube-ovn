/*
Copyright The Kubernetes Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/kubeovn/kube-ovn/pkg/apis/kubeovn/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// IPLister helps list IPs.
type IPLister interface {
	// List lists all IPs in the indexer.
	List(selector labels.Selector) (ret []*v1.IP, err error)
	// Get retrieves the IP from the index for a given name.
	Get(name string) (*v1.IP, error)
	IPListerExpansion
}

// iPLister implements the IPLister interface.
type iPLister struct {
	indexer cache.Indexer
}

// NewIPLister returns a new IPLister.
func NewIPLister(indexer cache.Indexer) IPLister {
	return &iPLister{indexer: indexer}
}

// List lists all IPs in the indexer.
func (s *iPLister) List(selector labels.Selector) (ret []*v1.IP, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.IP))
	})
	return ret, err
}

// Get retrieves the IP from the index for a given name.
func (s *iPLister) Get(name string) (*v1.IP, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("ip"), name)
	}
	return obj.(*v1.IP), nil
}
