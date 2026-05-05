CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE git_diff_app.credentials (
    user_id       SERIAL PRIMARY KEY,
    login         TEXT NOT NULL UNIQUE,
    salt          TEXT NOT NULL,
    password_hash TEXT NOT NULL
);

ALTER TABLE git_diff_app.users
ALTER COLUMN id TYPE INTEGER USING (id::integer);

ALTER TABLE git_diff_app.users
ADD CONSTRAINT credentials_id_fkey FOREIGN KEY (id)
REFERENCES git_diff_app.credentials (user_id);

CREATE TABLE git_diff_app.sessions (
    id         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id    INTEGER NOT NULL REFERENCES git_diff_app.users (id),
    created_at TIMESTAMPTZ NOT NULL,
    ttl        TIMESTAMPTZ NOT NULL
);

ALTER TABLE git_diff_app.users 
    ADD COLUMN role VARCHAR(5) DEFAULT 'user';