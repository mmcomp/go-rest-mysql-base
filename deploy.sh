#!/bin/bash

git pull

go build -o main .

./main migrate run

./main
