MAKEFLAGS += --silent

vault-checkout:
	./scripts/vault-checkout.sh

stop-docker:
	./scripts/stop-docker.sh

start-docker:
	$(MAKE) vault-checkout
	./scripts/start-docker.sh
	$(MAKE) service-up

service-up:
	./scripts/service-up.sh

acceptance-test:
	./scripts/acceptance-test.sh

verify:
	./scripts/verify.sh
