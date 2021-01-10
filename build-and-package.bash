#!/bin/bash

PKG_NAME="bouyguestelecom-sms"
DEBIAN_VERSIONS="stretch buster bullseye"

set -e
export GOPATH=$HOME/go

VERSION=$1
echo "Version to build: $VERSION"

PWD_BACKUP=$(pwd)
GOOS="linux"
VERSION_PATTERN="Version"
ARCHITECTURE_PATTERN="Architecture"
BOUYGUESTELECOM_SMS_GO_REPOSITORY_DIR=$GOPATH/src/github.com/cyosp/bouyguestelecom-sms
CONTROL_FILE_PATH=$BOUYGUESTELECOM_SMS_GO_REPOSITORY_DIR/DEBIAN/control
ARCHITECTURES=$(grep "$ARCHITECTURE_PATTERN" $CONTROL_FILE_PATH | cut -d ' ' -f 2-)

for architecture in $ARCHITECTURES
do
  echo "Build application for architecture: $architecture"
  cd $BOUYGUESTELECOM_SMS_GO_REPOSITORY_DIR
  case $architecture in
    amd64)
      env GOOS=$GOOS GOARCH=amd64 go build
    ;;
    armhf)
      env GOOS=$GOOS GOARCH=arm GOARM=7 go build
    ;;
  esac

  for version in $DEBIAN_VERSIONS
  do
    TMP_DIR=$(mktemp -d)
    echo "Package for Debian: $version"

    cp -a $BOUYGUESTELECOM_SMS_GO_REPOSITORY_DIR $TMP_DIR

    BOUYGUESTELECOM_SMS_SOURCES_DIR="$TMP_DIR/$$/$PKG_NAME"
    BOUYGUESTELECOM_SMS_PACKAGE_DIR="$BOUYGUESTELECOM_SMS_SOURCES_DIR/$PKG_NAME"
    PACKAGE_BIN_DIR=$BOUYGUESTELECOM_SMS_PACKAGE_DIR/usr/bin
    PACKAGE_DEBIAN_DIR=$BOUYGUESTELECOM_SMS_PACKAGE_DIR/DEBIAN
    CONTROL_FILE_PATH=$PACKAGE_DEBIAN_DIR/control
    mkdir -p $PACKAGE_BIN_DIR $PACKAGE_DEBIAN_DIR

    cp bouyguestelecom-sms $PACKAGE_BIN_DIR
    cp $BOUYGUESTELECOM_SMS_GO_REPOSITORY_DIR/DEBIAN/* $PACKAGE_DEBIAN_DIR

    PACKAGE_VERSION="$VERSION-0+deb"
    case $version in
      bullseye)
        PACKAGE_VERSION+="11"
      ;;
      buster)
        PACKAGE_VERSION+="10"
      ;;
      stretch)
        PACKAGE_VERSION+="9"
      ;;
    esac
    PACKAGE_VERSION+="u0"
    sed -i "s/$VERSION_PATTERN.*/$VERSION_PATTERN: $PACKAGE_VERSION/" $CONTROL_FILE_PATH
    sed -i "s/$ARCHITECTURE_PATTERN.*/$ARCHITECTURE_PATTERN: $architecture/" $CONTROL_FILE_PATH

    cd $BOUYGUESTELECOM_SMS_SOURCES_DIR
    sudo dpkg-deb --build $PKG_NAME 2>&1 >/dev/null

    cp $BOUYGUESTELECOM_SMS_SOURCES_DIR/$PKG_NAME.deb $PWD_BACKUP/${PKG_NAME}_${PACKAGE_VERSION}_${architecture}.deb

    rm -rf $TMP_DIR

    cd $BOUYGUESTELECOM_SMS_GO_REPOSITORY_DIR
  done
done

cd $PWD_BACKUP

exit 0
