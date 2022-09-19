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

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"
	"time"

	"github.com/crossplane-contrib/provider-jet-awssc/config/common"

	"github.com/crossplane-contrib/provider-jet-awssc/config/lambda"
	"github.com/crossplane-contrib/provider-jet-awssc/config/servicecatalog"

	tjconfig "github.com/crossplane/terrajet/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	resourcePrefix = "awssc"
	modulePath     = "github.com/crossplane-contrib/provider-jet-awssc"
)

// IncludedResources lists all resource patterns included in small set release.
var IncludedResources = []string{

	// Service Catalog
	"aws_servicecatalog_provisioning_artifact$",
	"aws_servicecatalog_provisioned_product$",
	"aws_lambda_function$",
	"aws_lambda_event_source_mapping$",
}

// Options for this Provider.
type Options struct {

	// Terraform ReadTimeout
	PollInterval time.Duration
}

//go:embed schema.json
var providerSchema string

// GetProvider returns provider configuration
func GetProvider() *tjconfig.Provider {
	var ot = tjconfig.OperationTimeouts{
		Read:   common.TerraformReadTimeout,
		Create: common.TerraformCreateTimeout,
		Update: common.TerraformUpdateTimeout,
		Delete: common.TerraformDeleteTimeout,
	}
	return GetProviderWithTimeouts(&ot)
}

// GetProviderWithTimeouts returns provider configuration
func GetProviderWithTimeouts(ot *tjconfig.OperationTimeouts) *tjconfig.Provider {
	defaultResourceFn := func(name string, terraformResource *schema.Resource, opts ...tjconfig.ResourceOption) *tjconfig.Resource {
		r := tjconfig.DefaultResource(name, terraformResource,
			GroupKindOverrides(),
			KindOverrides(),
			RegionAddition(),
			TagsAllRemoval(),
			IdentifierAssignedByAWS(),
			NamePrefixRemoval(),
			KnownReferencers(),
		)
		// Add any provider-specific defaulting here. For example:
		//   r.ExternalName = tjconfig.IdentifierFromProvider
		return r
	}

	pc := tjconfig.NewProviderWithSchema([]byte(providerSchema), resourcePrefix, modulePath,
		tjconfig.WithShortName("awsjet"),
		tjconfig.WithRootGroup("aws.jet.crossplane.io"),
		tjconfig.WithIncludeList(IncludedResources),
		tjconfig.WithDefaultResourceFn(defaultResourceFn))

	for _, configure := range []func(provider *tjconfig.Provider, ot *tjconfig.OperationTimeouts){
		servicecatalog.Configure,
		lambda.Configure,
	} {
		configure(pc, ot)
	}

	pc.ConfigureResources()
	return pc
}
