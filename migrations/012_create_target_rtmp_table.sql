-- Write your migrate up statements here
CREATE TABLE target_rtmp (
  target_id UUID PRIMARY KEY references target(id) ON DELETE CASCADE,
  url TEXT NOT NULL,
  stream_key TEXT NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);
---- create above / drop below ----
DROP TABLE IF EXISTS target_rtmp;
-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.