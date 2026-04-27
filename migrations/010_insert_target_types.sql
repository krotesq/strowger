-- Write your migrate up statements here
INSERT INTO target_type (name, description)
VALUES ('rtmp', 'Custom RTMP Server');
---- create above / drop below ----
DELETE FROM source_type WHERE name = 'rtmp';
-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.