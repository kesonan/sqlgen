#!/bin/bash

wd=$(pwd)

# generate create code
cd "$wd/create"
rm -rf "user_*.go"
sqlgen sql -f "example.sql" -o .

# generate delete code
cd "$wd/delete"
rm -rf "user_*.go"
sqlgen sql -f "example.sql" -o .

# generate read code
cd "$wd/read"
rm -rf "user_*.go"
sqlgen sql -f "example.sql" -o .

# generate update code
cd "$wd/update"
rm -rf "user_*.go"
sqlgen sql -f "example.sql" -o .

cd "$wd"

# go mod tidy
go mod tidy

go test ./...