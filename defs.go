package ciorba

const (
	QueryStringBidRequestID = "brid"
	QueryStringImpressionID = "iid"
	QueryStringBidID        = "bid"
	QueryStringBidTimestamp = "bidts"
	QueryStringCampaignID   = "cid"
	QueryStringCreativeID   = "crid"
	QueryStringExchange     = "xch"
	QueryStringSSP          = "ssp"
	// User cookie in Hex format (see CiorbaExt.HexUID) that was seen in 
	// bid request
	QueryStringHex          = "hex"
	
	QueryStringRedir        = "r"
)

type NurlCache interface {
}
