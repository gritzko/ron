#!/bin/bash

clear
cat sep2bits.txt | ./sep2bits.pl > bitsep.go
go fmt bitsep.go
ragel -Z -G2 -e dfa.rl
go build
go test

