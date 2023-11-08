#!/usr/bin/env bash

if [ "$1" = "cat" ]; then
  echo -n "dog2"
else
  echo -n "$1"
fi
