#!/bin/sh

set -e

host="$1"
shift
cmd="$@"

# Change "dbname=$POSTGRES_USER" to "dbname=$POSTGRES_DB"
until PGPASSWORD=$POSTGRES_PASSWORD psql -h "$host" -U "$POSTGRES_USER" -d "$POSTGRES_DB" -c '\q'; do
  echo "Postgres is unavailable - sleeping"
  sleep 1
done

echo "Postgres is up - executing command"
exec $cmd
