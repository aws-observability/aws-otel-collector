#!/usr/bin/env bash

# Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License").
# You may not use this file except in compliance with the License.
# A copy of the License is located at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# or in the "license" file accompanying this file. This file is distributed
# on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
# express or implied. See the License for the specific language governing
# permissions and limitations under the License.

set -e

echo "****************************************"
echo "Creating deb file for Debian Linux ${ARCH}"
echo "****************************************"

BUILD_ROOT="$(pwd)/build/linux/debian"
# remove "v" since deb only allow version name with digits.
VERSION=$(sed 's/v//g' < VERSION)
AOC_ROOT=${BUILD_ROOT}

echo "BASE_ROOT: ${BUILD_ROOT}  agent_version: ${VERSION} AGENT_BIN_DIR_DEB: ${AOC_ROOT}"

echo "Creating debian folders"
mkdir -p "${AOC_ROOT}/opt/aws/aws-otel-collector/logs"
mkdir -p "${AOC_ROOT}/opt/aws/aws-otel-collector/bin"
mkdir -p "${AOC_ROOT}/opt/aws/aws-otel-collector/etc"
mkdir -p "${AOC_ROOT}/opt/aws/aws-otel-collector/var"
mkdir -p "${AOC_ROOT}/opt/aws/aws-otel-collector/doc"

mkdir -p "${AOC_ROOT}/etc/init"
mkdir -p "${AOC_ROOT}/etc/systemd/system/"
mkdir -p "${AOC_ROOT}/bin"
mkdir -p "${AOC_ROOT}/usr/bin"
mkdir -p "${AOC_ROOT}/etc/amazon"
mkdir -p "${AOC_ROOT}/var/log/amazon"
mkdir -p "${AOC_ROOT}/var/run/amazon"

echo "Copying application files"
cp LICENSE "${AOC_ROOT}/opt/aws/aws-otel-collector/"
cp VERSION "${AOC_ROOT}/opt/aws/aws-otel-collector/bin/"
cp "build/linux/${ARCH}/aoc" "${AOC_ROOT}/opt/aws/aws-otel-collector/bin/aws-otel-collector"
cp tools/ctl/linux/aws-otel-collector-ctl.sh "${AOC_ROOT}/opt/aws/aws-otel-collector/bin/aws-otel-collector-ctl"
# default configuration
cp config.yaml "${AOC_ROOT}/opt/aws/aws-otel-collector/var/.config.yaml"
cp .env "${AOC_ROOT}/opt/aws/aws-otel-collector/etc"
cp extracfg.txt "${AOC_ROOT}/opt/aws/aws-otel-collector/etc"
cp tools/packaging/linux/aws-otel-collector.service "${AOC_ROOT}/etc/systemd/system/"
cp tools/packaging/linux/aws-otel-collector.conf "${AOC_ROOT}/etc/init/"

chmod ug+rx "${AOC_ROOT}/opt/aws/aws-otel-collector/bin/aws-otel-collector"
chmod ug+rx "${AOC_ROOT}/opt/aws/aws-otel-collector/bin/aws-otel-collector-ctl"
chmod ug+rx "${AOC_ROOT}/opt/aws/aws-otel-collector/etc/.env"
chmod ug+rx "${AOC_ROOT}/opt/aws/aws-otel-collector/etc/extracfg.txt"

echo "Creating symlinks"
# bin
ln -f -s /opt/aws/aws-otel-collector/bin/aws-otel-collector-ctl "${AOC_ROOT}/usr/bin/aws-otel-collector-ctl"
# etc
ln -f -s /opt/aws/aws-otel-collector/etc "${AOC_ROOT}/etc/amazon/aws-otel-collector"
# log
ln -f -s /opt/aws/aws-otel-collector/logs "${AOC_ROOT}/var/log/amazon/aws-otel-collector"
# pid
ln -f -s /opt/aws/aws-otel-collector/var "${AOC_ROOT}/var/run/amazon/aws-otel-collector"

echo "Copying debian build materials"
cp tools/packaging/debian/conffiles "${AOC_ROOT}/"
cp tools/packaging/debian/preinst "${AOC_ROOT}/"
cp tools/packaging/debian/postinst "${AOC_ROOT}/"
cp tools/packaging/debian/prerm "${AOC_ROOT}/"
cp tools/packaging/debian/debian-binary "${AOC_ROOT}/"

echo "Constructing the control file"
{
  echo 'Package: aws-otel-collector'
  echo "Architecture: ${ARCH}"
  echo -n 'Version: '
  echo -n "${VERSION}"
  echo '-1'
  cat tools/packaging/debian/control
} > "${BUILD_ROOT}/control"

echo "Setting permissioning as required by debian"
cd "${AOC_ROOT}/.."
find ./debian -type d -exec chmod 755 '{}' \;
cd ~-
# the below permissioning is required by debian
cd "${AOC_ROOT}"
tar czf data.tar.gz opt etc usr var --owner=0 --group=0
cd ~-
cd "${AOC_ROOT}"
tar czf control.tar.gz control conffiles preinst prerm postinst --owner=0 --group=0
cd ~-

echo "Creating the debian package"
ar r "${AOC_ROOT}/bin/aws-otel-collector-${ARCH}.deb" "${AOC_ROOT}/debian-binary"
ar r "${AOC_ROOT}/bin/aws-otel-collector-${ARCH}.deb" "${AOC_ROOT}/control.tar.gz"
ar r "${AOC_ROOT}/bin/aws-otel-collector-${ARCH}.deb" "${AOC_ROOT}/data.tar.gz"

echo "Copy debian file to ${DEST}"
mkdir -p "${DEST}"
mv "${AOC_ROOT}/bin/aws-otel-collector-${ARCH}.deb" "${DEST}/aws-otel-collector.deb"

echo "remove working directory"
rm -rf "${AOC_ROOT}"
