CREATE TABLE trademarks (
    id bigserial NOT NULL PRIMARY KEY,
    application_number varchar,
    application_date varchar,
    registration_date varchar,
    application_language_code varchar,
    second_language_code varchar,
    expiry_date varchar,
    name text NOT NULL
);