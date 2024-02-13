#!/bin/bash

go mod tidy && \
  go build -o go_practice && \
  ./go_practice
