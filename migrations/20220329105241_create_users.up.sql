-- CREATE DATABASE some_go_api_db;
-- CREATE DATABASE some_go_api_db_test;

CREATE TABLE IF NOT EXISTS users
(
    user_id            bigserial not null primary key,
    login              varchar   not null unique,
    encrypted_password varchar   not null
);

CREATE TABLE IF NOT EXISTS sessions
(
    session_id      bigserial                not null primary key,
    user_id         bigserial                not null,
    token           varchar                  not null unique,
    expiration_time timestamp with time zone not null,
    FOREIGN KEY (user_id) REFERENCES users (user_id)
);

CREATE TABLE IF NOT EXISTS authorization_events_names
(
    event_id smallint  not null primary key,
    event_name varchar not null
);

INSERT INTO authorization_events_names (event_id, event_name)
VALUES (0, 'success');

INSERT INTO authorization_events_names (event_id, event_name)
VALUES (1, 'wrong pass');

INSERT INTO authorization_events_names (event_id, event_name)
VALUES (2, 'user blocked');

CREATE TABLE IF NOT EXISTS authorization_events
(
    created_time timestamp with time zone not null primary key DEFAULT NOW(),
    user_id                bigserial not null,
    event_id                  smallint  not null,
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    FOREIGN KEY (event_id) REFERENCES authorization_events_names (event_id)
);

INSERT INTO users (login, encrypted_password)
VALUES ('user2', '$2a$04$naW5K8k.SIE9NlZwTbplzOHyHSilqnQ.PjY1QT2IYJgKsLO3KCCda');

INSERT INTO sessions (token, user_id, expiration_time)
VALUES ('simple_token', 1, TIMESTAMP '2011-05-16 15:36:38');

INSERT INTO users (login, encrypted_password)
VALUES ('user1', '$2a$04$ub8JnTuTcTLTROg8SDjiO.TEUzDBm.5HRNjapV0.9Yz0uUHamEFRa');

INSERT INTO authorization_events (user_id, event_id) VALUES (1, 1);