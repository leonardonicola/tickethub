CREATE EXTENSION IF NOT EXISTS pg_trgm;
CREATE EXTENSION IF NOT EXISTS unaccent;

CREATE TABLE IF NOT EXISTS events (
  id uuid primary key not null,
  title varchar(170) not null,
  description varchar(500) not null,
  address varchar(255) not null,
  date date not null,
  age_rating smallint not null,
  genre_id uuid references genres(id),
  poster varchar(200) not null,
  created_at timestamp default now(),
  updated_at timestamp default now()
);

ALTER TABLE events
ADD
  COLUMN searchable text
  GENERATED ALWAYS AS (
    COALESCE(title, '') || ' ' || COALESCE(address, '') 
    )
  STORED;

CREATE INDEX events_searchable_idx ON events USING gin(to_tsvector('portuguese', searchable));