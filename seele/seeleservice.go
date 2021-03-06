/**
*  @file
*  @copyright defined in go-seele/LICENSE
 */

package seele

import (
	"github.com/seeleteam/go-seele/common"
	"github.com/seeleteam/go-seele/core"
	"github.com/seeleteam/go-seele/core/types"
	"github.com/seeleteam/go-seele/log"
	"github.com/seeleteam/go-seele/p2p"
	"github.com/seeleteam/go-seele/rpc"
)

// SeeleService implements full node service.
type SeeleService struct {
	networkID     uint64
	seeleProtocol *SeeleProtocol
	log           *log.SeeleLog
	coinbase      common.Address // account address that mining rewards will be send to.

	txPool *core.TransactionPool
	chain  *core.Blockchain
}

func (s *SeeleService) TxPool() *core.TransactionPool { return s.txPool }
func (s *SeeleService) BlockChain() *core.Blockchain  { return s.chain }
func (s *SeeleService) NetVersion() uint64            { return s.networkID }

// ApplyTransaction applys a transaction
func (s *SeeleService) ApplyTransaction(coinbase common.Address, tx *types.Transaction) error {
	// TODO
	return nil
}

// NewSeeleService create SeeleService
func NewSeeleService(networkID uint64, log *log.SeeleLog) (s *SeeleService, err error) {
	s = &SeeleService{
		networkID: networkID,
		log:       log,
	}

	s.seeleProtocol, err = NewSeeleProtocol(networkID, log)
	return s, err
}

// Protocols implements node.Service, returning all the currently configured
// network protocols to start.
func (s *SeeleService) Protocols() (protos []p2p.ProtocolInterface) {
	protos = append(protos, s.seeleProtocol)
	return
}

// Start implements node.Service, starting goroutines needed by SeeleService.
func (s *SeeleService) Start(srvr *p2p.Server) error {
	return nil
}

// Stop implements node.Service, terminating all internal goroutines.
func (s *SeeleService) Stop() error {
	s.seeleProtocol.Stop()
	return nil
}

// APIs implements node.Service, returning the collection of RPC services the seele package offers.
func (s *SeeleService) APIs() (apis []rpc.API) {
	//TODO add other api interface, for example consensus engine
	return append(apis, []rpc.API{
		{
			Namespace: "seele",
			Version:   "1.0",
			Service:   NewPublicSeeleAPI(s),
			Public:    true,
		},
	}...)
}
