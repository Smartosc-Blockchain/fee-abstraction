package keeper_test

// go test -v -run ^TestKeeperTestSuite/TestIdentifyChain$ github.com/Smartosc-Blockchain/fa-chain/x/feeabstraction/keeper
func (s KeeperTestSuite) TestIdentifyChain() {
	s.SetupTest()

	s.App.FAKeeper.BeginBlocker(s.Ctx)
}
