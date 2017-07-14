create domain hermes_string text check (length(value) <= 255);
create domain hermes_timestamp timestamp without time zone default (now() at time zone 'utc');