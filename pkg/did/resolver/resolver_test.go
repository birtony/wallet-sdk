/*
Copyright Avast Software. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package resolver_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/trustbloc/wallet-sdk/pkg/did/resolver"
)

func TestDIDResolver(t *testing.T) {
	didResolver := resolver.NewDIDResolver()

	didDocResolution, err := didResolver.Resolve("did:key:z6MkjfbzWitsSUyFMTbBUSWNsJBHR7BefFp1WmABE3kRw8Qr")
	require.NoError(t, err)
	require.NotEmpty(t, didDocResolution)
}

func TestDIDResolver_InvalidDID(t *testing.T) {
	didResolver := resolver.NewDIDResolver()

	didDocResolution, err := didResolver.Resolve("did:example:abc")
	require.Error(t, err)
	require.EqualError(t, err, "resolve did:example:abc : did method example not supported for vdr")
	require.Empty(t, didDocResolution)
}
