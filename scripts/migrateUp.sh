#!/bin/bash

if [ -f .env ]; then
    source .env
fi

cd sql/schema

goose postgres $DIRECT_URL up