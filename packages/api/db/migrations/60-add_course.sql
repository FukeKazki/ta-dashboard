
-- +migrate Up
INSERT INTO course (name, thumbnail) VALUES ('マリオカートスタジアム', 'https://www.nintendo.co.jp/switch/aabpa/assets/images/course/thumbnail/1-1.jpg');
INSERT INTO course (name, thumbnail) VALUES ('ウォーターパーク', 'https://www.nintendo.co.jp/switch/aabpa/assets/images/course/thumbnail/1-2.jpg');
INSERT INTO course (name, thumbnail) VALUES ('スイーツキャニオン', 'https://www.nintendo.co.jp/switch/aabpa/assets/images/course/thumbnail/1-3.jpg');
INSERT INTO course (name, thumbnail) VALUES ('ドッスンいせき', 'https://www.nintendo.co.jp/switch/aabpa/assets/images/course/thumbnail/1-4.jpg');

-- +migrate Down
DELETE FROM course WHERE id IN (1, 2, 3, 4);
