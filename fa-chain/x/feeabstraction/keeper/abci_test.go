package keeper_test

func (s KeeperTestSuite) TestIdentifyChain() {
	s.SetupTest()

	s.App.FAKeeper.BeginBlocker(s.Ctx)
}
