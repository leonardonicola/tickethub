CREATE TABLE IF NOT EXISTS users (
  id uuid primary key,
  name varchar(100) not null,
  surname varchar(100) not null,
  address varchar(250) not null,
  email varchar(200) not null unique,
  password varchar(250) not null,
  created_at timestamp not null default now(),
  deleted_at timestamp
);