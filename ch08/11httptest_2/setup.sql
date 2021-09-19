-- psql -U gwp -f setup.sql -d gwp
-- postgres -D /usr/local/var/postgres

drop table posts;

create table posts (
  id      serial primary key,
  content text,
  author  varchar(255)
);