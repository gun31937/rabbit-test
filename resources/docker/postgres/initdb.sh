#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
	CREATE DATABASE rabbit_db;
	GRANT ALL PRIVILEGES ON DATABASE rabbit_db TO postgres;
EOSQL
