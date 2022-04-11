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

package servicecatalog

import (
	"fmt"

	"github.com/crossplane-contrib/provider-jet-awssc/config/common"
	"github.com/crossplane/terrajet/pkg/config"
)

// Configure adds configurations for servicecatalog group.
func Configure(p *config.Provider, ot *config.OperationTimeouts) {

	p.AddResourceConfigurator("aws_servicecatalog_provisioning_artifact", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		// r.ExternalName = config.IdentifierFromProvider
		r.UseAsync = true
		r.OperationTimeouts = *ot
	})

	p.AddResourceConfigurator("aws_servicecatalog_provisioned_product", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		// r.ExternalName = config.IdentifierFromProvider
		r.UseAsync = true
		r.OperationTimeouts = *ot
		/*		r.References["subnet_ids"] = config.Reference{
					//Type:              "github.com/crossplane/provider-aws/apis/ec2/v1beta1.Subnet",
					//RefFieldName:      "SubnetIdRefs",
					//SelectorFieldName: "SubnetIdSelector",
					Type:              "github.com/crossplane-contrib/provider-jet-aws/apis/ec2/v1alpha2.Subnet",
					RefFieldName:      "SubnetIdRefs",
					SelectorFieldName: "SubnetIdSelector",
				}
		*/

		// save the outputs as a map
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]interface{}) (map[string][]byte, error) {
			conn := map[string][]byte{}
			fmt.Println("attr in AdditionalConnectionDetailsFn", attr)
			/*			if o, ok := attr["outputs"].()
						if a, ok := attr["connection_name"].(string); ok {
							conn[CloudSQLSecretConnectionName] = []byte(a)
						}
						if a, ok := attr["private_ip_address"].(string); ok {
							conn[PrivateIPKey] = []byte(a)
						}
						if a, ok := attr["public_ip_address"].(string); ok {
							conn[PublicIPKey] = []byte(a)
						}
						if a, ok := attr["root_password"].(string); ok {
							conn[xpv1.ResourceCredentialsSecretPasswordKey] = []byte(a)
						}
						// map
						if certSlice, ok := attr["server_ca_cert"].([]interface{}); ok {
							if certattrs, ok := certSlice[0].(map[string]interface{}); ok {
								if a, ok := certattrs["cert"].(string); ok {
									conn[CloudSQLSecretServerCACertificateCertKey] = []byte(a)
								}
								if a, ok := certattrs["common_name"].(string); ok {
									conn[CloudSQLSecretServerCACertificateCommonNameKey] = []byte(a)
								}
								if a, ok := certattrs["create_time"].(string); ok {
									conn[CloudSQLSecretServerCACertificateCreateTimeKey] = []byte(a)
								}
								if a, ok := certattrs["expiration_time"].(string); ok {
									conn[CloudSQLSecretServerCACertificateExpirationTimeKey] = []byte(a)
								}
								if a, ok := certattrs["sha1_fingerprint"].(string); ok {
									conn[CloudSQLSecretServerCACertificateSha1FingerprintKey] = []byte(a)
								}
							}
						}
			*/return conn, nil
		}
	})

}
