#!/bin/bash

INSTALL_MODE=$1

SOURCE_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
TARGET_DIR=/usr/share/dude
TARGET_BINARY=/usr/bin/dude

function dev_install() {
  ln -s ${SOURCE_DIR}/assets ${TARGET_DIR}
  ln -s ${SOURCE_DIR}/dude ${TARGET_BINARY}
}

function system_install() {
  echo "System install not implemented yet"
  exit 1
}

CURRENT_UID=$(id -u)
if [ "${CURRENT_UID}" != 0 ]; then
  echo "You need to be 'root' to install"
  exit 2
fi

case ${INSTALL_MODE} in
    "dev" )
        dev_install ;;
    "system" )
        system ;;
    *)
      echo "Usage: $0 [dev|system]"
esac
