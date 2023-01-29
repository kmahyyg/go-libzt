package main

import (
	"encoding/json"
	"github.com/kmahyyg/go-libzt/pkg/common"
	"github.com/zerotier/go-ztidentity"
	"io"
	"log"
	"os"
)

func main() {
	fd, err := os.OpenFile("runtime.json", os.O_RDWR|os.O_SYNC|os.O_CREATE, 0640)
	if err != nil {
		panic(err)
	}
	defer fd.Close()
	defer fd.Sync()
	log.Println("file runtime.json opened.")
	fdData, err := io.ReadAll(fd)
	if err != nil {
		panic(err)
	}
	log.Println("data read from runtime.json.")
	var runtimeZTNodeId = &common.RuntimeNodeID{}
	if len(fdData) < 10 {
		ztId := ztidentity.NewZeroTierIdentity()
		log.Println("new zt identity generated.")
		runtimeZTNodeId.NodePriv = ztId.PrivateKeyString()
		runtimeZTNodeId.NodePub = ztId.PublicKeyString()
		jData, err := json.Marshal(runtimeZTNodeId)
		if err != nil {
			panic(err)
		}
		log.Println("data marshalled to json.")
		_, err = fd.Write(jData)
		if err != nil {
			panic(err)
		}
		log.Println("generated new identity, written to file.")
		return
	} else {
		err = json.Unmarshal(fdData, runtimeZTNodeId)
		if err != nil {
			panic(err)
		}
		log.Println("read identity successfully.")
		return
	}
}
