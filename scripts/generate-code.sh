#!/usr/bin/env bash

set -euo pipefail

# Remove generated files
for F in $(cat .openapi-generator/files) ; do
    rm -f "${F}"
done

# Generate library
docker run --rm --env JAVA_OPTS=-DmaxYamlCodePoints=9999999 -v "${PWD}:/local" openapitools/openapi-generator-cli:v7.8.0 generate --config /local/.openapi-generator/config.yaml --input-spec /local/api/openapi.yaml --output /local --inline-schema-options RESOLVE_INLINE_ENUMS=true --http-user-agent go-netbox/4.1.3
