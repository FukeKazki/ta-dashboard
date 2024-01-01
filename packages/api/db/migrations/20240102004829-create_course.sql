
-- +migrate Up
create table course (
  id int unsigned auto_increment primary key,
  created_at datetime not null default current_timestamp,
  updated_at datetime on update current_timestamp,
  name varchar(255) not null comment 'コース名',
  thumbnail varchar(255) not null comment 'サムネイル'
) comment="コース";

-- +migrate Down
drop table;
