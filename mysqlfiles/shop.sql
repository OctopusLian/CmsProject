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

 Date: 11/30/2018 16:25:36 PM
*/

SET NAMES utf8;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
--  Table structure for `shop`
-- ----------------------------
DROP TABLE IF EXISTS `shop`;
CREATE TABLE `shop` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) COLLATE utf8_bin NOT NULL DEFAULT '',
  `address` varchar(100) COLLATE utf8_bin NOT NULL DEFAULT '',
  `phone` varchar(11) COLLATE utf8_bin NOT NULL DEFAULT '',
  `status` int(11) NOT NULL DEFAULT '0',
  `recent_order_num` int(11) NOT NULL DEFAULT '0',
  `rating_count` int(11) NOT NULL DEFAULT '0',
  `rating` int(11) NOT NULL DEFAULT '0',
  `promotion_info` varchar(50) COLLATE utf8_bin NOT NULL DEFAULT '',
  `image_path` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `dele` int(11) NOT NULL DEFAULT '0',
  `latitude` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `longitude` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `description` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `float_delivery_fee` int(11) NOT NULL DEFAULT '0',
  `float_minimum_order_amount` int(11) NOT NULL DEFAULT '0',
  `is_premium` tinyint(1) NOT NULL DEFAULT '0',
  `delivery_mode` tinyint(1) NOT NULL DEFAULT '0',
  `new` tinyint(1) NOT NULL DEFAULT '0',
  `bao` tinyint(1) NOT NULL DEFAULT '0',
  `zhun` tinyint(1) NOT NULL DEFAULT '0',
  `piao` tinyint(1) NOT NULL DEFAULT '0',
  `start_time` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `end_time` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `business_license_image` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `catering_service_license_image` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `category` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
--  Records of `shop`
-- ----------------------------
BEGIN;
INSERT INTO `shop` VALUES ('1', '兰州牛肉拉面', '北京市海淀区西二旗北路', '13167583421', '1', '239', '5', '5', '', 'real.png', '1', '31.23146', '121.45722', '店铺简介-1', '0', '0', '0', '0', '0', '0', '0', '0', '05:45', '10:00', '', '', '异国料理/西餐'), ('2', '和平鸽饺子馆', '北京市朝阳区东湖渠', '18353123125', '1', '453', '6', '7', '', 'real.png', '0', '31.23146', '121.45722', '店铺简介1', '0', '0', '0', '0', '0', '0', '0', '0', '05:45', '10:00', '', '', '异国料理/西餐'), ('3', '店铺名称1', '上海市静安区北京西路', '13167582311', '1', '347', '887', '1', '店铺标语1', 'realWorld.png', '0', '31.23146', '121.45722', '店铺简介1', '5', '20', '1', '1', '1', '1', '1', '1', '05:45', '10:00', 'realWorld.png', 'realWorld.png', '异国料理/西餐'), ('4', '店铺名称2', '上海市静安区北京西路', '13167582311', '1', '347', '887', '1', '店铺标语2', 'realWorld.png', '0', '31.23146', '121.45722', '店铺简介2', '5', '20', '1', '1', '1', '1', '0', '1', '06:00', '13:15', 'realWorld.png', 'realWorld.png', '异国料理/披萨意面'), ('5', '店铺名称2', '上海市静安区北京西路', '13167582311', '1', '347', '887', '1', '店铺标语2', 'realWorld.png', '0', '31.23146', '121.45722', '店铺简介2', '5', '20', '1', '1', '1', '1', '0', '1', '06:00', '13:15', 'realWorld.png', 'realWorld.png', '异国料理/披萨意面'), ('6', '店铺名称2', '上海市静安区北京西路', '13167582311', '1', '347', '887', '1', '店铺标语2', 'realWorld.png', '0', '31.23146', '121.45722', '店铺简介2', '5', '20', '1', '1', '1', '1', '0', '1', '06:00', '13:15', 'realWorld.png', 'realWorld.png', '异国料理/披萨意面'), ('7', '店铺名称2', '上海市静安区北京西路', '13167582311', '1', '347', '887', '1', '店铺标语2', 'realWorld.png', '1', '31.23146', '121.45722', '店铺简介2', '5', '20', '1', '1', '1', '1', '0', '1', '06:00', '13:15', 'realWorld.png', 'realWorld.png', '异国料理/披萨意面'), ('8', '店铺名称2', '上海市静安区北京西路', '13167582311', '1', '347', '887', '1', '店铺标语2', 'realWorld.png', '0', '31.23146', '121.45722', '店铺简介2', '5', '20', '1', '1', '1', '1', '0', '1', '06:00', '13:15', 'realWorld.png', 'realWorld.png', '异国料理/披萨意面'), ('9', '店铺名称2', '上海市静安区北京西路', '13167582311', '1', '347', '887', '1', '店铺标语2', 'realWorld.png', '0', '31.23146', '121.45722', '店铺简介2', '5', '20', '1', '1', '1', '1', '0', '1', '06:00', '13:15', 'realWorld.png', 'realWorld.png', '异国料理/披萨意面'), ('10', '店铺名称3', '上海市静安区北京西路', '13167582311', '1', '347', '887', '1', '店铺标语3', 'realWorld.png', '0', '31.23146', '121.45722', '店铺简介3', '5', '20', '1', '1', '1', '1', '1', '1', '06:30', '09:00', 'realWorld.png', 'realWorld.png', '异国料理/披萨意面'), ('11', '店铺名称3', '上海市静安区北京西路', '13167582311', '1', '347', '887', '1', '店铺标语3', 'realWorld.png', '0', '31.23146', '121.45722', '店铺简介3', '5', '20', '1', '1', '1', '1', '1', '1', '06:30', '09:00', 'realWorld.png', 'realWorld.png', '异国料理/披萨意面'), ('12', '店铺名称3', '上海市静安区北京西路', '13167582311', '1', '347', '887', '1', '店铺标语3', 'realWorld.png', '0', '31.23146', '121.45722', '店铺简介3', '5', '20', '1', '1', '1', '1', '1', '1', '06:30', '09:00', 'realWorld.png', 'realWorld.png', '异国料理/披萨意面'), ('13', '店铺名称4', '上海市普陀区云岭西路583号', '13167582311', '1', '347', '887', '1', '店铺标语4', 'realWorld.png', '0', '31.23153', '121.3625', '店铺简介4', '5', '20', '1', '1', '1', '1', '1', '1', '06:30', '09:00', 'realWorld.png', 'realWorld.png', '异国料理/西餐'), ('14', '店铺名称4', '上海市普陀区云岭西路583号', '13167582311', '1', '347', '887', '1', '店铺标语4', 'realWorld.png', '0', '31.23153', '121.3625', '店铺简介4', '5', '20', '1', '1', '1', '1', '1', '1', '06:30', '09:00', 'realWorld.png', 'realWorld.png', '异国料理/西餐'), ('15', '店铺名称4', '上海市普陀区云岭西路583号', '13167582311', '1', '347', '887', '1', '店铺标语4', 'realWorld.png', '0', '31.23153', '121.3625', '店铺简介4', '5', '20', '1', '1', '1', '1', '1', '1', '06:30', '09:00', 'realWorld.png', 'realWorld.png', '异国料理/西餐'), ('16', '店铺名称6', '上海市浦东新区沪南路5191号-5', '13167584321', '1', '318', '81', '9', '店铺标语5', 'realWorld.png', '0', '31.0564', '121.58934', '店铺简介5', '6', '20', '1', '1', '1', '1', '1', '1', '05:45', '08:15', 'realWorld.png', 'realWorld.png', '异国料理/西餐'), ('17', '店铺名称7', '上海市长宁区水城南路19号古北鑫茂购物中心F1', '13167583421', '1', '347', '887', '1', '店铺标语7', 'realWorld.png', '0', '31.19412', '121.39312', '店铺简介7', '6', '20', '1', '1', '1', '1', '1', '1', '06:45', '10:00', 'realWorld.png', 'realWorld.png', '异国料理/披萨意面'), ('18', '店铺名称7', '上海市长宁区水城南路19号古北鑫茂购物中心F1', '13167583421', '1', '347', '887', '1', '店铺标语7', 'realWorld.png', '0', '31.19412', '121.39312', '店铺简介7', '6', '20', '1', '1', '1', '1', '1', '1', '06:45', '10:00', 'realWorld.png', 'realWorld.png', '异国料理/披萨意面'), ('19', '店铺名称7', '上海市长宁区水城南路19号古北鑫茂购物中心F1', '13167583421', '1', '347', '887', '1', '店铺标语7', 'realWorld.png', '0', '31.19412', '121.39312', '店铺简介7', '6', '20', '1', '1', '1', '1', '1', '1', '06:45', '10:00', 'realWorld.png', 'realWorld.png', '异国料理/披萨意面'), ('20', '店铺名称9', '上海市静安区北京西路', '13167582311', '1', '347', '887', '1', '店铺标语9', 'realWorld.png', '0', '31.23146', '121.45722', '店铺简介9', '5', '20', '1', '1', '1', '1', '1', '1', '06:15', '13:00', 'realWorld.png', 'realWorld.png', '异国料理/西餐'), ('21', '店铺名称10', '上海市黄浦区 ', '13167582345', '1', '347', '887', '1', '店铺标语10', 'realWorld.png', '0', '31.24007', '121.48185', '店铺简介10', '5', '20', '1', '1', '1', '1', '1', '1', '05:45', '06:30', 'realWorld.png', 'realWorld.png', '异国料理/披萨意面'), ('22', '店铺名称11', '上海市宝山区宝杨路1号国际邮轮码头海洋量子号5F', '13167582311', '1', '347', '887', '1', '店铺标语11', 'realWorld.png', '0', '31.403664', '121.5037', '店铺简介11', '5', '20', '1', '1', '1', '1', '1', '1', '05:30', '05:45', 'realWorld.png', 'realWorld.png', '异国料理/西餐'), ('23', '店铺名称12', '上海市黄浦区 ', '13167582341', '1', '347', '887', '1', '店铺标语12', 'realWorld.png', '0', '31.24007', '121.48185', '店铺简介12', '5', '20', '1', '1', '1', '1', '1', '1', '09:45', '13:45', 'realWorld.png', 'realWorld.png', '异国料理/西餐');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
