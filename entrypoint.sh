#!/bin/sh

set -e

host="$1"
shift
cmd="$@"

until PGPASSWORD=12345678 psql -h "$host" -U "postgres" -c '\q'; do
  sleep 1
done

make migrate_init
exec $cmd