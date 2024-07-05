-- ----------------------------
-- DataBase
-- ----------------------------
CREATE DATABASE lottery;

-- ----------------------------
-- Table structure for article
-- ----------------------------
DROP TABLE IF EXISTS `inventory`;
CREATE TABLE IF NOT EXISTS `inventory` (
    `id` int(11) NOT NULL AUTO_INCREMENT COMMENT "奖品id, 自增",
    `created_at` DATETIME(3) NULL DEFAULT NULL COMMENT "创建时间",
    `name` varchar(20) NOT NULL COMMENT "奖品名称",
    `description` varchar(100) NOT NULL DEFAULT "" COMMENT "奖品描述",
    `picture` int(11) NOT NULL DEFAULT "0" COMMENT,
)
