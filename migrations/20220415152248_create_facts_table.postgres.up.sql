CREATE TABLE facts (
  id uuid NOT NULL,
  created_at timestamp without time zone NOT NULL,
  updated_at timestamp without time zone NOT NULL,
  number SERIAL,
  content TEXT NOT NULL
);

ALTER TABLE facts ADD CONSTRAINT facts_pkey PRIMARY KEY (id);
