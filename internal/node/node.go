package node

import (
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/mxmkiv/go-blockchain/internal/miner"
	"github.com/mxmkiv/go-blockchain/internal/model"
	"github.com/mxmkiv/go-blockchain/internal/p2p"
	"go.etcd.io/bbolt"
	"go.uber.org/zap"
)

// global
//const GlobalSeedNodeHost = ""
//const GlobalSeedNodePort = ""
//const GlobalSeedNodeAddress = GlobalSeedNodeHost + ":" + GlobalSeedNodePort

// local
const LocalSeedNodeHost = "localhost"
const LocalSeedNodePort = "8888"
const LocalSeedNodeAddress = LocalSeedNodeHost + ":" + LocalSeedNodePort

const LocalNodeHost = "localhost"

type Node struct {
	NodeAddress      string // "host:port"
	NodePort         string
	MaxNodeConnected int

	Peers []net.Conn // connections to others node

	Logger *zap.Logger

	//services
	MsgManager *p2p.MsgManager
	Miner      *miner.Miner

	IsSeedNode  bool
	AddressList []string // list of node addresses that are already onchain

	Mempool []model.Transaction
	DB      *bbolt.DB // use bbolt as a database for recording blocks and wallet balances on disk
}

func IsSeedNodeExist(launchMode string) bool {

	var addr string
	if launchMode == "global" {
		addr = "" //Ip address from config // GlobalSeedNodeAddress
	} else {
		addr = LocalSeedNodeAddress
	}

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return false
	}

	manager := p2p.NewManager()
	manager.Init(conn, 1)
	manager.SendMessage(p2p.Message{Type: p2p.MsgEcho})
	defer conn.Close()

	return true
}

func GetAddressForNode(launchMode string) string {
	if launchMode == "global" {
		/*

			global ip from config + selected port (scan or user choise)

		*/

		return ""
	}

	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		fmt.Println("[error] allocating port for node")
		panic(err)
	}
	defer listener.Close()

	_, port, _ := net.SplitHostPort(listener.Addr().String())

	return LocalNodeHost + ":" + port
}

func NewNode(launchMode string, logger *zap.Logger, msgManager *p2p.MsgManager, miner *miner.Miner, MaxConn int, MemPoolSize int) *Node {

	if launchMode == "global" {

		/*

			TODO
			создание адреса ноды из введеного пользователем ip + порт либо выбранный сканированием либо указанный пользователем
			соаздние бд для записи (./data/node.db)

			проверить нужна ли функция GetAddressForNode и IsSeedNodeExist при глобальном запуске

		*/

		return &Node{}
	}

	Node := &Node{
		MaxNodeConnected: MaxConn,
		Mempool:          make([]model.Transaction, 0, MemPoolSize),
		Peers:            make([]net.Conn, 0, MaxConn),
	}

	if IsSeedNodeExist(launchMode) {
		// true - seed node already exist
		Node.IsSeedNode = false
		Node.NodeAddress = GetAddressForNode(launchMode)
		Node.AddressList = nil
	} else {
		// false - seed node is missing
		Node.IsSeedNode = true
		Node.NodeAddress = LocalSeedNodeAddress
		Node.AddressList = make([]string, 0, 256)
	}

	// get port for uniq db file
	port := func(addr string) string {
		mass := strings.Split(addr, ":")
		return mass[1]
	}(Node.NodeAddress)

	/*

		создание папки data в корне проекта

	*/

	pathDB := "./data/node_" + port + ".db"

	db, err := bbolt.Open(pathDB, 0600, nil)
	if err != nil {
		log.Fatal("[error] create db ", err)
	}
	// закрыть соединение при отключении ноды

	Node.DB = db
	Node.NodePort = port

	return Node

}

func (n *Node) Start() {

}
