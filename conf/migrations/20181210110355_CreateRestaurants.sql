
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE restaurants(
  id serial primary key not null,
  name text,
  type text,
  phone text,
  location point
);

CREATE INDEX on restaurants(name);
-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP INDEX restaurants_name_idx;

DROP TABLE restaurants;