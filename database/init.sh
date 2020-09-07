set -eu

psql -f /docker-entrypoint-initdb.d/tables.sql -U admin admin
