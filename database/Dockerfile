FROM mysql:8.0.33-debian
WORKDIR /app
COPY ./my.cnf /etc/mysql/conf.d/
COPY ./init.sql /docker-entrypoint-initdb.d/