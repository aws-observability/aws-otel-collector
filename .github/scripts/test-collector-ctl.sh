#!/bin/bash -x

set -e

# This script is supposed to run during the CI to validate that the control
# script is working as expected
# Important: this script requires root permission.

COLLECTOR_DEB=$1
COLLECTOR_CONFIG=$2

ADOT_CTL=/opt/aws/aws-otel-collector/bin/aws-otel-collector-ctl
ENV_FILE="/opt/aws/aws-otel-collector/etc/.env"
CONFIG_FILE="/opt/aws/aws-otel-collector/etc/config.yaml"

setup() {
    dpkg -i "$COLLECTOR_DEB"
}

test_collector_ctl_does_not_overwrite_env() {
    $ADOT_CTL -a start
    if [ ! -e $CONFIG_FILE ]; then
        echo "Config file does not exist"
        exit 1
    fi

    echo "EXTRA_ENV=\"some value\"" | tee -a $ENV_FILE > /dev/null

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

test_collector_ctl_with_sed_special_chars() {
    # ampersand is a special character in sed
    cfg="https://example.com?test=1&test=2"
    $ADOT_CTL -a start -c "$cfg"

    expected_string="^config=\"--config '${cfg}'\"$"

    if ! grep -q "${expected_string}" "$ENV_FILE";  then
        echo "Configuration is incorrect"
        exit 1
    fi

    cfg="./config.yaml"
    $ADOT_CTL -a start -c "$cfg"

    expected_string="^config=\"--config /opt/aws/aws-otel-collector/etc/config.yaml\"$"
    if ! grep -q "${expected_string}" "$ENV_FILE";  then
        echo "Configuration is incorrect"
        exit 1
    fi

    echo "${FUNCNAME[0]} ... OK"
}


test_collector_ctl_with_samecfg() {
    #ensure default cfg is cleared 
    rm /opt/aws/aws-otel-collector/etc/config.yaml
    
    $ADOT_CTL -a start -c "file:/opt/aws/aws-otel-collector/etc/config.yaml"

    echo "${FUNCNAME[0]} ... OK"
}


test_collector_ctl_with_samecfg_restart() {
    #populate default conf by starting without -c
    $ADOT_CTL -a start 

    $ADOT_CTL -a start -c "file:/opt/aws/aws-otel-collector/etc/config.yaml"

    echo "${FUNCNAME[0]} ... OK"
}

setup

## Tests
test_collector_ctl_does_not_overwrite_env
test_collector_ctl_with_sed_special_chars
test_collector_ctl_with_samecfg
test_collector_ctl_with_samecfg_restart
