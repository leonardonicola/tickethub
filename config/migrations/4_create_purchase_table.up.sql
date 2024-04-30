CREATE TYPE payment_status AS ENUM ('succeeded','processing', 'canceled');
CREATE TABLE IF NOT EXISTS purchase(
  id uuid primary key not null,
  ticket_id uuid references tickets(id),
  user_id uuid references users(id),
  quantity integer not null,
  status payment_status,
  payment_method varchar(150),
  created_at timestamp default now()
);