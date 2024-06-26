package kss

import (
	"encoding/json"
	"errors"

	"github.com/di-dao/sonr/crypto"
	"github.com/di-dao/sonr/crypto/core/protocol"
	"github.com/tink-crypto/tink-go/v2/daead"
	"github.com/tink-crypto/tink-go/v2/keyset"
)

// KssI is the interface for the keyshare set
type Set interface {
	Encrypt(key []byte, kh *keyset.Handle) (EncryptedSet, error)
	PublicKey() crypto.PublicKey
	Usr() User
	Val() Val
}

// keyshares is the set of keyshares for the protocol
type keyshares struct {
	val Val
	usr User

	valBz []byte
	usrBz []byte
}

// Encrypt encrypts the keyshares using a password
func (ks *keyshares) Encrypt(key []byte, kh *keyset.Handle) (EncryptedSet, error) {
	if kh == nil {
		return nil, errors.New("kh cannot be nil")
	}
	if key == nil {
		return nil, errors.New("key cannot be nil")
	}

	d, err := daead.New(kh)
	if err != nil {
		return nil, err
	}

	usrEncBz, err := d.EncryptDeterministically(ks.usrBz, key)
	if err != nil {
		return nil, err
	}

	valEncBz, err := d.EncryptDeterministically(ks.valBz, key)
	if err != nil {
		return nil, err
	}

	return &encryptedSet{
		publicKey: ks.PublicKey(),
		encValKey: usrEncBz,
		encUsrKey: valEncBz,
	}, nil
}

// Usr returns the user keyshare
func (ks *keyshares) Usr() User {
	return ks.usr
}

// Val returns the validator keyshare
func (ks *keyshares) Val() Val {
	return ks.val
}

// PublicKey returns the public key for the keyshare set
func (ks *keyshares) PublicKey() crypto.PublicKey {
	return ks.val.PublicKey()
}

// NewKeyshareSet creates a new KeyshareSet
func NewKeyshareSet(aliceResult *protocol.Message, bobResult *protocol.Message) (Set, error) {
	valBz, err := json.Marshal(aliceResult)
	if err != nil {
		return nil, err
	}
	usrBz, err := json.Marshal(bobResult)
	if err != nil {
		return nil, err
	}
	return &keyshares{
		val:   createValidatorKeyshare(aliceResult),
		usr:   createUserKeyshare(bobResult),
		valBz: valBz,
		usrBz: usrBz,
	}, nil
}

// LoadKeyshareSet loads bytes into a keyshare set
func LoadKeyshareSet(valBz []byte, usrBz []byte) (Set, error) {
	val, err := ValidatorKeyshareFromBytes(valBz)
	if err != nil {
		return nil, err
	}
	usr, err := UserKeyshareFromBytes(usrBz)
	if err != nil {
		return nil, err
	}
	return &keyshares{
		val:   val,
		usr:   usr,
		valBz: valBz,
		usrBz: usrBz,
	}, nil
}

// UserKeyshareFromBytes unmarshals a user keyshare from its bytes representation
func UserKeyshareFromBytes(bz []byte) (User, error) {
	val := &protocol.Message{}
	err := json.Unmarshal(bz, val)
	if err != nil {
		return nil, err
	}
	return createUserKeyshare(val), nil
}

// ValidatorKeyshareFromBytes unmarshals a validator keyshare from its bytes representation
func ValidatorKeyshareFromBytes(bz []byte) (Val, error) {
	val := &protocol.Message{}
	err := json.Unmarshal(bz, val)
	if err != nil {
		return nil, err
	}
	return createValidatorKeyshare(val), nil
}
