#!/bin/sh
set -e
rm -rf manpages
mkdir manpages
go run . man | gzip -c >manpages/cranlogs.1.gz
