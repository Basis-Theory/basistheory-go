# Contributing

## Prerequisites

First ensure you have Go installed:
```shell
brew install go
```

## Updating the SDK

Download the latest API schema to `./swagger.json` from `https://api.basistheory.com/swagger/v1/swagger.json`,
which can be done manually or with the command:

```shell
make update-api-spec
```

Once the latest api spec has been downloaded locally, you must regenerate the SDK using the open api generator,
which can be done with the command:

```shell
make generate-sdk
```

## Running tests

Copy the `.env.example` file to `.env.local` and enter valid Basis Theory
API keys for the dev environment into this file:

- `BASISTHEORY_MGMT_API_KEY`: API key issued to a `management` application with all permissions
- `BASISTHEORY_SRVR_API_KEY`: API key issued to a `private` application with all permissions

To run tests using this configuration, run:

```shell
make verify
```
