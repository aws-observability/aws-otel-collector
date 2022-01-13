#!/bin/sh
BINARY=aoc_linux_x86_64
if [[ $TARGETPLATFORM == "linux/arm64" ]]; then
  BINARY=aoc_linux_aarch64
fi
echo binary name is ${BINARY}
mkdir -p /workspace/build/linux
if [[ ! -f /workspace/${BINARY} ]]; then
  echo binary does not exsit
  exit 1
fi
echo copying binary
cp /workspace/${BINARY} /workspace/build/linux/aoc_linux
