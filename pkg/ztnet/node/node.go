/*
 *  SPDX-License-Identifier: AGPL-3.0-only
 *  Copyright (C) 2023 by kmahyyg in Patmeow Limited
 */

package node

import "bytes"

const (
	ZT_C25519_PUBLIC_KEY_LEN  = 64
	ZT_C25519_PRIVATE_KEY_LEN = 64
	ZT_C25519_SIGNATURE_LEN   = 96
)

type ZtNormalNode struct {
	privateKey    [ZT_C25519_PRIVATE_KEY_LEN]byte
	ZtNodeAddress uint64 // but only use big-endian high 40 bits
}

func (ztn ZtNormalNode) HasPrivateKey() bool {
	return bytes.Equal(ztn.privateKey[0:4], []byte{0, 0, 0, 0})
}
