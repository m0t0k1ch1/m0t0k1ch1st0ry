#!/bin/sh

set -eu

title=$1

if [ -z "${title}" ]; then
  echo "title is not specified"
  exit 1
fi

date=`date '+%Y/%m/%d'`
path="blog/${date}/${title}.md"

hugo new $path
