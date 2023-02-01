package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"github.com/kmahyyg/go-libzt/pkg/ztnet/node"
	"github.com/kmahyyg/go-libzt/pkg/ztnet/ztcrypto"
	"log"
	"math/rand"
	"os"
	"time"
)

var (
	ErrWorldSigningKeyIllegal = errors.New("world signing key current.c25519 / previous.c25519 is illegal")
	ErrPreflightCheckFailed   = errors.New("preflight check failed, internal requirement cannot be satisfied")
	errUseRecommendValue      = errors.New("potential risk of failed execution, use recommendation if possible")
)

var (
	prevkp    []byte
	curkp     []byte
	mConf     = &MkWorldConfig{}
	gConfFile = flag.String("c", "mkworld.json", "program config")
)

func init() {
	flag.Parse()
	log.Println("startup flag parsed: ", flag.Parsed())
}

func main() {
	/**
	// current.c25519: public key 64 bytes, private key 64 bytes
	// signature: must be signed by previous,
	// message is world after serialized, internal public key is current
	// if initial, previous=current
	// elliptic curve crypt operation are copied from NaCl
	**/
	// Now Start Preflight Check
	if err := Preflight(); err != nil {
		// if signing key is illegal, generate new
		switch err {
		case ErrWorldSigningKeyIllegal:
			log.Println("preflight check error occurred, but still can proceed.")
			prevkp = make([]byte, node.ZT_C25519_PUBLIC_KEY_LEN+node.ZT_C25519_PRIVATE_KEY_LEN)
			pub1, priv1 := ztcrypto.GenerateDualPair()
			copy(prevkp[:node.ZT_C25519_PUBLIC_KEY_LEN], pub1[:])
			copy(prevkp[node.ZT_C25519_PUBLIC_KEY_LEN:], priv1[:])
			copy(curkp, prevkp)
			err = os.WriteFile("current.c25519", prevkp, 0640)
			if err != nil {
				log.Println("failed to write generate c25519 key pair to disk.")
				panic(err)
			}
			err = os.WriteFile("previous.c25519", curkp, 0640)
			if err != nil {
				log.Println("failed to write generate c25519 key pair to disk.")
				panic(err)
			}
			log.Println("new world signing key generated.")
		case errUseRecommendValue:
			log.Println("!You've been warned! WARN! WARN! WARN!")
			if mConf.PlanetRecommend {
				log.Println("since you've set plRecommend to true, we will automatically choose a new value.")
				log.Println("which might be much suitable for you.")
				mConf.PlanetID = rand.Uint64()
				mConf.PlanetBirth = (uint64)(time.Now().UnixMilli())
			}
			log.Println("!You've been warned! WARN! WARN! WARN!")
		default:
			log.Println("preflight check failed.")
			// else panic
			panic(err)
		}
	}
	log.Println("preflight check successfully complete.")
	// Preflight check successfully completed
	ztW := &node.ZtWorld{
		Type:                            node.ZT_WORLD_TYPE_PLANET,
		ID:                              mConf.PlanetID,
		Timestamp:                       mConf.PlanetBirth,
		PublicKeyMustBeSignedByNextTime: [64]byte{},
		Nodes:                           nil,
	}
	var futurePubK [node.ZT_C25519_PUBLIC_KEY_LEN]byte
	copy(futurePubK[:], curkp[:node.ZT_C25519_PUBLIC_KEY_LEN])

	log.Println("generating pre-sign message.")
	toSignZtW, err := ztW.Serialize(true, futurePubK, [node.ZT_C25519_SIGNATURE_LEN]byte{})
	if err != nil {
		panic(err)
	}
	log.Println("pre-sign world generated and serialized successfully.")
	var sigPubK [node.ZT_C25519_PUBLIC_KEY_LEN]byte
	copy(sigPubK[:], prevkp[:node.ZT_C25519_PUBLIC_KEY_LEN])
	var sigPrivK [node.ZT_C25519_PRIVATE_KEY_LEN]byte
	copy(sigPrivK[:], prevkp[node.ZT_C25519_PUBLIC_KEY_LEN:])
	sig4NewWorld, err := ztcrypto.SignMessage(sigPubK, sigPrivK, toSignZtW)
	if err != nil {
		panic(err)
	}
	log.Println("world has been signed.")
	finalWorld, err := ztW.Serialize(false, futurePubK, sig4NewWorld)
	if err != nil {
		panic(err)
	}
	log.Println("new signed world are packed.")
	err = os.WriteFile(mConf.OutputFile, finalWorld, 0644)
	if err != nil {
		panic(err)
	}
	log.Println("packed new signed world has been written to file.")
	log.Println(" ")
	log.Println("now c language output: ")
	fmt.Println(" ")
	fmt.Println("#define ZT_DEFAULT_WORLD_LENGTH ", len(finalWorld))
	fmt.Printf("static const unsigned char ZT_DEFAULT_WORLD[ZT_DEFAULT_WORLD_LENGTH] = {")
	for i, v := range finalWorld {
		if i > 0 {
			fmt.Printf(",")
		}
		fmt.Printf("0x%02x", v)
	}
	fmt.Printf("};")
	fmt.Println(" ")
}

type MkWorldConfig struct {
	SigningKeyFiles []string      `json:"signing"`
	OutputFile      string        `json:"output"`
	RootNodes       []MkWorldNode `json:"rootNodes"`
	PlanetID        uint64        `json:"plID"`
	PlanetBirth     uint64        `json:"plBirth"`
	PlanetRecommend bool          `json:"plRecommend"`
}

type MkWorldNode struct {
	Comments    string   `json:"comments,omitempty"`
	IdentityStr string   `json:"identity"`
	Endpoints   []string `json:"endpoints"`
}

func Preflight() error {
	var nErr error
	gcfdata, err := os.ReadFile(*gConfFile)
	if err != nil {
		return err
	}
	log.Println("config file read.")
	err = json.Unmarshal(gcfdata, mConf)
	if err != nil {
		return err
	}
	log.Println("config file unmarshalled.")
	if len(mConf.SigningKeyFiles) != 2 {
		log.Println("signing key must have 2 files.")
		return ErrPreflightCheckFailed
	}
	if len(mConf.RootNodes) > node.ZT_WORLD_MAX_ROOTS {
		log.Println("root nodes are too many.")
		return ErrPreflightCheckFailed
	}
	for _, v := range mConf.RootNodes {
		if len(v.Endpoints) > node.ZT_WORLD_MAX_STABLE_ENDPOINTS_PER_ROOT {
			log.Println("stable endpoints for current root node are too many.")
			return ErrPreflightCheckFailed
		}
	}
	if mConf.PlanetID == node.ZT_WORLD_ID_EARTH || mConf.PlanetID == node.ZT_WORLD_ID_MARS || mConf.PlanetBirth == 1567191349589 {
		log.Println("!WARNING! You've specified a Planet ID / Birth that is currently in use.")
		nErr = errUseRecommendValue
	}
	if mConf.PlanetBirth <= 1567191349589 {
		log.Println("!WARNING! You've been created a world older than official, timestamp should be larger than 1567191349589.")
		nErr = errUseRecommendValue
	}
	var err1, err2 error
	prevkp, err1 = os.ReadFile(mConf.SigningKeyFiles[0])
	curkp, err2 = os.ReadFile(mConf.SigningKeyFiles[1])
	if err1 != nil || err2 != nil {
		log.Println("read world signing key failed: ", err1, err2)
		return ErrWorldSigningKeyIllegal
	}
	preqLen := node.ZT_C25519_PRIVATE_KEY_LEN + node.ZT_C25519_PUBLIC_KEY_LEN
	if len(prevkp) != preqLen || len(curkp) != preqLen {
		log.Println("world signing key does not satisfy required length.")
		return ErrWorldSigningKeyIllegal
	}
	return nErr
}
