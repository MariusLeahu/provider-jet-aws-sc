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

   Copied and adapter from [provider-jet-aws repo](https://github.com/crossplane-contrib/provider-jet-aws/blob/v0.4.2/internal/clients/aws.go).  
