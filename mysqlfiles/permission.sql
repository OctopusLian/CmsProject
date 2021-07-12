/*
 Navicat MySQL Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50724
 Source Host           : localhost
 Source Database       : elmcms

 Target Server Type    : MySQL
 Target Server Version : 50724
 File Encoding         : utf-8

 Date: 11/30/2018 16:25:24 PM
*/

SET NAMES utf8;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
--  Table structure for `permission`
-- ----------------------------
DROP TABLE IF EXISTS `permission`;
CREATE TABLE `permission` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `level` varchar(30) COLLATE utf8_bin NOT NULL DEFAULT '',
  `name` varchar(20) COLLATE utf8_bin NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
--  Records of `permission`
-- ----------------------------
BEGIN;
INSERT INTO `permission` VALUES ('1', 'ADD', '增加数据'), ('2', 'DELETE', '删除数据'), ('3', 'UPDATE', '修改数据'), ('4', 'QUERY', '查询数据');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
