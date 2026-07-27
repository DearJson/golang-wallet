-- 已执行过 solana-withdraw-split.sql 旧版本时，只需执行本增量脚本一次。

ALTER TABLE `currency`
  ADD COLUMN `withdraw_split_bps` smallint unsigned NOT NULL DEFAULT '0' COMMENT 'Solana提现分账比例（基点，200=2%）' AFTER `withdraw_split_amount`;

-- PYTHIA 按每笔提现的 2% 分账。
UPDATE `currency`
SET `withdraw_split_enabled` = 1,
    `withdraw_split_address` = '6caDFNP77yPCxB6zbRgLnRaQEAvxFzh9FTPu1d19bmJW',
    `withdraw_split_amount` = 0,
    `withdraw_split_bps` = 200
WHERE `id` = 103
  AND `main_chain` = 'solana'
  AND `contract_address` = 'CreiuhfwdWCN5mJbMJtA9bBpYQrQF2tCBuZwSPWfpump';
