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

package v1alpha1

import (
	kcpcache "github.com/kcp-dev/apimachinery/pkg/cache"
	"github.com/kcp-dev/logicalcluster/v2"

	nodev1alpha1 "k8s.io/api/node/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	nodev1alpha1listers "k8s.io/client-go/listers/node/v1alpha1"
	"k8s.io/client-go/tools/cache"
)

// RuntimeClassClusterLister can list RuntimeClasses across all workspaces, or scope down to a RuntimeClassLister for one workspace.
type RuntimeClassClusterLister interface {
	List(selector labels.Selector) (ret []*nodev1alpha1.RuntimeClass, err error)
	Cluster(cluster logicalcluster.Name) nodev1alpha1listers.RuntimeClassLister
}

type runtimeClassClusterLister struct {
	indexer cache.Indexer
}

// NewRuntimeClassClusterLister returns a new RuntimeClassClusterLister.
func NewRuntimeClassClusterLister(indexer cache.Indexer) *runtimeClassClusterLister {
	return &runtimeClassClusterLister{indexer: indexer}
}

// List lists all RuntimeClasses in the indexer across all workspaces.
func (s *runtimeClassClusterLister) List(selector labels.Selector) (ret []*nodev1alpha1.RuntimeClass, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*nodev1alpha1.RuntimeClass))
	})
	return ret, err
}

// Cluster scopes the lister to one workspace, allowing users to list and get RuntimeClasses.
func (s *runtimeClassClusterLister) Cluster(cluster logicalcluster.Name) nodev1alpha1listers.RuntimeClassLister {
	return &runtimeClassLister{indexer: s.indexer, cluster: cluster}
}

// runtimeClassLister implements the nodev1alpha1listers.RuntimeClassLister interface.
type runtimeClassLister struct {
	indexer cache.Indexer
	cluster logicalcluster.Name
}

// List lists all RuntimeClasses in the indexer for a workspace.
func (s *runtimeClassLister) List(selector labels.Selector) (ret []*nodev1alpha1.RuntimeClass, err error) {
	err = kcpcache.ListAllByCluster(s.indexer, s.cluster, selector, func(i interface{}) {
		ret = append(ret, i.(*nodev1alpha1.RuntimeClass))
	})
	return ret, err
}

// Get retrieves the RuntimeClass from the indexer for a given workspace and name.
func (s *runtimeClassLister) Get(name string) (*nodev1alpha1.RuntimeClass, error) {
	key := kcpcache.ToClusterAwareKey(s.cluster.String(), "", name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(nodev1alpha1.Resource("RuntimeClass"), name)
	}
	return obj.(*nodev1alpha1.RuntimeClass), nil
}
