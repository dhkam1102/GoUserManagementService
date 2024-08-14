CREATE DATABASE user_management;
USE user_management;

CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    role ENUM('admin', 'customer') NOT NULL DEFAULT 'customer' 
);


-- DROP TABLE users;

-- mysql> DESCRIBE users;
-- +----------+--------------------------+------+-----+----------+----------------+
-- | Field    | Type                     | Null | Key | Default  | Extra          |
-- +----------+--------------------------+------+-----+----------+----------------+
-- | id       | int                      | NO   | PRI | NULL     | auto_increment |
-- | email    | varchar(255)             | NO   | UNI | NULL     |                |
-- | password | varchar(255)             | NO   |     | NULL     |                |
-- | role     | enum('admin','customer') | NO   |     | customer |                |
-- +----------+--------------------------+------+-----+----------+----------------+



-- Current users
-- +----+-------------------+-----------+----------+
-- | id | email             | password  | role     |
-- +----+-------------------+-----------+----------+
-- |  1 | admin@example.com | adminpass | admin    |
-- |  2 | new@example.com   | newhello  | customer |
-- +----+-------------------+-----------+----------+