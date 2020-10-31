#!/usr/bin/env bash
set -o errexit
set -o nounset
set -o pipefail

##
## Integration testing script for the control plane library against the Envoy binary.
## This is a wrapper around the test app `pkg/test/main` that spawns/kills Envoy.
##

# Management server type. Valid values are "ads", "xds", "rest"
XDS=${XDS:-ads}

#Represents SUFFIX api version
SUFFIX=${SUFFIX:-}

# Number of RTDS layers.
if [ "$XDS" = "ads" ]; then
  RUNTIMES=2
else
  RUNTIMES=1
fi

(bin/test --xds=${XDS} --runtimes=${RUNTIMES} -debug "$@")&
SERVER_PID=$!

# Envoy start-up command
PIPE=${PIPE:-./bin/pipe-xds}
PIPE_LOG="pipe.${XDS}${SUFFIX}.log"
echo Pipe log: ${PIPE_LOG}

# Start pipe: important to keep drain time short
(${PIPE} -u 127.0.0.1:18000 -n "test-id" 2> ${PIPE_LOG})&
PIPE_PID=$!

function cleanup() {
  kill -QUIT ${PIPE_PID}
}
trap cleanup EXIT

wait ${SERVER_PID}
