FROM postgis/postgis:12-master

RUN apt-get update
RUN apt install postgresql-12-pgrouting -y
