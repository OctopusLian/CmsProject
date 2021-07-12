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

 Date: 11/30/2018 16:25:41 PM
*/

SET NAMES utf8;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
--  Table structure for `support_service`
-- ----------------------------
DROP TABLE IF EXISTS `support_service`;
CREATE TABLE `support_service` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `icon_name` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `icon_color` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `description` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
--  Records of `support_service`
-- ----------------------------
BEGIN;
INSERT INTO `support_service` VALUES ('1', '满减优惠', '', '', '满30减5，满60减8'), ('2', '满减优惠', '', '', 'fasd '), ('3', '满减优惠', '', '', '满30减5，满60减8'), ('4', '优惠大酬宾', '', '', '哈'), ('5', '新用户立减', '', '', '立马'), ('6', '进店领券', '', '', '灵犬');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
