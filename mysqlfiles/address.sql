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

 Date: 11/30/2018 16:24:49 PM
*/

SET NAMES utf8;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
--  Table structure for `address`
-- ----------------------------
DROP TABLE IF EXISTS `address`;
CREATE TABLE `address` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `address` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `phone` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `address_detail` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `is_valid` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
--  Records of `address`
-- ----------------------------
BEGIN;
INSERT INTO `address` VALUES ('1', '芜湖碧桂园', '13167582311', '2123', '1'), ('2', '北京市海淀区西二旗', '13167582311', '北京市海淀区西二旗智学院9号院', '1'), ('3', '泰安市', '13167582311', '泰安市新泰市', '1'), ('4', '济南市', '13167582311', '山东大学', '1'), ('5', '北京市', '13167582311', '领秀新硅谷', '1'), ('6', '芜湖碧桂园6', '13167582311', '2123', '1'), ('7', '北京市海淀区西二旗7', '13167582311', '北京市海淀区西二旗智学院9号院', '1'), ('8', '泰安市', '13167582311', '泰安市新泰市9', '1'), ('9', '济南市9', '13167582311', '山东大学', '1'), ('10', '北京市', '13167582311', '领秀新硅谷10', '1');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
