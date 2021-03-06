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

// This package imports things required by build scripts, to force `go mod` to see them as dependencies
package tools

import (
	_ "github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2019-07-01/compute/computeapi"       //nolint
	_ "github.com/Azure/azure-sdk-for-go/services/network/mgmt/2019-06-01/network/networkapi"       //nolint
	_ "github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2019-05-01/resources/resourcesapi" //nolint
	_ "k8s.io/code-generator/cmd/client-gen"                                                        //nolint
	_ "k8s.io/code-generator/cmd/conversion-gen"                                                    //nolint
	_ "k8s.io/code-generator/cmd/deepcopy-gen"                                                      //nolint
	_ "k8s.io/code-generator/cmd/defaulter-gen"                                                     //nolint
	_ "k8s.io/code-generator/cmd/informer-gen"                                                      //nolint
	_ "k8s.io/code-generator/cmd/lister-gen"                                                        //nolint
	_ "k8s.io/code-generator/cmd/register-gen"                                                      //nolint
	_ "k8s.io/code-generator/cmd/set-gen"                                                           //nolint
	_ "sigs.k8s.io/controller-tools/cmd/controller-gen"                                             //nolint
	_ "sigs.k8s.io/testing_frameworks/integration"                                                  //nolint
)
