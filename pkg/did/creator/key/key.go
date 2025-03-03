/*
Copyright Avast Software. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

// Package key contains a did:key creator implementation.
package key

import (
	"errors"

	"github.com/trustbloc/did-go/doc/did"
	"github.com/trustbloc/did-go/method/key"
	"github.com/trustbloc/kms-go/doc/jose/jwk"
	"github.com/trustbloc/wallet-sdk/pkg/walleterror"
)

// ErrorModule is the error module name used for errors relating to did:key creation.
const ErrorModule = "DIDKEY"

// Creator is used for creating did:key DID Documents.
type Creator struct {
	vdr *key.VDR
}

// NewCreator returns a new did:key document Creator.
// Deprecated: The standalone Create function should be used instead.
func NewCreator() *Creator {
	return &Creator{vdr: key.New()}
}

// Create creates a new did:key document using the given verification method.
// Deprecated: The standalone Create function should be used instead.
func (d *Creator) Create(vm *did.VerificationMethod) (*did.DocResolution, error) {
	didDocArgument := &did.Doc{VerificationMethod: []did.VerificationMethod{*vm}}

	return d.vdr.Create(didDocArgument)
}

// Create creates a new did:key document using the given verification method.
func Create(jsonWebKey *jwk.JWK) (*did.DocResolution, error) {
	if jsonWebKey == nil {
		return nil, walleterror.NewInvalidSDKUsageError(
			ErrorModule, errors.New("jwk object cannot be nil"))
	}

	var vm *did.VerificationMethod

	if jsonWebKey.Crv == "Ed25519" {
		// Workaround: when the did:key VDR creates a DID for ed25519, Ed25519VerificationKey2018 is the expected
		// verification method.
		publicKeyBytes, err := jsonWebKey.PublicKeyBytes()
		if err != nil {
			return nil, err
		}

		vm = &did.VerificationMethod{Value: publicKeyBytes, Type: "Ed25519VerificationKey2018"}
	} else {
		var err error

		vm, err = did.NewVerificationMethodFromJWK("", "JsonWebKey2020", "", jsonWebKey)
		if err != nil {
			return nil, err
		}
	}

	didDocArgument := &did.Doc{VerificationMethod: []did.VerificationMethod{*vm}}

	return key.New().Create(didDocArgument)
}
