CREATE DATABASE IF NOT EXISTS vote;

USE vote;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF NOT EXISTS `user`;  -- 用户表
CREATE TABLE IF NOT EXISTS user (
    id BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
    username VARCHAR(20) NOT NULL,
    `password` VARCHAR(500) NOT NULL,   -- 加密后的
    `role` int DEFAULT 2,  -- 1是管理员，2是普通用户
    created_time DATETIME NULL DEFAULT NULL,
    updated_time DATETIME NULL DEFAULT NULL,
    deleted_time DATETIME NULL DEFAULT NULL,
    PRIMARY KEY(`id`) USING BTREE,
)ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;
INSERT INTO user(id, username, password, role, created_time, updated_time, deleted_time) VALUES(0, "admin", "", 1, CURRENT_TIEMSTAMP(), CURRENT_TIEMSTAMP(), NULL);

-- ----------------------------
-- Table structure for vote
-- ----------------------------
DROP TABLE IF NOT EXISTS `vote`; -- 投票内容
CREATE TABLE IF NOT EXISTS vote (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    title VARCHAR(255),
    `type` INT(4), -- 0 表示单选，1表示多选
    `status` INT(4), -- 0正常，1表示超时 
    `time`  BIGINT, -- 有效时长
    user_id BIGINT, -- 创建人
    created_time DATETIME NULL DEFAULT NULL,
    updated_time DATETIME NULL DEFAULT NULL,
    deleted_time DATETIME NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE,
)ENGINE = InnoDB AUTO_INCREMENT = 0 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for vote_opt
-- ----------------------------
DROP TABLE IF NOT EXISTS `vote_opt`;  -- 投票选项
CREATE TABLE IF NOT EXISTS vote_opt(
    id BIGINT NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255),
    vote_id BIGINT,
    count INT,
    created_time DATETIME NULL DEFAULT NULL,
    updated_time DATETIME NULL DEFAULT NULL,
    deleted_time DATETIME NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE,
)ENGINE = InnoDB AUTO_INCREMENT = 0 CHARACTER SET = utf8 COLLATE utf8_general_ci ROW_FORMAT = Dynamic

-- ----------------------------
-- Table structure for vote_opt_user
-- ----------------------------
DROP TABLE IF NOT EXISTS `vote_opt_user`; -- 用户和投票的关系
CREATE TABLE IF NOT EXISTS (
    id BIGINT NOT NULL AUTO_INCREMENT,
    user_id BIGINT,
    vote_id BIGINT,
    vote_opt_id BIGINT,
    created_time DATETIME NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE,
)ENGINE = InnoDB AUTO_INCREMENT = 0 CHARACTER SET = utf8 COLLATE utf8_general_ci ROW_FORMAT = Dynamic
