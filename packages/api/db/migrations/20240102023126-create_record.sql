
-- +migrate Up
create table record (
  id int unsigned auto_increment primary key,
  created_at datetime not null default current_timestamp,
  updated_at datetime on update current_timestamp,
  user_id int unsigned not null comment 'ユーザーID',
  course_id int unsigned not null comment 'コースID',
  first_lap varchar(255) comment '1周目',
  second_lap varchar(255) comment '2周目',
  third_lap varchar(255) comment '3周目',
  total_lap varchar(255) not null comment '合計',
  FOREIGN KEY (user_id) REFERENCES user(id),
  FOREIGN KEY (course_id) REFERENCES course(id),
  UNIQUE KEY (user_id, course_id)
) comment="レコード";


-- +migrate Down
drop table record;
