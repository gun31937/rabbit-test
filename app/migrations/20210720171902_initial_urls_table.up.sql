CREATE TABLE urls (
      id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
      short_code varchar(15) unique,
      full_url text,
      expiry timestamptz,
      hits numeric(36) default 0,
      created_at timestamptz default CURRENT_TIMESTAMP
);