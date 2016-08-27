#!/bin/sh

TITLE=$1

if [ -z "${TITLE}" ]; then
  echo "title is not specified"
  exit 1
fi

DATE=`date '+%Y/%m/%d'`
PATH="blog/${DATE}/${TITLE}.md"

$GOPATH/bin/hugo new $PATH
