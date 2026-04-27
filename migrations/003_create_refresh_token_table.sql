-- Write your migrate up statements here
CREATE TABLE refresh_token (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  account_id UUID NOT NULL REFERENCES account(id) ON DELETE CASCADE,
  token_hash TEXT NOT NULL UNIQUE,
  expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
  revoked_at TIMESTAMP WITH TIME ZONE NULL,
  user_agent TEXT NULL,
  ip_address INET NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);
---- create above / drop below ----
DROP TABLE IF EXISTS refresh_token;
-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.