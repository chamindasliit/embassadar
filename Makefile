# file: Makefile

# Copyright 2018 Datawire. All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License

# Welcome to the Ambassador Makefile...

# We'll set REGISTRY_ERR in builder.mk
DEV_REGISTRY ?= $(REGISTRY_ERR)

# IS_PRIVATE: empty=false, nonempty=true
# Default is true if any of the git remotes have the string "private" in any of their URLs.
_git_remote_urls := $(shell git remote | xargs -n1 git remote get-url --all)
IS_PRIVATE ?= $(findstring private,$(_git_remote_urls))

RELEASE_DOCKER_REPO ?= quay.io/datawire/ambassador$(if $(IS_PRIVATE),-private)
BASE_DOCKER_REPO    ?= quay.io/datawire/ambassador-base$(if $(IS_PRIVATE),-private)
DEV_DOCKER_REPO     ?= $(DEV_REGISTRY)/dev

DOCKER_OPTS ?=

docker.tag.dev        = $(DEV_DOCKER_REPO):$(notdir $*)-$(shell tr : - < $<)
# By default, don't allow .release, .release-rc, .release-ea, or .base tags...
docker.tag.release    = $(error The 'release' tag is only valid for the 'ambassador' image)
docker.tag.release-rc = $(error The 'release-rc' tag is only valid for the 'ambassador' image)
docker.tag.release-ea = $(error The 'release-ea' tag is only valid for the 'ambassador' image)
docker.tag.base       = $(error The 'base' tag is only valid for the 'base-envoy' image)
# ... except for on specific images
# ... haha, jk, it's broken in CI?
#ambassador.docker.tag.release:
docker.tag.release    = $(RELEASE_DOCKER_REPO):$(RELEASE_VERSION)
#ambassador.docker.tag.release:
docker.tag.release-rc = $(RELEASE_DOCKER_REPO):$(RELEASE_VERSION) $(RELEASE_REPO):$(BUILD_VERSION)-latest-rc
#ambassador.docker.tag.release:
docker.tag.release-ea = $(RELEASE_DOCKER_REPO):$(RELEASE_VERSION)
#envoy-base.docker.tag.base:
docker.tag.base       = $(BASE_IMAGE.envoy)

# We'll set REGISTRY_ERR in builder.mk
docker.tag.dev = $(if $(DEV_REGISTRY),$(DEV_REGISTRY)/$*:$(patsubst sha256:%,%,$(shell cat $<)),$(REGISTRY_ERR))

# All Docker images that we know how to build
images.all =
# The subset of $(images.all) that we will deploy to the
# DEV_KUBECONFIG cluster.
images.cluster =

images.all += $(patsubst docker/%/Dockerfile,%,$(wildcard docker/*/Dockerfile)) test-auth-tls
images.cluster += $(filter test-%,$(images.all))
images.base += $(filter base-%,$(images.all))
# NB: images.all and images.cluster are used eagerly by
# builder/builder.mk, so we must set them before we include
# builder.mk.

OSS_HOME := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))
include $(OSS_HOME)/build-aux/prelude.mk
include $(OSS_HOME)/build-aux/var.mk
include $(OSS_HOME)/build-aux/docker.mk
include $(OSS_HOME)/builder/builder.mk
include $(OSS_HOME)/cxx/envoy.mk
include $(OSS_HOME)/build-aux-local/kat.mk
include $(OSS_HOME)/build-aux-local/docs.mk
include $(OSS_HOME)/build-aux-local/release.mk
include $(OSS_HOME)/build-aux-local/version.mk
.DEFAULT_GOAL = help

$(call module,ambassador,$(OSS_HOME))

sync: python/ambassador/VERSION.py

clean: _makefile_clean
clobber: _makefile_clobber
_makefile_clean:
	rm -f python/ambassador/VERSION.py
_makefile_clobber:
	rm -rf bin_*/
.PHONY: _makefile_clean _makefile_clobber

generate: ## Update generated sources that get committed to git
generate: pkg/api/kat/echo.pb.go
generate-clean: ## Delete generated sources that get committed to git (implies `make clobber`)
generate-clean: clobber
	rm -rf pkg/api
.PHONY: generate generate-clean
# NB: cxx/envoy.mk hooks in to 'generate' & 'generate-clean'

test-%.docker.stamp: docker/test-%/Dockerfile FORCE
	docker build --quiet --iidfile=$@ $(<D)
test-auth-tls.docker.stamp: docker/test-auth/Dockerfile FORCE
	docker build --quiet --build-arg TLS=--tls --iidfile=$@ $(<D)

# travis-script.sh needs to be able to know these variables
export-vars:
	@echo "export BASE_DOCKER_REPO='$(BASE_DOCKER_REPO)'"
	@echo "export RELEASE_DOCKER_REPO='$(RELEASE_DOCKER_REPO)'"
.PHONY: export-vars

# Configure GNU Make itself
SHELL = bash
.SECONDARY:
.DELETE_ON_ERROR:
