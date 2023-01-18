package model

import (
	"context"
	"fmt"
	log "lecture/WBA-BC-Project-04/logger"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	"github.com/holiman/uint256"
	"go.mongodb.org/mongo-driver/bson"
)

type Block struct {
	BlockHash    string        `bson:"blockHash"`
	BlockNumber  uint64        `bson:"blockNumber"`
	GasLimit     uint64        `bson:"gasLimit"`
	GasUsed      uint64        `bson:"gasUsed"`
	Time         uint64        `bson:"timestamp"`
	Nonce        uint64        `bson:"nonce"`
	Transactions []Transaction `bson:"transactions"`
}

type Transaction struct {
	TxHash      string `bson:"hash"`
	From        string `bson:"from"`
	To          string `bson:"to"` // 컨트랙트의 경우 nil 반환
	Nonce       uint64 `bson:"nonce"`
	GasPrice    uint64 `bson:"gasPrice"`
	GasLimit    uint64 `bson:"gasLimit"`
	Amount      uint64 `bson:"amount"`
	BlockHash   string `bson:"blockHash"`
	BlockNumber uint64 `bson:"blockNumber"`
}

type createMatchEvent struct {
	P1         string `bson:"p1"`
	P2         string `bson:"p2"`
	MatchId    string `bson:"matchId"`
	MatchState int    `bson:"matchState"`
}

func (p *Model) RunDaemon(quit chan bool) {
	fmt.Println("run daemon")
	// ethclint 초기화
	client, err := ethclient.Dial(p.daemon.url)
	if err != nil {
		log.Error(err.Error())
	}

	targetAddress := common.HexToAddress(p.game.contractAddress)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{targetAddress},
	}
	// subscribe
	logs := make(chan types.Log)
	sub := event.Resubscribe(10*time.Second, func(ctx context.Context) (event.Subscription, error) {
		return client.SubscribeFilterLogs(context.Background(), query, logs)
	})

	logCreateMatchSig := []byte("CreateMatchEvent(address,address,uint256)")
	logMatchEndSig := []byte("MatchEndEvent(uint256,int256)")
	logCreateMatchSigHash := crypto.Keccak256Hash(logCreateMatchSig)
	logMatchEndSigHash := crypto.Keccak256Hash(logMatchEndSig)

	for {
		select {
		case <-quit:
			fmt.Println("stop daemon")
			return
		case err := <-sub.Err():
			fmt.Println("sub.Err:", err.Error())
			log.Error(err.Error())
			return
		case scribelog := <-logs:
			switch scribelog.Topics[0].Hex() {
			case logCreateMatchSigHash.Hex():
				var createMatchEvent = createMatchEvent{}

				matchId, err := uint256.FromHex(scribelog.Topics[3].Hex())
				if err != nil {
					log.Error(err.Error())
				}
				createMatchEvent.P1 = scribelog.Topics[1].String()
				createMatchEvent.P2 = scribelog.Topics[2].String()
				createMatchEvent.MatchId = scribelog.Topics[3].String()
				createMatchEvent.MatchState = 1
				_, err = p.gameCol.InsertOne(context.TODO(), createMatchEvent)
				if err != nil {
					log.Error(err.Error())
				}
				fmt.Printf("SaveCreateMatchEvent: %v ", matchId)
			case logMatchEndSigHash.Hex():
				matchState := new(big.Int)
				matchState.SetBytes(scribelog.Topics[2].Bytes())

				updateFilter := bson.M{"matchId": scribelog.Topics[1].String(), "matchState": 1}
				update := bson.D{
					{Key: "$set", Value: bson.D{
						{Key: "matchState", Value: matchState.Int64()},
					}},
				}
				_, err = p.gameCol.UpdateOne(context.TODO(), updateFilter, update)
				if err != nil {
					log.Error(err.Error())
				}

				matchId, err := uint256.FromHex(scribelog.Topics[1].Hex())
				if err != nil {
					log.Error(err.Error())
				}
				fmt.Printf("SaveEndMatchEvent: %v ", matchId)
			}
		}
	}
}
func (p *Model) SaveBlock(block *Block) error {
	_, err := p.gameCol.InsertOne(context.TODO(), block)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	fmt.Println("Insert succeed")

	return nil
}

func (p *Model) SaveCreateMatchEvent(event *createMatchEvent) error {
	_, err := p.gameCol.InsertOne(context.TODO(), event)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	fmt.Println("Insert succeed")

	return nil
}
