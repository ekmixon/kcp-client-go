//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The KCP Authors.

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

// Code generated by kcp code-generator. DO NOT EDIT.

package v1

import (
	kcpcache "github.com/kcp-dev/apimachinery/pkg/cache"
	"github.com/kcp-dev/logicalcluster/v2"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	corev1listers "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
)

// EventClusterLister can list Events across all workspaces, or scope down to a EventLister for one workspace.
type EventClusterLister interface {
	List(selector labels.Selector) (ret []*corev1.Event, err error)
	Cluster(cluster logicalcluster.Name) corev1listers.EventLister
}

type eventClusterLister struct {
	indexer cache.Indexer
}

// NewEventClusterLister returns a new EventClusterLister.
func NewEventClusterLister(indexer cache.Indexer) *eventClusterLister {
	return &eventClusterLister{indexer: indexer}
}

// List lists all Events in the indexer across all workspaces.
func (s *eventClusterLister) List(selector labels.Selector) (ret []*corev1.Event, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*corev1.Event))
	})
	return ret, err
}

// Cluster scopes the lister to one workspace, allowing users to list and get Events.
func (s *eventClusterLister) Cluster(cluster logicalcluster.Name) corev1listers.EventLister {
	return &eventLister{indexer: s.indexer, cluster: cluster}
}

// eventLister implements the corev1listers.EventLister interface.
type eventLister struct {
	indexer cache.Indexer
	cluster logicalcluster.Name
}

// List lists all Events in the indexer for a workspace.
func (s *eventLister) List(selector labels.Selector) (ret []*corev1.Event, err error) {
	err = kcpcache.ListAllByCluster(s.indexer, s.cluster, selector, func(i interface{}) {
		ret = append(ret, i.(*corev1.Event))
	})
	return ret, err
}

// Events returns an object that can list and get Events in one namespace.
func (s *eventLister) Events(namespace string) corev1listers.EventNamespaceLister {
	return &eventNamespaceLister{indexer: s.indexer, cluster: s.cluster, namespace: namespace}
}

// eventNamespaceLister implements the corev1listers.EventNamespaceLister interface.
type eventNamespaceLister struct {
	indexer   cache.Indexer
	cluster   logicalcluster.Name
	namespace string
}

// List lists all Events in the indexer for a given workspace and namespace.
func (s *eventNamespaceLister) List(selector labels.Selector) (ret []*corev1.Event, err error) {
	err = kcpcache.ListAllByClusterAndNamespace(s.indexer, s.cluster, s.namespace, selector, func(i interface{}) {
		ret = append(ret, i.(*corev1.Event))
	})
	return ret, err
}

// Get retrieves the Event from the indexer for a given workspace, namespace and name.
func (s *eventNamespaceLister) Get(name string) (*corev1.Event, error) {
	key := kcpcache.ToClusterAwareKey(s.cluster.String(), s.namespace, name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(corev1.Resource("Event"), name)
	}
	return obj.(*corev1.Event), nil
}
