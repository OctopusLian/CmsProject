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

 Date: 11/30/2018 16:25:46 PM
*/

SET NAMES utf8;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
--  Table structure for `support_service_shops`
-- ----------------------------
DROP TABLE IF EXISTS `support_service_shops`;
CREATE TABLE `support_service_shops` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `support_service_id` int(11) NOT NULL,
  `shop_id` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
--  Records of `support_service_shops`
-- ----------------------------
BEGIN;
INSERT INTO `support_service_shops` VALUES ('4', '0', '19'), ('5', '0', '20'), ('6', '0', '20'), ('7', '0', '20'), ('8', '0', '21'), ('9', '0', '21'), ('10', '1', '22'), ('11', '2', '22'), ('12', '3', '23'), ('13', '4', '23'), ('14', '5', '23'), ('15', '6', '23');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
