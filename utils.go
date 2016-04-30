package ciorba

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"fmt"
)

const (
	// We are going to try these encodings in order
	encodings = []base64.Encoding{base64.RawURLEncoding, base64.URLEncoding, base64.StdEncoding, base64.RawStdEncoding}
)

// HexToB4UID method is an inverse of B64ToHexUID
func HexToB4UID(b64 string) (string, error) {
	panic("Not implemented")
	return "", nil
}

// B64ToHexUID converts Base64-safe (https://en.wikipedia.org/wiki/Base64#URL_applications)
// mod_uid (http://www.lexa.ru/programs/mod-uid-eng.html#format) cookie (example: rB4An1Z3WRwAwU3G_E4fAg) to its
// hexadecimal equivalent (example: 9F001EAC1C597756C64DC100021F4EFC).
func B64ToHexUID(b64 string) (string, error) {
	var cookie [4]uint32
	var data string
	var err0 error
	var err error
	
	for _, enc := range encodings {
		err, data = enc.DecodeString(b64)
		if err == nil {
			break
		}
		if err0 != nil {
			err0 = err
		}
	}
	
	if err0 != nil {
		return "", err0
	}
	
	b := []byte(data)
	buf := bytes.NewReader(b)
	err = binary.Read(buf, binary.LittleEndian, &cookie)
	if err != nil {
		return "", err
	}
	retval := fmt.Sprintf("%08X%08X%08X%08X\n", cookie[0], cookie[1], cookie[2], cookie[3])
	return retval, nil	
}
