/*
Copyright 2021 The Crossplane Authors.

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

package lambda

import (
	"github.com/crossplane/terrajet/pkg/config"
)

const (
	// VersionV1Alpha1 v1alpha1 means resource a resource wasn't modified
	VersionV1Alpha1 = "v1alpha1"
)

// Configure adds configurations for lambda group.
func Configure(p *config.Provider, ot *config.OperationTimeouts) {

	p.AddResourceConfigurator("aws_lambda_function", func(r *config.Resource) {
		r.Version = VersionV1Alpha1
		r.UseAsync = true
		r.OperationTimeouts = *ot
	})

	p.AddResourceConfigurator("aws_lambda_event_source_mapping", func(r *config.Resource) {
		r.Version = VersionV1Alpha1
		r.UseAsync = true
		r.OperationTimeouts = *ot
	})

}
