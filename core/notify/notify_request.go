// Copyright 2021 Tencent Inc. All rights reserved.

package notify

import (
	"net/http"
	"time"
)

// Request 微信支付通知请求结构
type Request struct {
	ID           string             `json:"id"`
	CreateTime   *time.Time         `json:"create_time"`
	EventType    string             `json:"event_type"`
	ResourceType string             `json:"resource_type"`
	Resource     *EncryptedResource `json:"resource"`
	Summary      string             `json:"summary"`

	// 原始通知请求
	RawRequest *http.Request
}

// EncryptedResource 微信支付通知请求中的内容
type EncryptedResource struct {
	Algorithm      string `json:"algorithm"`
	Ciphertext     string `json:"ciphertext"`
	AssociatedData string `json:"associated_data"`
	Nonce          string `json:"nonce"`
	OriginalType   string `json:"original_type"`

	Plaintext string // Ciphertext 解密后内容
}

// ContentMap 将微信支付通知请求内容Json解密后的内容Map
type ContentMap map[string]interface{}
