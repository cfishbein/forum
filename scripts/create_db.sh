#!/bin/bash

readonly dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

if [[ -f forum.db ]]; then
	echo "Database already exists"
	exit 0
fi

# TODO put db file somewhere more idiomatic
sqlite3 forum.db < "$dir"/create_tables.sql
