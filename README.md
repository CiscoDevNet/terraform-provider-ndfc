# Terraform Provider NDFC

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) >= 1.0
- [Go](https://golang.org/doc/install) >= 1.19


## Building The Provider

1. Clone the repository
2. Enter the repository directory
3. Build the provider using the Go Install command:

```shell
go install
```
The provider should be available in `$GOPATH/bin`

## Using the provider

The provider is not yet available in hashicorp terraform registry. So `terrform init` will not work

Refer https://developer.hashicorp.com/terraform/cli/config/config-file   
Section: Development Overrides for Provider Developers    
To add dev overrides to use a the plugin under development    


## Developing the Provider

See generator/generate.md


## Acceptance Tests

```shell
./run_accept_tests.sh // Run everything
./run_accept_tests.sh TestAcc<Pattern of tests> // Run Specific tests 
```


## Steps to run unit test with mockoon server

1. Install mockoon-cli server
```
npm install -g @mockoon/cli
```
2. Copy your mockoon environment data json file to `/terraform-provide-ndfc/mockoon_data.json`
3. Create unit test functions with prefix `TestUT_` and use `ut_client` as NDFC client variable.
4. Run `./run_unit_test.sh` to run the unit test cases.

## Provider plugin Documentation
[Provider](docs/index.md)   
[Resources](docs/resources)    
[Datasources](docs/data-sources)     
