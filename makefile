MAKEFLAGS += --silent

stop-docker:
	./scripts/stop-docker.sh

start-docker:
	./scripts/start-docker.sh

service-up:
	./scripts/service-up.sh

acceptance-test:
	./scripts/acceptance-test.sh

verify:
	./scripts/verify.sh
