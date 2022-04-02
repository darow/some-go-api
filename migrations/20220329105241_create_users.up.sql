-- CREATE DATABASE some_go_api_db;

CREATE TABLE IF NOT EXISTS users
(
    id
    bigserial
    not
    null
    primary
    key,
    login
    varchar
    not
    null
    unique,
    encrypted_password
    varchar
    not
    null
);