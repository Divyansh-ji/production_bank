#!/bin/sh
set -e

echo "run db migration"
/THEPROJECT/migrate -path /THEPROJECT/migration -database "$DB_SOURCE" -verbose up

echo "start the app"
exec "$@"
