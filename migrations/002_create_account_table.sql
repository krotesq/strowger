-- Write your migrate up statements here
CREATE TABLE account (
  uuid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  username TEXT NOT NULL UNIQUE,
  password_hash TEXT NOT NULL,
  active BOOLEAN NOT NULL DEFAULT TRUE,
  failed_login_attempts INT DEFAULT 0,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);
---- create above / drop below ----
DROP TABLE IF EXISTS account;
-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.