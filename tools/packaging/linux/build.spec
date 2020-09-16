Summary:    AWS Observability Collector
Name:       %{RPM_NAME}
Version:    %{VERSION}
Release:    1
License:    Amazon Software License. Copyright 2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.
Group:      Applications/AWS
Source0:    %{RPM_NAME}-%{VERSION}.tar.gz

%description
This package provides daemon of AWS Observability Collector

%prep
%setup -q

%install
rm -rf ${RPM_BUILD_ROOT}
mkdir -p ${RPM_BUILD_ROOT}
cp -fa * ${RPM_BUILD_ROOT}

%files
%dir /opt/aws
%dir /opt/aws/aws-observability-collector
%dir /opt/aws/aws-observability-collector/bin
%dir /opt/aws/aws-observability-collector/doc
%dir /opt/aws/aws-observability-collector/etc
%dir %attr(-, aoc, aoc) /opt/aws/aws-observability-collector/logs
%dir %attr(-, aoc, aoc) /opt/aws/aws-observability-collector/var
%dir %attr(-, aoc, aoc) /opt/aws/aws-observability-collector/etc

/opt/aws/aws-observability-collector/bin/aws-observability-collector
/opt/aws/aws-observability-collector/bin/aws-observability-collector-ctl
/opt/aws/aws-observability-collector/bin/VERSION
/opt/aws/aws-observability-collector/LICENSE
/opt/aws/aws-observability-collector/RELEASE_NOTE
/opt/aws/aws-observability-collector/var/.config.yaml
/opt/aws/aws-observability-collector/etc/.env

/etc/init/aws-observability-collector.conf
/etc/systemd/system/aws-observability-collector.service
/usr/bin/aws-observability-collector-ctl
/etc/amazon/aws-observability-collector
/var/log/amazon/aws-observability-collector
/var/run/amazon/aws-observability-collector

%pre
# Stop the agent before upgrades.
if [ $1 -ge 2 ]; then
    if [ -x /opt/aws/aws-observability-collector/bin/aws-observability-collector-ctl ]; then
        /opt/aws/aws-observability-collector/bin/aws-observability-collector-ctl -a stop
    fi
fi

# create group
if ! grep "^aoc:" /etc/group >/dev/null 2>&1; then
    groupadd -r aoc >/dev/null 2>&1
    echo "create group aoc, result: $?"
fi

# create user
if ! id aoc >/dev/null 2>&1; then
    useradd -r -M aoc -d /home/aoc -g aoc -c "AWS Observability Collector" -s $(test -x /sbin/nologin && echo /sbin/nologin || (test -x /usr/sbin/nologin && echo /usr/sbin/nologin || (test -x /bin/false && echo /bin/false || echo /bin/sh))) >/dev/null 2>&1
    echo "create user aoc, result: $?"
fi


%clean
rm -rf ${RPM_BUILD_ROOT}
