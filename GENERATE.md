# Provider generation
The provider is generated following the steps from https://github.com/crossplane/terrajet/blob/main/docs/generating-a-provider.md 

Following steps were executed.

## Steps

1. Generate a GitHub repository for the Crossplane provider by hitting the
   "**Use this awssc**" button in [provider-jet-awssc] repository.
2. Clone the repository to your local and `cd` into the repository directory.
   Fetch the [upbound/build] submodule by running the following:

    ```bash
    make submodules
    ```

3. Replace `template` with your provider name.

    1. Export `ProviderName`:

        ```bash
        export ProviderNameLower=awssc
        export ProviderNameUpper=AWSSC
        ```

    2. Run the `./hack/prepare.sh` script from repo root to prepare the repo, e.g., to
       replace all occurrences of `awssc` with your provider name:

        ```bash
       ./hack/prepare.sh
        ```
4. To configure the Terraform provider to generate from, update the following
   variables in `Makefile`:

    ```makefile
    export TERRAFORM_PROVIDER_SOURCE := hashicorp/aws
    export TERRAFORM_PROVIDER_VERSION := 4.8.0
    export TERRAFORM_PROVIDER_DOWNLOAD_NAME := terraform-provider-aws
    export TERRAFORM_PROVIDER_DOWNLOAD_URL_PREFIX := https://releases.hashicorp.com/terraform-provider-aws/4.8.0
    ```

   You can find `TERRAFORM_PROVIDER_SOURCE` and `TERRAFORM_PROVIDER_VERSION` in
   [Terraform AWS provider](https://registry.terraform.io/providers/hashicorp/aws/4.8.0/docs) documentation by hitting the "**USE PROVIDER**"
   button.

5. Implement `ProviderConfig` logic in [internal/clients/awssc.go](internal/clients/awssc.go). 

   Copied and adapter from [provider-jet-aws repo](https://github.com/crossplane-contrib/provider-jet-aws/blob/v0.5.0/internal/clients/aws.go).  

6. Updated [config/provider.go](config/provider.go) to include only the resources needed:
   * aws_servicecatalog_provisioning_artifact
   * aws_servicecatalog_provisioned_product
  
   Then copied and adapted the logic from [provider-jet-aws repo](https://github.com/crossplane-contrib/provider-jet-aws/blob/v0.4.2/config/provider.go).

7. Add custom configuration for the two resources in [config/servicecatalog/config.go](config/servicecatalog/config.go) (empty at the moment),
then call it in [config/provider.go](config/provider.go).  

8. Generate
   Corrected the dependency `github.com/google/go-cmp v0.5.7` in [go.mod](go.mod).

   ```bash
   # another dependency might be needed
   go get github.com/crossplane/crossplane-runtime/pkg/reconciler/managed@v0.15.1-0.20220106140106-428b7c390375
   
   # generate
   make generate
   ```

## Setup CI
Create a github environment in settings and add the following secrets:
* AWS_USR
* AWS_PSW
* DOCKER_USR
* DOCKER_PSW
* DOCKER_REGISTRY

## Tag the version
```bash
git tag v0.1.0
git push origin --tags
```