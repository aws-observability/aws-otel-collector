#!/bin/bash

# This script is supposed to run during the CI to validate that the control
# script is working as expected

COLLECTOR_DEB=$1
COLLECTOR_CONFIG=$2

ADOT_CTL=/opt/aws/aws-otel-collector/bin/aws-otel-collector-ctl
ENV_FILE="/opt/aws/aws-otel-collector/etc/.env"

setup() {
    sudo apt install $COLLECTOR_DEB
}

test_collector_ctl_does_not_overwrite_env() {
    # Prepare
    sudo $ADOT_CTL -a start -c $COLLECTOR_CONFIG

    # Act
    echo "EXTRA_ENV=\"some value\"" | sudo tee -a $ENV_FILE > /dev/null

    # Assert
    # We don't need the collector to succeed, just that the control script is execised
    sudo $ADOT_CTL -a start -c http://example.com

    sudo cat $ENV_FILE | grep -q EXTRA_ENV
    if [ $? -eq 0 ]; then
        echo "Env file was overwritten"
        exit 1
    fi

    sudo $ADOT_CTL -a start -c $COLLECTOR_CONFIG
    sudo cat $ENV_FILE | grep -q EXTRA_ENV
    if [ $? -ne 0 ]; then
        ehco "Env file was overwritten"
        exit 1
    fi
}


setup

## Tests
test_collector_ctl_does_not_overwrite_env
