#!/bin/bash

docker pull openapitools/openapi-generator-cli:v6.6.0
docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli:v6.6.0 generate \
  -i /local/swagger.json \
  -g go \
  -o /local \
  -c /local/openapi-config.yml \
  --remove-operation-id-prefix \
  apiTests=false \
  modelTests=false

cd $(dirname $0)
cd ../

go fmt