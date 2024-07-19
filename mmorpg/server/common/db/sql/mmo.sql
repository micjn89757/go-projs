-- ------------------------
-- DATABASE
-- ------------------------
CREATE DATABASE mmo;
USE mmo;

-- ------------------------
-- TABLE structure for user
-- ------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `account` varchar(15) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '用户名',
  `password` varchar(32) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '密码',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `deleted_at`  datetime DEFAULT NULL COMMENT '删除时间'
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `username` (`account`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC COMMENT='用户表';



-- ------------------------
-- TABLE structure for actor
-- ------------------------
DROP TABLE IF EXISTS `actor`;
CREATE TABLE `actor` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '角色id',
  `account` varchar(15) NOT NULL COMMENT '所属账号',
  `scene_id` int DEFAULT NULL COMMENT '场景id',
  `created_time` datetime DEFAULT NULL COMMENT '创建时间',
  `deleted_time` datetime DEFAULT NULL COMMENT '删除时间'
  `nickname` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '角色名',
  PRIMARY KEY (`id`),
  KEY `account` (`account`)
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
