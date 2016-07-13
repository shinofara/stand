#!/bin/sh
for os in darwin linux
do
    for arch in amd64
    do
	      GOOS=${os} GOARCH=${arch} go build -o ../../stand_${os}_${arch} -ldflags="-w -s"
    done
done
