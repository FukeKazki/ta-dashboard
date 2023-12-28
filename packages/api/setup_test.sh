#/bin/bash

set -eu

# データベースを作成
mysql -u root -e "create database if not exists \`${MYSQL_TEST_DATABASE}\` default character set utf8mb4 default collate utf8mb4_general_ci";
