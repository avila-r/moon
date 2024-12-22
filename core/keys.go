package core

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"log"
	"math/big"
)

type (
	PrivateKey struct {
		ECDSA *ecdsa.PrivateKey
	}

	PublicKey []byte

	PublicAddress [20]uint8

	Signature struct {
		S *big.Int
		R *big.Int
	}
)

func GeneratePrivateKey() PrivateKey {
	return GeneratePrivateKeyFrom(rand.Reader)
}

func GeneratePrivateKeyFrom(r io.Reader) PrivateKey {
	ecdsa, err := ecdsa.GenerateKey(elliptic.P256(), r)
	if err != nil {
		log.Fatalf("failed to generate private key - %v", err.Error())
	}

	return PrivateKey{
		ECDSA: ecdsa,
	}
}

func GeneratePublicAddressFromPublicKey(pub PublicKey) PublicAddress {
	h := sha256.Sum256(pub)
	b := h[len(h)-20:]

	if len(b) != 20 {
		log.Fatalf("given bytes with length %d should be 20", len(b))
	}

	var value [20]uint8
	for i := 0; i < 20; i++ {
		value[i] = b[i]
	}

	return PublicAddress(value)
}

func (sig Signature) ToString() string {
	bytes := append(sig.S.Bytes(), sig.R.Bytes()...)

	return hex.EncodeToString(bytes)
}

func (sig Signature) Verify(pub PublicKey, data []byte) bool {
	x, y := elliptic.UnmarshalCompressed(elliptic.P256(), pub)
	key := &ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     x,
		Y:     y,
	}

	return ecdsa.Verify(key, data, sig.R, sig.S)
}

func (prv PrivateKey) Sign(seed []byte) (*Signature, error) {
	r, s, err := ecdsa.Sign(rand.Reader, prv.ECDSA, seed)
	if err != nil {
		return nil, err
	}

	return &Signature{
		R: r,
		S: s,
	}, nil
}

func (prv PrivateKey) PublicKey() PublicKey {
	ecdsa := prv.ECDSA.PublicKey

	return elliptic.MarshalCompressed(ecdsa, ecdsa.X, ecdsa.Y)
}

func (pub PublicKey) ToString() string {
	return hex.EncodeToString(pub)
}

func (pub PublicKey) Address() PublicAddress {
	return GeneratePublicAddressFromPublicKey(pub)
}
