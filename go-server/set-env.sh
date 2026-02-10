#!bin/bash

while IFS='=' read -r key value; do
    echo "export $key=$value"
    export $key=$value
done < ../.env