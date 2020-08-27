package qiwi_payments

import (
	"time"
)

type Payload struct {
	Amount             Amount       `json:"amount"`
	Comment            string       `json:"comment"`
	ExpirationDateTime string       `json:"expirationDateTime"`
	CustomFields       CustomFields `json:"customFields"`
}
type CustomFields struct {
	ThemeCode string `json:"themeCode"`
}
type Amount struct {
	Currency string `json:"currency"`
	Value    string `json:"value"`
}
type QiwiStatus string
type ResponseQiwiPaymentBill struct {
	SiteID string `json:"siteId"`
	BillID string `json:"billId"`
	Amount struct {
		Value    float64 `json:"value,string"`
		Currency string  `json:"currency"`
	} `json:"amount"`
	Status struct {
		Value           QiwiStatus `json:"value"`
		ChangedDateTime time.Time  `json:"changedDateTime"`
	} `json:"status"`
	Comment            string    `json:"comment"`
	CreationDateTime   string    `json:"creationDateTime"`
	ExpirationDateTime time.Time `json:"expirationDateTime"`
	PayURL             string    `json:"payUrl"`
}
type ResponseQiwiPaymentBillError struct {
	ServiceName string    `json:"serviceName"`
	ErrorCode   string    `json:"errorCode"`
	Description string    `json:"description"`
	UserMessage string    `json:"userMessage"`
	Datetime    time.Time `json:"datetime"`
	TraceID     string    `json:"traceId"`
}
