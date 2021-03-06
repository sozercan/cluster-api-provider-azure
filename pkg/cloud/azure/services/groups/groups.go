/*
Copyright 2019 The Kubernetes Authors.

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

package groups

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2019-05-01/resources"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/pkg/errors"
	"k8s.io/klog"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/apis/azureprovider/v1alpha1"
)

// Get provides information about a resource group.
func (s *Service) Get(ctx context.Context, spec v1alpha1.ResourceSpec) (interface{}, error) {
	return s.Client.Get(ctx, s.Scope.ClusterConfig.ResourceGroup)
}

// // Reconcile gets/creates/updates a resource group.
func (s *Service) Reconcile(ctx context.Context, spec v1alpha1.ResourceSpec) error {
	klog.V(2).Infof("creating resource group %s", s.Scope.ClusterConfig.ResourceGroup)
	_, err := s.Client.CreateOrUpdate(ctx, s.Scope.ClusterConfig.ResourceGroup, resources.Group{Location: to.StringPtr(s.Scope.ClusterConfig.Location)})
	klog.V(2).Infof("successfully created resource group %s", s.Scope.ClusterConfig.ResourceGroup)
	return err
}

// Delete deletes the resource group with the provided name.
func (s *Service) Delete(ctx context.Context, spec v1alpha1.ResourceSpec) error {
	klog.V(2).Infof("deleting resource group %s", s.Scope.ClusterConfig.ResourceGroup)
	future, err := s.Client.Delete(ctx, s.Scope.ClusterConfig.ResourceGroup)
	if err != nil {
		return errors.Wrapf(err, "failed to delete resource group %s", s.Scope.ClusterConfig.ResourceGroup)
	}

	err = future.WaitForCompletionRef(ctx, s.Client.Client)
	if err != nil {
		return errors.Wrap(err, "cannot delete, future response")
	}

	_, err = future.Result(s.Client)

	klog.V(2).Infof("successfully deleted resource group %s", s.Scope.ClusterConfig.ResourceGroup)
	return err
}
