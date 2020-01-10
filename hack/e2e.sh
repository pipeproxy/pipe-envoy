#!/usr/bin/env bash
set -o errexit
set -o nounset
set -o pipefail

##
## Integration testing script for the control plane library against the Envoy binary.
## This is a wrapper around the test app `pkg/test/e2e` that spawns/kills Envoy.
##

# Management server type. Valid values are "ads", "xds", "rest"
XDS=${XDS:-ads}

# Number of RTDS layers.
if [ "$XDS" = "ads" ]; then
  RUNTIMES=2
else
  RUNTIMES=1
fi

(bin/e2e --xds=${XDS} --runtimes=${RUNTIMES} -debug "$@")&
SERVER_PID=$!

# Envoy start-up command
ENVOY=${ENVOY:-bin/envoy}
ENVOY_LOG="envoy.${XDS}.log"
echo Envoy log: ${ENVOY_LOG}

# Start envoy: important to keep drain time short
(${ENVOY} -c examples/${XDS}.yml -d> ${ENVOY_LOG})&
ENVOY_PID=$!

function cleanup() {
  kill ${ENVOY_PID}
}
trap cleanup EXIT

wait ${SERVER_PID}
