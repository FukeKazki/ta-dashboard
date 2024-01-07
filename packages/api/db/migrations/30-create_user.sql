
-- +migrate Up
create table user (
  id BINARY(16) NOT NULL DEFAULT (UUID_TO_BIN(UUID(), 1)) primary key,
  created_at datetime not null default current_timestamp,
  updated_at datetime on update current_timestamp,
  name varchar(255) not null comment 'ユーザー名',
  password varchar(255) not null comment 'パスワード'
) comment="ユーザー";
-- +migrate Down
drop table user;
