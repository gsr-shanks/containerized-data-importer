#!/usr/bin/env bash

#Copyright 2019 The CDI Authors.
#
#Licensed under the Apache License, Version 2.0 (the "License");
#you may not use this file except in compliance with the License.
#You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
#Unless required by applicable law or agreed to in writing, software
#distributed under the License is distributed on an "AS IS" BASIS,
#WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#See the License for the specific language governing permissions and
#limitations under the License.

set -e
script_dir="$(readlink -f $(dirname $0))"
source "${script_dir}"/common.sh
source "${script_dir}"/config.sh

BUILDER_SPEC="${BUILD_DIR}/docker/builder"

# When building and pushing a new image we do not provide the sha hash
# because docker assigns that for us.
BUILD_TAG=$(expr match $BUILDER_TAG '\(.*\)@sha256')

# Build the encapsulated compile and test container
(cd ${BUILDER_SPEC} && docker build --tag ${BUILD_TAG}:latest .)

DIGEST=$(docker images --digests | grep $BUILD_TAG | grep latest | awk '{ print $3 }')
echo "Image: ${BUILD_TAG}:latest"
echo "Digest: ${DIGEST}"

docker push ${BUILD_TAG}:latest
