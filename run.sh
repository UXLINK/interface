#!/bin/bash

env=$1

if [[  $env == "test" ]]; then
  ./uxuy-interface  -f ./uxuy-test.yaml
elif [[  $env == "pro" ]]; then
  ./uxuy-interface  -f ./uxuy.yaml
else
  ./bin/uxuy-interface -f ./src/etc/uxuy-dev.yaml
fi
