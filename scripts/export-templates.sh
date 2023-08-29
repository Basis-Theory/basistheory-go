#!/bin/bash

# this script will export the default templates to <repo_root>/templates/default, which can then be customized
# and committed. Do not commit all the default templates, as these can change when upgrading open api versions.
docker pull openapitools/openapi-generator-cli:v6.6.0
docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli:v6.6.0 author template \
  -g go \
  -o /local/templates/default \

cd $(dirname $0)
cd ../
