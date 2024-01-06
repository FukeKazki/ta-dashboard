
-- +migrate Up
create table record (
  id int unsigned auto_increment primary key,
  created_at datetime not null default current_timestamp,
  updated_at datetime on update current_timestamp,
  public_profile_id int unsigned not null comment '公開プロフィールID',
  course_id int unsigned not null comment 'コースID',
  first_lap varchar(255) comment '1周目',
  second_lap varchar(255) comment '2周目',
  third_lap varchar(255) comment '3周目',
  total_lap varchar(255) not null comment '合計',
  FOREIGN KEY (public_profile_id) REFERENCES public_profile(id),
  FOREIGN KEY (course_id) REFERENCES course(id),
  UNIQUE KEY (public_profile_id, course_id)
) comment="レコード";


-- +migrate Down
drop table record;
