## git describe --tags: tag 정보 출력
## git log -1 --format='$H': 최신 1개의 커밋 정보의 해시값 출력
VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='$H')
    
## library 설정 
ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=HelloChain \
	-X github.com/cosmos/cosmos-sdk/version.ServerName=hcd \
        -X github.com/cosmos/cosmos-sdk/version.ClientName=hccli \
        -X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
        -X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT)
    
BUILD_TAGS := -ldflags '$(ldflags)'
    
    
all: install
    
install: go.sum
	go install $(BUILD_FLAGS) ./cmd/hcd
	go install $(BUILD_FLAGS) ./cmd/hccli
    
go.sum: go.mod
		@echo "-> Ensuredependencies hab not been modufued"
		go mod verify
