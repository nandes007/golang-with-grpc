#!/bin/sh

set -e

echo "run db migration"
/app/migrate -path /app/migration -database "$DB_URL" -verbose up

echo "start the app"
exec "$@"