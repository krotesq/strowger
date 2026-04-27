-- Write your migrate up statements here
CREATE TABLE source_rtmp (
  source_id UUID PRIMARY KEY references source(id) ON DELETE CASCADE,
  url TEXT NOT NULL,
  stream_key TEXT NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);
---- create above / drop below ----
DROP TABLE IF EXISTS source_rtmp;
-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.