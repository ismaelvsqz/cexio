// Copyright 2021 Juan Ismael Vasquez <ismael.juan@gmail.com>.
// All rights reserved. Use of this source code is governed by
// a MIT style license that can be found in the LICENSE file.

package cexio

import (
	"encoding/json"
	"log"
)

// Request represents the request data
type Request struct {
	Method     string
	Parameters []string
	Data       map[string]interface{}
}

// URL returns the Request URL
func (req *Request) URL() string {
	u := CexURL + req.Method + "/"
	len := len(req.Parameters)
	for i, e := range req.Parameters {
		u = u + e
		if i+1 < len {
			u = u + "/"
		}
	}
	return u
}

// IsPrivate return true if the Request is a private resource
func (req *Request) IsPrivate() bool {
	switch req.Method {
	case
		MethodBalance,
		MethodOpenOrders,
		MethodActiveOrdersStatus,
		MethodArchivedOrders,
		MethodCancelOrder,
		MethodCancelOrders,
		MethodPlaceOrder,
		MethodGetOrder,
		MethodGetOrderTransactions,
		MethodGetAddress,
		MethodGetMyFee,
		MethodCancelReplaceOrder,
		MethodOpenPosition,
		MethodOpenPositions,
		MethodGetPosition,
		MethodClosePosition,
		MethodArchivedPositions,
		MethodGetMarginalFee:
		return true
	default:
		return false
	}
}

// IsPost return true if the Request is a HTTP Post message
func (req *Request) IsPost() bool {
	if req.IsPrivate() {
		return true
	}

	switch req.Method {
	case
		MethodConvert,
		MethodChart:
		return true
	default:
		return false
	}
}

// API storages the api authentication keys
type API struct {
	Username  string
	APIKey    string
	APISecret string
}

// Payload returns the message POST to the API
func (api *API) Payload(r Request) map[string]interface{} {
	result := make(map[string]interface{})
	for key, value := range r.Data {
		result[key] = value
	}
	if r.IsPrivate() && api.HasKeys() {
		signature, nonce := api.Signature()
		result["key"] = api.APIKey
		result["signature"] = signature
		result["nonce"] = nonce
	}
	return result
}

// HasKeys returns true if the API Keys are defined
func (api *API) HasKeys() bool {
	return (hasText(api.APIKey) &&
		hasText(api.APISecret) &&
		hasText(api.Username))
}

// Signature generates the authentication signature
func (api *API) Signature() (string, string) {
	nonce := nonce()
	message := nonce + api.Username + api.APIKey
	signature := toHmac256(message, api.APISecret)
	return signature, nonce
}

// Do executes the CEX.IO api
func (api *API) Do(r Request) ([]byte, error) {
	if !r.IsPost() {
		return getMethod(r.URL())
	}

	bodyJSON, err := json.Marshal(api.Payload(r))
	if err != nil {
		log.Printf(errorDecodingResponse, err)
		return nil, err
	}
	return postMethod(r.URL(), bodyJSON)
}
