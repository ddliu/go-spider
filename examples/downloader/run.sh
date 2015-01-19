#!/bin/bash
mkdir -p download
go run main.go -depth=3 -max=100 -follow=http://tooling.github.io/book-of-modern-frontend-tooling/* -target=./download/ http://tooling.github.io/book-of-modern-frontend-tooling/