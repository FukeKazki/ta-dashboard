
-- +migrate Up
INSERT INTO user (name, password) VALUES ('fukke', 'password');

-- +migrate Down
DELETE FROM user WHERE id IN (1);
