/*
 Navicat Premium Dump SQL

 Source Server         : Mysql
 Source Server Type    : MySQL
 Source Server Version : 80042 (8.0.42)
 Source Host           : localhost:3306
 Source Schema         : restaurant

 Target Server Type    : MySQL
 Target Server Version : 80042 (8.0.42)
 File Encoding         : 65001

 Date: 26/06/2025 16:49:31
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for customers
-- ----------------------------
DROP TABLE IF EXISTS `customers`;
CREATE TABLE `customers`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `cusname` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `password` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `cusphone` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `sex` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_customers_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of customers
-- ----------------------------
INSERT INTO `customers` VALUES (1, '2025-06-24 15:27:53.680', '2025-06-24 15:27:53.680', NULL, 'asd', '111111', '13182195527', '');
INSERT INTO `customers` VALUES (2, '2025-06-24 15:36:46.073', '2025-06-24 15:36:46.073', NULL, '123', '111111', '13160183439', '');
INSERT INTO `customers` VALUES (3, '2025-06-26 15:52:33.597', '2025-06-26 15:52:33.597', NULL, '王五', '123456', '13199990000', '');

-- ----------------------------
-- Table structure for dishes
-- ----------------------------
DROP TABLE IF EXISTS `dishes`;
CREATE TABLE `dishes`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `dishname` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `type` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `price` bigint NULL DEFAULT NULL,
  `picture` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_dishes_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 27 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of dishes
-- ----------------------------
INSERT INTO `dishes` VALUES (1, '2025-06-25 02:06:01.371', '2025-06-25 15:13:06.799', '2025-06-25 15:13:39.455', '红烧牛肉', '热菜', 30, '/uploads\\红烧牛肉.jpg');
INSERT INTO `dishes` VALUES (2, '2025-06-25 02:31:23.693', '2025-06-25 15:18:15.085', '2025-06-26 14:20:57.224', '炒饭', '主食', 10, '/uploads\\炒饭.jpg');
INSERT INTO `dishes` VALUES (4, '2025-06-26 14:22:54.856', '2025-06-26 14:22:54.856', NULL, '番茄炒蛋', '热菜', 12, '/uploads\\番茄炒蛋.jpg');
INSERT INTO `dishes` VALUES (5, '2025-06-26 14:25:05.345', '2025-06-26 14:25:05.345', NULL, '鱼香肉丝', '热菜', 24, '/uploads\\鱼香肉丝.jpg');
INSERT INTO `dishes` VALUES (6, '2025-06-26 14:25:48.934', '2025-06-26 14:25:48.934', NULL, '地三鲜', '热菜', 18, '/uploads\\地三鲜.jpg');
INSERT INTO `dishes` VALUES (7, '2025-06-26 14:26:39.743', '2025-06-26 14:26:39.743', NULL, '宫保鸡丁', '热菜', 22, '/uploads\\宫保鸡丁.jpg');
INSERT INTO `dishes` VALUES (8, '2025-06-26 14:27:06.933', '2025-06-26 14:27:06.933', NULL, '回锅肉', '热菜', 27, '/uploads\\回锅肉.jpg');
INSERT INTO `dishes` VALUES (9, '2025-06-26 14:27:29.218', '2025-06-26 14:27:29.218', NULL, '酸辣土豆丝', '热菜', 9, '/uploads\\酸辣土豆丝.jpg');
INSERT INTO `dishes` VALUES (10, '2025-06-26 14:27:58.246', '2025-06-26 14:27:58.246', NULL, '水煮肉片', '热菜', 35, '/uploads\\水煮肉片.jpg');
INSERT INTO `dishes` VALUES (11, '2025-06-26 14:28:28.271', '2025-06-26 14:28:28.271', NULL, '糖醋排骨', '热菜', 34, '/uploads\\糖醋排骨.jpg');
INSERT INTO `dishes` VALUES (12, '2025-06-26 14:28:49.223', '2025-06-26 14:28:49.223', NULL, '干煸豆角', '热菜', 13, '/uploads\\干煸豆角.jpg');
INSERT INTO `dishes` VALUES (13, '2025-06-26 14:29:31.253', '2025-06-26 14:29:31.253', NULL, '蚝油生菜', '热菜', 12, '/uploads\\蚝油生菜.jpg');
INSERT INTO `dishes` VALUES (14, '2025-06-26 14:29:55.191', '2025-06-26 14:29:55.191', NULL, '紫菜蛋花汤', '汤类', 14, '/uploads\\紫菜蛋花汤.jpg');
INSERT INTO `dishes` VALUES (15, '2025-06-26 14:30:15.479', '2025-06-26 14:30:15.479', NULL, '蘑菇鸡蛋汤', '汤类', 14, '/uploads\\蘑菇鸡蛋汤.jpg');
INSERT INTO `dishes` VALUES (16, '2025-06-26 14:39:37.423', '2025-06-26 14:39:37.423', NULL, '米饭', '主食', 2, '/uploads\\米饭.jpg');
INSERT INTO `dishes` VALUES (17, '2025-06-26 14:39:59.853', '2025-06-26 14:39:59.853', NULL, '馒头', '主食', 3, '/uploads\\馒头.jpg');
INSERT INTO `dishes` VALUES (18, '2025-06-26 14:40:17.987', '2025-06-26 14:40:17.987', NULL, '面条', '主食', 10, '/uploads\\面条.jpg');
INSERT INTO `dishes` VALUES (19, '2025-06-26 14:41:11.792', '2025-06-26 14:41:11.792', NULL, '凉拌藕片', '凉菜', 16, '/uploads\\凉拌藕片.jpg');
INSERT INTO `dishes` VALUES (20, '2025-06-26 14:41:47.684', '2025-06-26 14:41:47.684', NULL, '凉拌腐竹', '凉菜', 14, '/uploads\\凉拌腐竹.jpg');
INSERT INTO `dishes` VALUES (21, '2025-06-26 14:42:21.955', '2025-06-26 14:42:21.955', NULL, '夫妻肺片', '凉菜', 27, '/uploads\\夫妻肺片.jpg');
INSERT INTO `dishes` VALUES (22, '2025-06-26 14:43:02.489', '2025-06-26 14:43:02.489', NULL, '口水鸡', '凉菜', 21, '/uploads\\口水鸡.jpg');
INSERT INTO `dishes` VALUES (23, '2025-06-26 14:43:23.701', '2025-06-26 14:43:23.701', NULL, '双皮奶', '甜点', 14, '/uploads\\双皮奶.jpg');
INSERT INTO `dishes` VALUES (24, '2025-06-26 14:43:53.030', '2025-06-26 14:43:53.030', NULL, '驴打滚', '甜点', 18, '/uploads\\驴打滚.jpg');
INSERT INTO `dishes` VALUES (25, '2025-06-26 15:47:28.092', '2025-06-26 15:50:55.873', NULL, '皮蛋豆腐', '凉菜', 20, '/uploads\\皮蛋豆腐.jpg');
INSERT INTO `dishes` VALUES (26, '2025-06-26 15:47:49.827', '2025-06-26 15:51:27.627', '2025-06-26 15:51:37.436', '桂花糕', '甜点', 20, '/uploads\\桂花糕.jpg');

-- ----------------------------
-- Table structure for order_items
-- ----------------------------
DROP TABLE IF EXISTS `order_items`;
CREATE TABLE `order_items`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `order_id` bigint UNSIGNED NULL DEFAULT NULL,
  `dish_id` bigint UNSIGNED NULL DEFAULT NULL,
  `dishname` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `quantity` bigint NULL DEFAULT NULL,
  `price` bigint NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_order_items_deleted_at`(`deleted_at` ASC) USING BTREE,
  INDEX `fk_orders_dishorder`(`order_id` ASC) USING BTREE,
  CONSTRAINT `fk_orders_dishorder` FOREIGN KEY (`order_id`) REFERENCES `orders` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 17 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of order_items
-- ----------------------------
INSERT INTO `order_items` VALUES (13, '2025-06-26 16:03:46.985', '2025-06-26 16:03:46.985', NULL, 14, 4, '番茄炒蛋', 1, 12);
INSERT INTO `order_items` VALUES (14, '2025-06-26 16:03:46.985', '2025-06-26 16:03:46.985', NULL, 14, 7, '宫保鸡丁', 1, 22);
INSERT INTO `order_items` VALUES (15, '2025-06-26 16:03:46.985', '2025-06-26 16:03:46.985', NULL, 14, 16, '米饭', 2, 2);
INSERT INTO `order_items` VALUES (16, '2025-06-26 16:03:46.985', '2025-06-26 16:03:46.985', NULL, 14, 21, '夫妻肺片', 1, 27);

-- ----------------------------
-- Table structure for orders
-- ----------------------------
DROP TABLE IF EXISTS `orders`;
CREATE TABLE `orders`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `telephone` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `tableid` bigint NULL DEFAULT NULL,
  `customerid` bigint NULL DEFAULT NULL,
  `totalprice` bigint NULL DEFAULT NULL,
  `status` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_orders_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 15 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of orders
-- ----------------------------
INSERT INTO `orders` VALUES (14, '2025-06-26 16:03:46.984', '2025-06-26 16:04:28.553', NULL, '13182195527', 6, 1, 65, '已完成');

-- ----------------------------
-- Table structure for staffs
-- ----------------------------
DROP TABLE IF EXISTS `staffs`;
CREATE TABLE `staffs`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `staname` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `password` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `staphone` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `salary` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `age` bigint NULL DEFAULT NULL,
  `sex` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `role` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_staffs_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 17 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of staffs
-- ----------------------------
INSERT INTO `staffs` VALUES (1, '2025-06-24 16:06:20.000', '2025-06-26 14:53:42.242', NULL, '吕仕琳', '112233', '13182195527', '1000000', 20, '男', '管理员');
INSERT INTO `staffs` VALUES (2, '2025-06-24 16:58:10.671', '2025-06-24 16:58:10.671', '2025-06-24 22:10:51.650', '2222', '123456', '13182889909', '2000', 20, '男', '普通员工');
INSERT INTO `staffs` VALUES (3, '2025-06-24 22:17:46.366', '2025-06-25 01:10:49.050', '2025-06-25 21:16:45.423', '12121', '123456', '12121212121', '12222', 45, '女', '普通员工');
INSERT INTO `staffs` VALUES (4, '2025-06-26 14:45:36.642', '2025-06-26 14:49:05.600', NULL, '张三', '123456', '18994529859', '14000', 26, '男', '普通员工');
INSERT INTO `staffs` VALUES (5, '2025-06-26 14:46:23.862', '2025-06-26 14:49:01.090', NULL, '张玲玲', '123456', '19710512591', '15000', 21, '女', '普通员工');
INSERT INTO `staffs` VALUES (6, '2025-06-26 14:47:30.293', '2025-06-26 14:48:56.105', NULL, '李凤英', '123456', '18374023854', '12000', 34, '女', '普通员工');
INSERT INTO `staffs` VALUES (7, '2025-06-26 14:48:22.927', '2025-06-26 14:48:45.218', NULL, '孙笑川', '123456', '15368326502', '200', 38, '男', '普通员工');
INSERT INTO `staffs` VALUES (8, '2025-06-26 14:50:11.965', '2025-06-26 14:53:46.350', NULL, '陈一星', '123456', '18402349568', '8000', 21, '男', '普通员工');
INSERT INTO `staffs` VALUES (9, '2025-06-26 14:51:08.218', '2025-06-26 14:53:51.576', NULL, '胡兴启', '123456', '15036632405', '8500', 39, '男', '普通员工');
INSERT INTO `staffs` VALUES (10, '2025-06-26 14:51:48.261', '2025-06-26 14:53:56.667', NULL, '陈秋', '123456', '12947925344', '8500', 44, '女', '普通员工');
INSERT INTO `staffs` VALUES (11, '2025-06-26 14:52:48.197', '2025-06-26 14:54:03.647', NULL, '胡屿田', '123456', '18036665295', '500000', 21, '男', '管理员');
INSERT INTO `staffs` VALUES (12, '2025-06-26 14:53:12.396', '2025-06-26 14:54:08.775', NULL, '李金伦', '123456', '17396430889', '500000', 21, '男', '管理员');
INSERT INTO `staffs` VALUES (13, '2025-06-26 14:57:56.463', '2025-06-26 14:57:56.463', '2025-06-26 14:59:02.104', '12121', '123456', '222', '222', 22, '男', '');
INSERT INTO `staffs` VALUES (14, '2025-06-26 14:59:18.600', '2025-06-26 14:59:18.600', '2025-06-26 15:00:22.053', '11', '123456', '111', '22', 19, '男', '');
INSERT INTO `staffs` VALUES (15, '2025-06-26 15:00:29.194', '2025-06-26 15:00:29.194', '2025-06-26 15:00:35.219', '11', '123456', '111', '22', 19, '男', '普通员工');
INSERT INTO `staffs` VALUES (16, '2025-06-26 15:44:51.307', '2025-06-26 15:45:06.118', NULL, '蔡旭', '123456', '13155556666', '6000', 18, '女', '普通员工');

-- ----------------------------
-- Table structure for tables
-- ----------------------------
DROP TABLE IF EXISTS `tables`;
CREATE TABLE `tables`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `statu` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `maxnum` bigint NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_tables_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 16 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tables
-- ----------------------------
INSERT INTO `tables` VALUES (1, '2025-06-25 16:20:01.999', '2025-06-25 16:21:15.011', '2025-06-25 16:21:24.078', '空闲', 5);
INSERT INTO `tables` VALUES (2, '2025-06-25 16:21:32.323', '2025-06-26 15:22:57.279', NULL, '空闲', 4);
INSERT INTO `tables` VALUES (3, '2025-06-25 16:21:32.695', '2025-06-25 16:21:32.695', NULL, '空闲', 4);
INSERT INTO `tables` VALUES (4, '2025-06-25 16:21:32.875', '2025-06-25 16:21:32.875', NULL, '空闲', 4);
INSERT INTO `tables` VALUES (5, '2025-06-25 16:21:33.039', '2025-06-25 16:21:33.039', NULL, '空闲', 4);
INSERT INTO `tables` VALUES (6, '2025-06-25 16:21:33.206', '2025-06-26 16:04:28.563', NULL, '空闲', 4);
INSERT INTO `tables` VALUES (7, '2025-06-25 16:21:33.370', '2025-06-25 16:21:33.370', NULL, '空闲', 4);
INSERT INTO `tables` VALUES (8, '2025-06-25 16:21:33.533', '2025-06-25 16:21:33.533', NULL, '空闲', 4);
INSERT INTO `tables` VALUES (9, '2025-06-25 16:21:33.713', '2025-06-25 16:21:33.713', NULL, '空闲', 4);
INSERT INTO `tables` VALUES (10, '2025-06-25 16:21:33.879', '2025-06-25 16:21:33.879', NULL, '空闲', 4);
INSERT INTO `tables` VALUES (11, '2025-06-25 16:21:34.045', '2025-06-25 16:21:34.045', NULL, '空闲', 4);
INSERT INTO `tables` VALUES (12, '2025-06-25 16:21:34.208', '2025-06-25 16:21:34.208', NULL, '空闲', 4);
INSERT INTO `tables` VALUES (13, '2025-06-25 16:21:34.373', '2025-06-25 16:21:34.373', NULL, '空闲', 4);
INSERT INTO `tables` VALUES (14, '2025-06-25 16:21:34.561', '2025-06-25 16:21:34.561', '2025-06-26 15:45:58.651', '空闲', 4);
INSERT INTO `tables` VALUES (15, '2025-06-26 15:45:37.443', '2025-06-26 15:45:37.443', NULL, '空闲', 3);

SET FOREIGN_KEY_CHECKS = 1;
