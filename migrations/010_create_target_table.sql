-- Write your migrate up statements here
CREATE TABLE target (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name TEXT NOT NULL,
  description TEXT,
  target_type_id UUID NOT NULL REFERENCES target_type(id) ON DELETE RESTRICT,
  account_id UUID NOT NULL REFERENCES account(id) ON DELETE CASCADE,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),

  UNIQUE (account_id, name)
);
---- create above / drop below ----
DROP TABLE IF EXISTS target;
-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.