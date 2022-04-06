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

INSERT INTO users (login, encrypted_password)
VALUES ('user2', '$2a$04$naW5K8k.SIE9NlZwTbplzOHyHSilqnQ.PjY1QT2IYJgKsLO3KCCda');

INSERT INTO users (login, encrypted_password)
VALUES ('user1', '$2a$04$ub8JnTuTcTLTROg8SDjiO.TEUzDBm.5HRNjapV0.9Yz0uUHamEFRa');
