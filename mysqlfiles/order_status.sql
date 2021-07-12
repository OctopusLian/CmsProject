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

 Date: 11/30/2018 16:25:17 PM
*/

SET NAMES utf8;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
--  Table structure for `order_status`
-- ----------------------------
DROP TABLE IF EXISTS `order_status`;
CREATE TABLE `order_status` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `status_id` int(11) NOT NULL DEFAULT '0',
  `status_desc` varchar(100) COLLATE utf8_bin NOT NULL DEFAULT '',
  `user_order_id` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
--  Records of `order_status`
-- ----------------------------
BEGIN;
INSERT INTO `order_status` VALUES ('1', '1001', '未支付', '1'), ('2', '1002', '已支付', '1'), ('3', '1003', '已发货', '1'), ('4', '1004', '正在配送', '1'), ('5', '1005', '已接收', '1'), ('6', '1006', '发起退款', '1'), ('7', '1007', '正在退款', '1'), ('8', '1008', '取消订单', '1');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
