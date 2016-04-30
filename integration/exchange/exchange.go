package exchange

import (
	"github.com/Enremmeta/ciorba"
	"github.com/Enremmeta/ciorba/log"
	"github.com/Enremmeta/openrtb"
	"net/http"
	"strconv"
	"strings"
)

type ExchangeAdapter interface {
	ConvertRequest(http.Request) (ciorba.CiorbaRtbRequest, error)
	ConvertResponse(ciorba.CiorbaRtbResponse) (interface{}, error)
	DecryptWinningPrice(string) (uint64, error)
	Name() string
}

func ParseUserId(userID string, req *ciorba.CiorbaRtbRequest) {
	if req.User == nil {
		req.User = &openrtb.User{}
	}
	userID = strings.TrimSpace(userID)
	if userID == "" {
		return
	}
	_, err := strconv.ParseUint(userID, 16, 64)
	if err == nil {
		// We can parse a hex number, so it is a hex format UID
		req.CiorbaExt.HexUID = userID
		req.CiorbaExt.B64UID, err = ciorba.HexToB64UID(userID)
		if err != nil {
			// Just log the error, what else can we do...
			log.Errorf("Cannot convert %s: %s", userID, err)
		}
		return
	}
	req.CiorbaExt.B64UID = userID
	req.CiorbaExt.HexUID, err = ciorba.B64ToHexUID(userID)
	if err != nil {
		// Just log the error, what else can we do...
		log.Errorf("Cannot convert %s: %s", userID, err)
	}
}
