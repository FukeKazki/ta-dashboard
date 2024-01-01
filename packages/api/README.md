## 設計
![ER図](https://storage.googleapis.com/zenn-user-upload/4d905d50e165-20240101.jpg)

# メモ
## migration
apiコンテナに入って`sql-migrate`を使う。
migrationファイルは`db/migrations`に作成する。

```
# migrationファイルの作成
sql-migrate new create_users
# migration
sql-migrate up
```

