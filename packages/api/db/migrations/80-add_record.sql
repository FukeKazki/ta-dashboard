
-- +migrate Up
-- INSERT INTO record (user_id, course_id, first_lap, second_lap, third_lap, total_lap) VALUES (1, 1, '0:45.943', '0:44.434', '0:43.967', '2:14.344');
-- INSERT INTO record (public_profile_id, course_id, total_lap) VALUES (1, 2, '1:48:963');

-- +migrate Down
-- DELETE FROM record WHERE id IN (1);
