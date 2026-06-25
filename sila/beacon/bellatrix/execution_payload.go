package bellatrix

import (
	"context"
	"errors"
	"fmt"

	"github.com/sila-chain/zrnt/sila/beacon/common"
)

func ProcessSilaExecutionPayload(ctx context.Context, spec *common.Spec, state ExecutionTrackingBeaconState, executionPayload *SilaExecutionPayload, engine ExecutionEngine) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if engine == nil {
		return errors.New("nil execution engine")
	}

	slot, err := state.Slot()
	if err != nil {
		return err
	}

	completed := true
	if s, ok := state.(ExecutionUpgradeBeaconState); ok {
		var err error
		completed, err = s.IsTransitionCompleted()
		if err != nil {
			return err
		}
	}
	if completed {
		latestExecHeader, err := state.LatestSilaExecutionPayloadHeader()
		if err != nil {
			return err
		}
		parent, err := latestExecHeader.Raw()
		if err != nil {
			return fmt.Errorf("failed to read previous header: %v", err)
		}
		if executionPayload.ParentHash != parent.BlockHash {
			return fmt.Errorf("expected parent hash %s in execution payload, but got %s",
				parent.BlockHash, executionPayload.ParentHash)
		}
	}

	// verify random
	mixes, err := state.RandaoMixes()
	if err != nil {
		return err
	}
	expectedMix, err := mixes.GetRandomMix(spec.SlotToEpoch(slot))
	if err != nil {
		return err
	}
	if executionPayload.PrevRandao != expectedMix {
		return fmt.Errorf("invalid random data %s, expected %s", executionPayload.PrevRandao, expectedMix)
	}

	// verify timestamp
	genesisTime, err := state.GenesisTime()
	if err != nil {
		return err
	}
	if expectedTime, err := spec.TimeAtSlot(slot, genesisTime); err != nil {
		return fmt.Errorf("slot or genesis time in state is corrupt, cannot compute time: %v", err)
	} else if executionPayload.Timestamp != expectedTime {
		return fmt.Errorf("state at slot %d, genesis time %d, expected execution payload time %d, but got %d",
			slot, genesisTime, expectedTime, executionPayload.Timestamp)
	}

	if valid, err := VerifyAndNotifySilaNewPayload(ctx, engine, &SilaNewPayloadRequest{SilaExecutionPayload: executionPayload}); err != nil {
		return fmt.Errorf("unexpected problem in execution engine when inserting block %s (height %d), err: %v",
			executionPayload.BlockHash, executionPayload.BlockNumber, err)
	} else if !valid {
		return fmt.Errorf("execution engine says payload is invalid: %s (height %d)",
			executionPayload.BlockHash, executionPayload.BlockNumber)
	}

	return state.SetLatestSilaExecutionPayloadHeader(executionPayload.Header(spec))
}
