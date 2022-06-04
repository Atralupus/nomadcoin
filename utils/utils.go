package utils

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
)

func HandleErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func ToBytes(i interface{}) []byte {
	var blockBuffer bytes.Buffer

	encoder := gob.NewEncoder(&blockBuffer)
	err := encoder.Encode(i)
	HandleErr(err)

	return blockBuffer.Bytes()
}

func FromBytes(i interface{}, data []byte) {
	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(i)
	HandleErr(err)
}

func Hash(i interface{}) string {
	s := fmt.Sprintf("%v", i)
	hash := fmt.Sprintf("%x", sha256.Sum256([]byte(s)))
	return hash
}
