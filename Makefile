PACKAGES=$(shell go list ./...)
OUTPUT?=build/tendermint

BUILD_TAGS?=tendermint

# If building a release, please checkout the version tag to get the correct version setting
ifneq ($(shell git symbolic-ref -q --short HEAD),)
VERSION := unreleased-$(shell git symbolic-ref -q --short HEAD)-$(shell git rev-parse HEAD)
else
VERSION := $(shell git describe)
endif

LD_FLAGS = -X github.com/YunweiMao/tendermint/version.TMCoreSemVer=$(VERSION)
BUILD_FLAGS = -mod=readonly -ldflags "$(LD_FLAGS)"
HTTPS_GIT := https://github.com/YunweiMao/tendermint.git
CGO_ENABLED ?= 0

#if you want include TENDERMINT_BUILD_OPTIONS, then you need to run it as following:
#TENDERMINT_BUILD_OPTIONS="-tags cleveldb" make install

# handle nostrip
ifeq (,$(findstring nostrip,$(TENDERMINT_BUILD_OPTIONS)))
  BUILD_FLAGS += -trimpath
  LD_FLAGS += -s -w
endif

# handle race
ifeq (race,$(findstring race,$(TENDERMINT_BUILD_OPTIONS)))
  CGO_ENABLED=1
  BUILD_FLAGS += -race
endif

# handle cleveldb
ifeq (cleveldb,$(findstring cleveldb,$(TENDERMINT_BUILD_OPTIONS)))
  CGO_ENABLED=1
  BUILD_TAGS += cleveldb
endif

# handle badgerdb
ifeq (badgerdb,$(findstring badgerdb,$(TENDERMINT_BUILD_OPTIONS)))
  BUILD_TAGS += badgerdb
endif

# handle rocksdb
ifeq (rocksdb,$(findstring rocksdb,$(TENDERMINT_BUILD_OPTIONS)))
  CGO_ENABLED=1
  BUILD_TAGS += rocksdb
endif

# handle boltdb
ifeq (boltdb,$(findstring boltdb,$(TENDERMINT_BUILD_OPTIONS)))
  BUILD_TAGS += boltdb
endif

#LDFLAGS="-L/path/to/library -lmylib" make install
# allow users to pass additional flags via the conventional LDFLAGS variable
LD_FLAGS += $(LDFLAGS)

all: check build test install
.PHONY: all

###############################################################################
###                                Build Tendermint                        ###
###############################################################################

#The command line interface will print the following:
#CGO_ENABLED=0 go install -mod=readonly -ldflags "-X github.com/YunweiMao/tendermint/version.TMCoreSemVer=unreleased-main-be396f9de47ca223dd935598171d99c05b6c78b3 -s -w " -trimpath -tags tendermint ./cmd/tendermint

build:
	CGO_ENABLED=$(CGO_ENABLED) go build $(BUILD_FLAGS) -tags '$(BUILD_TAGS)' -o $(OUTPUT) ./cmd/tendermint/
.PHONY: build

install:
	CGO_ENABLED=$(CGO_ENABLED) go install $(BUILD_FLAGS) -tags $(BUILD_TAGS) ./cmd/tendermint
.PHONY: install