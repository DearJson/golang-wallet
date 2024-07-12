/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50736
 Source Host           : localhost:3306
 Source Schema         : gfast-v2-os

 Target Server Type    : MySQL
 Target Server Version : 50736
 File Encoding         : 65001

 Date: 03/03/2022 12:08:05
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule`  (
  `ptype` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v0` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v1` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v2` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v3` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v4` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v5` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for demo_data_auth
-- ----------------------------
DROP TABLE IF EXISTS `demo_data_auth`;
CREATE TABLE `demo_data_auth`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '标题',
  `created_by` int(10) UNSIGNED NULL DEFAULT 0 COMMENT '创建人',
  `updated_by` int(10) UNSIGNED NULL DEFAULT 0 COMMENT '修改人',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(0) NULL DEFAULT NULL COMMENT '修改时间',
  `deleted_at` datetime(0) NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for demo_gen
-- ----------------------------
DROP TABLE IF EXISTS `demo_gen`;
CREATE TABLE `demo_gen`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `demo_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '姓名',
  `demo_age` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '年龄',
  `classes` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '班级',
  `demo_born` datetime(0) NULL DEFAULT NULL COMMENT '出生年月',
  `demo_gender` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '性别',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '创建日期',
  `updated_at` datetime(0) NULL DEFAULT NULL COMMENT '修改日期',
  `deleted_at` datetime(0) NULL DEFAULT NULL COMMENT '删除日期',
  `created_by` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建人',
  `updated_by` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '修改人',
  `demo_status` tinyint(4) NOT NULL DEFAULT 0 COMMENT '状态',
  `demo_cate` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '分类',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '代码生成测试表' ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for demo_gen_class
-- ----------------------------
DROP TABLE IF EXISTS `demo_gen_class`;
CREATE TABLE `demo_gen_class`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '分类id',
  `class_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '分类名',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '代码生成关联测试表' ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for demo_gen_tree
-- ----------------------------
DROP TABLE IF EXISTS `demo_gen_tree`;
CREATE TABLE `demo_gen_tree`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `parent_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '父级ID',
  `demo_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '姓名',
  `demo_age` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '年龄',
  `classes` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '班级',
  `demo_born` datetime(0) NULL DEFAULT NULL COMMENT '出生年月',
  `demo_gender` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '性别',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '创建日期',
  `updated_at` datetime(0) NULL DEFAULT NULL COMMENT '修改日期',
  `deleted_at` datetime(0) NULL DEFAULT NULL COMMENT '删除日期',
  `created_by` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建人',
  `updated_by` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '修改人',
  `demo_status` tinyint(4) NOT NULL DEFAULT 0 COMMENT '状态',
  `demo_cate` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '分类',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '代码生成树形结构测试表' ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for plugins_manage
-- ----------------------------
DROP TABLE IF EXISTS `plugins_manage`;
CREATE TABLE `plugins_manage`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `store_id` int(11) NULL DEFAULT NULL COMMENT '插件在商城中的id',
  `p_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '插件名称英文',
  `p_title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '插件名称',
  `p_description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '插件介绍',
  `p_auth` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '作者',
  `is_install` tinyint(4) NOT NULL DEFAULT 0 COMMENT '是否安装',
  `status` tinyint(4) NOT NULL DEFAULT 0 COMMENT '状态',
  `version` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '当前版本',
  `price` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '价格',
  `download_times` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '下载次数',
  `install_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '安装路径（用于卸载）',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `storeIdUni`(`store_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 45 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '插件管理' ROW_FORMAT = Compact;

-- ----------------------------
-- Records of plugins_manage
-- ----------------------------
INSERT INTO `plugins_manage` VALUES (34, 1, 'mail', '邮件发送', '邮件发送', 'demo01', 1, 1, '1.0.0', 1, 1248, 'plugins');
INSERT INTO `plugins_manage` VALUES (35, 2, 'wechat', '微信公众号', '微信小程序管理接口，包含关注回复、关键字回复、图文消息、素材管理、菜单管理等', 'demo01', 0, 0, '1.0.2', 1, 456, NULL);
INSERT INTO `plugins_manage` VALUES (36, 3, 'cms', 'cms管理系统', 'cms管理系统', 'gfast01', 0, 0, '1.0.3', 9900, 278, NULL);
INSERT INTO `plugins_manage` VALUES (37, 4, 'demo2', '免费测试', '666', 'demo01', 0, 0, '1.2', 0, 2, NULL);
INSERT INTO `plugins_manage` VALUES (38, 5, 'mail', '邮件发送', 'qq smtp邮件发送功能', 'wilgx0', 0, 0, '1.0.1', 199, 42, NULL);
INSERT INTO `plugins_manage` VALUES (39, 6, 'wechat', '微信公众号', '微信小程序管理接口，包含关注回复、关键字回复、图文消息、素材管理、菜单管理等', 'gfast01', 0, 0, '1.0.1', 5900, 22, NULL);
INSERT INTO `plugins_manage` VALUES (40, 7, 'gencode', '代码生成', '代码生成功能完善，支持生成上传文件、图片及富文本编辑器功能', 'gfast01', 0, 0, '1.0.2', 900, 68, NULL);
INSERT INTO `plugins_manage` VALUES (41, 8, 'im', '在线客服', '在线客服', 'wilgx0', 0, 0, '1.0.1', 9900, 4, NULL);
INSERT INTO `plugins_manage` VALUES (42, 9, 'doorkeeper', 'IP过滤', 'IP地址过滤，支持ipv4 ipv6 ,可设置访问ip白名单和黑名单', 'wilgx0', 0, 0, '1.0.1', 199, 15, NULL);
INSERT INTO `plugins_manage` VALUES (43, 10, 'smsPlugin', '短信发送', '短信发送，支持网易云信、腾讯云、阿里云短信平台，可自定义配置。', 'wilgx0', 0, 0, '1.0.1', 1999, 2, NULL);
INSERT INTO `plugins_manage` VALUES (44, 11, 'wordFlow', '流程审批', '流程审批插件', 'gfast01', 0, 0, '1.0.1', 29900, 0, NULL);

-- ----------------------------
-- Table structure for sys_auth_rule
-- ----------------------------
DROP TABLE IF EXISTS `sys_auth_rule`;
CREATE TABLE `sys_auth_rule`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `pid` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '父ID',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '规则名称',
  `title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '规则名称',
  `icon` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '图标',
  `condition` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '条件',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `menu_type` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '类型 0目录 1菜单 2按钮',
  `weigh` int(10) NOT NULL DEFAULT 0 COMMENT '权重',
  `status` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态',
  `always_show` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '显示状态',
  `path` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '路由地址',
  `jump_path` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '跳转路由',
  `component` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '组件路径',
  `is_frame` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否外链 1是 0否',
  `module_type` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '所属模块',
  `model_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '模型ID',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '创建日期',
  `updated_at` datetime(0) NULL DEFAULT NULL COMMENT '修改日期',
  `deleted_at` datetime(0) NULL DEFAULT NULL COMMENT '删除日期',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `name`(`name`) USING BTREE,
  INDEX `pid`(`pid`) USING BTREE,
  INDEX `weigh`(`weigh`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 235 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '菜单节点表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_auth_rule
-- ----------------------------
INSERT INTO `sys_auth_rule` VALUES (1, 0, 'system/config', '系统配置', 'system', '', 'Admin tips', 0, 0, 1, 1, 'config', '', '', 0, 'sys_admin', 0, NULL, NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (2, 0, 'system/auth', '权限管理', 'peoples', '', '', 0, 0, 1, 1, 'system/auth', '', '', 0, 'sys_admin', 0, NULL, NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (3, 0, 'system/monitor', '系统监控', 'monitor', '', '', 0, 0, 1, 1, 'monitor', '', '', 0, 'sys_admin', 0, NULL, NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (5, 1, 'system/config/dict/type/list', '字典管理', 'dict', '', '', 1, 0, 1, 1, 'dict/list', '', 'system/config/dict/list', 0, 'sys_admin', 0, NULL, '2021-07-20 08:47:37', NULL);
INSERT INTO `sys_auth_rule` VALUES (6, 1, 'system/config/sysConfig/list', '参数管理', 'date-range', '', '', 1, 0, 1, 1, 'params/list', '', 'system/config/params/list', 0, 'sys_admin', 0, NULL, '2021-07-20 08:58:57', NULL);
INSERT INTO `sys_auth_rule` VALUES (8, 2, 'system/auth/menuList', '菜单管理', 'nested', '', '', 1, 0, 1, 1, 'menuList', '', 'system/auth/menuList', 0, 'sys_admin', 0, NULL, '2021-07-20 09:01:49', NULL);
INSERT INTO `sys_auth_rule` VALUES (21, 5, 'system/config/dict/type/delete', '删除字典', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, '2021-07-20 08:56:49', NULL);
INSERT INTO `sys_auth_rule` VALUES (16, 8, 'system/auth/addMenu', '添加菜单', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, '2021-07-20 09:02:20', NULL);
INSERT INTO `sys_auth_rule` VALUES (17, 8, 'system/auth/editMenu', '修改菜单', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (18, 8, 'system/auth/deleteMenu', '删除菜单', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (19, 5, 'system/config/dict/type/add', '添加字典', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, '2021-07-20 08:56:29', NULL);
INSERT INTO `sys_auth_rule` VALUES (20, 5, 'system/config/dict/type/edit', '修改字典', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, '2021-07-20 08:56:39', NULL);
INSERT INTO `sys_auth_rule` VALUES (22, 5, 'system/config/dict/data/list', '字典数据管理', '', '', '', 2, 0, 1, 1, 'dataList', '', '', 0, 'sys_admin', 0, NULL, '2021-07-20 08:57:25', NULL);
INSERT INTO `sys_auth_rule` VALUES (23, 5, 'system/config/dict/data/add', '添加字典数据', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, '2021-07-20 08:58:06', NULL);
INSERT INTO `sys_auth_rule` VALUES (24, 5, 'system/config/dict/data/edit', '修改字典数据', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, '2021-07-20 08:58:14', NULL);
INSERT INTO `sys_auth_rule` VALUES (25, 5, 'system/config/dict/data/delete', '删除字典数据', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, '2021-07-20 08:58:22', NULL);
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
INSERT INTO `sys_auth_rule` VALUES (75, 0, 'system/tools', '系统工具', 'server', '', '', 0, 0, 1, 1, 'system/tools', '', '', 0, 'sys_admin', 0, NULL, NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (76, 75, 'system/tools/build', '表单构建', 'build', '', '', 1, 0, 1, 1, 'build', '', 'system/tools/build', 0, 'sys_admin', 0, NULL, NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (86, 1, 'system/config/sysWebSet', '站点设置', 'system', '', '', 1, 0, 1, 1, '/webSet', '', 'system/config/webSet', 0, 'sys_admin', 0, NULL, '2021-07-20 08:59:41', NULL);
INSERT INTO `sys_auth_rule` VALUES (185, 86, 'system/config/sysWebSet/update', '站点配置更新', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, '2021-07-20 09:00:18', '2021-07-20 09:00:18', NULL);
INSERT INTO `sys_auth_rule` VALUES (186, 75, 'system/tools/gen/tableList', '代码生成', 'code', '', '', 1, 0, 1, 1, 'gen', '', 'system/tools/gen', 0, 'sys_admin', 0, '2021-08-16 10:36:54', '2021-08-16 12:12:49', NULL);
INSERT INTO `sys_auth_rule` VALUES (231, 227, 'system/pluginsManage/delete', '插件管理删除', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, '2021-12-22 10:12:09', '2021-12-22 10:12:09', NULL);
INSERT INTO `sys_auth_rule` VALUES (232, 227, 'system/pluginsManage/changeStatus', '插件管理状态修改', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, '2021-12-22 10:12:31', '2021-12-22 10:12:31', NULL);
INSERT INTO `sys_auth_rule` VALUES (230, 227, 'system/pluginsManage/edit', '插件管理修改', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, '2021-12-22 10:11:48', '2021-12-22 10:11:48', NULL);
INSERT INTO `sys_auth_rule` VALUES (229, 227, 'system/pluginsManage/add', '插件管理添加', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, '2021-12-22 10:11:29', '2021-12-22 10:11:29', NULL);
INSERT INTO `sys_auth_rule` VALUES (228, 227, 'system/pluginsManage/get', '插件管理查询', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, '2021-12-22 10:10:54', '2021-12-22 10:10:54', NULL);
INSERT INTO `sys_auth_rule` VALUES (227, 75, 'system/pluginsManage/list', '插件管理', 'list', '', '', 1, 0, 1, 1, 'pluginsManageList', '', 'system/pluginsManage/list', 0, 'sys_admin', 0, '2021-12-22 10:10:15', '2021-12-22 10:10:15', NULL);
INSERT INTO `sys_auth_rule` VALUES (226, 75, 'system/tools/api', '系统接口', 'guide', '', '', 1, 0, 1, 1, 'api', '', 'system/tools/api', 0, 'sys_admin', 0, '2021-12-22 09:42:42', '2021-12-22 09:42:42', NULL);
INSERT INTO `sys_auth_rule` VALUES (233, 0, 'demo/demoDataAuth', '数据权限测试管理', 'clipboard', '', '', 0, 0, 1, 1, 'demoDataAuth', '', '', 0, 'sys_admin', 0, '2022-03-03 11:52:48', '2022-03-03 11:52:48', NULL);
INSERT INTO `sys_auth_rule` VALUES (234, 233, 'demo/demoDataAuth/list', '数据权限测试列表', 'date', '', '', 1, 0, 1, 1, 'demoDataAuthList', '', 'demo/demoDataAuth/list', 0, 'sys_admin', 0, '2022-03-03 11:53:26', '2022-03-03 11:53:26', NULL);

-- ----------------------------
-- Table structure for sys_config
-- ----------------------------
DROP TABLE IF EXISTS `sys_config`;
CREATE TABLE `sys_config`  (
  `config_id` int(5) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '参数主键',
  `config_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '参数名称',
  `config_key` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '参数键名',
  `config_value` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '参数键值',
  `config_type` tinyint(1) NULL DEFAULT 0 COMMENT '系统内置（Y是 N否）',
  `create_by` int(64) UNSIGNED NULL DEFAULT 0 COMMENT '创建者',
  `update_by` int(64) UNSIGNED NULL DEFAULT 0 COMMENT '更新者',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(0) NULL DEFAULT NULL COMMENT '修改时间',
  `deleted_at` datetime(0) NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`config_id`) USING BTREE,
  UNIQUE INDEX `uni_config_key`(`config_key`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 13 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_config
-- ----------------------------

---基础配置
INSERT INTO `sys_config` (`config_name`, `config_key`, `config_value`, `config_type`, `create_by`, `update_by`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES ( '白名单IP', 'sys.whitelist', '127.0.0.1', 1, 1, 1, '白名单IP', NULL, '2022-04-13 02:49:33', NULL);
INSERT INTO `sys_config` (`config_name`, `config_key`, `config_value`, `config_type`, `create_by`, `update_by`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES ( '充值回调地址', 'sys.rechargeCallbackUrl', '', 1, 1, 1, '充值回调地址', NULL, NULL, NULL);
INSERT INTO `sys_config` (`config_name`, `config_key`, `config_value`, `config_type`, `create_by`, `update_by`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES ( '提币自动通过金额', 'sys.minWithdrawAudit', '', 1, 1, 1, '提币自动通过金额', NULL, NULL, NULL);
---基础配置

--币安充值合约配置
INSERT INTO `sys_config` (`config_name`, `config_key`, `config_value`, `config_type`, `create_by`, `update_by`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES ( '币安链提币地址', 'sys.bnbWithdrawAddress', '', 1, 1, 1, '币安链提币地址', NULL, NULL, NULL);
INSERT INTO `sys_config` (`config_name`, `config_key`, `config_value`, `config_type`, `create_by`, `update_by`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES ( '币安链提币地址密钥', 'sys.bnbWithdrawAddressPrivateKey', '', 1, 1, 1, '币安链提币地址密钥', NULL, NULL, NULL);
INSERT INTO `sys_config` (`config_name`, `config_key`, `config_value`, `config_type`, `create_by`, `update_by`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES ( '币安链gasPrice', 'sys.bnbGasPrice', '', 1, 1, 1, '币安链gasPrice', NULL, NULL, NULL);
INSERT INTO `sys_config` (`config_name`, `config_key`, `config_value`, `config_type`, `create_by`, `update_by`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES ( '币安链gasLimit', 'sys.bnbGasLimit', '', 1, 1, 1, '币安链gasLimit', NULL, NULL, NULL);
--币安充值合约配置

--币安获取地址额外配置
INSERT INTO `sys_config` (`config_name`, `config_key`, `config_value`, `config_type`, `create_by`, `update_by`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES ( '币安链手续费地址', 'sys.bnbFeeAddress', '', 1, 1, 1, '币安链手续费地址', NULL, NULL, NULL);
INSERT INTO `sys_config` (`config_name`, `config_key`, `config_value`, `config_type`, `create_by`, `update_by`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES ( '币安链手续费地址密钥', 'sys.bnbFeeAddressPrivateKey', '', 1, 1, 1, '币安链手续费地址密钥', NULL, NULL, NULL);
--币安获取地址额外配置


--波场充值合约配置
INSERT INTO `sys_config` (`config_name`, `config_key`, `config_value`, `config_type`, `create_by`, `update_by`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES ( '波场提币地址', 'sys.tronWithdrawAddress', '', 1, 1, 1, '波场提币地址', NULL, NULL, NULL);
INSERT INTO `sys_config` (`config_name`, `config_key`, `config_value`, `config_type`, `create_by`, `update_by`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES ( '波场提币地址密钥', 'sys.tronWithdrawAddressPrivateKey', '', 1, 1, 1, '波场提币地址密钥', NULL, NULL, NULL);
INSERT INTO `sys_config` (`config_name`, `config_key`, `config_value`, `config_type`, `create_by`, `update_by`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES ( '波场手续费几个trx', 'sys.tronFee', '8', 1, 1, 1, '波场手续费几个trx', NULL, NULL, NULL);
--波场充值合约配置


--波场获取地址额外配置
INSERT INTO `sys_config` (`config_name`, `config_key`, `config_value`, `config_type`, `create_by`, `update_by`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES ( '波场手续费地址', 'sys.tronFeeAddress', '', 1, 1, 1, '波场手续费地址', NULL, NULL, NULL);
INSERT INTO `sys_config` (`config_name`, `config_key`, `config_value`, `config_type`, `create_by`, `update_by`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES ( '波场手续费地址密钥', 'sys.tronFeeAddressPrivateKey', '', 1, 1, 1, '波场手续费地址密钥', NULL, NULL, NULL);
--波场获取地址额外配置

INSERT INTO `sys_config` (`config_name`, `config_key`, `config_value`, `config_type`, `create_by`, `update_by`, `remark`, `created_at`, `updated_at`) VALUES ('文件上传-文件大小', 'sys.uploadFile.fileSize', '50M', 1, 31, 31, '文件上传大小限制', NULL, '2021-07-06 14:57:35');
INSERT INTO `sys_config` ( `config_name`, `config_key`, `config_value`, `config_type`, `create_by`, `update_by`, `remark`, `created_at`, `updated_at`) VALUES ('文件上传-文件类型', 'sys.uploadFile.fileType', 'jpg,jpeg,gif,npm,png', 1, 31, 1, '文件上传后缀类型限制', NULL, '2022-04-11 03:11:49');
INSERT INTO `sys_config` ( `config_name`, `config_key`, `config_value`, `config_type`, `create_by`, `update_by`, `remark`, `created_at`, `updated_at`) VALUES ('图片上传-图片类型', 'sys.uploadFile.imageType', 'jpg,jpeg,gif,npm,png', 1, 31, 0, '图片上传后缀类型限制', NULL, NULL);
INSERT INTO `sys_config` ( `config_name`, `config_key`, `config_value`, `config_type`, `create_by`, `update_by`, `remark`, `created_at`, `updated_at`) VALUES ( '图片上传-图片大小', 'sys.uploadFile.imageSize', '50M', 1, 31, 31, '图片上传大小限制', NULL, NULL);
INSERT INTO `sys_config` (`config_name`, `config_key`, `config_value`, `config_type`, `create_by`, `update_by`, `remark`, `created_at`, `updated_at`) VALUES ( '静态资源', 'static.resource', '/', 1, 2, 0, '', NULL, NULL);

-- ----------------------------
-- Table structure for sys_dept
-- ----------------------------
DROP TABLE IF EXISTS `sys_dept`;
CREATE TABLE `sys_dept`  (
  `dept_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '部门id',
  `parent_id` bigint(20) NULL DEFAULT 0 COMMENT '父部门id',
  `ancestors` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '祖级列表',
  `dept_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '部门名称',
  `order_num` int(4) NULL DEFAULT 0 COMMENT '显示顺序',
  `leader` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '负责人',
  `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '联系电话',
  `email` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '邮箱',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '0' COMMENT '部门状态（0正常 1停用）',
  `created_by` bigint(20) UNSIGNED NULL DEFAULT 0 COMMENT '创建人',
  `updated_by` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(0) NULL DEFAULT NULL COMMENT '修改时间',
  `deleted_at` datetime(0) NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`dept_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 207 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '部门表' ROW_FORMAT = Compact;

-- ----------------------------
-- Records of sys_dept
-- ----------------------------
INSERT INTO `sys_dept` VALUES (100, 0, '0', 'CEO', 0, 'admin', '18000000000', 'wallet@qq.com', '1', 0, 31, '2021-07-13 15:56:52', '2021-07-13 15:57:05', NULL);

-- ----------------------------
-- Table structure for sys_dict_data
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_data`;
CREATE TABLE `sys_dict_data`  (
  `dict_code` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '字典编码',
  `dict_sort` int(4) NULL DEFAULT 0 COMMENT '字典排序',
  `dict_label` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '字典标签',
  `dict_value` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '字典键值',
  `dict_type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '字典类型',
  `css_class` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '样式属性（其他样式扩展）',
  `list_class` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '表格回显样式',
  `is_default` tinyint(1) NULL DEFAULT 0 COMMENT '是否默认（1是 0否）',
  `status` tinyint(1) NULL DEFAULT 0 COMMENT '状态（0正常 1停用）',
  `create_by` bigint(64) UNSIGNED NULL DEFAULT 0 COMMENT '创建者',
  `update_by` bigint(64) UNSIGNED NULL DEFAULT 0 COMMENT '更新者',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(0) NULL DEFAULT NULL COMMENT '修改时间',
  `deleted_at` datetime(0) NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`dict_code`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 88 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '字典数据表' ROW_FORMAT = Compact;

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
  `dict_id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '字典主键',
  `dict_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '字典名称',
  `dict_type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '字典类型',
  `status` tinyint(1) UNSIGNED NULL DEFAULT 0 COMMENT '状态（0正常 1停用）',
  `create_by` int(64) UNSIGNED NULL DEFAULT 0 COMMENT '创建者',
  `update_by` int(64) UNSIGNED NULL DEFAULT 0 COMMENT '更新者',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '创建日期',
  `updated_at` datetime(0) NULL DEFAULT NULL COMMENT '修改日期',
  `deleted_at` datetime(0) NULL DEFAULT NULL COMMENT '删除日期',
  PRIMARY KEY (`dict_id`) USING BTREE,
  UNIQUE INDEX `dict_type`(`dict_type`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 36 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '字典类型表' ROW_FORMAT = Compact;

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
  `job_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '任务ID',
  `job_name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '任务名称',
  `job_params` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '参数',
  `job_group` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'DEFAULT' COMMENT '任务组名',
  `invoke_target` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '调用目标字符串',
  `cron_expression` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT 'cron执行表达式',
  `misfire_policy` tinyint(4) NULL DEFAULT 1 COMMENT '计划执行策略（1多次执行 2执行一次）',
  `concurrent` tinyint(4) NULL DEFAULT 1 COMMENT '是否并发执行（0允许 1禁止）',
  `status` tinyint(4) NULL DEFAULT 0 COMMENT '状态（0正常 1暂停）',
  `create_by` bigint(64) UNSIGNED NULL DEFAULT 0 COMMENT '创建者',
  `update_by` bigint(64) UNSIGNED NULL DEFAULT 0 COMMENT '更新者',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '备注信息',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime(0) NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`job_id`, `job_name`, `job_group`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 13 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '定时任务调度表' ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for sys_login_log
-- ----------------------------
DROP TABLE IF EXISTS `sys_login_log`;
CREATE TABLE `sys_login_log`  (
  `info_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '访问ID',
  `login_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '登录账号',
  `ipaddr` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '登录IP地址',
  `login_location` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '登录地点',
  `browser` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '浏览器类型',
  `os` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '操作系统',
  `status` tinyint(4) NULL DEFAULT 0 COMMENT '登录状态（0成功 1失败）',
  `msg` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '提示消息',
  `login_time` datetime(0) NULL COMMENT '登录时间',
  `module` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '登录模块',
  PRIMARY KEY (`info_id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 25 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统访问记录' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sys_model_info
-- ----------------------------
DROP TABLE IF EXISTS `sys_model_info`;
CREATE TABLE `sys_model_info`  (
  `model_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '模型ID',
  `model_category_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '模板分类id',
  `model_name` char(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '模型标识',
  `model_title` char(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '模型名称',
  `model_pk` char(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '主键字段',
  `model_order` char(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '默认排序字段',
  `model_sort` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '表单字段排序',
  `model_list` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '列表显示字段，为空显示全部',
  `model_edit` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '可编辑字段，为空则除主键外均可以编辑',
  `model_indexes` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '索引字段',
  `search_list` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '高级搜索的字段',
  `create_time` bigint(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `update_time` bigint(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
  `model_status` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态',
  `model_engine` varchar(25) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'MyISAM' COMMENT '数据库引擎',
  `create_by` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建人',
  `update_by` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '修改人',
  PRIMARY KEY (`model_id`) USING BTREE,
  UNIQUE INDEX `name_uni`(`model_name`) USING BTREE COMMENT '模型名唯一'
) ENGINE = InnoDB AUTO_INCREMENT = 19 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '文档模型表' ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for sys_oper_log
-- ----------------------------
DROP TABLE IF EXISTS `sys_oper_log`;
CREATE TABLE `sys_oper_log`  (
  `oper_id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '日志主键',
  `title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '模块标题',
  `business_type` int(2) NULL DEFAULT 0 COMMENT '业务类型（0其它 1新增 2修改 3删除）',
  `method` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '方法名称',
  `request_method` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '请求方式',
  `operator_type` int(1) NULL DEFAULT 0 COMMENT '操作类别（0其它 1后台用户 2手机端用户）',
  `oper_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '操作人员',
  `dept_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '部门名称',
  `oper_url` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '请求URL',
  `oper_ip` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '主机地址',
  `oper_location` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '操作地点',
  `oper_param` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '请求参数',
  `json_result` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '返回参数',
  `status` int(1) NULL DEFAULT 0 COMMENT '操作状态（0正常 1异常）',
  `error_msg` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '错误消息',
  `oper_time` datetime(0) NULL DEFAULT NULL COMMENT '操作时间',
  PRIMARY KEY (`oper_id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 1183 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '操作日志记录' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sys_post
-- ----------------------------
DROP TABLE IF EXISTS `sys_post`;
CREATE TABLE `sys_post`  (
  `post_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '岗位ID',
  `post_code` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '岗位编码',
  `post_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '岗位名称',
  `post_sort` int(4) NOT NULL COMMENT '显示顺序',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '状态（0正常 1停用）',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  `created_by` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建人',
  `updated_by` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '修改人',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(0) NULL DEFAULT NULL COMMENT '修改时间',
  `deleted_at` datetime(0) NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`post_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '岗位信息表' ROW_FORMAT = Compact;

-- ----------------------------
-- Records of sys_post
-- ----------------------------
INSERT INTO `sys_post` VALUES (1, 'ceo', '超级管理员', 1, '1', '', 0, 0, '2021-07-11 11:32:58', NULL, NULL);

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `status` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态;0:禁用;1:正常',
  `list_order` float NOT NULL DEFAULT 0 COMMENT '排序',
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '角色名称',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `data_scope` tinyint(3) UNSIGNED NOT NULL DEFAULT 3 COMMENT '数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `status`(`status`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '角色表' ROW_FORMAT = Compact;

-- ----------------------------
-- Records of sys_role
-- ----------------------------
INSERT INTO `sys_role` VALUES (1, 1, 0, '超级管理员', '备注', 3);

-- ----------------------------
-- Table structure for sys_role_dept
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_dept`;
CREATE TABLE `sys_role_dept`  (
  `role_id` bigint(20) NOT NULL COMMENT '角色ID',
  `dept_id` bigint(20) NOT NULL COMMENT '部门ID',
  PRIMARY KEY (`role_id`, `dept_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '角色和部门关联表' ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_name` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户名',
  `mobile` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '中国手机不带国家代码，国际手机号格式为：国家代码-手机号',
  `user_nickname` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户昵称',
  `birthday` int(11) NOT NULL DEFAULT 0 COMMENT '生日',
  `user_password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '登录密码;cmf_password加密',
  `user_salt` char(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '加密盐',
  `user_status` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '用户状态;0:禁用,1:正常,2:未验证',
  `user_email` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户登录邮箱',
  `sex` tinyint(2) NOT NULL DEFAULT 0 COMMENT '性别;0:保密,1:男,2:女',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户头像',
  `dept_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '部门id',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `is_admin` tinyint(4) NOT NULL DEFAULT 1 COMMENT '是否后台管理员 1 是  0   否',
  `address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '联系地址',
  `describe` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT ' 描述信息',
  `phone_num` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '联系电话',
  `last_login_ip` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '最后登录ip',
  `last_login_time` datetime(0) NULL COMMENT '最后登录时间',
  `created_at` datetime(0) NULL COMMENT '创建时间',
  `updated_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime(0) NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `user_login`(`user_name`) USING BTREE,
  UNIQUE INDEX `mobile`(`mobile`) USING BTREE,
  INDEX `user_nickname`(`user_nickname`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 42 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户表' ROW_FORMAT = Compact;

-- ----------------------------
-- Records of sys_user
-- ----------------------------
INSERT INTO `sys_user` VALUES (1, 'admin', '18800000000', '超级管理员', 0, '39978de67915a11e94bfe9c879b2d9a1', 'gqwLs4n95E', 1, 'yxh669@qq.com', 1, 'https://yxh-1301841944.cos.ap-chongqing.myqcloud.com/gfast/2021-07-19/ccwpeuqz1i2s769hua.jpeg', 101, '', 1, 'asdasfdsaf大发放打发士大夫发按时', '描述信息', '18611111111', '[::1]', '2021-07-19 14:28:38', '2021-06-22 17:58:00', '2021-07-12 22:07:44', NULL);
-- ----------------------------
-- Table structure for sys_user_online
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_online`;
CREATE TABLE `sys_user_online`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `uuid` char(32) CHARACTER SET latin1 COLLATE latin1_general_ci NOT NULL DEFAULT '' COMMENT '用户标识',
  `token` varchar(255) CHARACTER SET latin1 COLLATE latin1_general_ci NOT NULL DEFAULT '' COMMENT '用户token',
  `create_time` datetime(0) NULL DEFAULT NULL COMMENT '登录时间',
  `user_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名',
  `ip` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '登录ip',
  `explorer` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '浏览器',
  `os` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '操作系统',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uni_token`(`token`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 16832 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户在线状态表' ROW_FORMAT = Dynamic;


-- ----------------------------
-- Table structure for sys_user_post
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_post`;
CREATE TABLE `sys_user_post`  (
  `user_id` bigint(20) NOT NULL COMMENT '用户ID',
  `post_id` bigint(20) NOT NULL COMMENT '岗位ID',
  PRIMARY KEY (`user_id`, `post_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户与岗位关联表' ROW_FORMAT = Compact;

-- ----------------------------
-- Records of sys_user_post
-- ----------------------------
INSERT INTO `sys_user_post` VALUES (1, 2);
-- ----------------------------
-- Table structure for sys_web_set
-- ----------------------------
DROP TABLE IF EXISTS `sys_web_set`;
CREATE TABLE `sys_web_set`  (
  `web_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `web_content` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '站点信息',
  PRIMARY KEY (`web_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of sys_web_set
-- ----------------------------
INSERT INTO `sys_web_set` VALUES (1, '{\"CopyrightInfo\":\"11122\",\"recordInfo\":\"111222\",\"statisticsCode\":\"11122\",\"webId\":1,\"webLogo\":\"https://yxh-1301841944.cos.ap-chongqing.myqcloud.com/gfast/2021-07-13/ccroz2q3sptczqwchk.jpg\",\"webName\":\"gfast\",\"webSite\":\"http://localhost/index#/webSet\"}');

-- ----------------------------
-- Table structure for tools_gen_table
-- ----------------------------
DROP TABLE IF EXISTS `tools_gen_table`;
CREATE TABLE `tools_gen_table`  (
  `table_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '编号',
  `table_name` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '表名称',
  `table_comment` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '表描述',
  `class_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '实体类名称',
  `tpl_category` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT 'crud' COMMENT '使用的模板（crud单表操作 tree树表操作）',
  `package_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '生成包路径',
  `module_name` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '生成模块名',
  `business_name` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '生成业务名',
  `function_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '生成功能名',
  `function_author` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '生成功能作者',
  `options` varchar(1000) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '其它生成选项',
  `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`table_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 70 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '代码生成业务表' ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for tools_gen_table_column
-- ----------------------------
DROP TABLE IF EXISTS `tools_gen_table_column`;
CREATE TABLE `tools_gen_table_column`  (
  `column_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '编号',
  `table_id` bigint(20) NULL DEFAULT NULL COMMENT '归属表编号',
  `column_name` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '列名称',
  `column_comment` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '列描述',
  `column_type` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '列类型',
  `go_type` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'Go类型',
  `go_field` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'Go字段名',
  `html_field` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'html字段名',
  `is_pk` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '是否主键（1是）',
  `is_increment` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '是否自增（1是）',
  `is_required` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '是否必填（1是）',
  `is_insert` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '是否为插入字段（1是）',
  `is_edit` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '是否编辑字段（1是）',
  `is_list` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '是否列表字段（1是）',
  `is_query` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '是否查询字段（1是）',
  `query_type` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT 'EQ' COMMENT '查询方式（等于、不等于、大于、小于、范围）',
  `html_type` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '显示类型（文本框、文本域、下拉框、复选框、单选框、日期控件）',
  `dict_type` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '字典类型',
  `sort` int(11) NULL DEFAULT NULL COMMENT '排序',
  `link_table_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '关联表名',
  `link_table_class` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '关联表类名',
  `link_table_package` varchar(150) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '关联表包名',
  `link_label_id` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '关联表键名',
  `link_label_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '关联表字段值',
  PRIMARY KEY (`column_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 720 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '代码生成业务表字段' ROW_FORMAT = Compact;

SET FOREIGN_KEY_CHECKS = 1;
