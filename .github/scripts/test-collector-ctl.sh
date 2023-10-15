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

test_collector_ctl_copy_cfg(){
    touch /tmpconf.yaml
    echo "testcfg" >> "/tmpconf.yaml"

    $ADOT_CTL -a start -c "/tmpconf.yaml"
    # we expect config file to be copied. Collector does not need to start correctly
    if cmp --silent -- "$CONFIG_FILE" "/tmpconf.yaml"; then
        echo "config files contents are identical"
    else
        echo "config files differ"
        echo "tmpconf: " && cat /tmpconf
        echo "config file: " && cat $CONFIG_FILE
        exit 1 
    fi

    $ADOT_CTL -a stop

    echo "${FUNCNAME[0]} ... OK"
}

test_start_with_default_cfg(){
    #clear default cfg
    rm /opt/aws/aws-otel-collector/etc/config.yaml

    $ADOT_CTL -a start
    #systemctl is-active will show "activating" if it fails to start
    # need to explicitly check for "active" output
    STATUS="$(systemctl is-active "aws-otel-collector.service")"
    if [ "${STATUS}" = "active" ] ; then
        echo "aws-otel-collecotr.service is started"
    else
        echo "aws-otel-collector.service is not running"
        exit 1
    fi

    echo "${FUNCNAME[0]} ... OK"
}

setup

## Tests
test_collector_ctl_does_not_overwrite_env
test_collector_ctl_with_sed_special_chars
test_collector_ctl_with_samecfg
test_collector_ctl_with_samecfg_restart
test_collector_ctl_copy_cfg
test_start_with_default_cfg
