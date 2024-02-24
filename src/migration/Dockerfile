FROM postgres:14-alpine

COPY table.sql /docker-entrypoint-initdb.d/
COPY /insert_data/. /docker-entrypoint-initdb.d/insert_data/
COPY init-user-db.sh /docker-entrypoint-initdb.d/
#COPY 00_init.sql /docker-entrypoint-initdb.d/

RUN chmod +x /docker-entrypoint-initdb.d/init-user-db.sh

EXPOSE 5432
