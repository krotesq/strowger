-- Write your migrate up statements here
INSERT INTO source_type (name, description)
VALUES ('rtmp', 'RTMP livestream input source');
---- create above / drop below ----
DELETE FROM source_type WHERE name = 'rtmp';
-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.