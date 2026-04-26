CREATE SCHEMA git_diff_app;

CREATE TABLE git_diff_app.users (
    id           SERIAL PRIMARY KEY,
    version      BIGINT NOT NULL DEFAULT 1,

    full_name    VARCHAR(100)
                     CHECK(char_length(full_name) BETWEEN 3 AND 100),

    phone_number VARCHAR(20)
                     CHECK(phone_number ~ '^\+[0-9]'
                         AND char_length(phone_number) BETWEEN 10 AND 20)
);

CREATE TABLE git_diff_app.files (
    id         SERIAL PRIMARY KEY,
    version    BIGINT NOT NULL DEFAULT 1,

    name       VARCHAR(100)   NOT NULL
                   CHECK(char_length(name) BETWEEN 1 AND 100),

    content    VARCHAR(10000) NOT NULL
                   CHECK(char_length(content) BETWEEN 1 AND 10000),

    created_at TIMESTAMPTZ    NOT NULL,

    owner_id   INTEGER REFERENCES git_diff_app.users(id),
    expiration TIMESTAMP
)