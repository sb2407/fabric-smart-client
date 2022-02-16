/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package implicitdb

import (
	"github.com/hyperledger-labs/fabric-smart-client/platform/fabric"
	view2 "github.com/hyperledger-labs/fabric-smart-client/platform/view"
	"github.com/pkg/errors"
)

type ImplicitDB struct {
	ch *fabric.Channel
}

func Get(sp view2.ServiceProvider, network, channel string) (*ImplicitDB, error) {
	fns := fabric.GetFabricNetworkService(sp, network)
	if fns == nil {
		return nil, errors.Errorf("fabric network service not found for network %s", network)
	}
	ch, err := fns.Channel(channel)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get channel %s", channel)
	}
	return &ImplicitDB{ch: ch}, nil
}

func (db *ImplicitDB) Put(key string, value []byte) error {
	_, _, err := db.ch.Chaincode(
		"implicitdb",
	).Invoke(
		"Put",
	).WithTransientEntry(
		key, value,
	).Call()
	if err != nil {
		return errors.Wrapf(err, "failed to put key %s", key)
	}
	return nil
}
