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

 Date: 11/30/2018 16:25:30 PM
*/

SET NAMES utf8;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
--  Table structure for `permission_admins`
-- ----------------------------
DROP TABLE IF EXISTS `permission_admins`;
CREATE TABLE `permission_admins` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `permission_id` int(11) NOT NULL,
  `admin_id` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
--  Records of `permission_admins`
-- ----------------------------
BEGIN;
INSERT INTO `permission_admins` VALUES ('1', '1', '1'), ('2', '2', '1'), ('3', '3', '1'), ('4', '4', '1'), ('5', '1', '2'), ('6', '4', '2'), ('7', '3', '2');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
