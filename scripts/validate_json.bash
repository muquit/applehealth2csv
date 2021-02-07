#!/bin/bash
# muquit@muquit.com Jan-10-2021 
for i in ./$1/*.json
do
    file="$i"
    echo -n "Validate $file: "
    if jq empty $file; then
        echo "OK"
    else
        echo "FAILED"
    fi
done
