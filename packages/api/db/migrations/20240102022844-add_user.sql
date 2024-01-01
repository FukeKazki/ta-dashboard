
-- +migrate Up
INSERT INTO user (name) VALUES ('ユーザー1');
INSERT INTO user (name) VALUES ('ユーザー2');

-- +migrate Down
DELETE FROM user WHERE id IN (1, 2);
