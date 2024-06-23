CREATE DATABASE IF NOT EXISTS vote;

-- admin  $2a$10$s1x4.Fx57DP7CU8GSia5KeEnFyYd9LTwUQXpF0tQHrV8nHFtwUJXi <nil>

USE vote;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;  
CREATE TABLE IF NOT EXISTS user (
    `id` BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
    `username` VARCHAR(20) NOT NULL,
    `password` VARCHAR(500) NOT NULL, 
    `role` int DEFAULT 2,  
    `created_time` DATETIME NULL DEFAULT NULL,
    `updated_time` DATETIME NULL DEFAULT NULL,
    `deleted_time` DATETIME NULL DEFAULT NULL,
  │  user.go
    PRIMARY KEY(`id`) USING BTREE
)ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;
-- INSERT INTO user(id, username, password, role, created_time, updated_time, deleted_time) VALUES(0, "admin", "", 1, CURRENT_TIEMSTAMP(), CURRENT_TIEMSTAMP(), NULL);

-- ----------------------------
-- Table structure for vote
-- ----------------------------
-- type 0 表示单选，1表示多选
-- status 0 正常， 1超时
-- time 有效时长
-- user_id 创建人
DROP TABLE IF EXISTS `vote`; 
CREATE TABLE IF NOT EXISTS vote (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `title` VARCHAR(255),
    `type` INT(4), 
    `status` INT(4), 
    `time`  BIGINT, 
    `user_id` BIGINT, 
    `created_time` DATETIME NULL DEFAULT NULL,
    `updated_time` DATETIME NULL DEFAULT NULL,
    `deleted_time` DATETIME NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
)ENGINE = InnoDB AUTO_INCREMENT = 0 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

insert into vote(title, type, status, time, user_id, created_time, updated_time) values("food", 0, 0, 24, 0, "2024-06-23 22:38:20", "2024-06-24 22:38:20");
-- ----------------------------
-- Table structure for vote_opt
-- ----------------------------
DROP TABLE IF EXISTS `vote_opt`;  
CREATE TABLE IF NOT EXISTS vote_opt(
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255),
    `vote_id` BIGINT,
    `count` INT,
    `created_time` DATETIME NULL DEFAULT NULL,
    `updated_time` DATETIME NULL DEFAULT NULL,
    `deleted_time` DATETIME NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
)ENGINE = InnoDB AUTO_INCREMENT = 0 CHARACTER SET = utf8 COLLATE utf8_general_ci ROW_FORMAT = Dynamic;

insert into vote_opt(name, vote_id, count, created_time, updated_time) values("fish", 1, 0, "2024-06-23 22:38:20", "2024-06-24 22:38:20"),
    ("malou", 1, 0, "2024-06-23 22:38:20", "2024-06-24 22:38:20");
-- ----------------------------
-- Table structure for vote_opt_user
-- ----------------------------
DROP TABLE IF EXISTS `vote_opt_user`; 
CREATE TABLE IF NOT EXISTS vote_opt_user(
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `user_id` BIGINT,
    `vote_id` BIGINT,
    `vote_opt_id` BIGINT,
    `created_time` DATETIME NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
)ENGINE = InnoDB AUTO_INCREMENT = 0 CHARACTER SET = utf8 COLLATE utf8_general_ci ROW_FORMAT = Dynamic;
