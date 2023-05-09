/*
 Navicat Premium Data Transfer

 Source Server         : Klaatoo
 Source Server Type    : MariaDB
 Source Server Version : 101102 (10.11.2-MariaDB-1:10.11.2+maria~ubu2204)
 Source Host           : 192.168.219.107:13306
 Source Schema         : Inae

 Target Server Type    : MariaDB
 Target Server Version : 101102 (10.11.2-MariaDB-1:10.11.2+maria~ubu2204)
 File Encoding         : 65001

 Date: 28/04/2023 10:40:17
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for btc_address
-- ----------------------------
DROP TABLE IF EXISTS `btc_address`;
CREATE TABLE `btc_address` (
  `address` varchar(70) NOT NULL DEFAULT '',
  `prikey` varchar(70) NOT NULL DEFAULT '',
  `pubkey` varchar(70) NOT NULL DEFAULT '',
  `create_dt` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' ON UPDATE current_timestamp(),
  `active` enum('Y','N') NOT NULL DEFAULT 'Y',
  PRIMARY KEY (`address`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of btc_address
-- ----------------------------
BEGIN;
INSERT INTO `btc_address` (`address`, `prikey`, `pubkey`, `create_dt`, `active`) VALUES ('18cBEMRxXHqzWWCxZNtU91F5sbUNKhL5PX', '1', '1', '2023-04-13 06:29:32', 'Y');
COMMIT;

-- ----------------------------
-- Table structure for btc_history
-- ----------------------------
DROP TABLE IF EXISTS `btc_history`;
CREATE TABLE `btc_history` (
  `id` int(1) unsigned NOT NULL AUTO_INCREMENT,
  `height` int(1) NOT NULL DEFAULT 0,
  `hash` varchar(70) NOT NULL DEFAULT '',
  `txid` varchar(70) NOT NULL DEFAULT '',
  `address` varchar(70) NOT NULL DEFAULT '',
  `value` varchar(20) NOT NULL,
  `create_dt` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of btc_history
-- ----------------------------
BEGIN;
INSERT INTO `btc_history` (`id`, `height`, `hash`, `txid`, `address`, `value`, `create_dt`) VALUES (1, 785178, '00000000000000000000fdbb2038835f82d501b2656f0107b67d0516d95ca036', '99be97228204673424c38a938f30f9ab2ee9faf402df65f90df91aa3ceab7427', '18cBEMRxXHqzWWCxZNtU91F5sbUNKhL5PX', '6.66252037', '2023-04-13 06:29:53');
INSERT INTO `btc_history` (`id`, `height`, `hash`, `txid`, `address`, `value`, `create_dt`) VALUES (2, 785184, '0000000000000000000524fff35568bcca15f497abf86c6bfff355872fb4dfa0', 'a206a77438acbe231d888b80c15960403c733eead712a45d901bd0458cf5c561', '18cBEMRxXHqzWWCxZNtU91F5sbUNKhL5PX', '6.35241745', '2023-04-13 06:55:21');
INSERT INTO `btc_history` (`id`, `height`, `hash`, `txid`, `address`, `value`, `create_dt`) VALUES (3, 785726, '0000000000000000000042b9210bc7f77fc46d8daa33e53d0ca662380001acf7', 'dd2662115c5d5ca6d5d046f96c6ad90f90abb98aa0b2cec60a11023db2e147c4', '18cBEMRxXHqzWWCxZNtU91F5sbUNKhL5PX', '6.29651552', '2023-04-17 02:02:40');
INSERT INTO `btc_history` (`id`, `height`, `hash`, `txid`, `address`, `value`, `create_dt`) VALUES (4, 785747, '00000000000000000002c1cb9572e495eb490094f6e41898b37337140f1f482f', '3c13f1b8b08adeb7adbfe773464f926db8807bb1fa7a42b0779a113499da1c8b', '18cBEMRxXHqzWWCxZNtU91F5sbUNKhL5PX', '6.29291011', '2023-04-17 06:05:39');
INSERT INTO `btc_history` (`id`, `height`, `hash`, `txid`, `address`, `value`, `create_dt`) VALUES (5, 785751, '000000000000000000008d8b18721c075580bbdf80ad98d313530f283acd595b', 'b0295efd94b056a82c23c6b2ca720263f4794d06bc3301a13cd096d02e2eedfd', '18cBEMRxXHqzWWCxZNtU91F5sbUNKhL5PX', '6.26181774', '2023-04-17 06:10:38');
INSERT INTO `btc_history` (`id`, `height`, `hash`, `txid`, `address`, `value`, `create_dt`) VALUES (6, 785753, '0000000000000000000394c9bd3b8de9228e778c77ba1790416d58f603f646e8', 'b60bb25f62f5cbcb513c5fbef664483c31307b457e4d7515bf5905ca58e50416', '18cBEMRxXHqzWWCxZNtU91F5sbUNKhL5PX', '6.41930095', '2023-04-17 06:30:38');
INSERT INTO `btc_history` (`id`, `height`, `hash`, `txid`, `address`, `value`, `create_dt`) VALUES (7, 786913, '000000000000000000030378d9d0631989c2d67695f695ca584f036957050fd4', '1b1825db522fd31292b1ba2e3bc6821f84fdbb5728b5dacf7815314ac8685403', '18cBEMRxXHqzWWCxZNtU91F5sbUNKhL5PX', '6.34624738', '2023-04-25 06:18:13');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
