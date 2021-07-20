CREATE TABLE urls (
      id UUID NOT NULL CONSTRAINT urls_pkey PRIMARY KEY,
      short_code varchar(15),
      full_url text,
      expiry timestamptz NOT NULL,
      hits numeric(36),
      created_at timestamptz NOT NULL
);