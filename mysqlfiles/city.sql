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

 Date: 11/30/2018 16:24:59 PM
*/

SET NAMES utf8;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
--  Table structure for `city`
-- ----------------------------
DROP TABLE IF EXISTS `city`;
CREATE TABLE `city` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `city_name` varchar(20) COLLATE utf8_bin NOT NULL DEFAULT '',
  `pin_yin` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `longitude` double NOT NULL DEFAULT '0',
  `latitude` double NOT NULL DEFAULT '0',
  `area_code` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `abbr` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
--  Records of `city`
-- ----------------------------
BEGIN;
INSERT INTO `city` VALUES ('1', '北京', 'Beijing', '116.407173', '39.90469', '010', 'BJ'), ('2', '天津', 'TianJin', '116.407173', '39.90469', '012', 'TJ'), ('3', '石家庄', 'ShiJiaZhuang', '116.407173', '39.90469', '0631', 'SJZ'), ('4', '唐山', 'TangShan', '116.407173', '39.90469', '0543', 'TS'), ('5', '大连', 'DaLian', '116.407173', '39.90469', '024', 'DL'), ('6', '长春', 'ChangChun', '116.407173', '39.90469', '0214', 'CHCH'), ('7', '哈尔滨', 'HaErBin', '116.407173', '39.90469', '0875', 'HEB'), ('8', '济南', 'JiNan', '114.407173', '37.90469', '0531', 'JN'), ('9', '烟台', 'YanTai', '114.407173', '37.90469', '0533', 'YT'), ('10', '淄博', 'ZiBo', '114.407173', '37.90469', '0539', 'ZB'), ('11', '青岛', 'QingDao', '114.407173', '41.90469', '0532', 'QD'), ('12', '泰安', 'TaiAn', '114.407173', '41.90469', '0538', 'TA');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
