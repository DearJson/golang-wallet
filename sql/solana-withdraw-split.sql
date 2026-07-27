-- Solana 提现固定分账升级脚本（执行一次）
-- 仅 Solana 自动提现任务读取 currency 表中的分账配置。

ALTER TABLE `currency`
  ADD COLUMN `withdraw_split_enabled` tinyint unsigned NOT NULL DEFAULT '0' COMMENT 'Solana提现固定分账开关：0关闭，1开启' AFTER `min_merge`,
  ADD COLUMN `withdraw_split_address` varchar(255) DEFAULT NULL COMMENT 'Solana提现固定分账地址' AFTER `withdraw_split_enabled`,
  ADD COLUMN `withdraw_split_amount` decimal(30,18) unsigned NOT NULL DEFAULT '0.000000000000000000' COMMENT 'Solana提现固定分账数量' AFTER `withdraw_split_address`;

ALTER TABLE `withdraw`
  ADD COLUMN `split_address` varchar(255) DEFAULT NULL COMMENT 'Solana提现实际分账地址' AFTER `url`,
  ADD COLUMN `split_amount` decimal(30,18) unsigned NOT NULL DEFAULT '0.000000000000000000' COMMENT 'Solana提现实际分账数量' AFTER `split_address`;

-- 示例：将 Solana USDT 每笔提现固定分出 2 USDT。
-- UPDATE `currency`
-- SET `withdraw_split_enabled` = 1,
--     `withdraw_split_address` = '替换为固定收款地址',
--     `withdraw_split_amount` = 2
-- WHERE `main_chain` = 'solana'
--   AND `contract_address` = 'Es9vMFrzaCERmJfrF4H2FYD4KCoNkY11McCe8BenwNYB';
