// Copyright (c) 2020 Clearmatics Technologies Ltd

package main

import (
	"math/big"

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

func (w worker) SubmitValuation(
	to commontypes.Address,
	tradesRootHash [32]byte,
	variationMargin *big.Int,
	payee commontypes.Address,
	exchangeRates []*big.Int) (*commontypes.Hash, error) {

	toCommon := conversion.ToCommonAddress(to)
	payeeCommon := conversion.ToCommonAddress(payee)
	transaction, err := w.session.SubmitValuation(toCommon, tradesRootHash, variationMargin, payeeCommon, exchangeRates)
	if err != nil {
		return nil, err
	}

	result := conversion.ToTransactionHash(transaction)
	return &result, nil
}
