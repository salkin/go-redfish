MYID="$(shell id -u $(whoami)):$(shell id -g $(whoami))"

.PHONY: code-gen
code-gen:
	rm -rf client
	docker run --rm \
		-u ${MYID} \
		-v ${PWD}:/local openapitools/openapi-generator-cli generate \
			      -i  /local/spec/openapi.yaml \
						-g go \
						-v \
						--package-name "client" \
						--git-repo-id "go-redfish/client" \
						--git-user-id Nordix \
						-p enumClassPrefix=true \
						-o /local/client

	mkdir -p api
	go run api_generator.go | gofmt > api/service_interface.go


.PHONY: deps
deps:
	go get github.com/stretchr/testify/assert
	go get golang.org/x/oauth2
	go get golang.org/x/net/context
	go get github.com/antihax/optional
