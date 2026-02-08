package rpc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gogf/gf/frame/g"
)

const HeliusBaseURL = "https://api-mainnet.helius-rpc.com"

// HeliusClient Helius API 客户端
type HeliusClient struct {
	ApiKey string
}

// NewHeliusClient 创建Helius客户端
func NewHeliusClient() *HeliusClient {
	apiKey := g.Config().GetString("helius.api_key")
	return &HeliusClient{ApiKey: apiKey}
}

// ==================== Webhook 管理 ====================

// HeliusWebhookConfig Webhook配置
type HeliusWebhookConfig struct {
	WebhookID        string   `json:"webhookID,omitempty"`
	WebhookURL       string   `json:"webhookURL"`
	TransactionTypes []string `json:"transactionTypes"`
	AccountAddresses []string `json:"accountAddresses"`
	WebhookType      string   `json:"webhookType"`
	AuthHeader       string   `json:"authHeader,omitempty"`
	Encoding         string   `json:"encoding,omitempty"`  // 编码格式（可选）
	TxnStatus        string   `json:"txnStatus,omitempty"` // 交易状态过滤: "success" 仅成功交易
}

// HeliusWebhookResponse Webhook API响应
type HeliusWebhookResponse struct {
	WebhookID        string   `json:"webhookID"`
	Wallet           string   `json:"wallet"`
	WebhookURL       string   `json:"webhookURL"`
	TransactionTypes []string `json:"transactionTypes"`
	AccountAddresses []string `json:"accountAddresses"`
	WebhookType      string   `json:"webhookType"`
	AuthHeader       string   `json:"authHeader"`
	Encoding         string   `json:"encoding,omitempty"`
	TxnStatus        string   `json:"txnStatus,omitempty"`
}

// CreateWebhook 创建Webhook
func (c *HeliusClient) CreateWebhook(config *HeliusWebhookConfig) (*HeliusWebhookResponse, error) {
	url := fmt.Sprintf("%s/v0/webhooks?api-key=%s", HeliusBaseURL, c.ApiKey)

	jsonData, err := json.Marshal(config)
	if err != nil {
		return nil, fmt.Errorf("marshal config error: %v", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("create webhook request error: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response error: %v", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("create webhook failed [%d]: %s", resp.StatusCode, string(body))
	}

	var result HeliusWebhookResponse
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response error: %v", err)
	}
	return &result, nil
}

// GetWebhook 获取Webhook信息
func (c *HeliusClient) GetWebhook(webhookID string) (*HeliusWebhookResponse, error) {
	url := fmt.Sprintf("%s/v0/webhooks/%s?api-key=%s", HeliusBaseURL, webhookID, c.ApiKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("get webhook request error: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response error: %v", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("get webhook failed [%d]: %s", resp.StatusCode, string(body))
	}

	var result HeliusWebhookResponse
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response error: %v", err)
	}
	return &result, nil
}

// UpdateWebhook 更新Webhook（用于增删监听地址）
func (c *HeliusClient) UpdateWebhook(webhookID string, config *HeliusWebhookConfig) (*HeliusWebhookResponse, error) {
	url := fmt.Sprintf("%s/v0/webhooks/%s?api-key=%s", HeliusBaseURL, webhookID, c.ApiKey)

	jsonData, err := json.Marshal(config)
	if err != nil {
		return nil, fmt.Errorf("marshal config error: %v", err)
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("create update request error: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("update webhook request error: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response error: %v", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("update webhook failed [%d]: %s", resp.StatusCode, string(body))
	}

	var result HeliusWebhookResponse
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response error: %v", err)
	}
	return &result, nil
}

// DeleteWebhook 删除Webhook
func (c *HeliusClient) DeleteWebhook(webhookID string) error {
	url := fmt.Sprintf("%s/v0/webhooks/%s?api-key=%s", HeliusBaseURL, webhookID, c.ApiKey)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return fmt.Errorf("create delete request error: %v", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("delete webhook request error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("delete webhook failed [%d]: %s", resp.StatusCode, string(body))
	}
	return nil
}

// AppendAddressesToWebhook 向已有Webhook追加监听地址
func (c *HeliusClient) AppendAddressesToWebhook(webhookID string, newAddresses []string) error {
	// 先获取现有配置
	existing, err := c.GetWebhook(webhookID)
	if err != nil {
		return fmt.Errorf("get existing webhook error: %v", err)
	}

	// 合并地址（去重）
	addressMap := make(map[string]bool)
	for _, addr := range existing.AccountAddresses {
		addressMap[addr] = true
	}
	for _, addr := range newAddresses {
		addressMap[addr] = true
	}

	allAddresses := make([]string, 0, len(addressMap))
	for addr := range addressMap {
		allAddresses = append(allAddresses, addr)
	}

	// 更新Webhook
	config := &HeliusWebhookConfig{
		WebhookURL:       existing.WebhookURL,
		TransactionTypes: existing.TransactionTypes,
		AccountAddresses: allAddresses,
		WebhookType:      existing.WebhookType,
	}
	_, err = c.UpdateWebhook(webhookID, config)
	return err
}

// RemoveAddressesFromWebhook 从Webhook移除监听地址
func (c *HeliusClient) RemoveAddressesFromWebhook(webhookID string, removeAddresses []string) error {
	// 先获取现有配置
	existing, err := c.GetWebhook(webhookID)
	if err != nil {
		return fmt.Errorf("get existing webhook error: %v", err)
	}

	// 移除指定地址
	removeMap := make(map[string]bool)
	for _, addr := range removeAddresses {
		removeMap[addr] = true
	}

	remaining := make([]string, 0)
	for _, addr := range existing.AccountAddresses {
		if !removeMap[addr] {
			remaining = append(remaining, addr)
		}
	}

	// 更新Webhook
	config := &HeliusWebhookConfig{
		WebhookURL:       existing.WebhookURL,
		TransactionTypes: existing.TransactionTypes,
		AccountAddresses: remaining,
		WebhookType:      existing.WebhookType,
	}
	_, err = c.UpdateWebhook(webhookID, config)
	return err
}

// ParseTransactions 通过Helius Enhanced API解析交易（返回与Webhook回调相同格式的增强数据）
// signatures: 交易签名列表（最多100个）
func (c *HeliusClient) ParseTransactions(signatures []string) ([]HeliusEnhancedTransaction, error) {
	url := fmt.Sprintf("%s/v0/transactions?api-key=%s", HeliusBaseURL, c.ApiKey)

	jsonData, err := json.Marshal(map[string]interface{}{
		"transactions": signatures,
	})
	if err != nil {
		return nil, fmt.Errorf("marshal request error: %v", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("parse transactions request error: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response error: %v", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("parse transactions failed [%d]: %s", resp.StatusCode, string(body))
	}

	var result []HeliusEnhancedTransaction
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response error: %v", err)
	}
	return result, nil
}

// EnsureWebhookExists 确保Webhook存在，不存在则创建
// webhookURL: 回调地址 (如 https://your-domain.com/webhook/solana)
// addresses: 需要监听的地址列表
func (c *HeliusClient) EnsureWebhookExists(webhookURL string, addresses []string) (string, error) {
	webhookID := g.Config().GetString("helius.webhook_id")
	webhookSecret := g.Config().GetString("helius.webhook_secret")

	if webhookID != "" {
		// 检查Webhook是否还存在
		_, err := c.GetWebhook(webhookID)
		if err == nil {
			// 更新地址列表
			config := &HeliusWebhookConfig{
				WebhookURL:       webhookURL,
				TransactionTypes: []string{"ANY"},
				AccountAddresses: addresses,
				WebhookType:      "enhanced",
				AuthHeader:       webhookSecret,
				TxnStatus:        "success",
			}
			_, err = c.UpdateWebhook(webhookID, config)
			if err != nil {
				g.Log().Printf("更新Solana Webhook失败: %v", err)
			}
			return webhookID, nil
		}
	}

	// 创建新的Webhook
	config := &HeliusWebhookConfig{
		WebhookURL:       webhookURL,
		TransactionTypes: []string{"ANY"},
		AccountAddresses: addresses,
		WebhookType:      "enhanced",
		AuthHeader:       webhookSecret,
		TxnStatus:        "success",
	}
	result, err := c.CreateWebhook(config)
	if err != nil {
		return "", fmt.Errorf("create webhook error: %v", err)
	}

	g.Log().Printf("Solana Webhook创建成功，ID: %s", result.WebhookID)
	return result.WebhookID, nil
}
