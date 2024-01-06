
-- +migrate Up
INSERT INTO public_profile (name) VALUES ('ユーザー1');
INSERT INTO public_profile (name) VALUES ('ユーザー2');
-- +migrate Down
DELETE FROM public_profile WHERE id IN (1, 2);
