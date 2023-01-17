package model

import (
	"context"
	"fmt"
	"io/ioutil"
	log "lecture/WBA-BC-Project-04/logger"
	"math/big"
	"path/filepath"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	"github.com/holiman/uint256"
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
	P1      string `bson:"p1"`
	P2      string `bson:"p2"`
	MatchId string `bson:"matchId"`
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
	abiPath, _ := filepath.Abs("./build/Game.abi")
	file, err := ioutil.ReadFile(abiPath)
	if err != nil {
		log.Error(err.Error())
	}

	sub := event.Resubscribe(10*time.Second, func(ctx context.Context) (event.Subscription, error) {
		return client.SubscribeFilterLogs(context.Background(), query, logs)
	})
	gameAbi, err := abi.JSON(strings.NewReader(string(file)))
	fmt.Println("subscribe: ", gameAbi)

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
				// DB insert
				err = p.SaveCreateMatchEvent(&createMatchEvent)
				if err != nil {
					log.Error(err.Error())
				}
				fmt.Printf("SaveCreateMatchEvent: %v ", matchId)
			case logMatchEndSigHash.Hex():
			default:

				block, err := client.BlockByNumber(context.Background(), new(big.Int).SetUint64(scribelog.BlockNumber))

				// 블록 구조체 생성
				b := Block{
					BlockHash:    block.Hash().Hex(),
					BlockNumber:  block.Number().Uint64(),
					GasLimit:     block.GasLimit(),
					GasUsed:      block.GasUsed(),
					Time:         block.Time(),
					Nonce:        block.Nonce(),
					Transactions: make([]Transaction, 0),
				}

				// 트랜잭션 추출
				txs := block.Transactions()
				if len(txs) > 0 {
					for _, tx := range txs {
						msg, err := tx.AsMessage(types.LatestSignerForChainID(tx.ChainId()), block.BaseFee())
						if err != nil {
							log.Error(err.Error())
						}

						// 트랜잭션 구조체 생성
						t := Transaction{
							TxHash:      tx.Hash().Hex(),
							To:          "", // 디폴트 값 처리
							From:        msg.From().Hex(),
							Nonce:       tx.Nonce(),
							GasPrice:    tx.GasPrice().Uint64(),
							GasLimit:    tx.Gas(),
							Amount:      tx.Value().Uint64(),
							BlockHash:   block.Hash().Hex(),
							BlockNumber: block.Number().Uint64(),
						}

						if tx.To() != nil {
							t.To = tx.To().Hex()
						}

						b.Transactions = append(b.Transactions, t)
					}
				}

				// DB insert
				err = p.SaveBlock(&b)
				if err != nil {
					log.Error(err.Error())
				}
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
