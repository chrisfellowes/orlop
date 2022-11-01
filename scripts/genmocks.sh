#!/usr/bin/env sh

# Code generated by shipbuilder init 1.21.0. DO NOT EDIT.

if [ ! -f "./scripts/check.sh" ]; then
  cd $(command dirname -- "$(command readlink -f "$(command -v -- "$0")")")/..
fi

. ./scripts/check.sh

if [ -d "./features/mocks" ]; then
  check go mockery

  if [ -f "./features/mocks/.env" ]; then
    . ./features/mocks/.env
  fi

  mockery="$mockery $mockery_args"

  echo "Generating mocks..."
  find mocks -type f -name \*.go -not -name module.go -delete

  set -e

  if [ -n "$mockery_pattern" ]; then
    $mockery --recursive --name "$mockery_pattern"
  else
    $mockery --all
  fi
fi