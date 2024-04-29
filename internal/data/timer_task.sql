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

 Date: 29/04/2024 14:02:16
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for timer_task
-- ----------------------------
DROP TABLE IF EXISTS `timer_task`;
CREATE TABLE `timer_task`  (
  `task_id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'taskId',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `modify_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `timer_id` bigint UNSIGNED NOT NULL COMMENT 'TimerId',
  `app` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'app',
  `output` varchar(1028) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'output',
  `status` tinyint NOT NULL DEFAULT 0 COMMENT '0新建，1成功，2失败',
  `run_timer` bigint NULL DEFAULT NULL COMMENT '运行时间',
  `cost_time` bigint UNSIGNED NOT NULL COMMENT '执行耗时',
  PRIMARY KEY (`task_id`) USING BTREE,
  UNIQUE INDEX `idx_timer_id_run_timer`(`timer_id` ASC, `run_timer` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 245 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'Timer Task任务信息' ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
