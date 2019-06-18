#!/bin/bash  
echo -e " \nOutPut from fund"
./fund

sleep 4
echo -e " \nOutPut from defer"
./defer

go run useof-defer.go funding.go
