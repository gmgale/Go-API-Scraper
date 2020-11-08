FROM postgres:latest

COPY init.sql /docker-entrypoint-initdb.d/
COPY ormconfig.docker.json ./ormconfig.json

RUN apt-get -q update && apt-get -qy install netcat

EXPOSE 5432