#!/bin/bash
proj_name=server
cd ../..
export GOPATH=$(pwd)
cd src/$proj_name/
export GOAPP=$proj_name
export GOENV=local
. .env
bee migrate -driver=postgres -conn="postgres://postgres:root@localhost:5432/$db_name?sslmode=disable"