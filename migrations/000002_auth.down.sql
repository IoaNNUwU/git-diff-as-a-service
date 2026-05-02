DROP EXTENSION pgcrypto;

DROP TABLE git_diff_app.credentials;

ALTER TABLE git_diff_app.users
DROP CONSTRAINT credentials_id_fkey;

ALTER TABLE git_diff_app.users
ALTER COLUMN id TYPE BIGINT USING (id::bigint);

