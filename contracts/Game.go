// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// GameMatch is an auto generated low-level Go binding around an user-defined struct.
type GameMatch struct {
	MatchStatus uint8
	MatchPrice  *big.Int
	MatchHash   [32]byte
	Name        string
	P1          common.Address
	P2          common.Address
	Winner      common.Address
}

// ContractsMetaData contains all meta data concerning the Contracts contract.
var ContractsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"p1\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"p2\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"matchId\",\"type\":\"uint256\"}],\"name\":\"CreateMatchEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"matchId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"int256\",\"name\":\"_matchState\",\"type\":\"int256\"}],\"name\":\"MatchEndEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"approveReset\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_p1Address\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_p2Address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_matchPrice\",\"type\":\"uint256\"}],\"name\":\"createMatchByOwner\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_matchId\",\"type\":\"uint256\"}],\"name\":\"getMatch\",\"outputs\":[{\"components\":[{\"internalType\":\"enumGame.MatchStatus\",\"name\":\"matchStatus\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"matchPrice\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"MatchHash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"p1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"p2\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"}],\"internalType\":\"structGame.Match\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_matchId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_winner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_losser\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"_matchState\",\"type\":\"int256\"}],\"name\":\"matchEnd\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferByOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040526040518060400160405280600c81526020017f43616e69736f20546f6b656e0000000000000000000000000000000000000000815250600390816200004a91906200040c565b506040518060400160405280600381526020017f43544b0000000000000000000000000000000000000000000000000000000000815250600490816200009191906200040c565b506012600560006101000a81548160ff021916908360ff1602179055506064600755348015620000c057600080fd5b506b033b2e3c9fd0803ce8000000600080828254620000e0919062000522565b925050819055506b033b2e3c9fd0803ce8000000600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825462000144919062000522565b9250508190555033600560016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506200055d565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806200021457607f821691505b6020821081036200022a5762000229620001cc565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302620002947fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8262000255565b620002a0868362000255565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b6000620002ed620002e7620002e184620002b8565b620002c2565b620002b8565b9050919050565b6000819050919050565b6200030983620002cc565b620003216200031882620002f4565b84845462000262565b825550505050565b600090565b6200033862000329565b62000345818484620002fe565b505050565b5b818110156200036d57620003616000826200032e565b6001810190506200034b565b5050565b601f821115620003bc57620003868162000230565b620003918462000245565b81016020851015620003a1578190505b620003b9620003b08562000245565b8301826200034a565b50505b505050565b600082821c905092915050565b6000620003e160001984600802620003c1565b1980831691505092915050565b6000620003fc8383620003ce565b9150826002028217905092915050565b620004178262000192565b67ffffffffffffffff8111156200043357620004326200019d565b5b6200043f8254620001fb565b6200044c82828562000371565b600060209050601f8311600181146200048457600084156200046f578287015190505b6200047b8582620003ee565b865550620004eb565b601f198416620004948662000230565b60005b82811015620004be5784890151825560018201915060208501945060208101905062000497565b86831015620004de5784890151620004da601f891682620003ce565b8355505b6001600288020188555050505b505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006200052f82620002b8565b91506200053c83620002b8565b9250828201905080821115620005575762000556620004f3565b5b92915050565b612495806200056d6000396000f3fe6080604052600436106100fe5760003560e01c80635a15c373116100955780638da5cb5b116100645780638da5cb5b1461037457806395d89b411461039f578063a0712d68146103ca578063a9059cbb146103f3578063dd62ed3e14610430576100fe565b80635a15c373146102a157806370a08231146102de5780637f2544491461031b57806387dadf8314610358576100fe565b806323b872dd116100d157806323b872dd146101d3578063313ce567146102105780633d092b3d1461023b57806342966c6814610278576100fe565b806306fdde0314610103578063095ea7b31461012e57806318160ddd1461016b57806322974e6414610196575b600080fd5b34801561010f57600080fd5b5061011861046d565b60405161012591906118e6565b60405180910390f35b34801561013a57600080fd5b50610155600480360381019061015091906119a6565b6104fb565b6040516101629190611a01565b60405180910390f35b34801561017757600080fd5b506101806105ed565b60405161018d9190611a2b565b60405180910390f35b3480156101a257600080fd5b506101bd60048036038101906101b89190611a46565b6105f3565b6040516101ca9190611a01565b60405180910390f35b3480156101df57600080fd5b506101fa60048036038101906101f59190611a86565b610741565b6040516102079190611a01565b60405180910390f35b34801561021c57600080fd5b506102256108f2565b6040516102329190611af5565b60405180910390f35b34801561024757600080fd5b50610262600480360381019061025d9190611b10565b610905565b60405161026f9190611cd1565b60405180910390f35b34801561028457600080fd5b5061029f600480360381019061029a9190611b10565b610b15565b005b3480156102ad57600080fd5b506102c860048036038101906102c39190611a86565b610c46565b6040516102d59190611a01565b60405180910390f35b3480156102ea57600080fd5b5061030560048036038101906103009190611cf3565b610dbe565b6040516103129190611a2b565b60405180910390f35b34801561032757600080fd5b50610342600480360381019061033d9190611d85565b610dd6565b60405161034f9190611a2b565b60405180910390f35b610372600480360381019061036d9190611e43565b6110ed565b005b34801561038057600080fd5b50610389611468565b6040516103969190611eb9565b60405180910390f35b3480156103ab57600080fd5b506103b461148e565b6040516103c191906118e6565b60405180910390f35b3480156103d657600080fd5b506103f160048036038101906103ec9190611b10565b61151c565b005b3480156103ff57600080fd5b5061041a600480360381019061041591906119a6565b61164d565b6040516104279190611a01565b60405180910390f35b34801561043c57600080fd5b5061045760048036038101906104529190611a46565b61176a565b6040516104649190611a2b565b60405180910390f35b6003805461047a90611f03565b80601f01602080910402602001604051908101604052809291908181526020018280546104a690611f03565b80156104f35780601f106104c8576101008083540402835291602001916104f3565b820191906000526020600020905b8154815290600101906020018083116104d657829003601f168201915b505050505081565b600081600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040516105db9190611a2b565b60405180910390a36001905092915050565b60005481565b60003373ffffffffffffffffffffffffffffffffffffffff16600560019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461064f57600080fd5b6000600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925600060405161072f9190611f79565b60405180910390a36001905092915050565b600081600260008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546107cf9190611fc3565b9250508190555081600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546108259190611fc3565b9250508190555081600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461087b9190611ff7565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516108df9190611a2b565b60405180910390a3600190509392505050565b600560009054906101000a900460ff1681565b61090d6117c2565b600660008381526020019081526020016000206040518060e00160405290816000820160009054906101000a900460ff1660058111156109505761094f611b3d565b5b600581111561096257610961611b3d565b5b8152602001600182015481526020016002820154815260200160038201805461098a90611f03565b80601f01602080910402602001604051908101604052809291908181526020018280546109b690611f03565b8015610a035780601f106109d857610100808354040283529160200191610a03565b820191906000526020600020905b8154815290600101906020018083116109e657829003601f168201915b505050505081526020016004820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016005820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016006820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815250509050919050565b3373ffffffffffffffffffffffffffffffffffffffff16600560019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610b6f57600080fd5b80600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610bbe9190611fc3565b9250508190555080600080828254610bd69190611fc3565b92505081905550600073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610c3b9190611a2b565b60405180910390a350565b60003373ffffffffffffffffffffffffffffffffffffffff16600560019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610ca257600080fd5b81600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610cf19190611fc3565b9250508190555081600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610d479190611ff7565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610dab9190611a2b565b60405180910390a3600190509392505050565b60016020528060005260406000206000915090505481565b60003373ffffffffffffffffffffffffffffffffffffffff16600560019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610e3257600080fd5b6000610e3c61178f565b9050610e466117c2565b600281600001906005811115610e5f57610e5e611b3d565b5b90816005811115610e7357610e72611b3d565b5b8152505087878080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050816060018190525085816080019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050848160a0019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505083816020018181525050806006600084815260200190815260200160002060008201518160000160006101000a81548160ff02191690836005811115610f7d57610f7c611b3d565b5b021790555060208201518160010155604082015181600201556060820151816003019081610fab91906121fc565b5060808201518160040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060a08201518160050160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060c08201518160060160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550905050818573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff167fd8a96bb89e5e9b116b9dbfd20464772a9bf578fb65b08f9ea1201216fc2e302960405160405180910390a4819250505095945050505050565b3373ffffffffffffffffffffffffffffffffffffffff16600560019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461114757600080fd5b600061115285610905565b90506002600581111561116857611167611b3d565b5b8160000151600581111561117f5761117e611b3d565b5b146111bf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111b69061231a565b60405180910390fd5b8373ffffffffffffffffffffffffffffffffffffffff16816080015173ffffffffffffffffffffffffffffffffffffffff16148061122c57508273ffffffffffffffffffffffffffffffffffffffff16816080015173ffffffffffffffffffffffffffffffffffffffff16145b80156112a157508373ffffffffffffffffffffffffffffffffffffffff168160a0015173ffffffffffffffffffffffffffffffffffffffff1614806112a057508273ffffffffffffffffffffffffffffffffffffffff168160a0015173ffffffffffffffffffffffffffffffffffffffff16145b5b6112e0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112d790612386565b60405180910390fd5b816003131580156112f2575060058212155b611331576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611328906123f2565b60405180910390fd5b6003820361138d5760036006600087815260200190815260200160002060000160006101000a81548160ff0219169083600581111561137357611372611b3d565b5b021790555061138783858360200151610c46565b50611433565b600482036113e95760046006600087815260200190815260200160002060000160006101000a81548160ff021916908360058111156113cf576113ce611b3d565b5b02179055506113e383858360200151610c46565b50611432565b600582036114315760056006600087815260200190815260200160002060000160006101000a81548160ff0219169083600581111561142b5761142a611b3d565b5b02179055505b5b5b81857f57167a618f86158dae3f0b66bdb917b5e14a83ca84f04ed67d799d030df0907060405160405180910390a35050505050565b600560019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6004805461149b90611f03565b80601f01602080910402602001604051908101604052809291908181526020018280546114c790611f03565b80156115145780601f106114e957610100808354040283529160200191611514565b820191906000526020600020905b8154815290600101906020018083116114f757829003601f168201915b505050505081565b3373ffffffffffffffffffffffffffffffffffffffff16600560019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461157657600080fd5b80600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546115c59190611ff7565b92505081905550806000808282546115dd9190611ff7565b925050819055503373ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516116429190611a2b565b60405180910390a350565b600081600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461169e9190611fc3565b9250508190555081600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546116f49190611ff7565b925050819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516117589190611a2b565b60405180910390a36001905092915050565b6002602052816000526040600020602052806000526040600020600091509150505481565b600044426040516020016117a4929190612433565b6040516020818303038152906040528051906020012060001c905090565b6040518060e00160405280600060058111156117e1576117e0611b3d565b5b8152602001600081526020016000801916815260200160608152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff1681525090565b600081519050919050565b600082825260208201905092915050565b60005b83811015611890578082015181840152602081019050611875565b60008484015250505050565b6000601f19601f8301169050919050565b60006118b882611856565b6118c28185611861565b93506118d2818560208601611872565b6118db8161189c565b840191505092915050565b6000602082019050818103600083015261190081846118ad565b905092915050565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061193d82611912565b9050919050565b61194d81611932565b811461195857600080fd5b50565b60008135905061196a81611944565b92915050565b6000819050919050565b61198381611970565b811461198e57600080fd5b50565b6000813590506119a08161197a565b92915050565b600080604083850312156119bd576119bc611908565b5b60006119cb8582860161195b565b92505060206119dc85828601611991565b9150509250929050565b60008115159050919050565b6119fb816119e6565b82525050565b6000602082019050611a1660008301846119f2565b92915050565b611a2581611970565b82525050565b6000602082019050611a406000830184611a1c565b92915050565b60008060408385031215611a5d57611a5c611908565b5b6000611a6b8582860161195b565b9250506020611a7c8582860161195b565b9150509250929050565b600080600060608486031215611a9f57611a9e611908565b5b6000611aad8682870161195b565b9350506020611abe8682870161195b565b9250506040611acf86828701611991565b9150509250925092565b600060ff82169050919050565b611aef81611ad9565b82525050565b6000602082019050611b0a6000830184611ae6565b92915050565b600060208284031215611b2657611b25611908565b5b6000611b3484828501611991565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60068110611b7d57611b7c611b3d565b5b50565b6000819050611b8e82611b6c565b919050565b6000611b9e82611b80565b9050919050565b611bae81611b93565b82525050565b611bbd81611970565b82525050565b6000819050919050565b611bd681611bc3565b82525050565b600082825260208201905092915050565b6000611bf882611856565b611c028185611bdc565b9350611c12818560208601611872565b611c1b8161189c565b840191505092915050565b611c2f81611932565b82525050565b600060e083016000830151611c4d6000860182611ba5565b506020830151611c606020860182611bb4565b506040830151611c736040860182611bcd565b5060608301518482036060860152611c8b8282611bed565b9150506080830151611ca06080860182611c26565b5060a0830151611cb360a0860182611c26565b5060c0830151611cc660c0860182611c26565b508091505092915050565b60006020820190508181036000830152611ceb8184611c35565b905092915050565b600060208284031215611d0957611d08611908565b5b6000611d178482850161195b565b91505092915050565b600080fd5b600080fd5b600080fd5b60008083601f840112611d4557611d44611d20565b5b8235905067ffffffffffffffff811115611d6257611d61611d25565b5b602083019150836001820283011115611d7e57611d7d611d2a565b5b9250929050565b600080600080600060808688031215611da157611da0611908565b5b600086013567ffffffffffffffff811115611dbf57611dbe61190d565b5b611dcb88828901611d2f565b95509550506020611dde8882890161195b565b9350506040611def8882890161195b565b9250506060611e0088828901611991565b9150509295509295909350565b6000819050919050565b611e2081611e0d565b8114611e2b57600080fd5b50565b600081359050611e3d81611e17565b92915050565b60008060008060808587031215611e5d57611e5c611908565b5b6000611e6b87828801611991565b9450506020611e7c8782880161195b565b9350506040611e8d8782880161195b565b9250506060611e9e87828801611e2e565b91505092959194509250565b611eb381611932565b82525050565b6000602082019050611ece6000830184611eaa565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680611f1b57607f821691505b602082108103611f2e57611f2d611ed4565b5b50919050565b6000819050919050565b6000819050919050565b6000611f63611f5e611f5984611f34565b611f3e565b611970565b9050919050565b611f7381611f48565b82525050565b6000602082019050611f8e6000830184611f6a565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000611fce82611970565b9150611fd983611970565b9250828203905081811115611ff157611ff0611f94565b5b92915050565b600061200282611970565b915061200d83611970565b925082820190508082111561202557612024611f94565b5b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026120bc7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261207f565b6120c6868361207f565b95508019841693508086168417925050509392505050565b60006120f96120f46120ef84611970565b611f3e565b611970565b9050919050565b6000819050919050565b612113836120de565b61212761211f82612100565b84845461208c565b825550505050565b600090565b61213c61212f565b61214781848461210a565b505050565b5b8181101561216b57612160600082612134565b60018101905061214d565b5050565b601f8211156121b0576121818161205a565b61218a8461206f565b81016020851015612199578190505b6121ad6121a58561206f565b83018261214c565b50505b505050565b600082821c905092915050565b60006121d3600019846008026121b5565b1980831691505092915050565b60006121ec83836121c2565b9150826002028217905092915050565b61220582611856565b67ffffffffffffffff81111561221e5761221d61202b565b5b6122288254611f03565b61223382828561216f565b600060209050601f8311600181146122665760008415612254578287015190505b61225e85826121e0565b8655506122c6565b601f1984166122748661205a565b60005b8281101561229c57848901518255600182019150602085019450602081019050612277565b868310156122b957848901516122b5601f8916826121c2565b8355505b6001600288020188555050505b505050505050565b7f6d61746368206973206e6f742073746172742100000000000000000000000000600082015250565b6000612304601383611861565b915061230f826122ce565b602082019050919050565b60006020820190508181036000830152612333816122f7565b9050919050565b7f61646472657373206973206e6f74206d61746368656421000000000000000000600082015250565b6000612370601783611861565b915061237b8261233a565b602082019050919050565b6000602082019050818103600083015261239f81612363565b9050919050565b7f6d617463682073746174652069732077726f6e67210000000000000000000000600082015250565b60006123dc601583611861565b91506123e7826123a6565b602082019050919050565b6000602082019050818103600083015261240b816123cf565b9050919050565b6000819050919050565b61242d61242882611970565b612412565b82525050565b600061243f828561241c565b60208201915061244f828461241c565b602082019150819050939250505056fea2646970667358221220e904b9ff2afdcf5b11b47d86ab2b56397a6432aa08982ac0a0d2716b2d5361f664736f6c63430008110033",
}

// ContractsABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractsMetaData.ABI instead.
var ContractsABI = ContractsMetaData.ABI

// ContractsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractsMetaData.Bin instead.
var ContractsBin = ContractsMetaData.Bin

// DeployContracts deploys a new Ethereum contract, binding an instance of Contracts to it.
func DeployContracts(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contracts, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// Contracts is an auto generated Go binding around an Ethereum contract.
type Contracts struct {
	ContractsCaller     // Read-only binding to the contract
	ContractsTransactor // Write-only binding to the contract
	ContractsFilterer   // Log filterer for contract events
}

// ContractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractsSession struct {
	Contract     *Contracts        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractsCallerSession struct {
	Contract *ContractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ContractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractsTransactorSession struct {
	Contract     *ContractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ContractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractsRaw struct {
	Contract *Contracts // Generic contract binding to access the raw methods on
}

// ContractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractsCallerRaw struct {
	Contract *ContractsCaller // Generic read-only contract binding to access the raw methods on
}

// ContractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractsTransactorRaw struct {
	Contract *ContractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContracts creates a new instance of Contracts, bound to a specific deployed contract.
func NewContracts(address common.Address, backend bind.ContractBackend) (*Contracts, error) {
	contract, err := bindContracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// NewContractsCaller creates a new read-only instance of Contracts, bound to a specific deployed contract.
func NewContractsCaller(address common.Address, caller bind.ContractCaller) (*ContractsCaller, error) {
	contract, err := bindContracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsCaller{contract: contract}, nil
}

// NewContractsTransactor creates a new write-only instance of Contracts, bound to a specific deployed contract.
func NewContractsTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractsTransactor, error) {
	contract, err := bindContracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsTransactor{contract: contract}, nil
}

// NewContractsFilterer creates a new log filterer instance of Contracts, bound to a specific deployed contract.
func NewContractsFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractsFilterer, error) {
	contract, err := bindContracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractsFilterer{contract: contract}, nil
}

// bindContracts binds a generic wrapper to an already deployed contract.
func bindContracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.ContractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Contracts *ContractsCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Contracts *ContractsSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Contracts.Contract.Allowance(&_Contracts.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Contracts *ContractsCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Contracts.Contract.Allowance(&_Contracts.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Contracts *ContractsCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Contracts *ContractsSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Contracts.Contract.BalanceOf(&_Contracts.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Contracts *ContractsCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Contracts.Contract.BalanceOf(&_Contracts.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Contracts *ContractsCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Contracts *ContractsSession) Decimals() (uint8, error) {
	return _Contracts.Contract.Decimals(&_Contracts.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Contracts *ContractsCallerSession) Decimals() (uint8, error) {
	return _Contracts.Contract.Decimals(&_Contracts.CallOpts)
}

// GetMatch is a free data retrieval call binding the contract method 0x3d092b3d.
//
// Solidity: function getMatch(uint256 _matchId) view returns((uint8,uint256,bytes32,string,address,address,address))
func (_Contracts *ContractsCaller) GetMatch(opts *bind.CallOpts, _matchId *big.Int) (GameMatch, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getMatch", _matchId)

	if err != nil {
		return *new(GameMatch), err
	}

	out0 := *abi.ConvertType(out[0], new(GameMatch)).(*GameMatch)

	return out0, err

}

// GetMatch is a free data retrieval call binding the contract method 0x3d092b3d.
//
// Solidity: function getMatch(uint256 _matchId) view returns((uint8,uint256,bytes32,string,address,address,address))
func (_Contracts *ContractsSession) GetMatch(_matchId *big.Int) (GameMatch, error) {
	return _Contracts.Contract.GetMatch(&_Contracts.CallOpts, _matchId)
}

// GetMatch is a free data retrieval call binding the contract method 0x3d092b3d.
//
// Solidity: function getMatch(uint256 _matchId) view returns((uint8,uint256,bytes32,string,address,address,address))
func (_Contracts *ContractsCallerSession) GetMatch(_matchId *big.Int) (GameMatch, error) {
	return _Contracts.Contract.GetMatch(&_Contracts.CallOpts, _matchId)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Contracts *ContractsCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Contracts *ContractsSession) Name() (string, error) {
	return _Contracts.Contract.Name(&_Contracts.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Contracts *ContractsCallerSession) Name() (string, error) {
	return _Contracts.Contract.Name(&_Contracts.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contracts *ContractsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contracts *ContractsSession) Owner() (common.Address, error) {
	return _Contracts.Contract.Owner(&_Contracts.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contracts *ContractsCallerSession) Owner() (common.Address, error) {
	return _Contracts.Contract.Owner(&_Contracts.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Contracts *ContractsCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Contracts *ContractsSession) Symbol() (string, error) {
	return _Contracts.Contract.Symbol(&_Contracts.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Contracts *ContractsCallerSession) Symbol() (string, error) {
	return _Contracts.Contract.Symbol(&_Contracts.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Contracts *ContractsCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Contracts *ContractsSession) TotalSupply() (*big.Int, error) {
	return _Contracts.Contract.TotalSupply(&_Contracts.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Contracts *ContractsCallerSession) TotalSupply() (*big.Int, error) {
	return _Contracts.Contract.TotalSupply(&_Contracts.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Contracts *ContractsTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Contracts *ContractsSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Approve(&_Contracts.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Contracts *ContractsTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Approve(&_Contracts.TransactOpts, spender, amount)
}

// ApproveReset is a paid mutator transaction binding the contract method 0x22974e64.
//
// Solidity: function approveReset(address target, address spender) returns(bool)
func (_Contracts *ContractsTransactor) ApproveReset(opts *bind.TransactOpts, target common.Address, spender common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "approveReset", target, spender)
}

// ApproveReset is a paid mutator transaction binding the contract method 0x22974e64.
//
// Solidity: function approveReset(address target, address spender) returns(bool)
func (_Contracts *ContractsSession) ApproveReset(target common.Address, spender common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.ApproveReset(&_Contracts.TransactOpts, target, spender)
}

// ApproveReset is a paid mutator transaction binding the contract method 0x22974e64.
//
// Solidity: function approveReset(address target, address spender) returns(bool)
func (_Contracts *ContractsTransactorSession) ApproveReset(target common.Address, spender common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.ApproveReset(&_Contracts.TransactOpts, target, spender)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Contracts *ContractsTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Contracts *ContractsSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Burn(&_Contracts.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Contracts *ContractsTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Burn(&_Contracts.TransactOpts, amount)
}

// CreateMatchByOwner is a paid mutator transaction binding the contract method 0x7f254449.
//
// Solidity: function createMatchByOwner(string _name, address _p1Address, address _p2Address, uint256 _matchPrice) returns(uint256)
func (_Contracts *ContractsTransactor) CreateMatchByOwner(opts *bind.TransactOpts, _name string, _p1Address common.Address, _p2Address common.Address, _matchPrice *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "createMatchByOwner", _name, _p1Address, _p2Address, _matchPrice)
}

// CreateMatchByOwner is a paid mutator transaction binding the contract method 0x7f254449.
//
// Solidity: function createMatchByOwner(string _name, address _p1Address, address _p2Address, uint256 _matchPrice) returns(uint256)
func (_Contracts *ContractsSession) CreateMatchByOwner(_name string, _p1Address common.Address, _p2Address common.Address, _matchPrice *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.CreateMatchByOwner(&_Contracts.TransactOpts, _name, _p1Address, _p2Address, _matchPrice)
}

// CreateMatchByOwner is a paid mutator transaction binding the contract method 0x7f254449.
//
// Solidity: function createMatchByOwner(string _name, address _p1Address, address _p2Address, uint256 _matchPrice) returns(uint256)
func (_Contracts *ContractsTransactorSession) CreateMatchByOwner(_name string, _p1Address common.Address, _p2Address common.Address, _matchPrice *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.CreateMatchByOwner(&_Contracts.TransactOpts, _name, _p1Address, _p2Address, _matchPrice)
}

// MatchEnd is a paid mutator transaction binding the contract method 0x87dadf83.
//
// Solidity: function matchEnd(uint256 _matchId, address _winner, address _losser, int256 _matchState) payable returns()
func (_Contracts *ContractsTransactor) MatchEnd(opts *bind.TransactOpts, _matchId *big.Int, _winner common.Address, _losser common.Address, _matchState *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "matchEnd", _matchId, _winner, _losser, _matchState)
}

// MatchEnd is a paid mutator transaction binding the contract method 0x87dadf83.
//
// Solidity: function matchEnd(uint256 _matchId, address _winner, address _losser, int256 _matchState) payable returns()
func (_Contracts *ContractsSession) MatchEnd(_matchId *big.Int, _winner common.Address, _losser common.Address, _matchState *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.MatchEnd(&_Contracts.TransactOpts, _matchId, _winner, _losser, _matchState)
}

// MatchEnd is a paid mutator transaction binding the contract method 0x87dadf83.
//
// Solidity: function matchEnd(uint256 _matchId, address _winner, address _losser, int256 _matchState) payable returns()
func (_Contracts *ContractsTransactorSession) MatchEnd(_matchId *big.Int, _winner common.Address, _losser common.Address, _matchState *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.MatchEnd(&_Contracts.TransactOpts, _matchId, _winner, _losser, _matchState)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns()
func (_Contracts *ContractsTransactor) Mint(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "mint", amount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns()
func (_Contracts *ContractsSession) Mint(amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Mint(&_Contracts.TransactOpts, amount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns()
func (_Contracts *ContractsTransactorSession) Mint(amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Mint(&_Contracts.TransactOpts, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Contracts *ContractsTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Contracts *ContractsSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Transfer(&_Contracts.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Contracts *ContractsTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Transfer(&_Contracts.TransactOpts, recipient, amount)
}

// TransferByOwner is a paid mutator transaction binding the contract method 0x5a15c373.
//
// Solidity: function transferByOwner(address target, address recipient, uint256 amount) returns(bool)
func (_Contracts *ContractsTransactor) TransferByOwner(opts *bind.TransactOpts, target common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "transferByOwner", target, recipient, amount)
}

// TransferByOwner is a paid mutator transaction binding the contract method 0x5a15c373.
//
// Solidity: function transferByOwner(address target, address recipient, uint256 amount) returns(bool)
func (_Contracts *ContractsSession) TransferByOwner(target common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.TransferByOwner(&_Contracts.TransactOpts, target, recipient, amount)
}

// TransferByOwner is a paid mutator transaction binding the contract method 0x5a15c373.
//
// Solidity: function transferByOwner(address target, address recipient, uint256 amount) returns(bool)
func (_Contracts *ContractsTransactorSession) TransferByOwner(target common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.TransferByOwner(&_Contracts.TransactOpts, target, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Contracts *ContractsTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Contracts *ContractsSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.TransferFrom(&_Contracts.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Contracts *ContractsTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.TransferFrom(&_Contracts.TransactOpts, sender, recipient, amount)
}

// ContractsApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Contracts contract.
type ContractsApprovalIterator struct {
	Event *ContractsApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsApproval represents a Approval event raised by the Contracts contract.
type ContractsApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Contracts *ContractsFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ContractsApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ContractsApprovalIterator{contract: _Contracts.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Contracts *ContractsFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ContractsApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsApproval)
				if err := _Contracts.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Contracts *ContractsFilterer) ParseApproval(log types.Log) (*ContractsApproval, error) {
	event := new(ContractsApproval)
	if err := _Contracts.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsCreateMatchEventIterator is returned from FilterCreateMatchEvent and is used to iterate over the raw logs and unpacked data for CreateMatchEvent events raised by the Contracts contract.
type ContractsCreateMatchEventIterator struct {
	Event *ContractsCreateMatchEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsCreateMatchEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsCreateMatchEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsCreateMatchEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsCreateMatchEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsCreateMatchEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsCreateMatchEvent represents a CreateMatchEvent event raised by the Contracts contract.
type ContractsCreateMatchEvent struct {
	P1      common.Address
	P2      common.Address
	MatchId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCreateMatchEvent is a free log retrieval operation binding the contract event 0xd8a96bb89e5e9b116b9dbfd20464772a9bf578fb65b08f9ea1201216fc2e3029.
//
// Solidity: event CreateMatchEvent(address indexed p1, address indexed p2, uint256 indexed matchId)
func (_Contracts *ContractsFilterer) FilterCreateMatchEvent(opts *bind.FilterOpts, p1 []common.Address, p2 []common.Address, matchId []*big.Int) (*ContractsCreateMatchEventIterator, error) {

	var p1Rule []interface{}
	for _, p1Item := range p1 {
		p1Rule = append(p1Rule, p1Item)
	}
	var p2Rule []interface{}
	for _, p2Item := range p2 {
		p2Rule = append(p2Rule, p2Item)
	}
	var matchIdRule []interface{}
	for _, matchIdItem := range matchId {
		matchIdRule = append(matchIdRule, matchIdItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "CreateMatchEvent", p1Rule, p2Rule, matchIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractsCreateMatchEventIterator{contract: _Contracts.contract, event: "CreateMatchEvent", logs: logs, sub: sub}, nil
}

// WatchCreateMatchEvent is a free log subscription operation binding the contract event 0xd8a96bb89e5e9b116b9dbfd20464772a9bf578fb65b08f9ea1201216fc2e3029.
//
// Solidity: event CreateMatchEvent(address indexed p1, address indexed p2, uint256 indexed matchId)
func (_Contracts *ContractsFilterer) WatchCreateMatchEvent(opts *bind.WatchOpts, sink chan<- *ContractsCreateMatchEvent, p1 []common.Address, p2 []common.Address, matchId []*big.Int) (event.Subscription, error) {

	var p1Rule []interface{}
	for _, p1Item := range p1 {
		p1Rule = append(p1Rule, p1Item)
	}
	var p2Rule []interface{}
	for _, p2Item := range p2 {
		p2Rule = append(p2Rule, p2Item)
	}
	var matchIdRule []interface{}
	for _, matchIdItem := range matchId {
		matchIdRule = append(matchIdRule, matchIdItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "CreateMatchEvent", p1Rule, p2Rule, matchIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsCreateMatchEvent)
				if err := _Contracts.contract.UnpackLog(event, "CreateMatchEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCreateMatchEvent is a log parse operation binding the contract event 0xd8a96bb89e5e9b116b9dbfd20464772a9bf578fb65b08f9ea1201216fc2e3029.
//
// Solidity: event CreateMatchEvent(address indexed p1, address indexed p2, uint256 indexed matchId)
func (_Contracts *ContractsFilterer) ParseCreateMatchEvent(log types.Log) (*ContractsCreateMatchEvent, error) {
	event := new(ContractsCreateMatchEvent)
	if err := _Contracts.contract.UnpackLog(event, "CreateMatchEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsMatchEndEventIterator is returned from FilterMatchEndEvent and is used to iterate over the raw logs and unpacked data for MatchEndEvent events raised by the Contracts contract.
type ContractsMatchEndEventIterator struct {
	Event *ContractsMatchEndEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsMatchEndEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsMatchEndEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsMatchEndEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsMatchEndEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsMatchEndEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsMatchEndEvent represents a MatchEndEvent event raised by the Contracts contract.
type ContractsMatchEndEvent struct {
	MatchId    *big.Int
	MatchState *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMatchEndEvent is a free log retrieval operation binding the contract event 0x57167a618f86158dae3f0b66bdb917b5e14a83ca84f04ed67d799d030df09070.
//
// Solidity: event MatchEndEvent(uint256 indexed matchId, int256 indexed _matchState)
func (_Contracts *ContractsFilterer) FilterMatchEndEvent(opts *bind.FilterOpts, matchId []*big.Int, _matchState []*big.Int) (*ContractsMatchEndEventIterator, error) {

	var matchIdRule []interface{}
	for _, matchIdItem := range matchId {
		matchIdRule = append(matchIdRule, matchIdItem)
	}
	var _matchStateRule []interface{}
	for _, _matchStateItem := range _matchState {
		_matchStateRule = append(_matchStateRule, _matchStateItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "MatchEndEvent", matchIdRule, _matchStateRule)
	if err != nil {
		return nil, err
	}
	return &ContractsMatchEndEventIterator{contract: _Contracts.contract, event: "MatchEndEvent", logs: logs, sub: sub}, nil
}

// WatchMatchEndEvent is a free log subscription operation binding the contract event 0x57167a618f86158dae3f0b66bdb917b5e14a83ca84f04ed67d799d030df09070.
//
// Solidity: event MatchEndEvent(uint256 indexed matchId, int256 indexed _matchState)
func (_Contracts *ContractsFilterer) WatchMatchEndEvent(opts *bind.WatchOpts, sink chan<- *ContractsMatchEndEvent, matchId []*big.Int, _matchState []*big.Int) (event.Subscription, error) {

	var matchIdRule []interface{}
	for _, matchIdItem := range matchId {
		matchIdRule = append(matchIdRule, matchIdItem)
	}
	var _matchStateRule []interface{}
	for _, _matchStateItem := range _matchState {
		_matchStateRule = append(_matchStateRule, _matchStateItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "MatchEndEvent", matchIdRule, _matchStateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsMatchEndEvent)
				if err := _Contracts.contract.UnpackLog(event, "MatchEndEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMatchEndEvent is a log parse operation binding the contract event 0x57167a618f86158dae3f0b66bdb917b5e14a83ca84f04ed67d799d030df09070.
//
// Solidity: event MatchEndEvent(uint256 indexed matchId, int256 indexed _matchState)
func (_Contracts *ContractsFilterer) ParseMatchEndEvent(log types.Log) (*ContractsMatchEndEvent, error) {
	event := new(ContractsMatchEndEvent)
	if err := _Contracts.contract.UnpackLog(event, "MatchEndEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Contracts contract.
type ContractsTransferIterator struct {
	Event *ContractsTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsTransfer represents a Transfer event raised by the Contracts contract.
type ContractsTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Contracts *ContractsFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ContractsTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ContractsTransferIterator{contract: _Contracts.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Contracts *ContractsFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ContractsTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsTransfer)
				if err := _Contracts.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Contracts *ContractsFilterer) ParseTransfer(log types.Log) (*ContractsTransfer, error) {
	event := new(ContractsTransfer)
	if err := _Contracts.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
