\c postgres;

DROP DATABASE handover;
CREATE DATABASE handover;

GRANT ALL PRIVILEGES ON DATABASE handover TO handover;

\c handover;

CREATE TABLE user_ (
    id SERIAL PRIMARY KEY,
    vk_id INT NOT NULL UNIQUE,
    name TEXT NOT NULL CHECK (length(name) >= 2),
    avatar TEXT NOT NULL CHECK (length(avatar) >= 10)
);

CREATE TABLE ads (
    id SERIAL PRIMARY KEY,
    user_author_id INT NOT NULL REFERENCES user_ (id) ON DELETE CASCADE,
    location_from TEXT NOT NULL,
    location_to TEXT NOT NULL,
    time_from TIMESTAMP NOT NULL,
    time_to TIMESTAMP NOT NULL,
    min_price INT NOT NULL CHECK (min_price >= 0),
    comment TEXT NOT NULL
);

CREATE INDEX ON user_ USING hash (vk_id);

CREATE INDEX ON ads (user_author_id);
