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

 Date: 11/30/2018 16:24:54 PM
*/

SET NAMES utf8;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
--  Table structure for `admin`
-- ----------------------------
DROP TABLE IF EXISTS `admin`;
CREATE TABLE `admin` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_name` varchar(12) COLLATE utf8_bin NOT NULL DEFAULT '',
  `create_time` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `status` int(11) NOT NULL DEFAULT '0',
  `avatar` varchar(50) COLLATE utf8_bin NOT NULL DEFAULT '',
  `pwd` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `city_id` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
--  Records of `admin`
-- ----------------------------
BEGIN;
INSERT INTO `admin` VALUES ('1', 'root', '2018-11-28 21:45', '0', 'avatar1543499205678897862.jpg', '123', '1'), ('2', 'yhw', '2018-03-23 12:32', '0', '', '123', '2');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
