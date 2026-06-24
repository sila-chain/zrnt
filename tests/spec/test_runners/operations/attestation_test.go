package operations

import (
	"fmt"
	"github.com/sila-chain/zrnt/sila/beacon/deneb"
	"testing"

	"github.com/sila-chain/zrnt/sila/beacon/altair"
	"github.com/sila-chain/zrnt/sila/beacon/common"
	"github.com/sila-chain/zrnt/sila/beacon/phase0"
	"github.com/sila-chain/zrnt/tests/spec/test_util"
)

type AttestationTestCase struct {
	test_util.BaseTransitionTest
	Attestation phase0.Attestation
}

func (c *AttestationTestCase) Load(t *testing.T, forkName test_util.ForkName, readPart test_util.TestPartReader) {
	c.BaseTransitionTest.Load(t, forkName, readPart)
	test_util.LoadSpecObj(t, "attestation", &c.Attestation, readPart)
}

func (c *AttestationTestCase) Run() error {
	epc, err := common.NewEpochsContext(c.Spec, c.Pre)
	if err != nil {
		return err
	}
	if s, ok := c.Pre.(phase0.Phase0PendingAttestationsBeaconState); ok {
		return phase0.ProcessAttestation(c.Spec, epc, s, &c.Attestation)
	} else if s, ok := c.Pre.(altair.AltairLikeBeaconState); ok {
		switch c.Fork {
		case "altair", "bellatrix", "capella":
			return altair.ProcessAttestation(c.Spec, epc, s, &c.Attestation)
		case "deneb":
			return deneb.ProcessAttestation(c.Spec, epc, s, &c.Attestation)
		default:
			return fmt.Errorf("unrecognized fork: %s", c.Fork)
		}
	} else {
		return fmt.Errorf("unrecognized state type: %T", s)
	}
}

func TestAttestation(t *testing.T) {
	test_util.RunTransitionTest(t, test_util.AllForks, "operations", "attestation",
		func() test_util.TransitionTest { return new(AttestationTestCase) })
}
