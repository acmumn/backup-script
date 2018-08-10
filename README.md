## Creating the `backup` user

```sql
CREATE USER 'backup'@'localhost' IDENTIFIED BY 'backup';
GRANT RELOAD, REPLICATION CLIENT ON *.* TO 'backup'@'localhost';
```
