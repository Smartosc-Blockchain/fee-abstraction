package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/Smartosc-Blockchain/fa-chain/app/apptesting"
	"github.com/Smartosc-Blockchain/fa-chain/x/interchainquery/keeper"
	"github.com/Smartosc-Blockchain/fa-chain/x/interchainquery/types"
)

const (
	HostChainId = "GAIA"
)

type KeeperTestSuite struct {
	apptesting.AppTestHelper
}

func (s *KeeperTestSuite) SetupTest() {
	s.Setup()
	s.CreateTransferChannel(HostChainId)
}

func (s *KeeperTestSuite) GetMsgServer() types.MsgServer {
	return keeper.NewMsgServerImpl(s.App.InterchainqueryKeeper)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
