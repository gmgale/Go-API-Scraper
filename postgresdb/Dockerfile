FROM postgres:latest
COPY init.sql /docker-entrypoint-initdb.d/
COPY ormconfig.docker.json ./ormconfig.json
EXPOSE 5432