#!/bin/bash

mygo_path=`pwd`
#export GOROOT=${mygo_path}/go
export GOPATH=${mygo_path}/gopath

inpath=`echo $PATH | grep "${GOROOT}" | wc -l`
if [ "${inpath}" != 1 ]; then
    export PATH=${GOROOT}/bin:$PATH
fi

gofmt -w ${GOPATH}/src/hellogo

# make version
VERSION=`git describe --all`
echo "version:$VERSION"
reversion=`git log -n 5 | grep commit | head -n1 | awk '{print $2}'`
last_author=`git log -n 5 | grep -A3 commit | grep Author: | head -n1 | awk '{print $2}'`
version="$VERSION ${reversion}
build at `date +"%Y-%m-%d %H:%M:%S"` `/sbin/ifconfig  | grep "inet addr" | head -n 1 | awk '{print substr($2,6)}'` `pwd`  
last changed author : $last_author"

tag=hello
if [ "$1" != "" ]; then
    tag=$1
fi

if [ "$2" == "debug" ]; then
    go build -ldflags "-X main._VERSION_=$VERSION.${reversion}" ${GOPATH}/src/hellogo/bin/${tag}.go
    if [ $? != 0 ]; then
        exit 1
    fi
else 
    go build -ldflags "-X main._VERSION_=$VERSION.${reversion}" ${GOPATH}/src/hellogo/bin/${tag}.go
    if [ $? != 0 ]; then
        exit 1
    fi
fi
