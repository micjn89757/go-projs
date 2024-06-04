CREATE DATABASE IF NOT EXISTS vote;

USE vote;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF NOT EXISTS `user`;
CREATE TABLE IF NOT EXISTS user (
    id INT NOT NULL,
    username VARCHAR(20) NOT NULL,
    password VARCHAR(500) NOT NULL,
    role int DEFAULT 2,
    created_time DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_time DATETIME DEFAULT CURRENT_TIMESTAMP,
);
