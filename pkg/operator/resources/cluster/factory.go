/*
Copyright 2018 The CDI Authors.

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

package cluster

import (
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
)

// FactoryArgs contains the required parameters to generate all cluster-scoped resources
type FactoryArgs struct {
	Namespace string
}

// Resource defines the interface of all cluster-wide components
type Resource runtime.Object

type factoryFunc func(*FactoryArgs) []Resource

var factoryFunctions = map[string]factoryFunc{
	"crd":        createCRDResources,
	"controller": createControllerResources,
	"apiserver":  createAPIServerResources,
}

// CreateAllResources creates all cluster-wide resources
func CreateAllResources(args *FactoryArgs) ([]Resource, error) {
	resources := []Resource{}
	for group := range factoryFunctions {
		rs, err := CreateResourceGroup(group, args)
		if err != nil {
			return nil, err
		}
		resources = append(resources, rs...)
	}
	return resources, nil
}

// CreateResourceGroup creates all cluster resources fr a specific group/component
func CreateResourceGroup(group string, args *FactoryArgs) ([]Resource, error) {
	f, ok := factoryFunctions[group]
	if !ok {
		return nil, fmt.Errorf("Group %s does not exist", group)
	}
	return f(args), nil
}