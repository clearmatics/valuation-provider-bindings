// Copyright (c) 2018 Clearmatics Technologies Ltd

package main

import (
	"crypto/ecdsa"
	"math/big"
	"strings"

	"github.com/clearmatics/autonity/accounts/abi"
	"github.com/clearmatics/autonity/accounts/abi/bind"
	"github.com/clearmatics/autonity/accounts/keystore"
	"github.com/clearmatics/autonity/common"
	"github.com/clearmatics/autonity/crypto/secp256k1"
	"github.com/clearmatics/autonity/ethclient"
	commoninterface "github.com/clearmatics/dcn-common/interfaces/valuationprovider"
	valuationprovider "github.com/clearmatics/valuation-provider-bindings/binding"
)

// ===== UTILS ====
// although these functions are implemented by libraries like cryptographyworker,
// this allows token factory to not have external dependencies
func dToPrivateKey(D *big.Int) (*ecdsa.PrivateKey, error) {
	dBytes := D.Bytes()
	var curve = secp256k1.S256()
	tkX, tkY := curve.ScalarBaseMult(dBytes)

	psk := new(ecdsa.PrivateKey)
	tkPub := new(ecdsa.PublicKey)
	tkPub.X = tkX
	tkPub.Y = tkY
	tkPub.Curve = curve
	psk.D = D
	psk.PublicKey = *tkPub
	return psk, nil
}

func getKeystorePrivateKey(keyFile, password string) (*ecdsa.PrivateKey, error) {
	key, err := keystore.DecryptKey([]byte(keyFile), password)
	if err != nil {
		return nil, err
	}
	return key.PrivateKey, nil
}

// ================

// ValuationProvider Factory type
type ValuationProvider int

// NewKeyed create session with private key
func (vp ValuationProvider) NewKeyed(uri string, address string, D *big.Int, gasLimit uint64, gasPrice *big.Int) (commoninterface.ValuationProviderInterface, error) {
	_ = ValuationProviderFactory

	contract, err := getContract(uri, address)
	if err != nil {
		return nil, err
	}

	key, err := dToPrivateKey(D)
	if err != nil {
		return nil, err
	}

	session, err := getSession(contract, key, gasLimit, gasPrice)
	if err != nil {
		return nil, err
	}

	contractABI, err := abi.JSON(strings.NewReader(valuationprovider.ValuationProviderABI))
	if err != nil {
		return nil, err
	}

	return worker{session, contractABI}, nil
}

// New create session with keystore
func (vp ValuationProvider) New(uri string, address, string, keyFile string, password string, gasLimit uint64, gasPrice *big.Int) (commoninterface.ValuationProviderInterface, error) {
	psk, err := getKeystorePrivateKey(keyFile, password)
	if err != nil {
		return nil, err
	}
	return vp.NewKeyed(uri, address, psk.D, gasLimit, gasPrice)
}

func getSession(contract *valuationprovider.ValuationProvider, key *ecdsa.PrivateKey, gasLimit uint64, gasPrice *big.Int) (*valuationprovider.ValuationProviderSession, error) {
	transactor := bind.NewKeyedTransactor(key)

	return &valuationprovider.ValuationProviderSession{
		Contract: contract,
		CallOpts: bind.CallOpts{
			Pending: false,
		},
		TransactOpts: bind.TransactOpts{
			From:     transactor.From,
			Signer:   transactor.Signer,
			GasLimit: gasLimit,
			GasPrice: gasPrice,
		},
	}, nil
}

func getContract(uri string, contractAddress string) (*valuationprovider.ValuationProvider, error) {
	connection, err := ethclient.Dial(uri)
	if err != nil {
		return nil, err
	}

	address := common.HexToAddress(contractAddress)
	contract, err := valuationprovider.NewValuationProvider(address, connection)
	if err != nil {
		return nil, err
	}

	return contract, nil
}

// ValuationProviderFactory exported
var ValuationProviderFactory ValuationProvider = 0
