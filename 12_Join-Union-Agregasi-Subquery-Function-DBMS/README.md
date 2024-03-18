## Lanjutan MySQL sebelumnya

JOIN (sebuah klausa untuk mengkombinasikan record dari dua atau lebih tabel)

- INNER
- LEFT
- RIGHT

```sql
# INNER JOIN
SELECT t.message FROM users u INNER JOIN tweets t ON u.id = t.user_id;

# LEFT JOIN
SELECT u.username, t.message FROM users u LEFT JOIN tweets t ON u.id = t.user_id;

# RIGHT JOIN
SELECT u.username, t.message FROM users u RIGHT JOIN tweets t ON u.id = t.user_id;
```

UNION

```sql
SELECT username, fullname FROM users WHERE id=1 UNION SELECT username, fullname FROM users WHERE id=2;
```

AGGREGATE \
Fungsi Agregasi adalah fungsi dimana nilai beberapa baris dikelompokkan bersama untuk membentuk nilai ringkasan tunggal.

- MIN
- MAX
- SUM
- AVG
- COUNT
- HAVING

```sql
# MIN
SELECT MIN(id) AS id FROM users;
SELECT MIN(created_at) AS created_at FROM users;

# MAX
SELECT MAX(id) FROM users
SELECT MAX(created_at) FROM users;

# SUM
SELECT SUM(favourite_count) FROM tweets WHERE user_id=1;

# AVG
SELECT AVG(favourite_count) FROM tweets WHERE user_id=1;

# COUNT
SELECT COUNT(1) FROM tweets WHERE user_id=1;

# HAVING
SELECT user_id FROM tweets GROUP BY user_id HAVING SUM(favourite_count) > 2;
```

FUNCTION \

```sql
DELIMITER $$
CREATE FUNCTION sf_count_tweet_peruser (user_id_p int) RETURNS INT DETERMINISTIC
BEGIN
DECLARE total INT;
SELECT COUNT(*) INTO total FROM tweets WHERE user_id = user_id_p AND type='tweets';
RETURN total;
END$$
DELIMITER;
```

TRIGGER FUNCTION

```sql
DELIMITER $$
CREATE TRIGGER delete_all_data_user
BEFORE DELETE ON users FOR EACH ROW
BEGIN
-- declare variables
DECLARE v_user_id INT;
SET v_user_id=OLD.id;
-- trigger code
DELETE FROM tweets WHERE user_id = v_user_id;
DELETE FROM user_followers WHERE user_id = v_user_id;
END$$
```
