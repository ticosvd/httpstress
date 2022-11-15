#!/bin/bash
# $1 - count of threads
# $2 - source IP address
go run httpclient.go -url http://10.199.100.100:9443/rfc4 -t $1 -s $2
