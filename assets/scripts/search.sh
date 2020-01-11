#!/usr/bin/env bash

mode=$1
keyword=$2
base_dir=${3:-$HOME}

function search_file_and_dirs() {
  /usr/bin/find ${base_dir} -xdev -iname "${keyword}*" -type f,d -printf "%A@, %p, %y\n" | sort -r
}

function search_text() {
  /usr/bin/grep -rIHni "${keyword}" "${base_dir}" | awk -F':' '{ print $1 ", " $2 ", " substr($0, index($0,$3))}'
}


if [ "${mode}" == "all" ]
then
  search_file_and_dirs
  search_text
elif [ "${mode}" == "files" ]
then
  search_file_and_dirs
elif [ "${mode}" == "text" ]
then
  search_text
else
  echo "Mode should be 'files' or 'text' or 'all'"
  exit 1
fi