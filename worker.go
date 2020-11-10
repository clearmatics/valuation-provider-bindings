// Copyright (c) 2020 Clearmatics Technologies Ltd

package main

import (
	"github.com/clearmatics/binding-helpers/conversion"
	commontypes "github.com/clearmatics/dcn-common/types"
	valuationprovider "github.com/clearmatics/valuation-provider-bindings/binding"

	"github.com/clearmatics/autonity/accounts/abi"
)

type worker struct {
	session *valuationprovider.ValuationProviderSession
	abi     abi.ABI
}

func (w worker) RequestValuation() (*commontypes.Hash, error) {
	transaction, err := w.session.RequestValuation()
	if err != nil {
		return nil, err
	}

	result := conversion.ToTransactionHash(transaction)
	return &result, nil
}

func (w worker) SubmitValuation(to commontypes.Address, valuation []byte) (*commontypes.Hash, error) {
	addrCommon := conversion.ToCommonAddress(to)
	transaction, err := w.session.SubmitValuation(addrCommon, valuation)
	if err != nil {
		return nil, err
	}

	result := conversion.ToTransactionHash(transaction)
	return &result, nil
}
