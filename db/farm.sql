-- MySQL dump 10.13  Distrib 5.7.34, for osx10.16 (x86_64)
--
-- Host: localhost    Database: farm
-- ------------------------------------------------------
-- Server version	5.7.34-log

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `farm_department`
--

DROP TABLE IF EXISTS `farm_department`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `farm_department` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `thirdparty_id` int(10) unsigned NOT NULL,
  `primary` varchar(128) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `hash` bigint(20) unsigned NOT NULL,
  `name` varchar(128) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `parent_id` int(10) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unq_thirdparty_primary` (`thirdparty_id`,`primary`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `farm_event`
--

DROP TABLE IF EXISTS `farm_event`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `farm_event` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `namespace` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '事件的所属命名空间',
  `payload` mediumtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '事件payload',
  `create_time` bigint(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_namespace` (`namespace`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `farm_filter`
--

DROP TABLE IF EXISTS `farm_filter`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `farm_filter` (
  `thirdparty_id` int(11) unsigned NOT NULL,
  `thirdparty_field` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `bucket1` varchar(4) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '一级桶做粗过滤',
  `bucket2` varchar(128) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '二级桶为bitmap确定是否存在1025bit（128字符）',
  `ids` json NOT NULL COMMENT '二级桶对应的ids',
  PRIMARY KEY (`thirdparty_id`,`thirdparty_field`,`bucket1`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `farm_subscriber`
--

DROP TABLE IF EXISTS `farm_subscriber`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `farm_subscriber` (
  `label` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '订阅者唯一标识',
  `offset` int(20) unsigned NOT NULL COMMENT '偏移(event_id)',
  `update_time` bigint(20) NOT NULL,
  PRIMARY KEY (`label`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `farm_thirdparty`
--

DROP TABLE IF EXISTS `farm_thirdparty`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `farm_thirdparty` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `label` varchar(128) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `namespace` varchar(16) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `enable` tinyint(4) NOT NULL DEFAULT '1',
  `config` json NOT NULL COMMENT 'config',
  `primary_attrs` json NOT NULL COMMENT '主键属性(唯一）',
  `index_attrs` json NOT NULL COMMENT '索引属性数组',
  `last_pull_users_hash` bigint(20) unsigned NOT NULL DEFAULT '0',
  `last_pull_depts_hash` bigint(20) unsigned NOT NULL DEFAULT '0',
  `update_time` int(11) NOT NULL,
  `create_time` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unq_label_namespace` (`label`,`namespace`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='三方配置表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `farm_user`
--

DROP TABLE IF EXISTS `farm_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `farm_user` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `thirdparty_id` int(11) unsigned NOT NULL,
  `primary` varchar(128) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '主键(通常由1或多个属性组成 )',
  `hash` bigint(20) unsigned NOT NULL,
  `column_num_1` bigint(20) DEFAULT NULL,
  `column_num_2` bigint(20) DEFAULT NULL,
  `column_num_3` bigint(20) DEFAULT NULL,
  `column_num_4` bigint(20) DEFAULT NULL,
  `column_num_5` bigint(20) DEFAULT NULL,
  `column_num_6` bigint(20) DEFAULT NULL,
  `column_num_7` bigint(20) DEFAULT NULL,
  `column_num_8` bigint(20) DEFAULT NULL,
  `column_num_9` bigint(20) DEFAULT NULL,
  `column_num_10` bigint(20) DEFAULT NULL,
  `column_num_11` bigint(20) DEFAULT NULL,
  `column_num_12` bigint(20) DEFAULT NULL,
  `column_num_13` bigint(20) DEFAULT NULL,
  `column_num_14` bigint(20) DEFAULT NULL,
  `column_num_15` bigint(20) DEFAULT NULL,
  `column_num_16` bigint(20) DEFAULT NULL,
  `column_num_17` bigint(20) DEFAULT NULL,
  `column_num_18` bigint(20) DEFAULT NULL,
  `column_num_19` bigint(20) DEFAULT NULL,
  `column_num_20` bigint(20) DEFAULT NULL,
  `column_num_21` bigint(20) DEFAULT NULL,
  `column_num_22` bigint(20) DEFAULT NULL,
  `column_num_23` bigint(20) DEFAULT NULL,
  `column_num_24` bigint(20) DEFAULT NULL,
  `column_num_25` bigint(20) DEFAULT NULL,
  `column_num_26` bigint(20) DEFAULT NULL,
  `column_num_27` bigint(20) DEFAULT NULL,
  `column_num_28` bigint(20) DEFAULT NULL,
  `column_num_29` bigint(20) DEFAULT NULL,
  `column_num_30` bigint(20) DEFAULT NULL,
  `column_text_1` text COLLATE utf8mb4_bin,
  `column_text_2` text COLLATE utf8mb4_bin,
  `column_text_3` text COLLATE utf8mb4_bin,
  `column_text_4` text COLLATE utf8mb4_bin,
  `column_text_5` text COLLATE utf8mb4_bin,
  `column_text_6` text COLLATE utf8mb4_bin,
  `column_text_7` text COLLATE utf8mb4_bin,
  `column_text_8` text COLLATE utf8mb4_bin,
  `column_text_9` text COLLATE utf8mb4_bin,
  `column_text_10` text COLLATE utf8mb4_bin,
  `column_text_11` text COLLATE utf8mb4_bin,
  `column_text_12` text COLLATE utf8mb4_bin,
  `column_text_13` text COLLATE utf8mb4_bin,
  `column_text_14` text COLLATE utf8mb4_bin,
  `column_text_15` text COLLATE utf8mb4_bin,
  `column_text_16` text COLLATE utf8mb4_bin,
  `column_text_17` text COLLATE utf8mb4_bin,
  `column_text_18` text COLLATE utf8mb4_bin,
  `column_text_19` text COLLATE utf8mb4_bin,
  `column_text_20` text COLLATE utf8mb4_bin,
  `column_text_21` text COLLATE utf8mb4_bin,
  `column_text_22` text COLLATE utf8mb4_bin,
  `column_text_23` text COLLATE utf8mb4_bin,
  `column_text_24` text COLLATE utf8mb4_bin,
  `column_text_25` text COLLATE utf8mb4_bin,
  `column_text_26` text COLLATE utf8mb4_bin,
  `column_text_27` text COLLATE utf8mb4_bin,
  `column_text_28` text COLLATE utf8mb4_bin,
  `column_text_29` text COLLATE utf8mb4_bin,
  `column_text_30` text COLLATE utf8mb4_bin,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unq_thirdparty_primary` (`thirdparty_id`,`primary`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='三方用户表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `farm_user_metadata`
--

DROP TABLE IF EXISTS `farm_user_metadata`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `farm_user_metadata` (
  `thirdparty_id` int(11) unsigned NOT NULL,
  `thirdparty_field` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '三方用属性名',
  `column_field` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '数据表内属性名',
  KEY `idx_thirdparty` (`thirdparty_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='三方用户属性与farm_user中的对应关系描述';
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-10-12 11:57:03
