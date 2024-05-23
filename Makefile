HOSTNAME=github.com
NAMESPACE=relaypro-open
NAME=dog_api_golang
BINARY=${NAME}
VERSION=v1.0.4
OS_ARCH=linux_amd64

default: install

build:
	go mod vendor
	go build -o ./bin/${BINARY}

debug-build:
	go build -gcflags="all=-N -l" -o ./bin/${BINARY}

debug:
	dlv exec --accept-multiclient --continue --headless ./bin/dog_api_golang -- -debug

github_release:
	git tag ${VERSION}
	git push --tags --force

delete_release:
	git tag -d ${VERSION}
	git push --delete origin ${VERSION}

test:
	go test --tags=integration github.com/relaypro-open/dog_api_golang/api -count=1
