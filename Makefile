_OS := $(shell uname -s)
_IPS := 1

_UID := $(shell id -u)
_GID := $(shell id -g)

export _UID
export _GID

do = docker exec -t $(1)
doco = docker-compose  -f docker/docker-compose.dev.yml $(1)

start:
	@echo -e '\e[1;31mGoing up...\e[0m'
	@$(call doco, up -d)
	@echo -e '\e[1;31mDone\e[0m'

status:
	@$(call doco, ps)

restart:
	@echo -e '\e[1;31mRestarting...\e[0m'
	@$(call doco, restart)
	@echo -e '\e[1;31mDone\e[0m'

build:
	@echo -e '\e[1;31mBuilding...\e[0m'
	@$(call doco, build --pull  --parallel)
	@echo -e '\e[1;31mDone\e[0m'

down:
	@echo -e '\e[1;31mStopping...\e[0m'
	@$(call doco, stop)
	@echo -e '\e[1;31mDone\e[0m'

stop: down

destroy:
	@echo -e '\e[1;31mDestroying...\e[0m'
	@$(call doco, down)
	@echo -e '\e[1;31mDone\e[0m'

### Client Commands

download:
	@$(call do, acronis-client go run /go/src/main.go)

shell:
	@$(call do, -ti acronis-client bash)

test:
	@$(call do, acronis-client bash -c "cd tests && go test ./...")