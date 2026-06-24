package random

import (
	"testing"

	"github.com/sila-chain/zrnt/tests/spec/test_util"
)

func TestRandomBlocks(t *testing.T) {
	test_util.RunTransitionTest(t, test_util.AllForks, "random", "random",
		func() test_util.TransitionTest { return new(test_util.BlocksTestCase) })
}
