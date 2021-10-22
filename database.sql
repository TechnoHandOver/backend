\c postgres;

DROP SCHEMA public CASCADE;
CREATE SCHEMA public;

CREATE TABLE user_ (
    id SERIAL PRIMARY KEY,
    vk_id INT NOT NULL UNIQUE,
    name VARCHAR(100) NOT NULL CHECK (length(name) >= 2), --TODO: точно 100?
    avatar TEXT CHECK (avatar IS NULL OR length(avatar) >= 10) --TODO: VARCHAR(...)
);

CREATE TABLE ads (
    id SERIAL PRIMARY KEY,
    user_author_id INT NOT NULL REFERENCES user_ (id) ON DELETE CASCADE,
    user_author_vk_id INT NOT NULL REFERENCES user_ (vk_id) ON DELETE CASCADE,
    loc_dep VARCHAR(100) NOT NULL,
    loc_arr VARCHAR(100) NOT NULL,
    date_arr TIMESTAMP NOT NULL, --TODO: not timestamp
    min_price INT NOT NULL CHECK (min_price >= 0),
    comment VARCHAR(100) NOT NULL
);

CREATE INDEX ON user_ USING hash (vk_id);

CREATE INDEX ON ads (user_author_id);
