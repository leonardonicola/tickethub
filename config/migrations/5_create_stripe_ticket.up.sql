CREATE TABLE IF NOT EXISTS stripe_ticket (
  id varchar(150) not null,
  price_id varchar(150) not null,
  ticket_id uuid references tickets(id),
  name varchar(150) not null,
  description text
);