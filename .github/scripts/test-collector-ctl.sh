#!/bin/bash

# This script is supposed to run during the CI to validate that the control
# script is working as expected
# Important: this script requires root permission.

COLLECTOR_DEB=$1
COLLECTOR_CONFIG=$2

ADOT_CTL=/opt/aws/aws-otel-collector/bin/aws-otel-collector-ctl
ENV_FILE="/opt/aws/aws-otel-collector/etc/.env"

setup() {
    apt-get install "$COLLECTOR_DEB"
}

test_collector_ctl_does_not_overwrite_env() {
    # Prepare
    $ADOT_CTL -a start -c "$COLLECTOR_CONFIG"

    # Act
    echo "EXTRA_ENV=\"some value\"" | tee -a $ENV_FILE > /dev/null

    # Assert
    # We don't need the collector to succeed, just that the control script is execised
    $ADOT_CTL -a start -c "http://example.com"

    if ! grep -q EXTRA_ENV "$ENV_FILE"; then
        echo "Env file was overwritten"
        exit 1
    fi

    $ADOT_CTL -a start -c "$COLLECTOR_CONFIG"
    if ! grep -q EXTRA_ENV "$ENV_FILE"; then
        ehco "Env file was overwritten"
        exit 1
    fi
    echo "${FUNCNAME[0]} ... OK"
}


setup

## Tests
test_collector_ctl_does_not_overwrite_env
