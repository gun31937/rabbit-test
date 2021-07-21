CREATE TABLE urls (
      id SERIAL PRIMARY KEY,
      short_code varchar(15) unique,
      full_url text,
      expiry timestamptz,
      hits numeric(36) default 0,
      created_at timestamptz,
      updated_at timestamptz,
      deleted_at timestamptz
);