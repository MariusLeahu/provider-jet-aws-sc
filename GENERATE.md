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