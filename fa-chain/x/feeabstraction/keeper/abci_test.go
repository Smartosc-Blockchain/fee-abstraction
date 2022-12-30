package keeper_test

// go test -v -run ^TestKeeperTestSuite/TestIdentifyChain$ github.com/Smartosc-Blockchain/fa-chain/x/feeabstraction/keeper
func (s KeeperTestSuite) TestIdentifyChain() {
	s.SetupTest()

	// Send token
	err := s.MockIBCTransfer()
	s.Suite.Require().NoError(err)

	// run begin blocker
	s.App.FAKeeper.BeginBlocker(s.Ctx)

	// identify correct denom
}
