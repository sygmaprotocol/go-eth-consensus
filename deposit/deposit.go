// Code generated by ethgo/abigen. DO NOT EDIT.
// Hash: c7ac9f62b1ca54bf9b6bcedde8041ba5961c2ef62683c4ff6b77a19ae09bab7a
// Version: 0.1.2
package deposit

import (
	"fmt"
	"math/big"

	"github.com/umbracle/ethgo"
	"github.com/umbracle/ethgo/contract"
	"github.com/umbracle/ethgo/jsonrpc"
)

var (
	_ = big.NewInt
	_ = jsonrpc.NewClient
)

// Deposit is a solidity contract
type Deposit struct {
	c *contract.Contract
}

// DeployDeposit deploys a new Deposit contract
func DeployDeposit(provider *jsonrpc.Client, from ethgo.Address, args []interface{}, opts ...contract.ContractOption) (contract.Txn, error) {
	return contract.DeployContract(abiDeposit, binDeposit, args, opts...)
}

// NewDeposit creates a new instance of the contract at a specific address
func NewDeposit(addr ethgo.Address, opts ...contract.ContractOption) *Deposit {
	return &Deposit{c: contract.NewContract(addr, abiDeposit, opts...)}
}

// calls

// GetDepositCount calls the get_deposit_count method in the solidity contract
func (d *Deposit) GetDepositCount(block ...ethgo.BlockNumber) (retval0 []byte, err error) {
	var out map[string]interface{}
	var ok bool

	out, err = d.c.Call("get_deposit_count", ethgo.EncodeBlock(block...))
	if err != nil {
		return
	}

	// decode outputs
	retval0, ok = out["0"].([]byte)
	if !ok {
		err = fmt.Errorf("failed to encode output at index 0")
		return
	}
	
	return
}

// GetDepositRoot calls the get_deposit_root method in the solidity contract
func (d *Deposit) GetDepositRoot(block ...ethgo.BlockNumber) (retval0 [32]byte, err error) {
	var out map[string]interface{}
	var ok bool

	out, err = d.c.Call("get_deposit_root", ethgo.EncodeBlock(block...))
	if err != nil {
		return
	}

	// decode outputs
	retval0, ok = out["0"].([32]byte)
	if !ok {
		err = fmt.Errorf("failed to encode output at index 0")
		return
	}
	
	return
}

// SupportsInterface calls the supportsInterface method in the solidity contract
func (d *Deposit) SupportsInterface(interfaceId [4]byte, block ...ethgo.BlockNumber) (retval0 bool, err error) {
	var out map[string]interface{}
	var ok bool

	out, err = d.c.Call("supportsInterface", ethgo.EncodeBlock(block...), interfaceId)
	if err != nil {
		return
	}

	// decode outputs
	retval0, ok = out["0"].(bool)
	if !ok {
		err = fmt.Errorf("failed to encode output at index 0")
		return
	}
	
	return
}

// txns

// Deposit sends a deposit transaction in the solidity contract
func (d *Deposit) Deposit(pubkey []byte, withdrawalCredentials []byte, signature []byte, depositDataRoot [32]byte) (contract.Txn, error) {
	return d.c.Txn("deposit", pubkey, withdrawalCredentials, signature, depositDataRoot)
}

// events

func (d *Deposit) DepositEventEventSig() ethgo.Hash {
	return d.c.GetABI().Events["DepositEvent"].ID()
}