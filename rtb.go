package ciorba

// This file contains Ciorba extensions to RTB (used internally only).

import (
	"github.com/Enremmeta/ciorba/exchange"
	"github.com/Enremmeta/openrtb"
	"net/http"
)

type CiorbaExt struct {
	// Hexadecimal mod_uid (http://www.lexa.ru/programs/mod-uid-eng.html#format) cookie.
	// Example: 9F001EAC1C597756C64DC100021F4EFC
	// In NGINX, the "decoded cookie" is what is written in
	// $uid_got/$uid_set (http://nginx.org/en/docs/http/ngx_http_userid_module.html#variables)
	// fields.
	HexUID string
	// Base64-safe (https://en.wikipedia.org/wiki/Base64#URL_applications)
	// mod_uid (http://www.lexa.ru/programs/mod-uid-eng.html#format) cookie.
	// Example: rB4An1Z3WRwAwU3G_E4fAg
	B64UID string
	// Exchange adapter
	ExchangeAdapter ExchangeAdapter
	// HTTP request that resulted in this bid request.
	HttpRequest http.Request
}

// Enhanced RTB request
type CiorbaRtbRequest struct {
	OpenRtbRequest *openrtb.BidRequest
	// CiorbaExt has all extensions.
	CiorbaExt *CiorbaExt
}

type CiorbaRtbResponse struct {
	openrtb.BidResponse
}
