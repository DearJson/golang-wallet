-- MySQL dump 10.13  Distrib 8.0.35, for Linux (x86_64)
--
-- Host: localhost    Database: wallet
-- ------------------------------------------------------
-- Server version	8.0.35

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `address`
--

DROP TABLE IF EXISTS `address`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `address` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `user_id` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '用户标识',
  `main_chain` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '主链',
  `address` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '地址',
  `private_key` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '私钥',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `chain_address_index` (`main_chain`,`address`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `address`
--

LOCK TABLES `address` WRITE;
/*!40000 ALTER TABLE `address` DISABLE KEYS */;
/*!40000 ALTER TABLE `address` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `casbin_rule`
--

DROP TABLE IF EXISTS `casbin_rule`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `casbin_rule` (
  `ptype` varchar(10) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  `v0` varchar(256) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  `v1` varchar(256) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  `v2` varchar(256) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  `v3` varchar(256) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  `v4` varchar(256) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  `v5` varchar(256) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `casbin_rule`
--

LOCK TABLES `casbin_rule` WRITE;
/*!40000 ALTER TABLE `casbin_rule` DISABLE KEYS */;
INSERT INTO `casbin_rule` VALUES ('g','1','1','','','',''),('g','2','2','','','',''),('p','2','295','All','','',''),('p','2','296','All','','',''),('p','2','297','All','','',''),('p','2','298','All','','',''),('p','2','299','All','','',''),('p','2','303','All','','',''),('p','2','304','All','','',''),('p','2','305','All','','',''),('p','2','306','All','','',''),('p','2','308','All','','',''),('p','2','309','All','','',''),('p','2','34','All','','',''),('p','2','36','All','','',''),('p','2','37','All','','',''),('p','2','38','All','','',''),('p','2','39','All','','',''),('g','3','2','','','','');
/*!40000 ALTER TABLE `casbin_rule` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `currency`
--

DROP TABLE IF EXISTS `currency`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `currency` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `main_chain` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '主链',
  `name` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '币种名称',
  `contract_address` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '合约地址',
  `decimals` int NOT NULL COMMENT '精度',
  `min_withdraw` decimal(15,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '最低提现金额',
  `min_merge` decimal(15,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '最低归集金额',
  `withdraw_split_enabled` tinyint unsigned NOT NULL DEFAULT '0' COMMENT 'Solana提现分账开关：0关闭，1开启',
  `withdraw_split_address` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT 'Solana提现分账地址',
  `withdraw_split_amount` decimal(30,18) unsigned NOT NULL DEFAULT '0.000000000000000000' COMMENT 'Solana提现固定分账数量',
  `withdraw_split_bps` smallint unsigned NOT NULL DEFAULT '0' COMMENT 'Solana提现分账比例（基点，200=2%）',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=103 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `currency`
--

LOCK TABLES `currency` WRITE;
/*!40000 ALTER TABLE `currency` DISABLE KEYS */;
INSERT INTO `currency` (`id`,`main_chain`,`name`,`contract_address`,`decimals`,`min_withdraw`,`min_merge`,`created_at`,`updated_at`) VALUES (1,'tron','TRX','TBRop8PopYu8atWWez3g3ueVtSpseW78b6',6,0.00,0.00,'2022-03-13 00:04:56','2022-04-11 20:31:27'),(2,'bsc','BNB','0xbb4cdb9cbd36b01bd1cbaebf2de08d9173bc095c',18,0.00,0.00,'2022-03-15 18:13:31','2022-08-23 15:22:39'),(3,'heco','HT','0x0000000000000000000000000000000000000000',18,0.00,0.00,NULL,'2022-05-17 17:49:23'),(50,'bsc','USDT','0x55d398326f99059ff775485246999027b3197955',18,0.00,0.00,'2022-08-03 15:59:40','2022-08-03 15:59:40'),(100,'eth','ETH','0x0000000000000000000000000000000000000000',6,0.00,0.00,'2022-08-03 15:59:40','2022-08-03 15:59:40'),(101,'solana','SOL','So11111111111111111111111111111111111111112',9,0.00,0.00,NULL,NULL),(102,'solana','USDT','Es9vMFrzaCERmJfrF4H2FYD4KCoNkY11McCe8BenwNYB',6,0.00,0.00,'2026-02-08 18:42:51','2026-02-08 18:42:51');
/*!40000 ALTER TABLE `currency` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `demo_data_auth`
--

DROP TABLE IF EXISTS `demo_data_auth`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `demo_data_auth` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT '' COMMENT '标题',
  `created_by` int unsigned DEFAULT '0' COMMENT '创建人',
  `updated_by` int unsigned DEFAULT '0' COMMENT '修改人',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '修改时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `demo_data_auth`
--

LOCK TABLES `demo_data_auth` WRITE;
/*!40000 ALTER TABLE `demo_data_auth` DISABLE KEYS */;
/*!40000 ALTER TABLE `demo_data_auth` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `demo_gen`
--

DROP TABLE IF EXISTS `demo_gen`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `demo_gen` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `demo_name` varchar(20) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '姓名',
  `demo_age` int unsigned NOT NULL DEFAULT '0' COMMENT '年龄',
  `classes` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '班级',
  `demo_born` datetime DEFAULT NULL COMMENT '出生年月',
  `demo_gender` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '性别',
  `created_at` datetime DEFAULT NULL COMMENT '创建日期',
  `updated_at` datetime DEFAULT NULL COMMENT '修改日期',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除日期',
  `created_by` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建人',
  `updated_by` bigint unsigned NOT NULL DEFAULT '0' COMMENT '修改人',
  `demo_status` tinyint NOT NULL DEFAULT '0' COMMENT '状态',
  `demo_cate` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '分类',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='代码生成测试表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `demo_gen`
--

LOCK TABLES `demo_gen` WRITE;
/*!40000 ALTER TABLE `demo_gen` DISABLE KEYS */;
/*!40000 ALTER TABLE `demo_gen` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `demo_gen_class`
--

DROP TABLE IF EXISTS `demo_gen_class`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `demo_gen_class` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '分类id',
  `class_name` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '分类名',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='代码生成关联测试表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `demo_gen_class`
--

LOCK TABLES `demo_gen_class` WRITE;
/*!40000 ALTER TABLE `demo_gen_class` DISABLE KEYS */;
/*!40000 ALTER TABLE `demo_gen_class` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `demo_gen_tree`
--

DROP TABLE IF EXISTS `demo_gen_tree`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `demo_gen_tree` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `parent_id` int unsigned NOT NULL DEFAULT '0' COMMENT '父级ID',
  `demo_name` varchar(20) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '姓名',
  `demo_age` int unsigned NOT NULL DEFAULT '0' COMMENT '年龄',
  `classes` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '班级',
  `demo_born` datetime DEFAULT NULL COMMENT '出生年月',
  `demo_gender` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '性别',
  `created_at` datetime DEFAULT NULL COMMENT '创建日期',
  `updated_at` datetime DEFAULT NULL COMMENT '修改日期',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除日期',
  `created_by` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建人',
  `updated_by` bigint unsigned NOT NULL DEFAULT '0' COMMENT '修改人',
  `demo_status` tinyint NOT NULL DEFAULT '0' COMMENT '状态',
  `demo_cate` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '分类',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='代码生成树形结构测试表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `demo_gen_tree`
--

LOCK TABLES `demo_gen_tree` WRITE;
/*!40000 ALTER TABLE `demo_gen_tree` DISABLE KEYS */;
/*!40000 ALTER TABLE `demo_gen_tree` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `fee_list`
--

DROP TABLE IF EXISTS `fee_list`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `fee_list` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `main_chain` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '主链',
  `coin_name` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '手续费币种',
  `withdraw_address` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '手续费地址',
  `address` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '地址',
  `amount` decimal(30,18) NOT NULL COMMENT '转移手续费',
  `hash` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT 'hash',
  `status` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '状态 1-待上链 2-上链成功 3-上链失败',
  `recharge_id` int unsigned NOT NULL DEFAULT '0' COMMENT '充值ID',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `address_index` (`main_chain`,`address`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `fee_list`
--

LOCK TABLES `fee_list` WRITE;
/*!40000 ALTER TABLE `fee_list` DISABLE KEYS */;
/*!40000 ALTER TABLE `fee_list` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `height`
--

DROP TABLE IF EXISTS `height`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `height` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `main_chain` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '主网',
  `height` int NOT NULL COMMENT '块号',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `height`
--

LOCK TABLES `height` WRITE;
/*!40000 ALTER TABLE `height` DISABLE KEYS */;
/*!40000 ALTER TABLE `height` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `liquidity_currency`
--

DROP TABLE IF EXISTS `liquidity_currency`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `liquidity_currency` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `main_chain` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '链',
  `name` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '币种名称',
  `contract_address` varchar(150) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '币种地址',
  `decimals` int NOT NULL COMMENT '精度',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `liquidity_currency`
--

LOCK TABLES `liquidity_currency` WRITE;
/*!40000 ALTER TABLE `liquidity_currency` DISABLE KEYS */;
/*!40000 ALTER TABLE `liquidity_currency` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `liquidity_swap_trade`
--

DROP TABLE IF EXISTS `liquidity_swap_trade`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `liquidity_swap_trade` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'id主键',
  `main_chain` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '主链',
  `block_hash` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '块hash',
  `from_address` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '交易方地址',
  `coin_token` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '币种1',
  `coin_token1` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '币种2',
  `coin_token_decimals` smallint unsigned NOT NULL DEFAULT '0' COMMENT '币种1精度',
  `coin_token1_decimals` smallint unsigned NOT NULL DEFAULT '0' COMMENT '币种2精度',
  `contract_address` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '币种1合约地址',
  `contract_address1` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '币种2合约地址',
  `amount` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '币种1数量',
  `amount1` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '币种2数量',
  `lp_address` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT 'LP合约地址',
  `fee_num` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '手续费流动性数量',
  `lp_num` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '流动性数量',
  `to_address` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '收币地址',
  `type` tinyint(1) NOT NULL COMMENT '类型 1-添加流动性 2-移除流动性',
  `hash` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '交易hash',
  `block_height` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '区块高度',
  `call_number` smallint unsigned DEFAULT '0' COMMENT '回调次数',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `hash_index` (`hash`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `liquidity_swap_trade`
--

LOCK TABLES `liquidity_swap_trade` WRITE;
/*!40000 ALTER TABLE `liquidity_swap_trade` DISABLE KEYS */;
/*!40000 ALTER TABLE `liquidity_swap_trade` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `monitor_currency`
--

DROP TABLE IF EXISTS `monitor_currency`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `monitor_currency` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `main_chain` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '链',
  `name` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '币种名称',
  `contract_address` varchar(150) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '币种地址',
  `decimals` int NOT NULL COMMENT '精度',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `monitor_currency`
--

LOCK TABLES `monitor_currency` WRITE;
/*!40000 ALTER TABLE `monitor_currency` DISABLE KEYS */;
/*!40000 ALTER TABLE `monitor_currency` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `monitor_swap_trade`
--

DROP TABLE IF EXISTS `monitor_swap_trade`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `monitor_swap_trade` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'id主键',
  `main_chain` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '主链',
  `block_hash` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '块hash',
  `from_address` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '交易方地址',
  `contract_address` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '兑换前币种',
  `contract_address1` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '兑换后币种',
  `amount` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '兑换前数量',
  `amount1` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '兑换后数量',
  `path` text CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci COMMENT '兑换路径',
  `to_address` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '收币地址',
  `type` tinyint(1) NOT NULL COMMENT '类型 1-买入 2-卖出',
  `hash` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '交易hash',
  `block_height` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '区块高度',
  `call_number` smallint unsigned DEFAULT '0' COMMENT '回调次数',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `hash_index` (`hash`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `monitor_swap_trade`
--

LOCK TABLES `monitor_swap_trade` WRITE;
/*!40000 ALTER TABLE `monitor_swap_trade` DISABLE KEYS */;
/*!40000 ALTER TABLE `monitor_swap_trade` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `plugins_manage`
--

DROP TABLE IF EXISTS `plugins_manage`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `plugins_manage` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `store_id` int DEFAULT NULL COMMENT '插件在商城中的id',
  `p_name` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '插件名称英文',
  `p_title` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '插件名称',
  `p_description` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '插件介绍',
  `p_auth` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '作者',
  `is_install` tinyint NOT NULL DEFAULT '0' COMMENT '是否安装',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态',
  `version` varchar(60) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '当前版本',
  `price` int unsigned NOT NULL DEFAULT '0' COMMENT '价格',
  `download_times` int unsigned NOT NULL DEFAULT '0' COMMENT '下载次数',
  `install_path` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '安装路径（用于卸载）',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `storeIdUni` (`store_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='插件管理';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `plugins_manage`
--

LOCK TABLES `plugins_manage` WRITE;
/*!40000 ALTER TABLE `plugins_manage` DISABLE KEYS */;
/*!40000 ALTER TABLE `plugins_manage` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `recharge`
--

DROP TABLE IF EXISTS `recharge`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `recharge` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'id主键',
  `main_chain` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '主链',
  `block_hash` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '块hash',
  `coin_token` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '币种1',
  `coin_token1` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '币种2',
  `from_address` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '发送方地址',
  `to_address` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '充币地址',
  `amount` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '充币数量',
  `amount1` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '充币2数量',
  `contract_address` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '币种1合约地址',
  `contract_address1` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '币种2合约地址',
  `hash` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '充币交易hashId',
  `block_height` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '区块高度',
  `call_number` smallint unsigned DEFAULT '0' COMMENT '回调次数',
  `status` smallint DEFAULT '1' COMMENT '状态，1充币成功待归集，2归集上链中，3归集成功, 4归集失败，5充值费用中',
  `imputation_hash` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '归集hash',
  `remarks` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '备注',
  `recharge_type` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '充币类型',
  `token_id` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '充值的卡牌tokenId  充值卡牌时使用',
  `custome_coin` text CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci COMMENT '自定义充值币种',
  `custome_user` text CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci COMMENT '自定义充值地址',
  `custome_amount` text CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci COMMENT '自定义充值数量',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `hash_index` (`hash`) USING BTREE,
  KEY `to_address_index` (`to_address`) USING BTREE,
  KEY `from_address_index` (`from_address`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=59 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `recharge`
--

LOCK TABLES `recharge` WRITE;
/*!40000 ALTER TABLE `recharge` DISABLE KEYS */;
INSERT INTO `recharge` VALUES (58,'solana','','USDT','','7Pp9D2oAjA59rvuFKVws5fcHr7zzQPb7uWa85x135k9z','B2dWxxKM6BFok7yeWVu8TGDDarWS6eUkz3zRZPBAHDJc','0.8','','Es9vMFrzaCERmJfrF4H2FYD4KCoNkY11McCe8BenwNYB','','5ueZxpQiQuL1qGiUA4QUFL2zuFsZejVFd3bWKms5YUcRo5xDo6FfRdMrMnWxAqzDLK6SXjEDdax2YKFtarK7PZte','398946292',0,1,NULL,'ORD-1770578549-1296',1,'','','','','2026-02-08 19:22:59','2026-02-08 19:22:59');
/*!40000 ALTER TABLE `recharge` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_auth_rule`
--

DROP TABLE IF EXISTS `sys_auth_rule`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_auth_rule` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `pid` int unsigned NOT NULL DEFAULT '0' COMMENT '父ID',
  `name` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '规则名称',
  `title` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '规则名称',
  `icon` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '图标',
  `condition` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '条件',
  `remark` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `menu_type` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '类型 0目录 1菜单 2按钮',
  `weigh` int NOT NULL DEFAULT '0' COMMENT '权重',
  `status` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  `always_show` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '显示状态',
  `path` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '路由地址',
  `jump_path` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '跳转路由',
  `component` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '组件路径',
  `is_frame` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '是否外链 1是 0否',
  `module_type` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '所属模块',
  `model_id` int unsigned NOT NULL DEFAULT '0' COMMENT '模型ID',
  `created_at` datetime DEFAULT NULL COMMENT '创建日期',
  `updated_at` datetime DEFAULT NULL COMMENT '修改日期',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除日期',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `name` (`name`) USING BTREE,
  KEY `pid` (`pid`) USING BTREE,
  KEY `weigh` (`weigh`) USING BTREE
) ENGINE=MyISAM AUTO_INCREMENT=419 DEFAULT CHARSET=utf8mb3 COMMENT='菜单节点表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_auth_rule`
--

LOCK TABLES `sys_auth_rule` WRITE;
/*!40000 ALTER TABLE `sys_auth_rule` DISABLE KEYS */;
INSERT INTO `sys_auth_rule` VALUES (1,0,'system/config','系统配置','system','','Admin tips',0,0,1,1,'config','','',0,'sys_admin',0,NULL,NULL,NULL),(2,0,'system/auth','权限管理','peoples','','',0,0,1,1,'system/auth','','',0,'sys_admin',0,NULL,NULL,NULL),(3,0,'system/monitor','系统监控','monitor','','',0,0,1,1,'monitor','','',0,'sys_admin',0,NULL,NULL,NULL),(6,1,'system/config/sysConfig/list','参数管理','date-range','','',1,1,1,1,'params/list','','system/config/params/list',0,'sys_admin',0,NULL,'2022-04-11 04:29:47',NULL),(8,2,'system/auth/menuList','菜单管理','nested','','',1,0,1,1,'menuList','','system/auth/menuList',0,'sys_admin',0,NULL,'2021-07-20 09:01:49',NULL),(16,8,'system/auth/addMenu','添加菜单','','','',2,0,1,1,'','','',0,'sys_admin',0,NULL,'2021-07-20 09:02:20',NULL),(17,8,'system/auth/editMenu','修改菜单','','','',2,0,1,1,'','','',0,'sys_admin',0,NULL,NULL,NULL),(18,8,'system/auth/deleteMenu','删除菜单','','','',2,0,1,1,'','','',0,'sys_admin',0,NULL,NULL,NULL),(26,2,'system/auth/roleList','角色管理','logininfor','','',1,0,1,1,'role','','system/auth/roleList',0,'sys_admin',0,NULL,NULL,NULL),(27,26,'system/auth/addRole','添加角色','','','',2,0,1,1,'','','',0,'sys_admin',0,NULL,NULL,NULL),(28,2,'system/auth/deptList','部门管理','peoples','','',1,0,1,1,'dept','','system/auth/dept',0,'sys_admin',0,NULL,'2021-07-20 09:03:46',NULL),(29,26,'system/auth/editRole','修改角色','','','',2,0,1,1,'','','',0,'sys_admin',0,NULL,NULL,NULL),(30,26,'system/auth/statusSetRole','设置角色状态','','','',2,0,1,1,'','','',0,'sys_admin',0,NULL,NULL,NULL),(31,26,'system/auth/deleteRole','删除角色','','','',2,0,1,1,'','','',0,'sys_admin',0,NULL,NULL,NULL),(33,2,'system/auth/postList','岗位管理','tab','','',1,0,1,1,'post','','system/auth/post',0,'sys_admin',0,NULL,'2021-07-20 09:04:51',NULL),(34,3,'system/monitor/online/list','在线用户','cascader','','',1,0,1,1,'online','','system/monitor/online/list',0,'sys_admin',0,NULL,NULL,NULL),(36,3,'system/monitor/job','定时任务','clipboard','','',1,0,1,1,'job','','system/monitor/job',0,'sys_admin',0,NULL,NULL,NULL),(37,3,'system/monitor/server/info','服务监控','dict','','',1,0,1,1,'server','','system/monitor/server',0,'sys_admin',0,NULL,'2021-07-19 16:07:23',NULL),(38,3,'system/monitor/loginLog','登录日志','chart','','',1,0,1,1,'logininfor','','system/monitor/logininfor',0,'sys_admin',0,NULL,'2021-07-20 09:08:00',NULL),(39,3,'system/monitor/operLog','操作日志','dashboard','','',1,0,1,1,'operlog','','system/monitor/operlog',0,'sys_admin',0,NULL,'2021-07-20 09:08:31',NULL),(40,2,'system/auth/userList','用户管理','user','','',1,0,1,1,'user','','system/auth/userList',0,'sys_admin',0,NULL,NULL,NULL),(41,6,'system/config/sysConfig/add','添加参数','','','',2,0,1,1,'','','',0,'sys_admin',0,NULL,'2021-07-20 08:59:09',NULL),(42,6,'system/config/sysConfig/edit','修改参数','','','',2,0,1,1,'','','',0,'sys_admin',0,NULL,'2021-07-20 08:59:17',NULL),(43,6,'system/config/sysConfig/delete','删除参数','','','',2,0,1,1,'','','',0,'sys_admin',0,NULL,'2021-07-20 08:59:25',NULL),(44,28,'system/auth/deptAdd','添加部门','','','',2,0,1,1,'','','',0,'sys_admin',0,NULL,'2021-07-20 09:04:05',NULL),(45,28,'system/auth/deptEdit','修改部门','','','',2,0,1,1,'','','',0,'sys_admin',0,NULL,'2021-07-20 09:04:19',NULL),(46,28,'system/auth/deptDelete','删除部门','','','',2,0,1,1,'','','',0,'sys_admin',0,NULL,'2021-07-20 09:04:35',NULL),(47,33,'system/auth/postAdd','添加岗位','','','',2,0,1,1,'','','',0,'sys_admin',0,NULL,'2021-07-20 09:05:04',NULL),(48,33,'system/auth/postEdit','修改岗位','','','',2,0,1,1,'','','',0,'sys_admin',0,NULL,'2021-07-20 09:05:15',NULL),(49,33,'system/auth/postDelete','删除岗位','','','',2,0,1,1,'','','',0,'sys_admin',0,NULL,'2021-07-20 09:05:25',NULL),(50,40,'system/auth/addUser','添加用户','','','',2,0,1,1,'','','',0,'sys_admin',0,NULL,NULL,NULL),(51,40,'system/auth/editUser','修改用户','','','',2,0,1,1,'','','',0,'sys_admin',0,NULL,NULL,NULL),(52,40,'system/auth/resetUserPwd','密码重置','','','',2,0,1,1,'','','',0,'sys_admin',0,NULL,NULL,NULL),(53,40,'system/auth/changeUserStatus','状态设置','','','',2,0,1,1,'','','',0,'sys_admin',0,NULL,NULL,NULL),(54,40,'system/auth/deleteUser','删除用户','','','',2,0,1,1,'','','',0,'sys_admin',0,NULL,'2021-07-20 09:06:16',NULL),(55,34,'system/monitor/online/forceLogout','强制退出','','','',2,0,1,1,'','','',0,'sys_admin',0,NULL,NULL,NULL),(56,36,'system/monitor/job/add','添加任务','','','',2,0,1,1,'','','',0,'sys_admin',0,NULL,NULL,NULL),(57,36,'system/monitor/job/edit','修改任务','','','',2,0,1,1,'','','',0,'sys_admin',0,NULL,NULL,NULL),(58,36,'system/monitor/job/start','开启任务','','','',2,0,1,1,'','','',0,'sys_admin',0,NULL,NULL,NULL),(59,36,'system/monitor/job/stop','停止任务','','','',2,0,1,1,'','','',0,'sys_admin',0,NULL,NULL,NULL),(60,36,'system/monitor/job/delete','删除任务','','','',2,0,1,1,'','','',0,'sys_admin',0,NULL,NULL,NULL),(61,38,'system/monitor/loginLog/delete','删除','','','',2,0,1,1,'','','',0,'sys_admin',0,NULL,'2021-07-20 09:08:10',NULL),(62,38,'system/monitor/loginLog/clear','清空','','','',2,0,1,1,'','','',0,'sys_admin',0,NULL,'2021-07-20 09:08:15',NULL),(63,39,'system/monitor/operLog/delete','删除','','','',2,0,1,1,'','','',0,'sys_admin',0,NULL,'2021-07-20 09:08:36',NULL),(64,39,'system/monitor/operLog/clear','清空','','','',2,0,1,1,'','','',0,'sys_admin',0,NULL,'2021-07-20 09:08:41',NULL),(400,0,'solanaWallet','Solana钱包管理','chart','','',0,2,1,1,'solanaWallet','','',0,'sys_admin',0,'2026-02-08 18:05:12','2026-02-08 18:05:12',NULL),(401,400,'system/solanaAddress/list','地址管理','','','',1,10,1,1,'solanaAddress','','system/solanaAddress/list',0,'sys_admin',0,'2026-02-08 18:05:12','2026-02-08 18:05:12',NULL),(402,401,'system/solanaAddress/get','查询','','','',2,0,1,1,'','','',0,'sys_admin',0,'2026-02-08 18:05:12','2026-02-08 18:05:12',NULL),(403,400,'system/solanaCurrency/list','币种管理','','','',1,9,1,1,'solanaCurrency','','system/solanaCurrency/list',0,'sys_admin',0,'2026-02-08 18:05:12','2026-02-08 18:05:12',NULL),(404,403,'system/solanaCurrency/get','查询','','','',2,0,1,1,'','','',0,'sys_admin',0,'2026-02-08 18:05:12','2026-02-08 18:05:12',NULL),(405,403,'system/solanaCurrency/add','添加','','','',2,0,1,1,'','','',0,'sys_admin',0,'2026-02-08 18:05:12','2026-02-08 18:05:12',NULL),(406,403,'system/solanaCurrency/edit','修改','','','',2,0,1,1,'','','',0,'sys_admin',0,'2026-02-08 18:05:12','2026-02-08 18:05:12',NULL),(407,403,'system/solanaCurrency/delete','删除','','','',2,0,1,1,'','','',0,'sys_admin',0,'2026-02-08 18:05:12','2026-02-08 18:05:12',NULL),(408,400,'system/solanaFeelist/list','手续费管理','','','',1,8,1,1,'solanaFeeList','','system/solanaFeeList/list',0,'sys_admin',0,'2026-02-08 18:05:12','2026-02-08 18:05:12',NULL),(409,408,'system/solanaFeelist/get','查询','','','',2,0,1,1,'','','',0,'sys_admin',0,'2026-02-08 18:05:12','2026-02-08 18:05:12',NULL),(410,400,'system/solanaRecharge/list','充值管理','','','',1,7,1,1,'solanaRechargeList','','system/solanaRecharge/list',0,'sys_admin',0,'2026-02-08 18:05:12','2026-02-08 18:05:12',NULL),(411,410,'system/solanaRecharge/get','查询','','','',2,0,1,1,'','','',0,'sys_admin',0,'2026-02-08 18:05:12','2026-02-08 18:05:12',NULL),(412,410,'system/solanaRecharge/callback','发起回调','','','',2,0,1,1,'','','',0,'sys_admin',0,'2026-02-08 18:05:12','2026-02-08 18:05:12',NULL),(413,400,'system/solanaWithdraw/list','提现管理','','','',1,6,1,1,'solanaWithdrawList','','system/solanaWithdraw/list',0,'sys_admin',0,'2026-02-08 18:05:12','2026-02-08 18:05:12',NULL),(414,413,'system/solanaWithdraw/get','查询','','','',2,0,1,1,'','','',0,'sys_admin',0,'2026-02-08 18:05:12','2026-02-08 18:05:12',NULL),(415,413,'system/solanaWithdraw/callback','发起回调','','','',2,0,1,1,'','','',0,'sys_admin',0,'2026-02-08 18:05:12','2026-02-08 18:05:12',NULL),(416,413,'system/solanaWithdraw/auditSuccess','审核通过','','','',2,0,1,1,'','','',0,'sys_admin',0,'2026-02-08 18:05:12','2026-02-08 18:05:12',NULL),(417,413,'system/solanaWithdraw/auditFail','审核拒绝','','','',2,0,1,1,'','','',0,'sys_admin',0,'2026-02-08 18:05:12','2026-02-08 18:05:12',NULL),(418,413,'system/solanaWithdraw/withdraw','手动提现','','','',2,0,1,1,'','','',0,'sys_admin',0,'2026-02-08 18:05:12','2026-02-08 18:05:12',NULL);
/*!40000 ALTER TABLE `sys_auth_rule` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_config`
--

DROP TABLE IF EXISTS `sys_config`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_config` (
  `config_id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '参数主键',
  `config_name` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT '' COMMENT '参数名称',
  `config_key` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT '' COMMENT '参数键名',
  `config_value` varchar(500) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT '' COMMENT '参数键值',
  `config_type` tinyint(1) DEFAULT '0' COMMENT '系统内置（Y是 N否）',
  `create_by` int unsigned DEFAULT '0' COMMENT '创建者',
  `update_by` int unsigned DEFAULT '0' COMMENT '更新者',
  `remark` varchar(500) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '备注',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '修改时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`config_id`) USING BTREE,
  UNIQUE KEY `uni_config_key` (`config_key`) USING BTREE
) ENGINE=MyISAM AUTO_INCREMENT=36 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_config`
--

LOCK TABLES `sys_config` WRITE;
/*!40000 ALTER TABLE `sys_config` DISABLE KEYS */;
INSERT INTO `sys_config` VALUES (1,'白名单IP','sys.whitelist','127.0.0.1',1,1,1,'白名单IP',NULL,'2022-08-23 15:26:31',NULL),(2,'充值回调地址','sys.rechargeCallbackUrl','http://127.0.0.1:8000/api/v1/wallRechargeCallback',1,1,1,'充值回调地址',NULL,'2026-02-08 19:44:41',NULL),(3,'提币自动通过金额','sys.minWithdrawAudit','0',1,1,1,'提币自动通过金额',NULL,'2023-05-17 21:33:35',NULL),(31,'Solana提现地址','sys.solanaWithdrawAddress','',0,0,0,'Solana出金钱包地址',NULL,NULL,NULL),(32,'Solana提现私钥','sys.solanaWithdrawAddressPrivateKey','',0,0,0,'Solana出金钱包私钥',NULL,NULL,NULL),(33,'Solana手续费地址','sys.solanaFeeAddress','',0,0,0,'Solana手续费代付钱包地址',NULL,NULL,NULL),(34,'Solana手续费私钥','sys.solanaFeeAddressPrivateKey','',0,0,0,'Solana手续费代付钱包私钥',NULL,NULL,NULL),(35,'Solana归集地址','sys.solanaMergeAddress','',0,0,0,'Solana归集目标钱包地址',NULL,NULL,NULL);
/*!40000 ALTER TABLE `sys_config` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_dept`
--

DROP TABLE IF EXISTS `sys_dept`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_dept` (
  `dept_id` bigint NOT NULL AUTO_INCREMENT COMMENT '部门id',
  `parent_id` bigint DEFAULT '0' COMMENT '父部门id',
  `ancestors` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT '' COMMENT '祖级列表',
  `dept_name` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT '' COMMENT '部门名称',
  `order_num` int DEFAULT '0' COMMENT '显示顺序',
  `leader` varchar(20) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '负责人',
  `phone` varchar(11) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '联系电话',
  `email` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '邮箱',
  `status` char(1) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT '0' COMMENT '部门状态（0正常 1停用）',
  `created_by` bigint unsigned DEFAULT '0' COMMENT '创建人',
  `updated_by` bigint DEFAULT NULL COMMENT '修改人',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '修改时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`dept_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=101 DEFAULT CHARSET=utf8mb3 COMMENT='部门表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_dept`
--

LOCK TABLES `sys_dept` WRITE;
/*!40000 ALTER TABLE `sys_dept` DISABLE KEYS */;
INSERT INTO `sys_dept` VALUES (100,0,'0','CEO',0,'admin','18000000000','wallet@qq.com','1',0,31,'2021-07-13 15:56:52','2021-07-13 15:57:05',NULL);
/*!40000 ALTER TABLE `sys_dept` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_dict_data`
--

DROP TABLE IF EXISTS `sys_dict_data`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_dict_data` (
  `dict_code` bigint NOT NULL AUTO_INCREMENT COMMENT '字典编码',
  `dict_sort` int DEFAULT '0' COMMENT '字典排序',
  `dict_label` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT '' COMMENT '字典标签',
  `dict_value` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT '' COMMENT '字典键值',
  `dict_type` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT '' COMMENT '字典类型',
  `css_class` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '样式属性（其他样式扩展）',
  `list_class` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '表格回显样式',
  `is_default` tinyint(1) DEFAULT '0' COMMENT '是否默认（1是 0否）',
  `status` tinyint(1) DEFAULT '0' COMMENT '状态（0正常 1停用）',
  `create_by` bigint unsigned DEFAULT '0' COMMENT '创建者',
  `update_by` bigint unsigned DEFAULT '0' COMMENT '更新者',
  `remark` varchar(500) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '备注',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '修改时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`dict_code`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=88 DEFAULT CHARSET=utf8mb3 COMMENT='字典数据表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_dict_data`
--

LOCK TABLES `sys_dict_data` WRITE;
/*!40000 ALTER TABLE `sys_dict_data` DISABLE KEYS */;
INSERT INTO `sys_dict_data` VALUES (1,0,'男','1','sys_user_sex','','',0,1,31,2,'备注信息',NULL,NULL,NULL),(2,0,'女','2','sys_user_sex','','',0,1,31,31,'备注信息',NULL,NULL,NULL),(3,0,'保密','0','sys_user_sex','','',1,1,31,31,'备注信息',NULL,NULL,NULL),(28,0,'正常','0','sys_job_status','','default',1,1,31,0,'',NULL,NULL,NULL),(29,0,'暂停','1','sys_job_status','','default',0,1,31,31,'',NULL,NULL,NULL),(30,0,'默认','DEFAULT','sys_job_group','','default',1,1,31,0,'',NULL,NULL,NULL),(31,0,'系统','SYSTEM','sys_job_group','','default',0,1,31,0,'',NULL,NULL,NULL),(32,0,'成功','1','admin_login_status','','default',0,1,31,31,'',NULL,NULL,NULL),(33,0,'失败','0','admin_login_status','','default',0,1,31,0,'',NULL,NULL,NULL),(34,0,'成功','1','sys_oper_log_status','','default',0,1,31,0,'',NULL,NULL,NULL),(35,0,'失败','0','sys_oper_log_status','','default',0,1,31,0,'',NULL,NULL,NULL),(36,0,'重复执行','1','sys_job_policy','','default',1,1,31,0,'',NULL,NULL,NULL),(37,0,'执行一次','2','sys_job_policy','','default',1,1,31,0,'',NULL,NULL,NULL),(38,0,'显示','1','sys_show_hide',NULL,'default',1,1,31,0,NULL,NULL,NULL,NULL),(39,0,'隐藏','0','sys_show_hide',NULL,'default',0,1,31,0,NULL,NULL,NULL,NULL),(40,0,'正常','1','sys_normal_disable','','default',1,1,31,0,'',NULL,NULL,NULL),(41,0,'停用','0','sys_normal_disable','','default',0,1,31,0,'',NULL,NULL,NULL),(49,0,'是','1','sys_yes_no','','',1,1,31,0,'',NULL,NULL,NULL),(50,0,'否','0','sys_yes_no','','',0,1,31,0,'',NULL,NULL,NULL),(61,0,'政府工作目标','1','gov_cate_models','','',0,1,2,0,'',NULL,NULL,NULL),(62,0,'系统后台','sys_admin','menu_module_type','','',1,1,2,0,'',NULL,NULL,NULL),(63,0,'政务工作','gov_work','menu_module_type','','',0,1,2,0,'',NULL,NULL,NULL),(65,0,'[work]测试业务表','wf_news','flow_type','','',0,1,2,2,'',NULL,NULL,NULL),(66,0,'回退修改','-1','flow_status','','',0,1,31,0,'',NULL,NULL,NULL),(67,0,'保存中','0','flow_status','','',0,1,31,0,'',NULL,NULL,NULL),(68,0,'流程中','1','flow_status','','',0,1,31,0,'',NULL,NULL,NULL),(69,0,'审批通过','2','flow_status','','',0,1,31,2,'',NULL,NULL,NULL),(70,2,'发布栏目','2','sys_blog_sign','','',0,1,31,31,'',NULL,NULL,NULL),(71,3,'跳转栏目','3','sys_blog_sign','','',0,1,31,31,'',NULL,NULL,NULL),(72,4,'单页栏目','4','sys_blog_sign','','',0,1,31,31,'',NULL,NULL,NULL),(73,2,'置顶','1','sys_log_sign','','',0,1,31,31,'',NULL,NULL,NULL),(74,3,'幻灯','2','sys_log_sign','','',0,1,31,31,'',NULL,NULL,NULL),(75,4,'推荐','3','sys_log_sign','','',0,1,31,31,'',NULL,NULL,NULL),(76,1,'一般','0','sys_log_sign','','',0,1,31,31,'',NULL,NULL,NULL),(77,1,'频道页','1','sys_blog_sign','','',0,1,31,31,'',NULL,NULL,NULL),(78,0,'普通','0','flow_level','','',0,1,31,0,'',NULL,'2021-07-20 08:55:20',NULL),(79,0,'加急','1','flow_level','','',0,1,31,0,'',NULL,'2021-07-20 08:55:20',NULL),(80,0,'紧急','2','flow_level','','',0,1,31,0,'',NULL,'2021-07-20 08:55:20',NULL),(81,0,'特急','3','flow_level','','',0,1,31,31,'',NULL,'2021-07-20 08:55:25',NULL),(82,0,'频道页','1','sys_blog_type','','',0,1,31,0,'',NULL,NULL,NULL),(83,0,'发布栏目','2','sys_blog_type','','',0,1,31,0,'',NULL,NULL,NULL),(84,0,'跳转栏目','3','sys_blog_type','','',0,1,31,31,'',NULL,NULL,NULL),(85,0,'单页栏目','4','sys_blog_type','','',0,1,31,0,'',NULL,NULL,NULL),(87,0,'[cms]文章表','cms_news','flow_type','','',0,1,31,0,'',NULL,NULL,NULL);
/*!40000 ALTER TABLE `sys_dict_data` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_dict_type`
--

DROP TABLE IF EXISTS `sys_dict_type`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_dict_type` (
  `dict_id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '字典主键',
  `dict_name` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT '' COMMENT '字典名称',
  `dict_type` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT '' COMMENT '字典类型',
  `status` tinyint unsigned DEFAULT '0' COMMENT '状态（0正常 1停用）',
  `create_by` int unsigned DEFAULT '0' COMMENT '创建者',
  `update_by` int unsigned DEFAULT '0' COMMENT '更新者',
  `remark` varchar(500) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '备注',
  `created_at` datetime DEFAULT NULL COMMENT '创建日期',
  `updated_at` datetime DEFAULT NULL COMMENT '修改日期',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除日期',
  PRIMARY KEY (`dict_id`) USING BTREE,
  UNIQUE KEY `dict_type` (`dict_type`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=36 DEFAULT CHARSET=utf8mb3 COMMENT='字典类型表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_dict_type`
--

LOCK TABLES `sys_dict_type` WRITE;
/*!40000 ALTER TABLE `sys_dict_type` DISABLE KEYS */;
INSERT INTO `sys_dict_type` VALUES (1,'用户性别','sys_user_sex',1,31,1,'用于选择用户性别',NULL,NULL,NULL),(3,'任务状态','sys_job_status',1,31,31,'任务状态列表',NULL,NULL,NULL),(13,'任务分组','sys_job_group',1,31,0,'',NULL,NULL,NULL),(14,'管理员登录状态','admin_login_status',1,31,0,'',NULL,NULL,NULL),(15,'操作日志状态','sys_oper_log_status',1,31,0,'',NULL,NULL,NULL),(16,'任务策略','sys_job_policy',1,31,0,'',NULL,NULL,NULL),(17,'菜单状态','sys_show_hide',1,31,0,'菜单状态',NULL,NULL,NULL),(18,'系统开关','sys_normal_disable',1,31,31,'系统开关',NULL,NULL,NULL),(24,'系统内置','sys_yes_no',1,31,0,'',NULL,NULL,NULL),(29,'政务工作模型分类','gov_cate_models',1,2,0,'',NULL,NULL,NULL),(30,'菜单模块类型','menu_module_type',1,2,0,'',NULL,NULL,NULL),(31,'工作流程类型','flow_type',1,2,0,'',NULL,NULL,NULL),(32,'工作流程审批状态','flow_status',1,31,0,'工作流程审批状态',NULL,NULL,NULL),(33,'博客分类类型','sys_blog_type',1,31,31,'博客分类中的标志',NULL,NULL,NULL),(34,'博客日志标志','sys_log_sign',1,31,0,'博客日志管理中的标志数据字典',NULL,NULL,NULL),(35,'工作流紧急状态','flow_level',1,31,31,'',NULL,'2021-07-20 08:55:20',NULL);
/*!40000 ALTER TABLE `sys_dict_type` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_job`
--

DROP TABLE IF EXISTS `sys_job`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_job` (
  `job_id` bigint NOT NULL AUTO_INCREMENT COMMENT '任务ID',
  `job_name` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '任务名称',
  `job_params` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT '' COMMENT '参数',
  `job_group` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT 'DEFAULT' COMMENT '任务组名',
  `invoke_target` varchar(500) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '调用目标字符串',
  `cron_expression` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT '' COMMENT 'cron执行表达式',
  `misfire_policy` tinyint DEFAULT '1' COMMENT '计划执行策略（1多次执行 2执行一次）',
  `concurrent` tinyint DEFAULT '1' COMMENT '是否并发执行（0允许 1禁止）',
  `status` tinyint DEFAULT '0' COMMENT '状态（0正常 1暂停）',
  `create_by` bigint unsigned DEFAULT '0' COMMENT '创建者',
  `update_by` bigint unsigned DEFAULT '0' COMMENT '更新者',
  `remark` varchar(500) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT '' COMMENT '备注信息',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`job_id`,`job_name`,`job_group`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb3 COMMENT='定时任务调度表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_job`
--

LOCK TABLES `sys_job` WRITE;
/*!40000 ALTER TABLE `sys_job` DISABLE KEYS */;
INSERT INTO `sys_job` VALUES (1,'币安扫描任务','','','bscSweepTask','*/10 * * * * ?',1,1,0,0,0,'',NULL,NULL,NULL),(4,'检查归集出金手续费状态','','','checkStatusTask','*/30 * * * * ?',1,1,0,1,0,'',NULL,NULL,NULL),(5,'出金任务','','','withdrawTask','*/30 * * * * ?',1,1,0,1,0,'',NULL,NULL,NULL),(6,'归集任务','','','rechargeTask','*/30 * * * * ?',1,1,1,1,0,'',NULL,NULL,'2023-05-15 02:06:58'),(7,'检查用户登录状态','','DEFAULT','checkUserOnline','20 */1 * * * ?',1,1,0,1,1,'',NULL,NULL,NULL),(9,'更新节点','','DEFAULT','updateNode','20 */1 * * * ?',1,1,0,0,0,'',NULL,NULL,NULL);
/*!40000 ALTER TABLE `sys_job` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_login_log`
--

DROP TABLE IF EXISTS `sys_login_log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_login_log` (
  `info_id` bigint NOT NULL AUTO_INCREMENT COMMENT '访问ID',
  `login_name` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT '' COMMENT '登录账号',
  `ipaddr` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT '' COMMENT '登录IP地址',
  `login_location` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT '' COMMENT '登录地点',
  `browser` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT '' COMMENT '浏览器类型',
  `os` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT '' COMMENT '操作系统',
  `status` tinyint DEFAULT '0' COMMENT '登录状态（0成功 1失败）',
  `msg` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT '' COMMENT '提示消息',
  `login_time` datetime DEFAULT NULL COMMENT '登录时间',
  `module` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT '' COMMENT '登录模块',
  PRIMARY KEY (`info_id`) USING BTREE
) ENGINE=MyISAM AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb3 COMMENT='系统访问记录';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_login_log`
--

LOCK TABLES `sys_login_log` WRITE;
/*!40000 ALTER TABLE `sys_login_log` DISABLE KEYS */;
INSERT INTO `sys_login_log` VALUES (12,'12tK4ofQyhfu1bt','45.62.109.106, 45.62.109.106','','Chrome','Intel Mac OS X 10_15_7',1,'登录成功','2026-02-08 18:42:08','系统后台');
/*!40000 ALTER TABLE `sys_login_log` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_model_info`
--

DROP TABLE IF EXISTS `sys_model_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_model_info` (
  `model_id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '模型ID',
  `model_category_id` int unsigned NOT NULL DEFAULT '0' COMMENT '模板分类id',
  `model_name` char(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '模型标识',
  `model_title` char(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '模型名称',
  `model_pk` char(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '主键字段',
  `model_order` char(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '默认排序字段',
  `model_sort` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '表单字段排序',
  `model_list` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '列表显示字段，为空显示全部',
  `model_edit` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT '' COMMENT '可编辑字段，为空则除主键外均可以编辑',
  `model_indexes` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '索引字段',
  `search_list` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT '' COMMENT '高级搜索的字段',
  `create_time` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` bigint unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  `model_status` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  `model_engine` varchar(25) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT 'MyISAM' COMMENT '数据库引擎',
  `create_by` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建人',
  `update_by` bigint unsigned NOT NULL DEFAULT '0' COMMENT '修改人',
  PRIMARY KEY (`model_id`) USING BTREE,
  UNIQUE KEY `name_uni` (`model_name`) USING BTREE COMMENT '模型名唯一'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='文档模型表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_model_info`
--

LOCK TABLES `sys_model_info` WRITE;
/*!40000 ALTER TABLE `sys_model_info` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_model_info` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_oper_log`
--

DROP TABLE IF EXISTS `sys_oper_log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_oper_log` (
  `oper_id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '日志主键',
  `title` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT '' COMMENT '模块标题',
  `business_type` int DEFAULT '0' COMMENT '业务类型（0其它 1新增 2修改 3删除）',
  `method` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT '' COMMENT '方法名称',
  `request_method` varchar(10) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT '' COMMENT '请求方式',
  `operator_type` int DEFAULT '0' COMMENT '操作类别（0其它 1后台用户 2手机端用户）',
  `oper_name` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT '' COMMENT '操作人员',
  `dept_name` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT '' COMMENT '部门名称',
  `oper_url` varchar(500) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT '' COMMENT '请求URL',
  `oper_ip` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT '' COMMENT '主机地址',
  `oper_location` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT '' COMMENT '操作地点',
  `oper_param` text CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci COMMENT '请求参数',
  `json_result` text CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci COMMENT '返回参数',
  `status` int DEFAULT '0' COMMENT '操作状态（0正常 1异常）',
  `error_msg` varchar(2000) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT '' COMMENT '错误消息',
  `oper_time` datetime DEFAULT NULL COMMENT '操作时间',
  PRIMARY KEY (`oper_id`) USING BTREE
) ENGINE=MyISAM AUTO_INCREMENT=220 DEFAULT CHARSET=utf8mb3 COMMENT='操作日志记录';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_oper_log`
--

LOCK TABLES `sys_oper_log` WRITE;
/*!40000 ALTER TABLE `sys_oper_log` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_oper_log` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_post`
--

DROP TABLE IF EXISTS `sys_post`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_post` (
  `post_id` bigint NOT NULL AUTO_INCREMENT COMMENT '岗位ID',
  `post_code` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '岗位编码',
  `post_name` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '岗位名称',
  `post_sort` int NOT NULL COMMENT '显示顺序',
  `status` char(1) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '状态（0正常 1停用）',
  `remark` varchar(500) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '备注',
  `created_by` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建人',
  `updated_by` bigint unsigned NOT NULL DEFAULT '0' COMMENT '修改人',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '修改时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`post_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb3 COMMENT='岗位信息表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_post`
--

LOCK TABLES `sys_post` WRITE;
/*!40000 ALTER TABLE `sys_post` DISABLE KEYS */;
INSERT INTO `sys_post` VALUES (1,'ceo','超级管理员',1,'1','',0,0,'2021-07-11 11:32:58',NULL,NULL);
/*!40000 ALTER TABLE `sys_post` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_role`
--

DROP TABLE IF EXISTS `sys_role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_role` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `status` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '状态;0:禁用;1:正常',
  `list_order` float NOT NULL DEFAULT '0' COMMENT '排序',
  `name` varchar(20) CHARACTER SET utf8mb3 COLLATE utf8mb3_unicode_ci NOT NULL DEFAULT '' COMMENT '角色名称',
  `remark` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `data_scope` tinyint unsigned NOT NULL DEFAULT '3' COMMENT '数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `status` (`status`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb3 COMMENT='角色表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_role`
--

LOCK TABLES `sys_role` WRITE;
/*!40000 ALTER TABLE `sys_role` DISABLE KEYS */;
INSERT INTO `sys_role` VALUES (1,1,0,'超级管理员','备注',3),(2,0,0,'观察者','',3);
/*!40000 ALTER TABLE `sys_role` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_role_dept`
--

DROP TABLE IF EXISTS `sys_role_dept`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_role_dept` (
  `role_id` bigint NOT NULL COMMENT '角色ID',
  `dept_id` bigint NOT NULL COMMENT '部门ID',
  PRIMARY KEY (`role_id`,`dept_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='角色和部门关联表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_role_dept`
--

LOCK TABLES `sys_role_dept` WRITE;
/*!40000 ALTER TABLE `sys_role_dept` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_role_dept` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_user`
--

DROP TABLE IF EXISTS `sys_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_user` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `user_name` varchar(60) CHARACTER SET utf8mb3 COLLATE utf8mb3_unicode_ci NOT NULL DEFAULT '' COMMENT '用户名',
  `mobile` varchar(20) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '中国手机不带国家代码，国际手机号格式为：国家代码-手机号',
  `user_nickname` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_unicode_ci NOT NULL DEFAULT '' COMMENT '用户昵称',
  `birthday` int NOT NULL DEFAULT '0' COMMENT '生日',
  `user_password` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '登录密码;cmf_password加密',
  `user_salt` char(10) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '加密盐',
  `user_status` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '用户状态;0:禁用,1:正常,2:未验证',
  `user_email` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '用户登录邮箱',
  `sex` tinyint NOT NULL DEFAULT '0' COMMENT '性别;0:保密,1:男,2:女',
  `avatar` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '用户头像',
  `dept_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '部门id',
  `remark` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `is_admin` tinyint NOT NULL DEFAULT '1' COMMENT '是否后台管理员 1 是  0   否',
  `address` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '联系地址',
  `describe` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT ' 描述信息',
  `phone_num` varchar(20) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '联系电话',
  `last_login_ip` varchar(15) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '最后登录ip',
  `last_login_time` datetime DEFAULT NULL COMMENT '最后登录时间',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `user_login` (`user_name`) USING BTREE,
  UNIQUE KEY `mobile` (`mobile`) USING BTREE,
  KEY `user_nickname` (`user_nickname`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb3 COMMENT='用户表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_user`
--

LOCK TABLES `sys_user` WRITE;
/*!40000 ALTER TABLE `sys_user` DISABLE KEYS */;
INSERT INTO `sys_user` VALUES (1,'12tK4ofQyhfu1bt','18800000000','超级管理员',0,'52e64d9fd45edb45e36a04f17971b5da','bynrGbgkiK',1,'admin@qq.com',1,'pub_upload/2022-04-11/cj6sdjf021iw6samde.jpeg',100,'',1,'asdasfdsaf大发放打发士大夫发按时','描述信息','18611111111','113.90.13.247','2023-05-17 21:27:22','2021-06-22 17:58:00','2022-04-11 03:10:21',NULL);
/*!40000 ALTER TABLE `sys_user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_user_online`
--

DROP TABLE IF EXISTS `sys_user_online`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_user_online` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `uuid` char(32) CHARACTER SET latin1 COLLATE latin1_general_ci NOT NULL DEFAULT '' COMMENT '用户标识',
  `token` varchar(255) CHARACTER SET latin1 COLLATE latin1_general_ci NOT NULL DEFAULT '' COMMENT '用户token',
  `create_time` datetime DEFAULT NULL COMMENT '登录时间',
  `user_name` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '用户名',
  `ip` varchar(120) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '登录ip',
  `explorer` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '浏览器',
  `os` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '操作系统',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uni_token` (`token`) USING BTREE
) ENGINE=MyISAM AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb3 COMMENT='用户在线状态表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_user_online`
--

LOCK TABLES `sys_user_online` WRITE;
/*!40000 ALTER TABLE `sys_user_online` DISABLE KEYS */;
INSERT INTO `sys_user_online` VALUES (12,'ec7d4d45bd582a78d290ab2cd5ea8b78','3kyRknVwOHYBlwxirj29vjKtr5kRYxUgOVId9jPGV2YwNkDk80ftVg6Gxuv+XNtmN/rcmr4BkqrZ1yYGG8JB7clG1V05R8cpNdrSdczBjCOgZZM2RdHuTJ2p9XVHxkZRGkpkTi4rjSolHSG3F0CXWg==','2026-02-08 18:42:07','12tK4ofQyhfu1bt','45.62.109.106, 45.62.109.106','Chrome','Intel Mac OS X 10_15_7');
/*!40000 ALTER TABLE `sys_user_online` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_user_post`
--

DROP TABLE IF EXISTS `sys_user_post`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_user_post` (
  `user_id` bigint NOT NULL COMMENT '用户ID',
  `post_id` bigint NOT NULL COMMENT '岗位ID',
  PRIMARY KEY (`user_id`,`post_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='用户与岗位关联表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_user_post`
--

LOCK TABLES `sys_user_post` WRITE;
/*!40000 ALTER TABLE `sys_user_post` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_user_post` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_web_set`
--

DROP TABLE IF EXISTS `sys_web_set`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_web_set` (
  `web_id` int NOT NULL AUTO_INCREMENT COMMENT '主键',
  `web_content` text CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci COMMENT '站点信息',
  PRIMARY KEY (`web_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_web_set`
--

LOCK TABLES `sys_web_set` WRITE;
/*!40000 ALTER TABLE `sys_web_set` DISABLE KEYS */;
INSERT INTO `sys_web_set` VALUES (1,'{\"CopyrightInfo\":\"json\",\"recordInfo\":\"111222\",\"statisticsCode\":\"11122\",\"webId\":1,\"webLogo\":\"https://yxh-1301841944.cos.ap-chongqing.myqcloud.com/gfast/2021-07-13/ccroz2q3sptczqwchk.jpg\",\"webName\":\"钱包管理系统\",\"webSite\":\"http://localhost/index#/webSet\"}');
/*!40000 ALTER TABLE `sys_web_set` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tools_gen_table`
--

DROP TABLE IF EXISTS `tools_gen_table`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tools_gen_table` (
  `table_id` bigint NOT NULL AUTO_INCREMENT COMMENT '编号',
  `table_name` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT '' COMMENT '表名称',
  `table_comment` varchar(500) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT '' COMMENT '表描述',
  `class_name` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT '' COMMENT '实体类名称',
  `tpl_category` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT 'crud' COMMENT '使用的模板（crud单表操作 tree树表操作）',
  `package_name` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '生成包路径',
  `module_name` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '生成模块名',
  `business_name` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '生成业务名',
  `function_name` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '生成功能名',
  `function_author` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '生成功能作者',
  `options` varchar(1000) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '其它生成选项',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`table_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=76 DEFAULT CHARSET=utf8mb3 COMMENT='代码生成业务表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tools_gen_table`
--

LOCK TABLES `tools_gen_table` WRITE;
/*!40000 ALTER TABLE `tools_gen_table` DISABLE KEYS */;
INSERT INTO `tools_gen_table` VALUES (70,'address','','Address','crud','gfast/app/system','system','address','','gfast','','2022-04-11 03:32:45','2022-04-11 03:32:45',''),(71,'currency','','Currency','crud','gfast/app/system','system','currency','','gfast','','2022-04-11 04:18:33','2022-04-11 04:18:33',''),(72,'fee_list','','FeeList','crud','gfast/app/system','system','fee_list','','gfast','','2022-04-11 04:18:33','2022-04-11 04:18:33',''),(73,'height','','Height','crud','gfast/app/system','system','height','','gfast','','2022-04-11 04:18:33','2022-04-11 04:18:33',''),(74,'recharge','','Recharge','crud','gfast/app/system','system','recharge','','gfast','','2022-04-11 04:18:33','2022-04-11 04:18:33',''),(75,'withdraw','','Withdraw','crud','gfast/app/system','system','withdraw','','gfast','','2022-04-11 04:18:34','2022-04-11 04:18:34','');
/*!40000 ALTER TABLE `tools_gen_table` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tools_gen_table_column`
--

DROP TABLE IF EXISTS `tools_gen_table_column`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tools_gen_table_column` (
  `column_id` bigint NOT NULL AUTO_INCREMENT COMMENT '编号',
  `table_id` bigint DEFAULT NULL COMMENT '归属表编号',
  `column_name` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '列名称',
  `column_comment` varchar(500) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '列描述',
  `column_type` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '列类型',
  `go_type` varchar(500) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT 'Go类型',
  `go_field` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT 'Go字段名',
  `html_field` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT 'html字段名',
  `is_pk` char(1) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '是否主键（1是）',
  `is_increment` char(1) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '是否自增（1是）',
  `is_required` char(1) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '是否必填（1是）',
  `is_insert` char(1) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '是否为插入字段（1是）',
  `is_edit` char(1) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '是否编辑字段（1是）',
  `is_list` char(1) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '是否列表字段（1是）',
  `is_query` char(1) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '是否查询字段（1是）',
  `query_type` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT 'EQ' COMMENT '查询方式（等于、不等于、大于、小于、范围）',
  `html_type` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '显示类型（文本框、文本域、下拉框、复选框、单选框、日期控件）',
  `dict_type` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT '' COMMENT '字典类型',
  `sort` int DEFAULT NULL COMMENT '排序',
  `link_table_name` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '关联表名',
  `link_table_class` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '关联表类名',
  `link_table_package` varchar(150) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '关联表包名',
  `link_label_id` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '关联表键名',
  `link_label_name` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '关联表字段值',
  PRIMARY KEY (`column_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=781 DEFAULT CHARSET=utf8mb3 COMMENT='代码生成业务表字段';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tools_gen_table_column`
--

LOCK TABLES `tools_gen_table_column` WRITE;
/*!40000 ALTER TABLE `tools_gen_table_column` DISABLE KEYS */;
INSERT INTO `tools_gen_table_column` VALUES (720,70,'id','','int unsigned','','Id','id','1','1','0','0','0','0','0','EQ','','',1,'','','','',''),(721,70,'user_id','用户标识','int','int','UserId','userId','0','0','1','1','1','1','1','EQ','input','',2,'','','','',''),(722,70,'main_chain','主链','varchar(50)','string','MainChain','mainChain','0','0','1','1','1','1','1','EQ','input','',3,'','','','',''),(723,70,'address','地址','varchar(100)','string','Address','address','0','0','1','1','1','1','1','EQ','input','',4,'','','','',''),(724,70,'private_key','私钥','varchar(255)','string','PrivateKey','privateKey','0','0','1','1','1','1','1','EQ','input','',5,'','','','',''),(725,70,'created_at','','timestamp','Time','CreatedAt','createdAt','0','0','0','0','0','0','0','EQ','datetime','',6,'','','','',''),(726,70,'updated_at','','timestamp','Time','UpdatedAt','updatedAt','0','0','0','0','0','0','0','EQ','datetime','',7,'','','','',''),(727,71,'id','','int unsigned','','Id','id','1','1','0','0','0','0','0','EQ','','',1,'','','','',''),(728,71,'main_chain','主链','varchar(50)','string','MainChain','mainChain','0','0','1','1','1','1','1','EQ','input','',2,'','','','',''),(729,71,'name','币种名称','varchar(50)','string','Name','name','0','0','1','1','1','1','1','LIKE','input','',3,'','','','',''),(730,71,'contract_address','合约地址','varchar(255)','string','ContractAddress','contractAddress','0','0','','1','1','1','1','EQ','input','',4,'','','','',''),(731,71,'decimals','精度','int','int','Decimals','decimals','0','0','1','1','1','1','1','EQ','input','',5,'','','','',''),(732,71,'created_at','','timestamp','Time','CreatedAt','createdAt','0','0','0','0','0','0','0','EQ','datetime','',6,'','','','',''),(733,71,'updated_at','','timestamp','Time','UpdatedAt','updatedAt','0','0','0','0','0','0','0','EQ','datetime','',7,'','','','',''),(734,72,'id','','int unsigned','','Id','id','1','1','0','0','0','0','0','EQ','','',1,'','','','',''),(735,72,'main_chain','主链','varchar(50)','string','MainChain','mainChain','0','0','1','1','1','1','1','EQ','input','',2,'','','','',''),(736,72,'coin_name','手续费币种','varchar(50)','string','CoinName','coinName','0','0','1','1','1','1','1','LIKE','input','',3,'','','','',''),(737,72,'address','地址','varchar(255)','string','Address','address','0','0','1','1','1','1','1','EQ','input','',4,'','','','',''),(738,72,'amount','转移手续费','decimal(30,18)','float64','Amount','amount','0','0','1','1','1','1','1','EQ','input','',5,'','','','',''),(739,72,'hash','hash','varchar(255)','string','Hash','hash','0','0','1','1','1','1','1','EQ','input','',6,'','','','',''),(740,72,'status','状态 1-待上链 2-上链成功 3-上链失败','tinyint unsigned','','Status','status','0','0','1','1','1','1','1','EQ','radio','',7,'','','','',''),(741,72,'recharge_id','充值ID','int unsigned','','RechargeId','rechargeId','0','0','1','1','1','1','1','EQ','','',8,'','','','',''),(742,72,'created_at','','timestamp','Time','CreatedAt','createdAt','0','0','0','0','0','0','0','EQ','datetime','',9,'','','','',''),(743,72,'updated_at','','timestamp','Time','UpdatedAt','updatedAt','0','0','0','0','0','0','0','EQ','datetime','',10,'','','','',''),(744,73,'id','','int unsigned','','Id','id','1','1','0','0','0','0','0','EQ','','',1,'','','','',''),(745,73,'main_chain','主网','varchar(50)','string','MainChain','mainChain','0','0','1','1','1','1','1','EQ','input','',2,'','','','',''),(746,73,'height','块号','int','int','Height','height','0','0','1','1','1','1','1','EQ','input','',3,'','','','',''),(747,74,'id','id主键','bigint','int64','Id','id','1','1','0','0','0','0','0','EQ','input','',1,'','','','',''),(748,74,'main_chain','主链','varchar(50)','string','MainChain','mainChain','0','0','','1','1','1','1','EQ','input','',2,'','','','',''),(749,74,'block_hash','块hash','varchar(255)','string','BlockHash','blockHash','0','0','','1','1','1','1','EQ','input','',3,'','','','',''),(750,74,'coin_token','币种1','varchar(50)','string','CoinToken','coinToken','0','0','','1','1','1','1','EQ','input','',4,'','','','',''),(751,74,'coin_token1','币种2','varchar(50)','string','CoinToken1','coinToken1','0','0','','1','1','1','1','EQ','input','',5,'','','','',''),(752,74,'from_address','发送方地址','varchar(255)','string','FromAddress','fromAddress','0','0','','1','1','1','1','EQ','input','',6,'','','','',''),(753,74,'to_address','充币地址','varchar(255)','string','ToAddress','toAddress','0','0','','1','1','1','1','EQ','input','',7,'','','','',''),(754,74,'amount','充币数量','decimal(30,18)','float64','Amount','amount','0','0','','1','1','1','1','EQ','input','',8,'','','','',''),(755,74,'amount1','充币2数量','decimal(30,18) unsigned','float64','Amount1','amount1','0','0','','1','1','1','1','EQ','input','',9,'','','','',''),(756,74,'contract_address','币种1合约地址','varchar(255)','string','ContractAddress','contractAddress','0','0','','1','1','1','1','EQ','input','',10,'','','','',''),(757,74,'contract_address1','币种2合约地址','varchar(255)','string','ContractAddress1','contractAddress1','0','0','','1','1','1','1','EQ','input','',11,'','','','',''),(758,74,'hash','充币交易hashId','varchar(255)','string','Hash','hash','0','0','1','1','1','1','1','EQ','input','',12,'','','','',''),(759,74,'block_height','区块高度','varchar(255)','string','BlockHeight','blockHeight','0','0','','1','1','1','1','EQ','input','',13,'','','','',''),(760,74,'call_number','回调次数','smallint unsigned','','CallNumber','callNumber','0','0','','1','1','1','1','EQ','','',14,'','','','',''),(761,74,'status','状态，1充币成功待归集，2归集上链中，3归集成功, 4归集失败，5充值费用中','smallint','int','Status','status','0','0','1','1','1','1','1','EQ','radio','',15,'','','','',''),(762,74,'imputation_hash','归集hash','varchar(255)','string','ImputationHash','imputationHash','0','0','','1','1','1','1','EQ','input','',16,'','','','',''),(763,74,'remarks','备注','varchar(255)','string','Remarks','remarks','0','0','','1','1','1','1','EQ','input','',17,'','','','',''),(764,74,'recharge_type','充币数量 ','tinyint unsigned','','RechargeType','rechargeType','0','0','1','1','1','1','1','EQ','select','',18,'','','','',''),(765,74,'created_at','','timestamp','Time','CreatedAt','createdAt','0','0','0','0','0','0','0','EQ','datetime','',19,'','','','',''),(766,74,'updated_at','','timestamp','Time','UpdatedAt','updatedAt','0','0','0','0','0','0','0','EQ','datetime','',20,'','','','',''),(767,75,'id','','int unsigned','','Id','id','1','1','0','0','0','0','0','EQ','','',1,'','','','',''),(768,75,'main_chain','主链','varchar(50)','string','MainChain','mainChain','0','0','1','1','1','1','1','EQ','input','',2,'','','','',''),(769,75,'coin_token','币种 ','varchar(100)','string','CoinToken','coinToken','0','0','1','1','1','1','1','EQ','input','',3,'','','','',''),(770,75,'address','转出地址','varchar(255)','string','Address','address','0','0','1','1','1','1','1','EQ','input','',4,'','','','',''),(771,75,'amount','提币数量','decimal(30,18)','float64','Amount','amount','0','0','1','1','1','1','1','EQ','input','',5,'','','','',''),(772,75,'contract_address','合约地址','varchar(255)','string','ContractAddress','contractAddress','0','0','','1','1','1','1','EQ','input','',6,'','','','',''),(773,75,'status','状态 1待审核，2待上链，3上链中，4上链失败，5上链成功，6已拒绝','tinyint unsigned','','Status','status','0','0','1','1','1','1','1','EQ','radio','',7,'','','','',''),(774,75,'nonce','nonce','bigint unsigned','','Nonce','nonce','0','0','1','1','1','1','1','EQ','','',8,'','','','',''),(775,75,'confirmation_number','上链确认次数','smallint unsigned','','ConfirmationNumber','confirmationNumber','0','0','1','1','1','1','1','EQ','','',9,'','','','',''),(776,75,'hash','交易Hash','varchar(255)','string','Hash','hash','0','0','','1','1','1','1','EQ','input','',10,'','','','',''),(777,75,'remarks','交易备注','varchar(255)','string','Remarks','remarks','0','0','','1','1','1','1','EQ','input','',11,'','','','',''),(778,75,'notify_url','交易回调地址','varchar(255)','string','NotifyUrl','notifyUrl','0','0','','1','1','1','1','EQ','input','',12,'','','','',''),(779,75,'created_at','','timestamp','Time','CreatedAt','createdAt','0','0','0','0','0','0','0','EQ','datetime','',13,'','','','',''),(780,75,'updated_at','','timestamp','Time','UpdatedAt','updatedAt','0','0','0','0','0','0','0','EQ','datetime','',14,'','','','','');
/*!40000 ALTER TABLE `tools_gen_table_column` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `withdraw`
--

DROP TABLE IF EXISTS `withdraw`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `withdraw` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `main_chain` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '主链',
  `coin_token` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '币种 ',
  `withdraw_address` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '转出地址',
  `address` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '转入地址',
  `amount` decimal(30,18) NOT NULL COMMENT '提币数量',
  `contract_address` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '合约地址',
  `status` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '状态 1待审核，2待上链，3上链中，4上链失败，5上链成功，6已拒绝',
  `nonce` bigint unsigned NOT NULL DEFAULT '0' COMMENT 'nonce',
  `confirmation_number` smallint unsigned NOT NULL DEFAULT '0' COMMENT '上链确认次数',
  `nonce1` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT 'nonce1',
  `hashKey` text CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT 'hashKey',
  `hash` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '交易Hash',
  `remarks` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '交易备注',
  `notify_url` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '交易回调地址',
  `token_id` bigint DEFAULT NULL COMMENT 'tokenID (nft使用)',
  `url` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT 'url (nft使用)',
  `split_address` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT 'Solana提现实际分账地址',
  `split_amount` decimal(30,18) unsigned NOT NULL DEFAULT '0.000000000000000000' COMMENT 'Solana提现实际分账数量',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `withdraw`
--

LOCK TABLES `withdraw` WRITE;
/*!40000 ALTER TABLE `withdraw` DISABLE KEYS */;
/*!40000 ALTER TABLE `withdraw` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping events for database 'wallet'
--

--
-- Dumping routines for database 'wallet'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2026-02-08 19:46:55
