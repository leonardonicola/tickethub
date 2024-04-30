CREATE TABLE IF NOT EXISTS tickets (
  id uuid primary key not null,
  name varchar(100) not null,
  price decimal(7,2) not null,
  total_qty integer not null,
  available_qty integer not null,
  description text,
  max_per_user integer not null,
  event_id uuid references events(id)
);