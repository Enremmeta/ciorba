package ciorba

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
)

func failIfError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func Run(configURL string) {
	u, err := url.Parse(configURL)
	failIfError(err)
	var contents []byte
	switch u.Scheme {
	case "":
		fallthrough
	case "file":
		contents, err = ioutil.ReadFile(u.Path)
		failIfError(err)
	default:
		panic(fmt.Sprintf("Cannot handle %s for now.", u.Scheme))
	}
	Config = &CiorbaConfig{}
	
	failIfError(json.Unmarshal(contents, Config))
	if err != nil {
		panic(err.Error())
	}
	
	err = initLogging()
	if err != nil {
		panic(err.Error())
	}
	
	
}
