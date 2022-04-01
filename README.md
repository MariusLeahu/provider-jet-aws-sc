# Terrajet AWS Service Catalog Provider

`provider-jet-awssc` is a [Crossplane](https://crossplane.io/) provider that is built
using [Terrajet](https://github.com/crossplane/terrajet) code generation tools and exposes XRM-compliant managed
resources for the AWSSC API.

It is based on [Terraform AWS provider 4.8.0](https://registry.terraform.io/providers/hashicorp/aws/4.8.0/docs) and
contains the following resources for AWS Service Catalog:

* [aws_servicecatalog_provisioning_artifact](https://registry.terraform.io/providers/hashicorp/aws/4.8.0/docs/resources/servicecatalog_provisioning_artifact)
* [aws_servicecatalog_provisioned_product](https://registry.terraform.io/providers/hashicorp/aws/4.8.0/docs/resources/servicecatalog_provisioned_product)

## Getting Started

Install the provider by using the following command after changing the image tag to
the [latest release](https://github.com/mleahu/provider-jet-awssc/releases):

```
kubectl crossplane install provider mleahu/provider-jet-awssc:v0.1.0
```

You can see the CRDs [here](package/crds).

## Usage

See [examples](examples) folder.

## Developing

How it was generated: see [here](GENERATE.md)

## Licensing

provider-jet-awssc is under the Apache 2.0 license.
