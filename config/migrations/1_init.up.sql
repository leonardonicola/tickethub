CREATE TABLE IF NOT EXISTS users (
  id uuid primary key,
  name varchar(100) not null,
  surname varchar(100) not null,
  address varchar(250) not null,
  email varchar(200) not null unique,
  password varchar(250) not null,
  created_at timestamp default now(),
  updated_at timestamp default now(),
  deleted_at timestamp
);

CREATE TABLE IF NOT EXISTS genres(
  id uuid primary key not null,
  name varchar(150) not null
);

INSERT INTO genres (id, name) VALUES 
('25e227a3-97f2-437d-a61f-9388df1ebf2a', 'Música'),
('92c3a1e9-fde5-49a5-ba71-10c82f76b566', 'Tecnologia'),
('7810867a-1f60-442d-9c7d-93a78b19a695', 'Comédia'),
('1ee9bc94-51ad-405f-8d92-ff4aa2c25317', 'Agropecuária'),
('b90f0f04-3813-4edc-8fd2-169f90f865e9', 'Esportes'),
('daa3791d-750d-4a38-96e8-0f1a36cbbdcb', 'Culinária'),
('4c0bc617-17ee-4de3-8a61-21e78eefad85', 'Arte e Cultura'),
('3b54e160-4bb8-492a-af33-6e4da62d7a05', 'Negócios e Empreendedorismo'),
('b440a212-0b20-44f5-aee6-92c8c9f35261', 'Saúde e Bem-estar'),
('3cf84b6d-312e-40ec-8e9c-3a00a13d145e', 'Ciência e Educação');
