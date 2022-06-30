MAKEFLAGS += --silent

verify:
	./scripts/verify.sh

generate-sdk:
	./scripts/generate-sdk.sh