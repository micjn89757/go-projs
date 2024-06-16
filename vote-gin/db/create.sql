CREATE DATABASE IF NOT EXISTS vote;

USE vote;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF NOT EXISTS `user`;
CREATE TABLE IF NOT EXISTS user (
    id BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
    username VARCHAR(20) NOT NULL,
    password VARCHAR(500) NOT NULL,
    role int DEFAULT 2,
    created_time DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_time DATETIME DEFAULT CURRENT_TIMESTAMP,
);


-- ----------------------------
-- Table structure for vote
-- ----------------------------
DROP TABLE IF NOT EXISTS `vote`
CREATE TABLE IF NOT EXISTS vote (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    title VARCHAR(255),
    `type` INT(4), -- 0 表示单选，1表示多选
    `status` INT(4), -- 0正常，1表示超时 
    `time`  BIGINT, -- 有效时长
    userid BIGINT, -- 创建人
    created_time DATETIME NULL DEFAULT NULL,
    updated_time DATETIME NULL DEFAULT NULL,
    deleted_time DATETIME NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE,
)ENGINE = InnoDB AUTO_INCREMENT = 0 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;