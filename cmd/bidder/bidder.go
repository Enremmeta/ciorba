package ciorba

import (
	"encoding/json"
	"io/ioutil"
	"net/url"
	"fmt"
)

func Run(configURL string) {
	u, err := url.Parse(configURL)	
	if err != nil {
		panic(err.Error())
	}
	var contents []byte
	switch u.Scheme {
		case "":
		fallthrough;
		case "file":
			contents, err = ioutil.ReadFile(configURL)
			if err != nil {
				panic(err.Error()) 	
			 }		
		default:	
			panic(fmt.Sprintf("Unsupported protocol for now: %s", u.Scheme))
	}
	config := &CiorbaConfig{}
	if err != nil {
		panic(err.Error())
	}
	err = json.Unmarshal(contents, config)
	if err != nil {
		panic(err.Error())
	}

}