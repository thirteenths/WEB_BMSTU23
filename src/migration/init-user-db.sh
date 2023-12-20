#!/bin/bash

set -e

##psql -v ON_ERROR_STOP=1 --username standup <<-EOSQL
    #CREATE USER standup WITH PASSWORD 'password';
    #CREATE DATABASE standup;
    #GRANT ALL PRIVILEGES ON SCHEMA public TO standup;
    #GRANT ALL PRIVILEGES ON DATABASE standup TO standup;
    #ALTER DATABASE standup OWNER TO standup;
#EOSQL

psql -U standup -d standup -f /docker-entrypoint-initdb.d/table.sql

psql -U standup -d standup -f /docker-entrypoint-initdb.d/insert_data/insert_place.sql
psql -U standup -d standup -f /docker-entrypoint-initdb.d/insert_data/insert_event.sql
psql -U standup -d standup -f /docker-entrypoint-initdb.d/insert_data/insert_comic.sql
psql -U standup -d standup -f /docker-entrypoint-initdb.d/insert_data/insert_poster.sql