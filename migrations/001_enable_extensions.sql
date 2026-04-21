-- Write your migrate up statements here
CREATE EXTENSION IF NOT EXISTS pgcrypto;
CREATE EXTENSION IF NOT EXISTS pg_trgm;
---- create above / drop below ----
DROP EXTENSION IF EXISTS pgcrypto;
DROP EXTENSION IF EXISTS pg_trgm;
-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.