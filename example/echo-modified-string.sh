#!/usr/bin/env bash

if [ "$1" = "cat" ]; then
  echo -n "dog"
else
  echo -n "$1"
fi
