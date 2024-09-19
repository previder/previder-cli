# previder-cli
Previder CLI is the command line client for the Previder Portal

**Release:**

[![Release Version](https://img.shields.io/github/v/release/previder/previder-cli?label=previder-cli)](https://github.com/previder/previder-cli/releases/latest)

**Last build:**

![Last build](https://github.com/previder/previder-cli/actions/workflows/go.yml/badge.svg)

**Last release:**

![Last publish](https://github.com/previder/previder-cli/actions/workflows/goreleaser.yml/badge.svg)

# Getting started
Previder-cli is a stand-alone binary to use with the Previder Portal

To see all usages, run
```shell
./previder-cli --help
```

## Token
Use the token directly from the command-line or define the PREVIDER_TOKEN environment variable.

## Usage example
```shell
./previder-cli -t <insert-token> virtualserver list
```
Will print all Virtual servers in the tenant belonging to the token

```shell
export PREVIDER_TOKEN="insert-token"
./previder-cli version
```
Will print the current version of the client

## Output
The default output format is `json`. Lists of environments, tokens and secrets can also be pretty-printed with the `-o pretty` parameter.
