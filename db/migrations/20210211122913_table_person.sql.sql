-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE person (
    pk serial,
    first_name text,
    primary key(pk)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE person;
