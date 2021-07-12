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

 Date: 11/30/2018 16:25:56 PM
*/

SET NAMES utf8;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
--  Table structure for `user_order`
-- ----------------------------
DROP TABLE IF EXISTS `user_order`;
CREATE TABLE `user_order` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `sum_money` int(11) NOT NULL DEFAULT '0',
  `del_flag` int(11) NOT NULL DEFAULT '0',
  `order_status_id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `shop_id` int(11) NOT NULL,
  `time` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `order_time` bigint(20) unsigned NOT NULL DEFAULT '0',
  `address_id` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
--  Records of `user_order`
-- ----------------------------
BEGIN;
INSERT INTO `user_order` VALUES ('1', '2018', '0', '1', '1', '1', '2018-11-28 20:28', '1543334500000', '1'), ('4', '3421', '0', '2', '2', '2', '2018-11-18 20:28', '1542544118420', '1');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
