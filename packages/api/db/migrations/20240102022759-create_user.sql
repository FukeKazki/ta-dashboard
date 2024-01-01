
-- +migrate Up
create table user (
  id int unsigned auto_increment primary key,
  created_at datetime not null default current_timestamp,
  updated_at datetime on update current_timestamp,
  name varchar(255) not null comment 'ユーザー名'
) comment="ユーザー";
-- +migrate Down
drop table user;
