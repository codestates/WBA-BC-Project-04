package model

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	cont "lecture/WBA-BC-Project-04/contracts"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/holiman/uint256"
)

type CreateMatch struct {
	RoomName   string `json:"roomName" bson:"roomName"`
	P1Address  string `json:"p1Address" bson:"p1Address"`
	P2Address  string `json:"p2Address" bson:"p2Address"`
	MatchPrice int    `json:"matchPrice" bson:"matchPrice"`
}

type EndMatch struct {
	MatchId    string `json:"matchId" bson:"matchId"`
	Winner     string `json:"winner" bson:"winner"`
	Losser     string `json:"losser" bson:"losser"`
	MatchState int    `json:"matchState" bson:"matchState"`
}

func (p *Model) CreateMatchModel(match CreateMatch) error {

	// 블록체인 네트워크와 연결할 클라이언트를 생성하기 위한 rpc url 연결
	client, err := ethclient.Dial(p.game.netUrl)
	if err != nil {
		log.Error("client 에러", err.Error())
	}

	// 기본키 지정
	privateKey, err := crypto.HexToECDSA(p.game.privateKey)
	if err != nil {
		log.Error("HexToECDSA 에러", err.Error())
		return err
	}

	// privatekey로부터 publickey를 거쳐 자신의 address 변환
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Error("fail convert, publickey")
		return errors.New("fail convert")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// 컨트랙트 어드레스
	tokenAddress := common.HexToAddress(p.game.contractAddress)
	instance, err := cont.NewContracts(tokenAddress, client)
	if err != nil {
		log.Error("NewContractsCaller 에러", err.Error())
	}

	targetAddress := common.HexToAddress(match.P1Address)
	opponentAddress := common.HexToAddress(match.P2Address)
	// 현재 계정의 nonce를 가져옴. 다음 트랜잭션에서 사용할 nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Error("PendingNonceAt 에러", err.Error())
		return err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Error("SuggestGasPrice 에러", err.Error())
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	price := big.NewInt(int64(match.MatchPrice))
	tx, err := instance.CreateMatchByOwner(auth, match.RoomName, targetAddress, opponentAddress, price)
	if err != nil {
		log.Error("CreateMatchByOwner", err.Error())
		return err
	}
	fmt.Printf("tx sent: %s", tx.Hash().Hex()) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870

	return nil
}

func (p *Model) EndMatchModel(match EndMatch) error {

	// 블록체인 네트워크와 연결할 클라이언트를 생성하기 위한 rpc url 연결
	client, err := ethclient.Dial(p.game.netUrl)
	if err != nil {
		log.Error("client 에러", err.Error())
	}

	// 기본키 지정
	privateKey, err := crypto.HexToECDSA(p.game.privateKey)
	if err != nil {
		log.Error("HexToECDSA 에러", err.Error())
		return err
	}

	// privatekey로부터 publickey를 거쳐 자신의 address 변환
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Error("fail convert, publickey")
		return errors.New("fail convert")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// 컨트랙트 어드레스
	tokenAddress := common.HexToAddress(p.game.contractAddress)
	instance, err := cont.NewContracts(tokenAddress, client)
	if err != nil {
		log.Error("NewContractsCaller 에러", err.Error())
	}

	winnerAddress := common.HexToAddress(match.Winner)
	losserAddress := common.HexToAddress(match.Losser)
	// 현재 계정의 nonce를 가져옴. 다음 트랜잭션에서 사용할 nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Error("PendingNonceAt 에러", err.Error())
		return err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Error("SuggestGasPrice 에러", err.Error())
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice
	matchState := big.NewInt(int64(match.MatchState))
	matchId := new(big.Int)
	matchIdtemp, _ := uint256.FromHex(match.MatchId)
	matchId.Set(matchIdtemp.ToBig())

	tx, err := instance.MatchEnd(auth, matchId, winnerAddress, losserAddress, matchState)
	if err != nil {
		log.Error("CreateMatchByOwner", err.Error())
		return err
	}

	fmt.Printf("tx sent: %s", tx.Hash().Hex()) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870

	return nil
}
