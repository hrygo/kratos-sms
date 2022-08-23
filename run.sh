#!/bin/sh

go mod tidy
go clean 
make build 
./bin/kratos-sms --conf bin/configs 
