#!/bin/bash
cd $GOPATH/src/github.com/nearmap/goreportcardlite
make install
make start &
sleep 5s
curl "http://localhost:8000/report?repo=$GOPATH/src/$1" > $GOPATH/src/$1/goreportcard.htm