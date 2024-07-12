/*
 Navicat Premium Data Transfer

 Source Server         : mysql8.0本地
 Source Server Type    : MySQL
 Source Server Version : 80033 (8.0.33)
 Source Host           : localhost:3306
 Source Schema         : wallet

 Target Server Type    : MySQL
 Target Server Version : 80033 (8.0.33)
 File Encoding         : 65001

 Date: 23/05/2023 19:20:58
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for address
-- ----------------------------
DROP TABLE IF EXISTS `address`;
CREATE TABLE `address`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '用户标识',
  `main_chain` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '主链',
  `address` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '地址',
  `private_key` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '私钥',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `chain_address_index`(`main_chain` ASC, `address` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of address
-- ----------------------------

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule`  (
  `ptype` varchar(10) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `v0` varchar(256) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `v1` varchar(256) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `v2` varchar(256) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `v3` varchar(256) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `v4` varchar(256) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `v5` varchar(256) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
INSERT INTO `casbin_rule` VALUES ('g', '1', '1', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('g', '2', '2', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '295', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '296', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '297', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '298', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '299', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '303', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '304', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '305', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '306', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '308', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '309', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '34', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '36', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '37', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '38', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '39', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('g', '3', '2', '', '', '', '');

-- ----------------------------
-- Table structure for currency
-- ----------------------------
DROP TABLE IF EXISTS `currency`;
CREATE TABLE `currency`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `main_chain` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '主链',
  `name` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '币种名称',
  `contract_address` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '合约地址',
  `decimals` int NOT NULL COMMENT '精度',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 64 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of currency
-- ----------------------------
INSERT INTO `currency` VALUES (1, 'tron', 'TRX', 'TBRop8PopYu8atWWez3g3ueVtSpseW78b6', 6, '2022-03-13 00:04:56', '2022-04-11 20:31:27');
INSERT INTO `currency` VALUES (2, 'bsc', 'BNB', '0xbb4cdb9cbd36b01bd1cbaebf2de08d9173bc095c', 18, '2022-03-15 18:13:31', '2022-08-23 15:22:39');
INSERT INTO `currency` VALUES (3, 'heco', 'HT', '0x1000000000000000000000000000000000000000', 18, NULL, '2022-05-17 17:49:23');
INSERT INTO `currency` VALUES (50, 'bsc', 'USDT', '0x55d398326f99059ff775485246999027b3197955', 18, '2022-08-03 15:59:40', '2022-08-03 15:59:40');
INSERT INTO `currency` VALUES (100, 'eth', 'ETH', '0x1000000000000000000000000000000000000000', 6, '2022-08-03 15:59:40', '2022-08-03 15:59:40');
-- ----------------------------
-- Table structure for demo_data_auth
-- ----------------------------
DROP TABLE IF EXISTS `demo_data_auth`;
CREATE TABLE `demo_data_auth`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `title` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '标题',
  `created_by` int UNSIGNED NULL DEFAULT 0 COMMENT '创建人',
  `updated_by` int UNSIGNED NULL DEFAULT 0 COMMENT '修改人',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '修改时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of demo_data_auth
-- ----------------------------

-- ----------------------------
-- Table structure for demo_gen
-- ----------------------------
DROP TABLE IF EXISTS `demo_gen`;
CREATE TABLE `demo_gen`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `demo_name` varchar(20) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '姓名',
  `demo_age` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '年龄',
  `classes` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '班级',
  `demo_born` datetime NULL DEFAULT NULL COMMENT '出生年月',
  `demo_gender` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '性别',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建日期',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '修改日期',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除日期',
  `created_by` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建人',
  `updated_by` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '修改人',
  `demo_status` tinyint NOT NULL DEFAULT 0 COMMENT '状态',
  `demo_cate` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '分类',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '代码生成测试表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of demo_gen
-- ----------------------------

-- ----------------------------
-- Table structure for demo_gen_class
-- ----------------------------
DROP TABLE IF EXISTS `demo_gen_class`;
CREATE TABLE `demo_gen_class`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '分类id',
  `class_name` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '分类名',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '代码生成关联测试表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of demo_gen_class
-- ----------------------------

-- ----------------------------
-- Table structure for demo_gen_tree
-- ----------------------------
DROP TABLE IF EXISTS `demo_gen_tree`;
CREATE TABLE `demo_gen_tree`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `parent_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '父级ID',
  `demo_name` varchar(20) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '姓名',
  `demo_age` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '年龄',
  `classes` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '班级',
  `demo_born` datetime NULL DEFAULT NULL COMMENT '出生年月',
  `demo_gender` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '性别',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建日期',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '修改日期',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除日期',
  `created_by` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建人',
  `updated_by` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '修改人',
  `demo_status` tinyint NOT NULL DEFAULT 0 COMMENT '状态',
  `demo_cate` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '分类',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '代码生成树形结构测试表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of demo_gen_tree
-- ----------------------------

-- ----------------------------
-- Table structure for fee_list
-- ----------------------------
DROP TABLE IF EXISTS `fee_list`;
CREATE TABLE `fee_list`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `main_chain` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '主链',
  `coin_name` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '手续费币种',
  `withdraw_address` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '手续费地址',
  `address` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '地址',
  `amount` decimal(30, 18) NOT NULL COMMENT '转移手续费',
  `hash` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT 'hash',
  `status` tinyint UNSIGNED NOT NULL DEFAULT 1 COMMENT '状态 1-待上链 2-上链成功 3-上链失败',
  `recharge_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值ID',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `address_index`(`main_chain` ASC, `address` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of fee_list
-- ----------------------------

-- ----------------------------
-- Table structure for height
-- ----------------------------
DROP TABLE IF EXISTS `height`;
CREATE TABLE `height`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `main_chain` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '主网',
  `height` int NOT NULL COMMENT '块号',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of height
-- ----------------------------

-- ----------------------------
-- Table structure for liquidity_currency
-- ----------------------------
DROP TABLE IF EXISTS `liquidity_currency`;
CREATE TABLE `liquidity_currency`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `main_chain` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '链',
  `name` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '币种名称',
  `contract_address` varchar(150) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '币种地址',
  `decimals` int NOT NULL COMMENT '精度',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of liquidity_currency
-- ----------------------------

-- ----------------------------
-- Table structure for liquidity_swap_trade
-- ----------------------------
DROP TABLE IF EXISTS `liquidity_swap_trade`;
CREATE TABLE `liquidity_swap_trade`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'id主键',
  `main_chain` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '主链',
  `block_hash` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '块hash',
  `from_address` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '交易方地址',
  `coin_token` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '币种1',
  `coin_token1` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '币种2',
  `coin_token_decimals` smallint UNSIGNED NOT NULL DEFAULT 0 COMMENT '币种1精度',
  `coin_token1_decimals` smallint UNSIGNED NOT NULL DEFAULT 0 COMMENT '币种2精度',
  `contract_address` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '币种1合约地址',
  `contract_address1` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '币种2合约地址',
  `amount` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '币种1数量',
  `amount1` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '币种2数量',
  `lp_address` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT 'LP合约地址',
  `fee_num` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '手续费流动性数量',
  `lp_num` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '流动性数量',
  `to_address` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '收币地址',
  `type` tinyint(1) NOT NULL COMMENT '类型 1-添加流动性 2-移除流动性',
  `hash` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '交易hash',
  `block_height` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '区块高度',
  `call_number` smallint UNSIGNED NULL DEFAULT 0 COMMENT '回调次数',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `hash_index`(`hash` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of liquidity_swap_trade
-- ----------------------------

-- ----------------------------
-- Table structure for monitor_currency
-- ----------------------------
DROP TABLE IF EXISTS `monitor_currency`;
CREATE TABLE `monitor_currency`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `main_chain` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '链',
  `name` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '币种名称',
  `contract_address` varchar(150) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '币种地址',
  `decimals` int NOT NULL COMMENT '精度',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of monitor_currency
-- ----------------------------

-- ----------------------------
-- Table structure for monitor_swap_trade
-- ----------------------------
DROP TABLE IF EXISTS `monitor_swap_trade`;
CREATE TABLE `monitor_swap_trade`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'id主键',
  `main_chain` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '主链',
  `block_hash` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '块hash',
  `from_address` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '交易方地址',
  `contract_address` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '兑换前币种',
  `contract_address1` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '兑换后币种',
  `amount` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '兑换前数量',
  `amount1` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '兑换后数量',
  `path` text CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL COMMENT '兑换路径',
  `to_address` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '收币地址',
  `type` tinyint(1) NOT NULL COMMENT '类型 1-买入 2-卖出',
  `hash` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '交易hash',
  `block_height` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '区块高度',
  `call_number` smallint UNSIGNED NULL DEFAULT 0 COMMENT '回调次数',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `hash_index`(`hash` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of monitor_swap_trade
-- ----------------------------

-- ----------------------------
-- Table structure for plugins_manage
-- ----------------------------
DROP TABLE IF EXISTS `plugins_manage`;
CREATE TABLE `plugins_manage`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `store_id` int NULL DEFAULT NULL COMMENT '插件在商城中的id',
  `p_name` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '插件名称英文',
  `p_title` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '插件名称',
  `p_description` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '插件介绍',
  `p_auth` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '作者',
  `is_install` tinyint NOT NULL DEFAULT 0 COMMENT '是否安装',
  `status` tinyint NOT NULL DEFAULT 0 COMMENT '状态',
  `version` varchar(60) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '当前版本',
  `price` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '价格',
  `download_times` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '下载次数',
  `install_path` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '安装路径（用于卸载）',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `storeIdUni`(`store_id` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '插件管理' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of plugins_manage
-- ----------------------------

-- ----------------------------
-- Table structure for recharge
-- ----------------------------
DROP TABLE IF EXISTS `recharge`;
CREATE TABLE `recharge`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'id主键',
  `main_chain` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '主链',
  `block_hash` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '块hash',
  `coin_token` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '币种1',
  `coin_token1` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '币种2',
  `from_address` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '发送方地址',
  `to_address` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '充币地址',
  `amount` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '充币数量',
  `amount1` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '充币2数量',
  `contract_address` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '币种1合约地址',
  `contract_address1` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '币种2合约地址',
  `hash` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '充币交易hashId',
  `block_height` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '区块高度',
  `call_number` smallint UNSIGNED NULL DEFAULT 0 COMMENT '回调次数',
  `status` smallint NULL DEFAULT 1 COMMENT '状态，1充币成功待归集，2归集上链中，3归集成功, 4归集失败，5充值费用中',
  `imputation_hash` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '归集hash',
  `remarks` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '备注',
  `recharge_type` tinyint UNSIGNED NOT NULL DEFAULT 1 COMMENT '充币类型',
  `token_id` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '充值的卡牌tokenId  充值卡牌时使用',
  `custome_coin` text COMMENT '自定义充值币种',
  `custome_user` text CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci COMMENT '自定义充值地址',
  `custome_amount` text COMMENT '自定义充值数量',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `hash_index`(`hash` ASC) USING BTREE,
  INDEX `to_address_index`(`to_address` ASC) USING BTREE,
  INDEX `from_address_index`(`from_address` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 58 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of recharge
-- ----------------------------

-- ----------------------------
-- Table structure for sys_auth_rule
-- ----------------------------
DROP TABLE IF EXISTS `sys_auth_rule`;
CREATE TABLE `sys_auth_rule`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `pid` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '父ID',
  `name` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '规则名称',
  `title` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '规则名称',
  `icon` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '图标',
  `condition` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '条件',
  `remark` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `menu_type` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '类型 0目录 1菜单 2按钮',
  `weigh` int NOT NULL DEFAULT 0 COMMENT '权重',
  `status` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态',
  `always_show` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '显示状态',
  `path` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '路由地址',
  `jump_path` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '跳转路由',
  `component` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '组件路径',
  `is_frame` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否外链 1是 0否',
  `module_type` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '所属模块',
  `model_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '模型ID',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建日期',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '修改日期',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除日期',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `name`(`name`) USING BTREE,
  INDEX `pid`(`pid`) USING BTREE,
  INDEX `weigh`(`weigh`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 331 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '菜单节点表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_auth_rule
-- ----------------------------
INSERT INTO `sys_auth_rule` VALUES (1, 0, 'system/config', '系统配置', 'system', '', 'Admin tips', 0, 0, 1, 1, 'config', '', '', 0, 'sys_admin', 0, NULL, NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (2, 0, 'system/auth', '权限管理', 'peoples', '', '', 0, 0, 1, 1, 'system/auth', '', '', 0, 'sys_admin', 0, NULL, NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (3, 0, 'system/monitor', '系统监控', 'monitor', '', '', 0, 0, 1, 1, 'monitor', '', '', 0, 'sys_admin', 0, NULL, NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (6, 1, 'system/config/sysConfig/list', '参数管理', 'date-range', '', '', 1, 1, 1, 1, 'params/list', '', 'system/config/params/list', 0, 'sys_admin', 0, NULL, '2022-04-11 04:29:47', NULL);
INSERT INTO `sys_auth_rule` VALUES (8, 2, 'system/auth/menuList', '菜单管理', 'nested', '', '', 1, 0, 1, 1, 'menuList', '', 'system/auth/menuList', 0, 'sys_admin', 0, NULL, '2021-07-20 09:01:49', NULL);
INSERT INTO `sys_auth_rule` VALUES (16, 8, 'system/auth/addMenu', '添加菜单', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, '2021-07-20 09:02:20', NULL);
INSERT INTO `sys_auth_rule` VALUES (17, 8, 'system/auth/editMenu', '修改菜单', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (18, 8, 'system/auth/deleteMenu', '删除菜单', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (26, 2, 'system/auth/roleList', '角色管理', 'logininfor', '', '', 1, 0, 1, 1, 'role', '', 'system/auth/roleList', 0, 'sys_admin', 0, NULL, NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (27, 26, 'system/auth/addRole', '添加角色', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (28, 2, 'system/auth/deptList', '部门管理', 'peoples', '', '', 1, 0, 1, 1, 'dept', '', 'system/auth/dept', 0, 'sys_admin', 0, NULL, '2021-07-20 09:03:46', NULL);
INSERT INTO `sys_auth_rule` VALUES (29, 26, 'system/auth/editRole', '修改角色', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (30, 26, 'system/auth/statusSetRole', '设置角色状态', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (31, 26, 'system/auth/deleteRole', '删除角色', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (33, 2, 'system/auth/postList', '岗位管理', 'tab', '', '', 1, 0, 1, 1, 'post', '', 'system/auth/post', 0, 'sys_admin', 0, NULL, '2021-07-20 09:04:51', NULL);
INSERT INTO `sys_auth_rule` VALUES (34, 3, 'system/monitor/online/list', '在线用户', 'cascader', '', '', 1, 0, 1, 1, 'online', '', 'system/monitor/online/list', 0, 'sys_admin', 0, NULL, NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (36, 3, 'system/monitor/job', '定时任务', 'clipboard', '', '', 1, 0, 1, 1, 'job', '', 'system/monitor/job', 0, 'sys_admin', 0, NULL, NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (37, 3, 'system/monitor/server/info', '服务监控', 'dict', '', '', 1, 0, 1, 1, 'server', '', 'system/monitor/server', 0, 'sys_admin', 0, NULL, '2021-07-19 16:07:23', NULL);
INSERT INTO `sys_auth_rule` VALUES (38, 3, 'system/monitor/loginLog', '登录日志', 'chart', '', '', 1, 0, 1, 1, 'logininfor', '', 'system/monitor/logininfor', 0, 'sys_admin', 0, NULL, '2021-07-20 09:08:00', NULL);
INSERT INTO `sys_auth_rule` VALUES (39, 3, 'system/monitor/operLog', '操作日志', 'dashboard', '', '', 1, 0, 1, 1, 'operlog', '', 'system/monitor/operlog', 0, 'sys_admin', 0, NULL, '2021-07-20 09:08:31', NULL);
INSERT INTO `sys_auth_rule` VALUES (40, 2, 'system/auth/userList', '用户管理', 'user', '', '', 1, 0, 1, 1, 'user', '', 'system/auth/userList', 0, 'sys_admin', 0, NULL, NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (41, 6, 'system/config/sysConfig/add', '添加参数', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, '2021-07-20 08:59:09', NULL);
INSERT INTO `sys_auth_rule` VALUES (42, 6, 'system/config/sysConfig/edit', '修改参数', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, '2021-07-20 08:59:17', NULL);
INSERT INTO `sys_auth_rule` VALUES (43, 6, 'system/config/sysConfig/delete', '删除参数', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, '2021-07-20 08:59:25', NULL);
INSERT INTO `sys_auth_rule` VALUES (44, 28, 'system/auth/deptAdd', '添加部门', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, '2021-07-20 09:04:05', NULL);
INSERT INTO `sys_auth_rule` VALUES (45, 28, 'system/auth/deptEdit', '修改部门', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, '2021-07-20 09:04:19', NULL);
INSERT INTO `sys_auth_rule` VALUES (46, 28, 'system/auth/deptDelete', '删除部门', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, '2021-07-20 09:04:35', NULL);
INSERT INTO `sys_auth_rule` VALUES (47, 33, 'system/auth/postAdd', '添加岗位', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, '2021-07-20 09:05:04', NULL);
INSERT INTO `sys_auth_rule` VALUES (48, 33, 'system/auth/postEdit', '修改岗位', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, '2021-07-20 09:05:15', NULL);
INSERT INTO `sys_auth_rule` VALUES (49, 33, 'system/auth/postDelete', '删除岗位', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, '2021-07-20 09:05:25', NULL);
INSERT INTO `sys_auth_rule` VALUES (50, 40, 'system/auth/addUser', '添加用户', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (51, 40, 'system/auth/editUser', '修改用户', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (52, 40, 'system/auth/resetUserPwd', '密码重置', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (53, 40, 'system/auth/changeUserStatus', '状态设置', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (54, 40, 'system/auth/deleteUser', '删除用户', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, '2021-07-20 09:06:16', NULL);
INSERT INTO `sys_auth_rule` VALUES (55, 34, 'system/monitor/online/forceLogout', '强制退出', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (56, 36, 'system/monitor/job/add', '添加任务', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (57, 36, 'system/monitor/job/edit', '修改任务', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (58, 36, 'system/monitor/job/start', '开启任务', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (59, 36, 'system/monitor/job/stop', '停止任务', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (60, 36, 'system/monitor/job/delete', '删除任务', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (61, 38, 'system/monitor/loginLog/delete', '删除', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, '2021-07-20 09:08:10', NULL);
INSERT INTO `sys_auth_rule` VALUES (62, 38, 'system/monitor/loginLog/clear', '清空', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, '2021-07-20 09:08:15', NULL);
INSERT INTO `sys_auth_rule` VALUES (63, 39, 'system/monitor/operLog/delete', '删除', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, '2021-07-20 09:08:36', NULL);
INSERT INTO `sys_auth_rule` VALUES (64, 39, 'system/monitor/operLog/clear', '清空', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, '2021-07-20 09:08:41', NULL);
INSERT INTO `sys_auth_rule` VALUES (300, 0, 'bscWallet', '币安钱包管理', 'chart', '', '', 0, 2, 1, 1, 'bscWallet', '', '', 0, 'sys_admin', 0, '2022-04-11 04:12:38', '2022-04-16 17:23:52', NULL);

INSERT INTO `sys_auth_rule` VALUES (301, 300, 'system/bscAddress/list', '地址管理', '', '', '', 1, 10, 1, 1, 'bscAddress', '', 'system/bscAddress/list', 0, 'sys_admin', 0, NULL, '2022-04-16 03:34:25', NULL);
INSERT INTO `sys_auth_rule` VALUES (302, 301, 'system/bscAddress/get', '查询', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, '2022-04-16 01:22:04', NULL);

INSERT INTO `sys_auth_rule` VALUES (303, 300, 'system/bscCurrency/list', '币种管理', '', '', '', 1, 9, 1, 1, 'bscCurrency', '', 'system/bscCurrency/list', 0, 'sys_admin', 0, NULL, '2022-04-16 03:34:29', NULL);
INSERT INTO `sys_auth_rule` VALUES (304, 303, 'system/bscCurrency/get', '查询', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, '2022-04-16 02:05:46', NULL);
INSERT INTO `sys_auth_rule` VALUES (305, 303, 'system/bscCurrency/add', '添加', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, '2022-04-16 02:05:53', NULL);
INSERT INTO `sys_auth_rule` VALUES (306, 303, 'system/bscCurrency/edit', '修改', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, '2022-04-16 02:05:58', NULL);
INSERT INTO `sys_auth_rule` VALUES (307, 303, 'system/bscCurrency/delete', '删除', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, '2022-04-16 02:06:04', NULL);

INSERT INTO `sys_auth_rule` VALUES (308, 300, 'system/bscFeeList/list', '手续费管理', '', '', '', 1, 8, 1, 1, 'bscFeeList', '', 'system/bscFeeList/list', 0, 'sys_admin', 0, NULL, '2022-04-16 03:37:40', NULL);
INSERT INTO `sys_auth_rule` VALUES (309, 308, 'system/bscFeeList/get', '查询', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, '2022-04-16 02:06:52', NULL);


INSERT INTO `sys_auth_rule` VALUES (310, 300, 'system/bscRecharge/list', '充值管理', '', '', '', 1, 7, 1, 1, 'bscRechargeList', '', 'system/bscRecharge/list', 0, 'sys_admin', 0, NULL, '2022-04-16 03:34:38', NULL);
INSERT INTO `sys_auth_rule` VALUES (311, 310, 'system/bscRecharge/get', '查询', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, '2022-04-16 02:08:25', NULL);
INSERT INTO `sys_auth_rule` VALUES (312, 310, 'system/bscRecharge/callback', '发起回调', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, '2022-04-16 03:41:55', NULL);


INSERT INTO `sys_auth_rule` VALUES (313, 300, 'system/bscWithdraw/list', '提现管理', '', '', '', 1, 6, 1, 1, 'bscWithdrawList', '', 'system/bscWithdraw/list', 0, 'sys_admin', 0, NULL, '2022-04-16 03:34:42', NULL);
INSERT INTO `sys_auth_rule` VALUES (314, 313, 'system/bscWithdraw/callback', '发起回调', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, '2022-04-16 03:45:50', '2022-04-16 03:45:50', NULL);
INSERT INTO `sys_auth_rule` VALUES (315, 313, 'system/bscWithdraw/auditSuccess', '审核通过', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, '2022-04-16 17:24:22', '2022-04-16 17:24:22', NULL);
INSERT INTO `sys_auth_rule` VALUES (316, 313, 'system/bscWithdraw/get', '查询', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, '2022-04-16 02:08:35', NULL);
INSERT INTO `sys_auth_rule` VALUES (317, 313, 'system/bscWithdraw/auditFail', '审核拒绝', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, '2022-04-16 17:24:49', '2022-04-16 17:24:49', NULL);

-- ----------------------------
-- Table structure for sys_config
-- ----------------------------
DROP TABLE IF EXISTS `sys_config`;
CREATE TABLE `sys_config`  (
  `config_id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '参数主键',
  `config_name` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '参数名称',
  `config_key` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '参数键名',
  `config_value` varchar(500) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '参数键值',
  `config_type` tinyint(1) NULL DEFAULT 0 COMMENT '系统内置（Y是 N否）',
  `create_by` int UNSIGNED NULL DEFAULT 0 COMMENT '创建者',
  `update_by` int UNSIGNED NULL DEFAULT 0 COMMENT '更新者',
  `remark` varchar(500) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '备注',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '修改时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`config_id`) USING BTREE,
  UNIQUE INDEX `uni_config_key`(`config_key`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 31 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_config
-- ----------------------------
INSERT INTO `sys_config` VALUES (1, '白名单IP', 'sys.whitelist', '127.0.0.1', 1, 1, 1, '白名单IP', NULL, '2022-08-23 15:26:31', NULL);
INSERT INTO `sys_config` VALUES (2, '充值回调地址', 'sys.rechargeCallbackUrl', '', 1, 1, 1, '充值回调地址', NULL, '2023-05-15 01:50:15', NULL);
INSERT INTO `sys_config` VALUES (3, '提币自动通过金额', 'sys.minWithdrawAudit', '0', 1, 1, 1, '提币自动通过金额', NULL, '2023-05-17 21:33:35', NULL);
INSERT INTO `sys_config` VALUES (4, '币安归集地址', 'sys.bnbMergeAddress', '', 0, 0, 1, '币安归集地址', NULL, '2023-04-11 14:39:46', NULL);
INSERT INTO `sys_config` VALUES (5, '币安链gasPrice', 'sys.bnbGasPrice', '5', 1, 1, 1, '币安链gasPrice', NULL, '2023-02-15 13:36:27', NULL);
INSERT INTO `sys_config` VALUES (6, '币安链gasLimit', 'sys.bnbGasLimit', '350000', 1, 1, 1, '币安链gasLimit', NULL, NULL, NULL);
INSERT INTO `sys_config` VALUES (7, '币安链提币地址', 'sys.bnbWithdrawAddress', '', 1, 1, 1, '币安链提币地址', NULL, '2023-05-15 20:36:45', NULL);
INSERT INTO `sys_config` VALUES (8, '币安链提币地址密钥', 'sys.bnbWithdrawAddressPrivateKey', '', 1, 1, 1, '币安链提币地址密钥', NULL, '2023-05-16 16:55:41', NULL);
INSERT INTO `sys_config` VALUES (9, '币安链手续费地址', 'sys.bnbFeeAddress', '', 1, 1, 1, '币安链手续费地址', NULL, '2023-04-11 14:44:28', NULL);
INSERT INTO `sys_config` VALUES (10, '币安链手续费地址密钥', 'sys.bnbFeeAddressPrivateKey', '', 1, 1, 1, '币安链手续费地址密钥', NULL, '2023-04-11 14:45:48', NULL);

-- ----------------------------
-- Table structure for sys_dept
-- ----------------------------
DROP TABLE IF EXISTS `sys_dept`;
CREATE TABLE `sys_dept`  (
  `dept_id` bigint NOT NULL AUTO_INCREMENT COMMENT '部门id',
  `parent_id` bigint NULL DEFAULT 0 COMMENT '父部门id',
  `ancestors` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '祖级列表',
  `dept_name` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '部门名称',
  `order_num` int NULL DEFAULT 0 COMMENT '显示顺序',
  `leader` varchar(20) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '负责人',
  `phone` varchar(11) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '联系电话',
  `email` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '邮箱',
  `status` char(1) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '0' COMMENT '部门状态（0正常 1停用）',
  `created_by` bigint UNSIGNED NULL DEFAULT 0 COMMENT '创建人',
  `updated_by` bigint NULL DEFAULT NULL COMMENT '修改人',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '修改时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`dept_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 101 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '部门表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_dept
-- ----------------------------
INSERT INTO `sys_dept` VALUES (100, 0, '0', 'CEO', 0, 'admin', '18000000000', 'wallet@qq.com', '1', 0, 31, '2021-07-13 15:56:52', '2021-07-13 15:57:05', NULL);

-- ----------------------------
-- Table structure for sys_dict_data
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_data`;
CREATE TABLE `sys_dict_data`  (
  `dict_code` bigint NOT NULL AUTO_INCREMENT COMMENT '字典编码',
  `dict_sort` int NULL DEFAULT 0 COMMENT '字典排序',
  `dict_label` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '字典标签',
  `dict_value` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '字典键值',
  `dict_type` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '字典类型',
  `css_class` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '样式属性（其他样式扩展）',
  `list_class` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '表格回显样式',
  `is_default` tinyint(1) NULL DEFAULT 0 COMMENT '是否默认（1是 0否）',
  `status` tinyint(1) NULL DEFAULT 0 COMMENT '状态（0正常 1停用）',
  `create_by` bigint UNSIGNED NULL DEFAULT 0 COMMENT '创建者',
  `update_by` bigint UNSIGNED NULL DEFAULT 0 COMMENT '更新者',
  `remark` varchar(500) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '备注',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '修改时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`dict_code`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 88 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '字典数据表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_dict_data
-- ----------------------------
INSERT INTO `sys_dict_data` VALUES (1, 0, '男', '1', 'sys_user_sex', '', '', 0, 1, 31, 2, '备注信息', NULL, NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (2, 0, '女', '2', 'sys_user_sex', '', '', 0, 1, 31, 31, '备注信息', NULL, NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (3, 0, '保密', '0', 'sys_user_sex', '', '', 1, 1, 31, 31, '备注信息', NULL, NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (28, 0, '正常', '0', 'sys_job_status', '', 'default', 1, 1, 31, 0, '', NULL, NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (29, 0, '暂停', '1', 'sys_job_status', '', 'default', 0, 1, 31, 31, '', NULL, NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (30, 0, '默认', 'DEFAULT', 'sys_job_group', '', 'default', 1, 1, 31, 0, '', NULL, NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (31, 0, '系统', 'SYSTEM', 'sys_job_group', '', 'default', 0, 1, 31, 0, '', NULL, NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (32, 0, '成功', '1', 'admin_login_status', '', 'default', 0, 1, 31, 31, '', NULL, NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (33, 0, '失败', '0', 'admin_login_status', '', 'default', 0, 1, 31, 0, '', NULL, NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (34, 0, '成功', '1', 'sys_oper_log_status', '', 'default', 0, 1, 31, 0, '', NULL, NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (35, 0, '失败', '0', 'sys_oper_log_status', '', 'default', 0, 1, 31, 0, '', NULL, NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (36, 0, '重复执行', '1', 'sys_job_policy', '', 'default', 1, 1, 31, 0, '', NULL, NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (37, 0, '执行一次', '2', 'sys_job_policy', '', 'default', 1, 1, 31, 0, '', NULL, NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (38, 0, '显示', '1', 'sys_show_hide', NULL, 'default', 1, 1, 31, 0, NULL, NULL, NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (39, 0, '隐藏', '0', 'sys_show_hide', NULL, 'default', 0, 1, 31, 0, NULL, NULL, NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (40, 0, '正常', '1', 'sys_normal_disable', '', 'default', 1, 1, 31, 0, '', NULL, NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (41, 0, '停用', '0', 'sys_normal_disable', '', 'default', 0, 1, 31, 0, '', NULL, NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (49, 0, '是', '1', 'sys_yes_no', '', '', 1, 1, 31, 0, '', NULL, NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (50, 0, '否', '0', 'sys_yes_no', '', '', 0, 1, 31, 0, '', NULL, NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (61, 0, '政府工作目标', '1', 'gov_cate_models', '', '', 0, 1, 2, 0, '', NULL, NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (62, 0, '系统后台', 'sys_admin', 'menu_module_type', '', '', 1, 1, 2, 0, '', NULL, NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (63, 0, '政务工作', 'gov_work', 'menu_module_type', '', '', 0, 1, 2, 0, '', NULL, NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (65, 0, '[work]测试业务表', 'wf_news', 'flow_type', '', '', 0, 1, 2, 2, '', NULL, NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (66, 0, '回退修改', '-1', 'flow_status', '', '', 0, 1, 31, 0, '', NULL, NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (67, 0, '保存中', '0', 'flow_status', '', '', 0, 1, 31, 0, '', NULL, NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (68, 0, '流程中', '1', 'flow_status', '', '', 0, 1, 31, 0, '', NULL, NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (69, 0, '审批通过', '2', 'flow_status', '', '', 0, 1, 31, 2, '', NULL, NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (70, 2, '发布栏目', '2', 'sys_blog_sign', '', '', 0, 1, 31, 31, '', NULL, NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (71, 3, '跳转栏目', '3', 'sys_blog_sign', '', '', 0, 1, 31, 31, '', NULL, NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (72, 4, '单页栏目', '4', 'sys_blog_sign', '', '', 0, 1, 31, 31, '', NULL, NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (73, 2, '置顶', '1', 'sys_log_sign', '', '', 0, 1, 31, 31, '', NULL, NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (74, 3, '幻灯', '2', 'sys_log_sign', '', '', 0, 1, 31, 31, '', NULL, NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (75, 4, '推荐', '3', 'sys_log_sign', '', '', 0, 1, 31, 31, '', NULL, NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (76, 1, '一般', '0', 'sys_log_sign', '', '', 0, 1, 31, 31, '', NULL, NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (77, 1, '频道页', '1', 'sys_blog_sign', '', '', 0, 1, 31, 31, '', NULL, NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (78, 0, '普通', '0', 'flow_level', '', '', 0, 1, 31, 0, '', NULL, '2021-07-20 08:55:20', NULL);
INSERT INTO `sys_dict_data` VALUES (79, 0, '加急', '1', 'flow_level', '', '', 0, 1, 31, 0, '', NULL, '2021-07-20 08:55:20', NULL);
INSERT INTO `sys_dict_data` VALUES (80, 0, '紧急', '2', 'flow_level', '', '', 0, 1, 31, 0, '', NULL, '2021-07-20 08:55:20', NULL);
INSERT INTO `sys_dict_data` VALUES (81, 0, '特急', '3', 'flow_level', '', '', 0, 1, 31, 31, '', NULL, '2021-07-20 08:55:25', NULL);
INSERT INTO `sys_dict_data` VALUES (82, 0, '频道页', '1', 'sys_blog_type', '', '', 0, 1, 31, 0, '', NULL, NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (83, 0, '发布栏目', '2', 'sys_blog_type', '', '', 0, 1, 31, 0, '', NULL, NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (84, 0, '跳转栏目', '3', 'sys_blog_type', '', '', 0, 1, 31, 31, '', NULL, NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (85, 0, '单页栏目', '4', 'sys_blog_type', '', '', 0, 1, 31, 0, '', NULL, NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (87, 0, '[cms]文章表', 'cms_news', 'flow_type', '', '', 0, 1, 31, 0, '', NULL, NULL, NULL);

-- ----------------------------
-- Table structure for sys_dict_type
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_type`;
CREATE TABLE `sys_dict_type`  (
  `dict_id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '字典主键',
  `dict_name` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '字典名称',
  `dict_type` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '字典类型',
  `status` tinyint UNSIGNED NULL DEFAULT 0 COMMENT '状态（0正常 1停用）',
  `create_by` int UNSIGNED NULL DEFAULT 0 COMMENT '创建者',
  `update_by` int UNSIGNED NULL DEFAULT 0 COMMENT '更新者',
  `remark` varchar(500) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '备注',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建日期',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '修改日期',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除日期',
  PRIMARY KEY (`dict_id`) USING BTREE,
  UNIQUE INDEX `dict_type`(`dict_type` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 36 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '字典类型表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_dict_type
-- ----------------------------
INSERT INTO `sys_dict_type` VALUES (1, '用户性别', 'sys_user_sex', 1, 31, 1, '用于选择用户性别', NULL, NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (3, '任务状态', 'sys_job_status', 1, 31, 31, '任务状态列表', NULL, NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (13, '任务分组', 'sys_job_group', 1, 31, 0, '', NULL, NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (14, '管理员登录状态', 'admin_login_status', 1, 31, 0, '', NULL, NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (15, '操作日志状态', 'sys_oper_log_status', 1, 31, 0, '', NULL, NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (16, '任务策略', 'sys_job_policy', 1, 31, 0, '', NULL, NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (17, '菜单状态', 'sys_show_hide', 1, 31, 0, '菜单状态', NULL, NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (18, '系统开关', 'sys_normal_disable', 1, 31, 31, '系统开关', NULL, NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (24, '系统内置', 'sys_yes_no', 1, 31, 0, '', NULL, NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (29, '政务工作模型分类', 'gov_cate_models', 1, 2, 0, '', NULL, NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (30, '菜单模块类型', 'menu_module_type', 1, 2, 0, '', NULL, NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (31, '工作流程类型', 'flow_type', 1, 2, 0, '', NULL, NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (32, '工作流程审批状态', 'flow_status', 1, 31, 0, '工作流程审批状态', NULL, NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (33, '博客分类类型', 'sys_blog_type', 1, 31, 31, '博客分类中的标志', NULL, NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (34, '博客日志标志', 'sys_log_sign', 1, 31, 0, '博客日志管理中的标志数据字典', NULL, NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (35, '工作流紧急状态', 'flow_level', 1, 31, 31, '', NULL, '2021-07-20 08:55:20', NULL);

-- ----------------------------
-- Table structure for sys_job
-- ----------------------------
DROP TABLE IF EXISTS `sys_job`;
CREATE TABLE `sys_job`  (
  `job_id` bigint NOT NULL AUTO_INCREMENT COMMENT '任务ID',
  `job_name` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '任务名称',
  `job_params` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '参数',
  `job_group` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT 'DEFAULT' COMMENT '任务组名',
  `invoke_target` varchar(500) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '调用目标字符串',
  `cron_expression` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT 'cron执行表达式',
  `misfire_policy` tinyint NULL DEFAULT 1 COMMENT '计划执行策略（1多次执行 2执行一次）',
  `concurrent` tinyint NULL DEFAULT 1 COMMENT '是否并发执行（0允许 1禁止）',
  `status` tinyint NULL DEFAULT 0 COMMENT '状态（0正常 1暂停）',
  `create_by` bigint UNSIGNED NULL DEFAULT 0 COMMENT '创建者',
  `update_by` bigint UNSIGNED NULL DEFAULT 0 COMMENT '更新者',
  `remark` varchar(500) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '备注信息',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`job_id`, `job_name`, `job_group`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '定时任务调度表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_job
-- ----------------------------
INSERT INTO `sys_job` VALUES (1, '币安扫描任务', '', '', 'bscSweepTask', '*/10 * * * * ?', 1, 1, 0, 0, 0, '', NULL, NULL, NULL);
INSERT INTO `sys_job` VALUES (4, '检查归集出金手续费状态', '', '', 'checkStatusTask', '*/30 * * * * ?', 1, 1, 0, 1, 0, '', NULL, NULL, NULL);
INSERT INTO `sys_job` VALUES (5, '出金任务', '', '', 'withdrawTask', '*/30 * * * * ?', 1, 1, 0, 1, 0, '', NULL, NULL, NULL);
INSERT INTO `sys_job` VALUES (6, '归集任务', '', '', 'rechargeTask', '*/30 * * * * ?', 1, 1, 1, 1, 0, '', NULL, NULL, '2023-05-15 02:06:58');
INSERT INTO `sys_job` VALUES (7, '检查用户登录状态', '', 'DEFAULT', 'checkUserOnline', '20 */1 * * * ?', 1, 1, 0, 1, 1, '', NULL, NULL, NULL);
INSERT INTO `sys_job` VALUES (9, '更新节点', '', 'DEFAULT', 'updateNode', '20 */1 * * * ?', 1, 1, 0, 0, 0, '', NULL, NULL, NULL);

-- ----------------------------
-- Table structure for sys_login_log
-- ----------------------------
DROP TABLE IF EXISTS `sys_login_log`;
CREATE TABLE `sys_login_log`  (
  `info_id` bigint NOT NULL AUTO_INCREMENT COMMENT '访问ID',
  `login_name` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '登录账号',
  `ipaddr` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '登录IP地址',
  `login_location` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '登录地点',
  `browser` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '浏览器类型',
  `os` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '操作系统',
  `status` tinyint NULL DEFAULT 0 COMMENT '登录状态（0成功 1失败）',
  `msg` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '提示消息',
  `login_time` datetime NULL DEFAULT NULL COMMENT '登录时间',
  `module` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '登录模块',
  PRIMARY KEY (`info_id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 12 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '系统访问记录' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_login_log
-- ----------------------------

-- ----------------------------
-- Table structure for sys_model_info
-- ----------------------------
DROP TABLE IF EXISTS `sys_model_info`;
CREATE TABLE `sys_model_info`  (
  `model_id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '模型ID',
  `model_category_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '模板分类id',
  `model_name` char(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '模型标识',
  `model_title` char(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '模型名称',
  `model_pk` char(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '主键字段',
  `model_order` char(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '默认排序字段',
  `model_sort` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '表单字段排序',
  `model_list` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '列表显示字段，为空显示全部',
  `model_edit` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '可编辑字段，为空则除主键外均可以编辑',
  `model_indexes` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '索引字段',
  `search_list` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '高级搜索的字段',
  `create_time` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `update_time` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
  `model_status` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态',
  `model_engine` varchar(25) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT 'MyISAM' COMMENT '数据库引擎',
  `create_by` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建人',
  `update_by` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '修改人',
  PRIMARY KEY (`model_id`) USING BTREE,
  UNIQUE INDEX `name_uni`(`model_name` ASC) USING BTREE COMMENT '模型名唯一'
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '文档模型表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_model_info
-- ----------------------------

-- ----------------------------
-- Table structure for sys_oper_log
-- ----------------------------
DROP TABLE IF EXISTS `sys_oper_log`;
CREATE TABLE `sys_oper_log`  (
  `oper_id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '日志主键',
  `title` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '模块标题',
  `business_type` int NULL DEFAULT 0 COMMENT '业务类型（0其它 1新增 2修改 3删除）',
  `method` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '方法名称',
  `request_method` varchar(10) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '请求方式',
  `operator_type` int NULL DEFAULT 0 COMMENT '操作类别（0其它 1后台用户 2手机端用户）',
  `oper_name` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '操作人员',
  `dept_name` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '部门名称',
  `oper_url` varchar(500) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '请求URL',
  `oper_ip` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '主机地址',
  `oper_location` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '操作地点',
  `oper_param` text CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL COMMENT '请求参数',
  `json_result` text CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL COMMENT '返回参数',
  `status` int NULL DEFAULT 0 COMMENT '操作状态（0正常 1异常）',
  `error_msg` varchar(2000) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '错误消息',
  `oper_time` datetime NULL DEFAULT NULL COMMENT '操作时间',
  PRIMARY KEY (`oper_id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 220 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '操作日志记录' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_oper_log
-- ----------------------------

-- ----------------------------
-- Table structure for sys_post
-- ----------------------------
DROP TABLE IF EXISTS `sys_post`;
CREATE TABLE `sys_post`  (
  `post_id` bigint NOT NULL AUTO_INCREMENT COMMENT '岗位ID',
  `post_code` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '岗位编码',
  `post_name` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '岗位名称',
  `post_sort` int NOT NULL COMMENT '显示顺序',
  `status` char(1) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '状态（0正常 1停用）',
  `remark` varchar(500) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '备注',
  `created_by` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建人',
  `updated_by` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '修改人',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '修改时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`post_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '岗位信息表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_post
-- ----------------------------
INSERT INTO `sys_post` VALUES (1, 'ceo', '超级管理员', 1, '1', '', 0, 0, '2021-07-11 11:32:58', NULL, NULL);

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `status` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态;0:禁用;1:正常',
  `list_order` float NOT NULL DEFAULT 0 COMMENT '排序',
  `name` varchar(20) CHARACTER SET utf8mb3 COLLATE utf8mb3_unicode_ci NOT NULL DEFAULT '' COMMENT '角色名称',
  `remark` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `data_scope` tinyint UNSIGNED NOT NULL DEFAULT 3 COMMENT '数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `status`(`status` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '角色表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_role
-- ----------------------------
INSERT INTO `sys_role` VALUES (1, 1, 0, '超级管理员', '备注', 3);
INSERT INTO `sys_role` VALUES (2, 0, 0, '观察者', '', 3);

-- ----------------------------
-- Table structure for sys_role_dept
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_dept`;
CREATE TABLE `sys_role_dept`  (
  `role_id` bigint NOT NULL COMMENT '角色ID',
  `dept_id` bigint NOT NULL COMMENT '部门ID',
  PRIMARY KEY (`role_id`, `dept_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '角色和部门关联表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_role_dept
-- ----------------------------

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_name` varchar(60) CHARACTER SET utf8mb3 COLLATE utf8mb3_unicode_ci NOT NULL DEFAULT '' COMMENT '用户名',
  `mobile` varchar(20) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '中国手机不带国家代码，国际手机号格式为：国家代码-手机号',
  `user_nickname` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_unicode_ci NOT NULL DEFAULT '' COMMENT '用户昵称',
  `birthday` int NOT NULL DEFAULT 0 COMMENT '生日',
  `user_password` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '登录密码;cmf_password加密',
  `user_salt` char(10) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '加密盐',
  `user_status` tinyint UNSIGNED NOT NULL DEFAULT 1 COMMENT '用户状态;0:禁用,1:正常,2:未验证',
  `user_email` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '用户登录邮箱',
  `sex` tinyint NOT NULL DEFAULT 0 COMMENT '性别;0:保密,1:男,2:女',
  `avatar` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '用户头像',
  `dept_id` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '部门id',
  `remark` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `is_admin` tinyint NOT NULL DEFAULT 1 COMMENT '是否后台管理员 1 是  0   否',
  `address` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '联系地址',
  `describe` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT ' 描述信息',
  `phone_num` varchar(20) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '联系电话',
  `last_login_ip` varchar(15) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '最后登录ip',
  `last_login_time` datetime NULL DEFAULT NULL COMMENT '最后登录时间',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `user_login`(`user_name` ASC) USING BTREE,
  UNIQUE INDEX `mobile`(`mobile` ASC) USING BTREE,
  INDEX `user_nickname`(`user_nickname` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '用户表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_user
-- ----------------------------
INSERT INTO `sys_user` VALUES (1, '12tK4ofQyhfu1bt', '18800000000', '超级管理员', 0, '52e64d9fd45edb45e36a04f17971b5da', 'bynrGbgkiK', 1, 'admin@qq.com', 1, 'pub_upload/2022-04-11/cj6sdjf021iw6samde.jpeg', 100, '', 1, 'asdasfdsaf大发放打发士大夫发按时', '描述信息', '18611111111', '113.90.13.247', '2023-05-17 21:27:22', '2021-06-22 17:58:00', '2022-04-11 03:10:21', NULL);

-- ----------------------------
-- Table structure for sys_user_online
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_online`;
CREATE TABLE `sys_user_online`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `uuid` char(32) CHARACTER SET latin1 COLLATE latin1_general_ci NOT NULL DEFAULT '' COMMENT '用户标识',
  `token` varchar(255) CHARACTER SET latin1 COLLATE latin1_general_ci NOT NULL DEFAULT '' COMMENT '用户token',
  `create_time` datetime NULL DEFAULT NULL COMMENT '登录时间',
  `user_name` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '用户名',
  `ip` varchar(120) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '登录ip',
  `explorer` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '浏览器',
  `os` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '操作系统',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uni_token`(`token`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 12 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '用户在线状态表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_user_online
-- ----------------------------

-- ----------------------------
-- Table structure for sys_user_post
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_post`;
CREATE TABLE `sys_user_post`  (
  `user_id` bigint NOT NULL COMMENT '用户ID',
  `post_id` bigint NOT NULL COMMENT '岗位ID',
  PRIMARY KEY (`user_id`, `post_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '用户与岗位关联表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_user_post
-- ----------------------------

-- ----------------------------
-- Table structure for sys_web_set
-- ----------------------------
DROP TABLE IF EXISTS `sys_web_set`;
CREATE TABLE `sys_web_set`  (
  `web_id` int NOT NULL AUTO_INCREMENT COMMENT '主键',
  `web_content` text CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL COMMENT '站点信息',
  PRIMARY KEY (`web_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_web_set
-- ----------------------------
INSERT INTO `sys_web_set` VALUES (1, '{\"CopyrightInfo\":\"json\",\"recordInfo\":\"111222\",\"statisticsCode\":\"11122\",\"webId\":1,\"webLogo\":\"https://yxh-1301841944.cos.ap-chongqing.myqcloud.com/gfast/2021-07-13/ccroz2q3sptczqwchk.jpg\",\"webName\":\"钱包管理系统\",\"webSite\":\"http://localhost/index#/webSet\"}');

-- ----------------------------
-- Table structure for tools_gen_table
-- ----------------------------
DROP TABLE IF EXISTS `tools_gen_table`;
CREATE TABLE `tools_gen_table`  (
  `table_id` bigint NOT NULL AUTO_INCREMENT COMMENT '编号',
  `table_name` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '表名称',
  `table_comment` varchar(500) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '表描述',
  `class_name` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '实体类名称',
  `tpl_category` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT 'crud' COMMENT '使用的模板（crud单表操作 tree树表操作）',
  `package_name` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '生成包路径',
  `module_name` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '生成模块名',
  `business_name` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '生成业务名',
  `function_name` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '生成功能名',
  `function_author` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '生成功能作者',
  `options` varchar(1000) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '其它生成选项',
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`table_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 76 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '代码生成业务表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tools_gen_table
-- ----------------------------
INSERT INTO `tools_gen_table` VALUES (70, 'address', '', 'Address', 'crud', 'gfast/app/system', 'system', 'address', '', 'gfast', '', '2022-04-11 03:32:45', '2022-04-11 03:32:45', '');
INSERT INTO `tools_gen_table` VALUES (71, 'currency', '', 'Currency', 'crud', 'gfast/app/system', 'system', 'currency', '', 'gfast', '', '2022-04-11 04:18:33', '2022-04-11 04:18:33', '');
INSERT INTO `tools_gen_table` VALUES (72, 'fee_list', '', 'FeeList', 'crud', 'gfast/app/system', 'system', 'fee_list', '', 'gfast', '', '2022-04-11 04:18:33', '2022-04-11 04:18:33', '');
INSERT INTO `tools_gen_table` VALUES (73, 'height', '', 'Height', 'crud', 'gfast/app/system', 'system', 'height', '', 'gfast', '', '2022-04-11 04:18:33', '2022-04-11 04:18:33', '');
INSERT INTO `tools_gen_table` VALUES (74, 'recharge', '', 'Recharge', 'crud', 'gfast/app/system', 'system', 'recharge', '', 'gfast', '', '2022-04-11 04:18:33', '2022-04-11 04:18:33', '');
INSERT INTO `tools_gen_table` VALUES (75, 'withdraw', '', 'Withdraw', 'crud', 'gfast/app/system', 'system', 'withdraw', '', 'gfast', '', '2022-04-11 04:18:34', '2022-04-11 04:18:34', '');

-- ----------------------------
-- Table structure for tools_gen_table_column
-- ----------------------------
DROP TABLE IF EXISTS `tools_gen_table_column`;
CREATE TABLE `tools_gen_table_column`  (
  `column_id` bigint NOT NULL AUTO_INCREMENT COMMENT '编号',
  `table_id` bigint NULL DEFAULT NULL COMMENT '归属表编号',
  `column_name` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '列名称',
  `column_comment` varchar(500) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '列描述',
  `column_type` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '列类型',
  `go_type` varchar(500) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT 'Go类型',
  `go_field` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT 'Go字段名',
  `html_field` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT 'html字段名',
  `is_pk` char(1) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '是否主键（1是）',
  `is_increment` char(1) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '是否自增（1是）',
  `is_required` char(1) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '是否必填（1是）',
  `is_insert` char(1) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '是否为插入字段（1是）',
  `is_edit` char(1) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '是否编辑字段（1是）',
  `is_list` char(1) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '是否列表字段（1是）',
  `is_query` char(1) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '是否查询字段（1是）',
  `query_type` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT 'EQ' COMMENT '查询方式（等于、不等于、大于、小于、范围）',
  `html_type` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '显示类型（文本框、文本域、下拉框、复选框、单选框、日期控件）',
  `dict_type` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '字典类型',
  `sort` int NULL DEFAULT NULL COMMENT '排序',
  `link_table_name` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '关联表名',
  `link_table_class` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '关联表类名',
  `link_table_package` varchar(150) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '关联表包名',
  `link_label_id` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '关联表键名',
  `link_label_name` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '关联表字段值',
  PRIMARY KEY (`column_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 781 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '代码生成业务表字段' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tools_gen_table_column
-- ----------------------------
INSERT INTO `tools_gen_table_column` VALUES (720, 70, 'id', '', 'int unsigned', '', 'Id', 'id', '1', '1', '0', '0', '0', '0', '0', 'EQ', '', '', 1, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (721, 70, 'user_id', '用户标识', 'int', 'int', 'UserId', 'userId', '0', '0', '1', '1', '1', '1', '1', 'EQ', 'input', '', 2, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (722, 70, 'main_chain', '主链', 'varchar(50)', 'string', 'MainChain', 'mainChain', '0', '0', '1', '1', '1', '1', '1', 'EQ', 'input', '', 3, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (723, 70, 'address', '地址', 'varchar(100)', 'string', 'Address', 'address', '0', '0', '1', '1', '1', '1', '1', 'EQ', 'input', '', 4, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (724, 70, 'private_key', '私钥', 'varchar(255)', 'string', 'PrivateKey', 'privateKey', '0', '0', '1', '1', '1', '1', '1', 'EQ', 'input', '', 5, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (725, 70, 'created_at', '', 'timestamp', 'Time', 'CreatedAt', 'createdAt', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 6, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (726, 70, 'updated_at', '', 'timestamp', 'Time', 'UpdatedAt', 'updatedAt', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 7, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (727, 71, 'id', '', 'int unsigned', '', 'Id', 'id', '1', '1', '0', '0', '0', '0', '0', 'EQ', '', '', 1, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (728, 71, 'main_chain', '主链', 'varchar(50)', 'string', 'MainChain', 'mainChain', '0', '0', '1', '1', '1', '1', '1', 'EQ', 'input', '', 2, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (729, 71, 'name', '币种名称', 'varchar(50)', 'string', 'Name', 'name', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 3, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (730, 71, 'contract_address', '合约地址', 'varchar(255)', 'string', 'ContractAddress', 'contractAddress', '0', '0', '', '1', '1', '1', '1', 'EQ', 'input', '', 4, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (731, 71, 'decimals', '精度', 'int', 'int', 'Decimals', 'decimals', '0', '0', '1', '1', '1', '1', '1', 'EQ', 'input', '', 5, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (732, 71, 'created_at', '', 'timestamp', 'Time', 'CreatedAt', 'createdAt', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 6, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (733, 71, 'updated_at', '', 'timestamp', 'Time', 'UpdatedAt', 'updatedAt', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 7, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (734, 72, 'id', '', 'int unsigned', '', 'Id', 'id', '1', '1', '0', '0', '0', '0', '0', 'EQ', '', '', 1, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (735, 72, 'main_chain', '主链', 'varchar(50)', 'string', 'MainChain', 'mainChain', '0', '0', '1', '1', '1', '1', '1', 'EQ', 'input', '', 2, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (736, 72, 'coin_name', '手续费币种', 'varchar(50)', 'string', 'CoinName', 'coinName', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 3, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (737, 72, 'address', '地址', 'varchar(255)', 'string', 'Address', 'address', '0', '0', '1', '1', '1', '1', '1', 'EQ', 'input', '', 4, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (738, 72, 'amount', '转移手续费', 'decimal(30,18)', 'float64', 'Amount', 'amount', '0', '0', '1', '1', '1', '1', '1', 'EQ', 'input', '', 5, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (739, 72, 'hash', 'hash', 'varchar(255)', 'string', 'Hash', 'hash', '0', '0', '1', '1', '1', '1', '1', 'EQ', 'input', '', 6, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (740, 72, 'status', '状态 1-待上链 2-上链成功 3-上链失败', 'tinyint unsigned', '', 'Status', 'status', '0', '0', '1', '1', '1', '1', '1', 'EQ', 'radio', '', 7, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (741, 72, 'recharge_id', '充值ID', 'int unsigned', '', 'RechargeId', 'rechargeId', '0', '0', '1', '1', '1', '1', '1', 'EQ', '', '', 8, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (742, 72, 'created_at', '', 'timestamp', 'Time', 'CreatedAt', 'createdAt', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 9, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (743, 72, 'updated_at', '', 'timestamp', 'Time', 'UpdatedAt', 'updatedAt', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 10, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (744, 73, 'id', '', 'int unsigned', '', 'Id', 'id', '1', '1', '0', '0', '0', '0', '0', 'EQ', '', '', 1, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (745, 73, 'main_chain', '主网', 'varchar(50)', 'string', 'MainChain', 'mainChain', '0', '0', '1', '1', '1', '1', '1', 'EQ', 'input', '', 2, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (746, 73, 'height', '块号', 'int', 'int', 'Height', 'height', '0', '0', '1', '1', '1', '1', '1', 'EQ', 'input', '', 3, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (747, 74, 'id', 'id主键', 'bigint', 'int64', 'Id', 'id', '1', '1', '0', '0', '0', '0', '0', 'EQ', 'input', '', 1, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (748, 74, 'main_chain', '主链', 'varchar(50)', 'string', 'MainChain', 'mainChain', '0', '0', '', '1', '1', '1', '1', 'EQ', 'input', '', 2, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (749, 74, 'block_hash', '块hash', 'varchar(255)', 'string', 'BlockHash', 'blockHash', '0', '0', '', '1', '1', '1', '1', 'EQ', 'input', '', 3, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (750, 74, 'coin_token', '币种1', 'varchar(50)', 'string', 'CoinToken', 'coinToken', '0', '0', '', '1', '1', '1', '1', 'EQ', 'input', '', 4, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (751, 74, 'coin_token1', '币种2', 'varchar(50)', 'string', 'CoinToken1', 'coinToken1', '0', '0', '', '1', '1', '1', '1', 'EQ', 'input', '', 5, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (752, 74, 'from_address', '发送方地址', 'varchar(255)', 'string', 'FromAddress', 'fromAddress', '0', '0', '', '1', '1', '1', '1', 'EQ', 'input', '', 6, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (753, 74, 'to_address', '充币地址', 'varchar(255)', 'string', 'ToAddress', 'toAddress', '0', '0', '', '1', '1', '1', '1', 'EQ', 'input', '', 7, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (754, 74, 'amount', '充币数量', 'decimal(30,18)', 'float64', 'Amount', 'amount', '0', '0', '', '1', '1', '1', '1', 'EQ', 'input', '', 8, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (755, 74, 'amount1', '充币2数量', 'decimal(30,18) unsigned', 'float64', 'Amount1', 'amount1', '0', '0', '', '1', '1', '1', '1', 'EQ', 'input', '', 9, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (756, 74, 'contract_address', '币种1合约地址', 'varchar(255)', 'string', 'ContractAddress', 'contractAddress', '0', '0', '', '1', '1', '1', '1', 'EQ', 'input', '', 10, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (757, 74, 'contract_address1', '币种2合约地址', 'varchar(255)', 'string', 'ContractAddress1', 'contractAddress1', '0', '0', '', '1', '1', '1', '1', 'EQ', 'input', '', 11, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (758, 74, 'hash', '充币交易hashId', 'varchar(255)', 'string', 'Hash', 'hash', '0', '0', '1', '1', '1', '1', '1', 'EQ', 'input', '', 12, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (759, 74, 'block_height', '区块高度', 'varchar(255)', 'string', 'BlockHeight', 'blockHeight', '0', '0', '', '1', '1', '1', '1', 'EQ', 'input', '', 13, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (760, 74, 'call_number', '回调次数', 'smallint unsigned', '', 'CallNumber', 'callNumber', '0', '0', '', '1', '1', '1', '1', 'EQ', '', '', 14, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (761, 74, 'status', '状态，1充币成功待归集，2归集上链中，3归集成功, 4归集失败，5充值费用中', 'smallint', 'int', 'Status', 'status', '0', '0', '1', '1', '1', '1', '1', 'EQ', 'radio', '', 15, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (762, 74, 'imputation_hash', '归集hash', 'varchar(255)', 'string', 'ImputationHash', 'imputationHash', '0', '0', '', '1', '1', '1', '1', 'EQ', 'input', '', 16, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (763, 74, 'remarks', '备注', 'varchar(255)', 'string', 'Remarks', 'remarks', '0', '0', '', '1', '1', '1', '1', 'EQ', 'input', '', 17, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (764, 74, 'recharge_type', '充币数量 ', 'tinyint unsigned', '', 'RechargeType', 'rechargeType', '0', '0', '1', '1', '1', '1', '1', 'EQ', 'select', '', 18, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (765, 74, 'created_at', '', 'timestamp', 'Time', 'CreatedAt', 'createdAt', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 19, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (766, 74, 'updated_at', '', 'timestamp', 'Time', 'UpdatedAt', 'updatedAt', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 20, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (767, 75, 'id', '', 'int unsigned', '', 'Id', 'id', '1', '1', '0', '0', '0', '0', '0', 'EQ', '', '', 1, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (768, 75, 'main_chain', '主链', 'varchar(50)', 'string', 'MainChain', 'mainChain', '0', '0', '1', '1', '1', '1', '1', 'EQ', 'input', '', 2, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (769, 75, 'coin_token', '币种 ', 'varchar(100)', 'string', 'CoinToken', 'coinToken', '0', '0', '1', '1', '1', '1', '1', 'EQ', 'input', '', 3, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (770, 75, 'address', '转出地址', 'varchar(255)', 'string', 'Address', 'address', '0', '0', '1', '1', '1', '1', '1', 'EQ', 'input', '', 4, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (771, 75, 'amount', '提币数量', 'decimal(30,18)', 'float64', 'Amount', 'amount', '0', '0', '1', '1', '1', '1', '1', 'EQ', 'input', '', 5, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (772, 75, 'contract_address', '合约地址', 'varchar(255)', 'string', 'ContractAddress', 'contractAddress', '0', '0', '', '1', '1', '1', '1', 'EQ', 'input', '', 6, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (773, 75, 'status', '状态 1待审核，2待上链，3上链中，4上链失败，5上链成功，6已拒绝', 'tinyint unsigned', '', 'Status', 'status', '0', '0', '1', '1', '1', '1', '1', 'EQ', 'radio', '', 7, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (774, 75, 'nonce', 'nonce', 'bigint unsigned', '', 'Nonce', 'nonce', '0', '0', '1', '1', '1', '1', '1', 'EQ', '', '', 8, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (775, 75, 'confirmation_number', '上链确认次数', 'smallint unsigned', '', 'ConfirmationNumber', 'confirmationNumber', '0', '0', '1', '1', '1', '1', '1', 'EQ', '', '', 9, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (776, 75, 'hash', '交易Hash', 'varchar(255)', 'string', 'Hash', 'hash', '0', '0', '', '1', '1', '1', '1', 'EQ', 'input', '', 10, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (777, 75, 'remarks', '交易备注', 'varchar(255)', 'string', 'Remarks', 'remarks', '0', '0', '', '1', '1', '1', '1', 'EQ', 'input', '', 11, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (778, 75, 'notify_url', '交易回调地址', 'varchar(255)', 'string', 'NotifyUrl', 'notifyUrl', '0', '0', '', '1', '1', '1', '1', 'EQ', 'input', '', 12, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (779, 75, 'created_at', '', 'timestamp', 'Time', 'CreatedAt', 'createdAt', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 13, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (780, 75, 'updated_at', '', 'timestamp', 'Time', 'UpdatedAt', 'updatedAt', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 14, '', '', '', '', '');

-- ----------------------------
-- Table structure for withdraw
-- ----------------------------
DROP TABLE IF EXISTS `withdraw`;
CREATE TABLE `withdraw`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `main_chain` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '主链',
  `coin_token` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '币种 ',
  `withdraw_address` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '转出地址',
  `address` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '转入地址',
  `amount` decimal(30, 18) NOT NULL COMMENT '提币数量',
  `contract_address` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '合约地址',
  `status` tinyint UNSIGNED NOT NULL DEFAULT 1 COMMENT '状态 1待审核，2待上链，3上链中，4上链失败，5上链成功，6已拒绝',
  `nonce` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT 'nonce',
  `confirmation_number` smallint UNSIGNED NOT NULL DEFAULT 0 COMMENT '上链确认次数',
  `nonce1` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT 'nonce1',
  `hashKey` text CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT 'hashKey',
  `hash` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '交易Hash',
  `remarks` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '交易备注',
  `notify_url` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '交易回调地址',
  `token_id` bigint NULL DEFAULT NULL COMMENT 'tokenID (nft使用)',
  `url` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT 'url (nft使用)',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of withdraw
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;
