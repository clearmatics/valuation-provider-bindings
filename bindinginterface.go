// Copyright (c) 2020 Clearmatics Technologies Ltd

package main

import (
	commontypes "github.com/clearmatics/dcn-common/types"
)

// BindingInterface defines all the functions of the ValuationProvider contract bindings
type BindingInterface interface {
	RequestValuation() (*commontypes.Hash, error)
	SubmitValuation(to commontypes.Address, valuation []byte) (*commontypes.Hash, error)
}
