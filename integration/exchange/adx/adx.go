package adx

import (
	"github.com/Enremmeta/ciorba"
	"github.com/Enremmeta/ciorba/integration/exchange"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"net/http"
)

type Config struct {
	GeoTablePath  string
	encryptionKey []byte
	integrityKey  []byte
}

type AdxExchangeAdapter struct {
}

// Name provides name of this exchange.
func (adx AdxExchangeAdapter) Name() string {
	return "adx"
}

// See
// https://developers.google.com/ad-exchange/rtb/response-guide/decrypt-price#encryption-scheme
func (adx AdxExchangeAdapter) DecryptWinningPrice(price string) uint64 {

}

func (adx AdxExchangeAdapter) ConvertRequest(req http.Request) (ciorba.CiorbaRtbRequest, error) {
	retval := ciorba.CiorbaRtbRequest{}
	adxReq := &BidRequest{}
	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, er
	}
	err = proto.Unmarshal(buf, adxReq)
	if err != nil {
		return nil, err
	}
	// Let the fun begin...
	return retval, nil	
}

func (adx AdxExchangeAdapter) ConvertResponse(ciorba.CiorbaRtbResponse) interface{} {

}
