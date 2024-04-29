/*
 Navicat Premium Data Transfer

 Source Server         : xtimer_docker
 Source Server Type    : MySQL
 Source Server Version : 80036 (8.0.36)
 Source Host           : localhost:8086
 Source Schema         : bitstorm-svr-go

 Target Server Type    : MySQL
 Target Server Version : 80036 (8.0.36)
 File Encoding         : 65001

 Date: 29/04/2024 14:02:31
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for xtimer
-- ----------------------------
DROP TABLE IF EXISTS `xtimer`;
CREATE TABLE `xtimer`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'TimerId',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `modify_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `app` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'app',
  `name` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'name',
  `status` tinyint NOT NULL DEFAULT 0 COMMENT '0新建，1激活，2未激活',
  `cron` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'cron表达式',
  `notify_http_param` varchar(8192) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '回调上下文',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 83671 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'Timer 信息' ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
