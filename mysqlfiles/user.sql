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

 Date: 11/30/2018 16:25:51 PM
*/

SET NAMES utf8;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
--  Table structure for `user`
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_name` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `registe_time` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `mobile` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `is_active` int(11) NOT NULL DEFAULT '0',
  `balance` int(11) NOT NULL DEFAULT '0',
  `avatar` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `city_id` int(11) NOT NULL,
  `user_order_id` int(11) NOT NULL,
  `pwd` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `del_flag` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
--  Records of `user`
-- ----------------------------
BEGIN;
INSERT INTO `user` VALUES ('1', 'yuxinburen', '2018-09-10 10:31', '13165872311', '1', '23', '11.png', '2', '0', '123', '0'), ('2', 'yhw', '2018-08-20 23:41', '13406208437', '1', '456', 'realworld.png', '4', '0', '123', '0');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
