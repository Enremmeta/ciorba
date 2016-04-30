package ciorba

import (
	"net/url"
	"github.com/Enremmeta/ciorba/integration/exchange/adx"
)

// ```
//{
//   "LogDirectory" : "/opt/ciorba/log",
//   "Port" : 10000,
//   "Exchanges" : {},
//   "Geo" : { "MaxMindDir" : "/opt/ciorba/data/geo/maxmind" }
//}
// ```
var Config *CiorbaConfig

type Exchange map[string]interface{}

type GeoConfig struct {
	MaxMindFiles string
}

type CiorbaURL struct {
	*url.URL
}

func (u *CiorbaURL) UnmarshalJSON(data []byte) error {
	s := string(data)
	u2, err := url.Parse(s)
	if err != nil {
		return err
	}
	u.URL = u2
	return nil
}

func (u *CiorbaURL) MarshalJSON() ([]byte, error) {
	return []byte(u.URL.String()), nil
}

type DBConfig struct {
	CiorbaURL
	Name string
}

type SQLiteConfig struct {
	DBConfig
}

type MySQLConfig struct {
	DBConfig
}

type AerospikeConfig struct {
	DBConfig
}

type Databases struct {
	Aerospikes []AerospikeConfig
	MySQLs     []MySQLConfig
	SQLites    []SQLiteConfig
}

type LogConfig struct {
}

type LogsConfig struct {
	LogDirectory string
	MainStdout   bool
	Main         LogConfig
	Request      LogConfig
	RawRequest   LogConfig
	Bid          LogConfig
	Win          LogConfig
	Loss         LogConfig
	Impression   LogConfig
	Click        LogConfig
}

const (
	LoadSheddingBehaviorOptout = "optout"
	LoadSheddingBehavior503 = "503"
)

type WebConfig struct {
	Port      uint
	AdminPort uint
	KeepAliveTimeoutMillis uint
	Worker1Count: uint
	LoadSheddingBehavior: string
}

type Exchanges struct {
	AdX adx.Config
}

type CiorbaConfig struct {
	Web WebConfig
	Exchanges Exchanges `json:"exchanges,omitempty"`
	Databases Databases
	Logs      LogConfig
}
