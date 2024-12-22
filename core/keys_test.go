package core_test

import (
	"crypto/rand"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/avila-r/moon/core"
)

func GenerateRandomSeed() ([]byte, error) {
	seed := make([]byte, 32)

	_, err := rand.Read(seed)
	if err != nil {
		return nil, err
	}

	return seed, nil
}

func Test_KeypairSign(t *testing.T) {
	t.Run("expect success", func(t *testing.T) {
		var (
			assert = assert.New(t)

			prv = core.GeneratePrivateKey()
			pub = prv.PublicKey()
		)

		seed, err := GenerateRandomSeed()
		assert.Nil(err, "should generate random seed without error")

		sig, err := prv.Sign(seed)
		assert.Nil(err, "should sign seed without error")

		valid := sig.Verify(pub, seed)
		assert.True(valid, "should verify the signature with the correct public key and seed")
	})

	t.Run("expect failure", func(t *testing.T) {
		var (
			assert = assert.New(t)

			prv = core.GeneratePrivateKey()
			pub = prv.PublicKey()
		)

		seed, err := GenerateRandomSeed()
		assert.Nil(err, "should generate random seed without error")

		sig, err := prv.Sign(seed)
		assert.Nil(err, "should sign seed without error")

		valid := sig.Verify(core.GeneratePrivateKey().PublicKey(), seed)
		assert.False(valid, "should not verify the signature with a different public key")

		valid = sig.Verify(pub, []byte("invalid seed"))
		assert.False(valid, "should not verify the signature with an incorrect seed")
	})

	t.Run("expect failure - invalid private key", func(t *testing.T) {
		var (
			assert = assert.New(t)
			prv    = core.GeneratePrivateKey()
		)

		seed, err := GenerateRandomSeed()
		assert.Nil(err, "should generate random seed without error")

		sig, err := prv.Sign(seed)
		assert.Nil(err, "should sign seed without error")

		valid := sig.Verify(core.GeneratePrivateKey().PublicKey(), seed)
		assert.False(valid, "should not verify the signature with an invalid private key")
	})
}
