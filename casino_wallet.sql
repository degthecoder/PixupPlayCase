/*
 Navicat Premium Dump SQL

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 80405 (8.4.5)
 Source Host           : localhost:3306
 Source Schema         : casino_wallet

 Target Server Type    : MySQL
 Target Server Version : 80405 (8.4.5)
 File Encoding         : 65001

 Date: 19/06/2025 16:53:12
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for players
-- ----------------------------
DROP TABLE IF EXISTS `players`;
CREATE TABLE `players` (
  `player_id` varchar(64) NOT NULL,
  `wallet_id` varchar(64) NOT NULL,
  `balance` decimal(18,2) NOT NULL DEFAULT '0.00',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`player_id`),
  UNIQUE KEY `wallet_id` (`wallet_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of players
-- ----------------------------
BEGIN;
INSERT INTO `players` (`player_id`, `wallet_id`, `balance`, `created_at`) VALUES ('0ffeab7', '2ff31b4', 845.00, '2025-06-19 16:39:58');
INSERT INTO `players` (`player_id`, `wallet_id`, `balance`, `created_at`) VALUES ('67fff5c', '684038d', 968.00, '2025-06-19 15:54:56');
INSERT INTO `players` (`player_id`, `wallet_id`, `balance`, `created_at`) VALUES ('7da93e2', '334bac0', 1030.00, '2025-06-19 16:39:58');
INSERT INTO `players` (`player_id`, `wallet_id`, `balance`, `created_at`) VALUES ('a99eb01', 'd1fbe42', 1120.00, '2025-06-19 16:39:58');
INSERT INTO `players` (`player_id`, `wallet_id`, `balance`, `created_at`) VALUES ('b12c7d3', '71f94ee', 875.00, '2025-06-19 16:39:58');
INSERT INTO `players` (`player_id`, `wallet_id`, `balance`, `created_at`) VALUES ('beaa718', '1b1c324', 915.00, '2025-06-19 16:39:58');
INSERT INTO `players` (`player_id`, `wallet_id`, `balance`, `created_at`) VALUES ('c6de48a', '8dffa72', 780.00, '2025-06-19 16:39:58');
INSERT INTO `players` (`player_id`, `wallet_id`, `balance`, `created_at`) VALUES ('dff0ca4', 'cf2d9e7', 1100.00, '2025-06-19 16:39:58');
INSERT INTO `players` (`player_id`, `wallet_id`, `balance`, `created_at`) VALUES ('e3fa1b9', '9aa321c', 1025.00, '2025-06-19 16:39:58');
INSERT INTO `players` (`player_id`, `wallet_id`, `balance`, `created_at`) VALUES ('f4bc912', 'ab183d1', 990.00, '2025-06-19 16:39:58');
INSERT INTO `players` (`player_id`, `wallet_id`, `balance`, `created_at`) VALUES ('player1', 'wallet1', 1000.00, '2025-06-19 15:02:26');
INSERT INTO `players` (`player_id`, `wallet_id`, `balance`, `created_at`) VALUES ('player2', 'wallet2', 850.00, '2025-06-19 15:02:26');
INSERT INTO `players` (`player_id`, `wallet_id`, `balance`, `created_at`) VALUES ('player3', 'wallet3', 1200.00, '2025-06-19 15:02:26');
INSERT INTO `players` (`player_id`, `wallet_id`, `balance`, `created_at`) VALUES ('player4', 'wallet4', 500.00, '2025-06-19 15:02:26');
INSERT INTO `players` (`player_id`, `wallet_id`, `balance`, `created_at`) VALUES ('player5', 'wallet5', 1300.00, '2025-06-19 15:02:26');
COMMIT;

-- ----------------------------
-- Table structure for transactions
-- ----------------------------
DROP TABLE IF EXISTS `transactions`;
CREATE TABLE `transactions` (
  `id` int NOT NULL AUTO_INCREMENT,
  `req_id` varchar(64) NOT NULL,
  `type` enum('bet','result') NOT NULL,
  `player_id` varchar(64) NOT NULL,
  `wallet_id` varchar(64) NOT NULL,
  `round_id` varchar(64) NOT NULL,
  `session_id` varchar(64) NOT NULL,
  `game_code` varchar(64) DEFAULT NULL,
  `amount` decimal(18,2) NOT NULL,
  `currency` varchar(8) NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `req_id` (`req_id`),
  KEY `player_id` (`player_id`),
  KEY `idx_round_type` (`round_id`,`type`),
  CONSTRAINT `transactions_ibfk_1` FOREIGN KEY (`player_id`) REFERENCES `players` (`player_id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of transactions
-- ----------------------------
BEGIN;
INSERT INTO `transactions` (`id`, `req_id`, `type`, `player_id`, `wallet_id`, `round_id`, `session_id`, `game_code`, `amount`, `currency`, `created_at`) VALUES (1, '42af0cb0-a016-4b55-a02f', 'bet', '67fff5c', '684038d', '3bb83e87-e540-4dd1-a3e8', '37717449-fa2f-4f4f-9f03', 'ntn_aloha', 10.00, 'INR', '2025-06-19 16:00:34');
INSERT INTO `transactions` (`id`, `req_id`, `type`, `player_id`, `wallet_id`, `round_id`, `session_id`, `game_code`, `amount`, `currency`, `created_at`) VALUES (2, 'c300ff2d-f7d1-43dc-b054', 'result', '67fff5c', '684038d', '3bb83e87-e540-4dd1-a3e8', '37717449-fa2f-4f4f-9f03', 'ntn_aloha', 8.00, 'INR', '2025-06-19 16:03:01');
INSERT INTO `transactions` (`id`, `req_id`, `type`, `player_id`, `wallet_id`, `round_id`, `session_id`, `game_code`, `amount`, `currency`, `created_at`) VALUES (3, 'req-bet-1', 'bet', '67fff5c', '684038d', 'round-123', 'session-abc', 'ntn_spin', 50.00, 'INR', '2025-06-19 16:45:31');
INSERT INTO `transactions` (`id`, `req_id`, `type`, `player_id`, `wallet_id`, `round_id`, `session_id`, `game_code`, `amount`, `currency`, `created_at`) VALUES (4, 'req-res-1', 'result', '67fff5c', '684038d', 'round-123', 'session-abc', 'ntn_spin', 20.00, 'INR', '2025-06-19 16:45:36');
INSERT INTO `transactions` (`id`, `req_id`, `type`, `player_id`, `wallet_id`, `round_id`, `session_id`, `game_code`, `amount`, `currency`, `created_at`) VALUES (5, 'req-bet-2', 'bet', 'e3fa1b9', '9aa321c', 'round-456', 'session-def', 'ntn_aloha', 75.00, 'INR', '2025-06-19 16:45:59');
INSERT INTO `transactions` (`id`, `req_id`, `type`, `player_id`, `wallet_id`, `round_id`, `session_id`, `game_code`, `amount`, `currency`, `created_at`) VALUES (6, 'req-res-2', 'result', 'e3fa1b9', '9aa321c', 'round-456', 'session-def', 'ntn_aloha', 150.00, 'INR', '2025-06-19 16:46:17');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
