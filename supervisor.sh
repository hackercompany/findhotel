#!/bin/bash

source /usr/local/application/source/geolocation/gopath.rc
go test model
echo "Making binary"
make
echo "Make Finished"
/usr/bin/supervisord -n -c /etc/supervisord.d/process.conf