CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  email hermes_string NOT NULL,
  first_name hermes_string,
  last_name hermes_string,
  created_at hermes_timestamp,
  updated_at hermes_timestamp
);

CREATE UNIQUE INDEX users_email_idx ON users (lower(email));