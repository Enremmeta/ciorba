// log package is concerned with logging information about the bidding process. 
// We can distinguish the following types of log files managed by this package.
// 
// Type I. Diagnostics. 
// 
// These logs are for diagnostics/troubleshooting
// purposes, and generally log level here can be regulated.
//
// 1. main -- this is the main application log, and is intended for general 
//    troubleshooting. This is what most people think about when they think
//    about logs. 
// 2. rawrequest -- logs requests as they come in (see ... below on how requests
//    that are binary are logged). Because of high volume of data this can generate
//    it is not recommended to enable it in production except on a small percentage
//    of traffic.
// 3. rawresponse -- similar to the above but for response. 
// 
// Type II. Event. 
// 
// Each of these logs represents an event in the life of an RTB session (request,
// bid, win, loss, error, impression, click, possibly others). These logs are structured 
// (tab-separated tables, really), with well-defined metadata (columns) and a rule
// that future releases will never change/remove columns, but only
// add new ones to the end. These logs can easily be ingested into various 
// databases/data warehouses for reporting and analytics.         

package log

import (
	"log"
	"net/http"
	"github.com/Enremmeta/openrtb"
	"github.com/Enremmeta/ciorba"
)

const (
	ciorbaLogFlagsMain = log.Ldate | log.Ltime | log.LUTC | log.Llongfile
)

// TODO do some stuff around synchronizing it when reloading

type CiorbaLogger struct {
	log.Logger
}

//	QueryStringBidRequestID = "brid"
//	QueryStringImpressionID = "iid"
//	QueryStringBidID        = "bid"
//	QueryStringBidTimestamp = "bidts"
//	QueryStringCampaignID   = "cid"
//	QueryStringCreativeID   = "crid"
//	QueryStringExchange     = "xch"
//	QueryStringSSP          = "ssp"
//	// User cookie in Hex format (see CiorbaExt.HexUID) that was seen in 
//	// bid request. 
//	QueryStringHex          = "hex"
//	

func Errorf(msg string, args...interface{}) {
	
}

// commonCols extracts the common columns from the appropriate entity
// (e.g., bid request or query string) common to all event logs. These are:
// 
func commonCols() []string {
	
}

func commonColsRtbRequest(req ciorba.CiorbaRtbRequest, imp.openrtb.Impression) string {

}

// commonColsWebRequest
func commonColsWebRequest(req http.Request) ([]string) {
	return []string{req.FormValue
}

func initLogging() error {
	logConfig := Config.Logs

}
