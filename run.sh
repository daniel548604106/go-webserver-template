#!/bin/bash

## Go build only focuses on the files that are not test files
go build -o bookings cmd/web/*.go && ./bookings