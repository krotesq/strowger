-- Write your migrate up statements here
INSERT INTO account (username, password_hash)
VALUES ('admin', crypt('admin', gen_salt('bf', 12)))
ON CONFLICT (username) DO NOTHING;
---- create above / drop below ----
DELETE FROM account WHERE username = 'admin';
-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.