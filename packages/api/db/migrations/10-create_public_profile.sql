
-- +migrate Up
create table public_profile (
  id int unsigned auto_increment primary key,
  created_at datetime not null default current_timestamp,
  updated_at datetime on update current_timestamp,
  name varchar(255) not null comment 'ユーザー名'
) comment="公開プロフィール";
-- +migrate Down
drop table public_profile;
