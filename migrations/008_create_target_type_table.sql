-- Write your migrate up statements here
CREATE TABLE target_type (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name TEXT NOT NULL UNIQUE,
  description TEXT,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);
---- create above / drop below ----
DROP TABLE IF EXISTS target_type;
-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.