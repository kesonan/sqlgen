#!/bin/bash

wd=$(pwd)

function generate() {
    dir=$1
    name=$2
    output="$dir/$name"
    if [ "$output" == "/" ]; then
       exit 1
    fi

    rm -rf "$output"
    mkdir -p "$output"

    cd "$output"
    sqlgen $name -f "$dir/example.sql" -o .
}


cd "$wd"

# generate bun code
generate "$wd" "bun"

# generate gorm code
generate "$wd" "gorm"

# generate sql code
generate "$wd" "sql"

# generate sqlx code
generate "$wd" "sqlx"

# generate xorm code
generate "$wd" "xorm"

# go mod tidy
go mod tidy

go test ./...