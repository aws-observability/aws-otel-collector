Summary:    ADOT Collector
Name:       %{RPM_NAME}
Version:    %{VERSION}
Release:    1
License:    Apache 2.0
Group:      Applications/AWS
Source0:    %{RPM_NAME}-%{VERSION}.tar.gz

%description
This package provides daemon of ADOT Collector

%prep
%setup -q

%install
rm -rf ${RPM_BUILD_ROOT}
mkdir -p ${RPM_BUILD_ROOT}
cp -fa * ${RPM_BUILD_ROOT}

%files
%dir /opt/aws/aws-otel-collector
%dir /opt/aws/aws-otel-collector/bin
%dir /opt/aws/aws-otel-collector/doc
%dir /opt/aws/aws-otel-collector/etc
%dir %attr(750, aoc, aoc) /opt/aws/aws-otel-collector/logs
%dir %attr(750, aoc, aoc) /opt/aws/aws-otel-collector/var
%dir %attr(750, aoc, aoc) /opt/aws/aws-otel-collector/etc

/opt/aws/aws-otel-collector/bin/aws-otel-collector
/opt/aws/aws-otel-collector/bin/aws-otel-collector-ctl
/opt/aws/aws-otel-collector/bin/VERSION
/opt/aws/aws-otel-collector/LICENSE
/opt/aws/aws-otel-collector/RELEASE_NOTE
%attr(640, aoc, aoc) /opt/aws/aws-otel-collector/var/.config.yaml
%attr(640, aoc, aoc) /opt/aws/aws-otel-collector/etc/.env
%attr(640, aoc, aoc) /opt/aws/aws-otel-collector/etc/extracfg.txt

/etc/init/aws-otel-collector.conf
/etc/systemd/system/aws-otel-collector.service
/usr/bin/aws-otel-collector-ctl
/etc/amazon/aws-otel-collector
/var/log/amazon/aws-otel-collector
/var/run/amazon/aws-otel-collector

%pre
# Stop the agent before upgrades.
if [ $1 -ge 2 ]; then
    if [ -x /opt/aws/aws-otel-collector/bin/aws-otel-collector-ctl ]; then
        /opt/aws/aws-otel-collector/bin/aws-otel-collector-ctl -a stop
    fi
fi

# create group
if ! grep "^aoc:" /etc/group >/dev/null 2>&1; then
    groupadd -r aoc >/dev/null 2>&1
    echo "create group aoc, result: $?"
fi

# create user
if ! id aoc >/dev/null 2>&1; then
    useradd -r -M aoc -d /home/aoc -g aoc -c "ADOT Collector" -s $(test -x /sbin/nologin && echo /sbin/nologin || (test -x /usr/sbin/nologin && echo /usr/sbin/nologin || (test -x /bin/false && echo /bin/false || echo /bin/sh))) >/dev/null 2>&1
    echo "create user aoc, result: $?"
fi


%clean
rm -rf ${RPM_BUILD_ROOT}
