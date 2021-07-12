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

 Date: 11/30/2018 16:25:06 PM
*/

SET NAMES utf8;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
--  Table structure for `food`
-- ----------------------------
DROP TABLE IF EXISTS `food`;
CREATE TABLE `food` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `description` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `rating` int(11) NOT NULL DEFAULT '0',
  `month_sales` int(11) NOT NULL DEFAULT '0',
  `image_path` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `activity` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `attributes` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `specs` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `category_id` int(11) NOT NULL,
  `restaurant_id` int(11) NOT NULL,
  `del_flag` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
--  Records of `food`
-- ----------------------------
BEGIN;
INSERT INTO `food` VALUES ('1', '小炒肉盖饭', '非常好吃，性价比高', '5', '123', 'real.png', '满20减10', '', '', '1', '1', '0'), ('2', '韭菜鸡蛋盖饭', '性价比高', '6', '243', 'real.png', '满30减10', '', '', '1', '1', '0'), ('3', '食品名称1', '食品详情1', '1', '0', 'realWorld.png', '食品活动1', '', '', '1', '1', '0');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
