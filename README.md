# bouyguestelecom-sms

⚠️ Not maintained because of cell phone provider change ⚠️

[![BSD-3 license](https://img.shields.io/badge/license-BSD--3--Clause-428F7E.svg)](https://tldrlegal.com/license/bsd-3-clause-license-%28revised%29)

`bouyguestelecom-sms` is based on [bouyguessms](https://github.com/tomsquest/bouyguessms) whose:
 * Development has been stopped
 * Original license is here: [bouyguessms/LICENSE](bouyguessms/LICENSE)

## Setup
```bash
sudo apt install golang-go
cat << EOF >> ~/.bashrc

# Go environment
export GOPATH=$HOME/go
EOF
```

## Download sources
```bash
go get github.com/cyosp/bouyguestelecom-sms
```

## Download dependencies
```bash
go get github.com/pkg/errors
go get golang.org/x/net/publicsuffix
```

## Build
 * Linux AMD 64 bits
```bash
env GOOS=linux GOARCH=amd64 go build
```
 * Linux ARM v7
```bash
env GOOS=linux GOARCH=arm GOARM=7 go build
```

## Reduce binary size
```bash
strip bouyguestelecom-sms
```

## Installation
```bash
sudo mv bouyguestelecom-sms /usr/local/bin
```
