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

	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	rbacv1listers "k8s.io/client-go/listers/rbac/v1"
	"k8s.io/client-go/tools/cache"
)

// ClusterRoleClusterLister can list ClusterRoles across all workspaces, or scope down to a ClusterRoleLister for one workspace.
type ClusterRoleClusterLister interface {
	List(selector labels.Selector) (ret []*rbacv1.ClusterRole, err error)
	Cluster(cluster logicalcluster.Name) rbacv1listers.ClusterRoleLister
}

type clusterRoleClusterLister struct {
	indexer cache.Indexer
}

// NewClusterRoleClusterLister returns a new ClusterRoleClusterLister.
func NewClusterRoleClusterLister(indexer cache.Indexer) *clusterRoleClusterLister {
	return &clusterRoleClusterLister{indexer: indexer}
}

// List lists all ClusterRoles in the indexer across all workspaces.
func (s *clusterRoleClusterLister) List(selector labels.Selector) (ret []*rbacv1.ClusterRole, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*rbacv1.ClusterRole))
	})
	return ret, err
}

// Cluster scopes the lister to one workspace, allowing users to list and get ClusterRoles.
func (s *clusterRoleClusterLister) Cluster(cluster logicalcluster.Name) rbacv1listers.ClusterRoleLister {
	return &clusterRoleLister{indexer: s.indexer, cluster: cluster}
}

// clusterRoleLister implements the rbacv1listers.ClusterRoleLister interface.
type clusterRoleLister struct {
	indexer cache.Indexer
	cluster logicalcluster.Name
}

// List lists all ClusterRoles in the indexer for a workspace.
func (s *clusterRoleLister) List(selector labels.Selector) (ret []*rbacv1.ClusterRole, err error) {
	err = kcpcache.ListAllByCluster(s.indexer, s.cluster, selector, func(i interface{}) {
		ret = append(ret, i.(*rbacv1.ClusterRole))
	})
	return ret, err
}

// Get retrieves the ClusterRole from the indexer for a given workspace and name.
func (s *clusterRoleLister) Get(name string) (*rbacv1.ClusterRole, error) {
	key := kcpcache.ToClusterAwareKey(s.cluster.String(), "", name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(rbacv1.Resource("ClusterRole"), name)
	}
	return obj.(*rbacv1.ClusterRole), nil
}
