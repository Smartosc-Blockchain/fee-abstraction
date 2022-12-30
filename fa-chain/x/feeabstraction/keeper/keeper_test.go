package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"

	transfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
	ibctesting "github.com/cosmos/ibc-go/v3/testing"

	"github.com/Smartosc-Blockchain/fa-chain/app/apptesting"
	appparams "github.com/Smartosc-Blockchain/fa-chain/app/params"
	"github.com/Smartosc-Blockchain/fa-chain/x/interchainquery/keeper"
	"github.com/Smartosc-Blockchain/fa-chain/x/interchainquery/types"
)

const (
	HostChainId = "OSMO"
)

type KeeperTestSuite struct {
	apptesting.AppTestHelper
}

func (s *KeeperTestSuite) SetupTest() {
	s.Setup()
	s.CreateTransferChannel(HostChainId)
}

// sending ufa from fachain to osmosis
func (s *KeeperTestSuite) MockIBCTransfer() error {
	timeoutHeight := clienttypes.NewHeight(0, 110)

	amount, _ := sdk.NewIntFromString("1000000000") // 2^63 (one above int64)
	coinToSendToB := sdk.NewCoin(appparams.DefaultBondDenom, amount)

	// send from chainA to chainB
	msg := transfertypes.NewMsgTransfer(s.TransferPath.EndpointA.ChannelConfig.PortID, s.TransferPath.EndpointA.ChannelID, coinToSendToB, s.Chain.SenderAccount.GetAddress().String(), s.HostChain.SenderAccount.GetAddress().String(), timeoutHeight, 0)
	res, err := s.Chain.SendMsgs(msg)
	if err != nil {
		return err
	}

	packet, err := ibctesting.ParsePacketFromEvents(res.GetEvents())
	if err != nil {
		return err
	}

	// relay send
	if err = s.TransferPath.RelayPacket(packet); err != nil {
		return err
	}

	return nil
}

func (s *KeeperTestSuite) GetMsgServer() types.MsgServer {
	return keeper.NewMsgServerImpl(s.App.InterchainqueryKeeper)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
