
-- +migrate Up
INSERT INTO user (name, password, public_profile_id) VALUES ('ユーザー1', 'password', 1);
INSERT INTO user (name, password, public_profile_id) VALUES ('ユーザー2', 'password', 2);

-- +migrate Down
DELETE FROM user WHERE id IN (1, 2);
