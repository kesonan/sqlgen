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

    # generate create code
    cd "$output/create"
    sqlgen $name -f "example.sql" -o .

    # generate delete code
    cd "$output/delete"
    sqlgen $name -f "example.sql" -o .

    # generate read code
    cd "$output/read"
    sqlgen $name -f "example.sql" -o .

    # generate update code
    cd "$output/update"
    sqlgen $name -f "example.sql" -o .
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