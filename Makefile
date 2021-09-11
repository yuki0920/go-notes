## OpenAPI Gen

generate-api:
	docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate \
	-i local/openapi/openapi.yaml \
	-g go-echo-server \
	-o local/api
