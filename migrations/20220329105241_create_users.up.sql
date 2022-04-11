-- CREATE DATABASE some_go_api_db;
-- CREATE DATABASE some_go_api_db_test;

CREATE TABLE IF NOT EXISTS users
(
    user_id bigserial not null primary key,
    login varchar not null unique,
    encrypted_password varchar not null,
    login_attempts smallint not null
);

CREATE TABLE IF NOT EXISTS sessions
(
    session_id bigserial not null primary key,
    user_id bigserial not null,
    token varchar not null unique,
    expiration_time timestamp with time zone not null,
    CONSTRAINT fk_user
        FOREIGN KEY(user_id)
            REFERENCES users(user_id)
);

INSERT INTO users (login, encrypted_password, login_attempts)
VALUES ('user2', '$2a$04$naW5K8k.SIE9NlZwTbplzOHyHSilqnQ.PjY1QT2IYJgKsLO3KCCda', 0);

INSERT INTO sessions (token, user_id, expiration_time)
VALUES ('simple_token', 0, TIMESTAMP '2011-05-16 15:36:38');

INSERT INTO users (login, encrypted_password)
VALUES ('user1', '$2a$04$ub8JnTuTcTLTROg8SDjiO.TEUzDBm.5HRNjapV0.9Yz0uUHamEFRa');