CREATE TABLE IF NOT EXISTS events (
  id uuid primary key not null,
  title varchar(170) not null,
  description varchar(500) not null,
  address varchar(255) not null,
  date date not null,
  age_rating smallint not null,
  genre varchar(50) not null,
  poster varchar(200) not null
);