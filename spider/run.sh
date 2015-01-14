#!/bin/bash
mkdir -p download
go run main.go -depth=3 -max=100 -follow=*github.com/* -target=./download/ https://github.com/