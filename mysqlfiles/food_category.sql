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

 Date: 11/30/2018 16:25:13 PM
*/

SET NAMES utf8;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
--  Table structure for `food_category`
-- ----------------------------
DROP TABLE IF EXISTS `food_category`;
CREATE TABLE `food_category` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `category_name` varchar(32) COLLATE utf8_bin NOT NULL DEFAULT '',
  `category_desc` varchar(200) COLLATE utf8_bin NOT NULL DEFAULT '',
  `restaurant_id` int(11) NOT NULL,
  `level` int(11) NOT NULL DEFAULT '0',
  `parent_category_id` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
--  Records of `food_category`
-- ----------------------------
BEGIN;
INSERT INTO `food_category` VALUES ('1', '异国料理', '外国人吃饭的口味', '1', '1', '0'), ('2', '日韩料理', '日本和韩国的料理', '2', '2', '1'), ('3', '西餐', '西方人吃的正餐叫西餐', '3', '2', '1'), ('4', '披萨意面', '意大利的一种面', '1', '2', '1'), ('5', '食品种类11', '种类描述11', '0', '0', '0');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
