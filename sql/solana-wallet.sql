-- Solana 链初始化SQL
-- 执行前请确认数据库已存在 currency、sys_config 等表

-- 添加SOL原生代币
INSERT INTO `currency` (`main_chain`, `name`, `contract_address`, `decimals`, `min_withdraw`)
VALUES ('solana', 'SOL', 'So11111111111111111111111111111111111111112', 9, 0);

-- 添加系统配置（提现地址、手续费地址等，需手动配置实际值）
INSERT INTO `sys_config` (`config_key`, `config_value`, `config_name`, `config_type`, `remark`)
VALUES ('sys.solanaWithdrawAddress', '', 'Solana提现地址',1, 'Solana出金钱包地址');

INSERT INTO `sys_config` (`config_key`, `config_value`, `config_name`, `config_type`, `remark`)
VALUES ('sys.solanaWithdrawAddressPrivateKey', '', 'Solana提现私钥',1, 'Solana出金钱包私钥');

INSERT INTO `sys_config` (`config_key`, `config_value`, `config_name`, `config_type`, `remark`)
VALUES ('sys.solanaFeeAddress', '', 'Solana手续费地址',1, 'Solana手续费代付钱包地址');

INSERT INTO `sys_config` (`config_key`, `config_value`, `config_name`, `config_type`, `remark`)
VALUES ('sys.solanaFeeAddressPrivateKey', '', 'Solana手续费私钥',1, 'Solana手续费代付钱包私钥');

INSERT INTO `sys_config` (`config_key`, `config_value`, `config_name`, `config_type`, `remark`)
VALUES ('sys.solanaMergeAddress', '', 'Solana归集地址',1, 'Solana归集目标钱包地址');

-- ==================== 后台菜单 sys_auth_rule ====================
-- ID从400开始，避免与现有链菜单冲突
-- 字段顺序: id, pid, name, title, icon, condition, remark, menu_type, weigh, is_hide, path, component, is_link, module_type, model_id, is_iframe, is_cached, redirect, is_affix, link_url

-- Solana钱包管理 (顶级目录 menu_type=0)
INSERT INTO `sys_auth_rule` VALUES (400, 0, 'solanaWallet', 'Solana钱包管理', 'chart', '', '', 0, 2, 1, 1, 'solanaWallet', '', '', 0, 'sys_admin', 0, NOW(), NOW(), NULL);

-- 地址管理 (菜单 menu_type=1, pid=400)
INSERT INTO `sys_auth_rule` VALUES (401, 400, 'system/solanaAddress/list', '地址管理', '', '', '', 1, 10, 1, 1, 'solanaAddress', '', 'system/solanaAddress/list', 0, 'sys_admin', 0, NOW(), NOW(), NULL);
INSERT INTO `sys_auth_rule` VALUES (402, 401, 'system/solanaAddress/get', '查询', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NOW(), NOW(), NULL);

-- 币种管理 (菜单 menu_type=1, pid=400)
INSERT INTO `sys_auth_rule` VALUES (403, 400, 'system/solanaCurrency/list', '币种管理', '', '', '', 1, 9, 1, 1, 'solanaCurrency', '', 'system/solanaCurrency/list', 0, 'sys_admin', 0, NOW(), NOW(), NULL);
INSERT INTO `sys_auth_rule` VALUES (404, 403, 'system/solanaCurrency/get', '查询', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NOW(), NOW(), NULL);
INSERT INTO `sys_auth_rule` VALUES (405, 403, 'system/solanaCurrency/add', '添加', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NOW(), NOW(), NULL);
INSERT INTO `sys_auth_rule` VALUES (406, 403, 'system/solanaCurrency/edit', '修改', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NOW(), NOW(), NULL);
INSERT INTO `sys_auth_rule` VALUES (407, 403, 'system/solanaCurrency/delete', '删除', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NOW(), NOW(), NULL);

-- 手续费管理 (菜单 menu_type=1, pid=400)
INSERT INTO `sys_auth_rule` VALUES (408, 400, 'system/solanaFeelist/list', '手续费管理', '', '', '', 1, 8, 1, 1, 'solanaFeeList', '', 'system/solanaFeeList/list', 0, 'sys_admin', 0, NOW(), NOW(), NULL);
INSERT INTO `sys_auth_rule` VALUES (409, 408, 'system/solanaFeelist/get', '查询', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NOW(), NOW(), NULL);

-- 充值管理 (菜单 menu_type=1, pid=400)
INSERT INTO `sys_auth_rule` VALUES (410, 400, 'system/solanaRecharge/list', '充值管理', '', '', '', 1, 7, 1, 1, 'solanaRechargeList', '', 'system/solanaRecharge/list', 0, 'sys_admin', 0, NOW(), NOW(), NULL);
INSERT INTO `sys_auth_rule` VALUES (411, 410, 'system/solanaRecharge/get', '查询', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NOW(), NOW(), NULL);
INSERT INTO `sys_auth_rule` VALUES (412, 410, 'system/solanaRecharge/callback', '发起回调', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NOW(), NOW(), NULL);

-- 提现管理 (菜单 menu_type=1, pid=400)
INSERT INTO `sys_auth_rule` VALUES (413, 400, 'system/solanaWithdraw/list', '提现管理', '', '', '', 1, 6, 1, 1, 'solanaWithdrawList', '', 'system/solanaWithdraw/list', 0, 'sys_admin', 0, NOW(), NOW(), NULL);
INSERT INTO `sys_auth_rule` VALUES (414, 413, 'system/solanaWithdraw/get', '查询', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NOW(), NOW(), NULL);
INSERT INTO `sys_auth_rule` VALUES (415, 413, 'system/solanaWithdraw/callback', '发起回调', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NOW(), NOW(), NULL);
INSERT INTO `sys_auth_rule` VALUES (416, 413, 'system/solanaWithdraw/auditSuccess', '审核通过', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NOW(), NOW(), NULL);
INSERT INTO `sys_auth_rule` VALUES (417, 413, 'system/solanaWithdraw/auditFail', '审核拒绝', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NOW(), NOW(), NULL);
INSERT INTO `sys_auth_rule` VALUES (418, 413, 'system/solanaWithdraw/withdraw', '手动提现', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NOW(), NOW(), NULL);

-- ==================== 常用SPL Token（按需取消注释）====================
-- USDC
-- INSERT INTO `currency` (`main_chain`, `name`, `contract_address`, `decimals`, `min_withdraw`, `status`)
-- VALUES ('solana', 'USDC', 'EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v', 6, 0, 1);

-- USDT
-- INSERT INTO `currency` (`main_chain`, `name`, `contract_address`, `decimals`, `min_withdraw`, `status`)
-- VALUES ('solana', 'USDT', 'Es9vMFrzaCERmJfrF4H2FYD4KCoNkY11McCe8BenwNYB', 6, 0, 1);
