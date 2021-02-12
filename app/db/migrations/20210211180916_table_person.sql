
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE client (
	pk serial NOT NULL,
	created_at timestamptz NOT NULL,
	update_at timestamptz NOT NULL,
	first_name varchar(140) NOT NULL,
	last_name varchar(140) NOT NULL,
	ci varchar(12) NOT NULL,
	birthday date NULL,
	sex varchar(1) NULL,
	nationality varchar(140) NULL,
	des_type varchar(140) NULL,
	code1 varchar(20) NULL,
	code2 varchar(20) NULL,
	code3 varchar(20) NULL,
	email varchar(254) NULL,
	direction varchar(500) NULL,
	phone varchar(10) NULL,
	CONSTRAINT clients_client_pkey PRIMARY KEY (pk)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE client;
