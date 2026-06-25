package execution

import (
	"context"

	"github.com/sila-chain/zrnt/sila/beacon/bellatrix"
	"github.com/sila-chain/zrnt/sila/beacon/capella"
	"github.com/sila-chain/zrnt/sila/beacon/common"
	"github.com/sila-chain/zrnt/sila/beacon/deneb"
)

type NoOpExecutionEngine struct{}

func (n NoOpExecutionEngine) DenebNotifySilaNewPayload(ctx context.Context, executionPayload *deneb.SilaExecutionPayload, parentBeaconBlockRoot common.Root) (valid bool, err error) {
	return true, nil
}

func (n NoOpExecutionEngine) DenebIsValidVersionedHashes(ctx context.Context, payload *deneb.SilaExecutionPayload, versionedHashes []common.Hash32) (bool, error) {
	return true, nil
}

func (n NoOpExecutionEngine) DenebIsValidBlockHash(ctx context.Context, payload *deneb.SilaExecutionPayload, parentBeaconBlockRoot common.Root) (bool, error) {
	return true, nil
}

func (n NoOpExecutionEngine) CapellaNotifySilaNewPayload(ctx context.Context, executionPayload *capella.SilaExecutionPayload) (valid bool, err error) {
	return true, nil
}

func (n NoOpExecutionEngine) CapellaIsValidBlockHash(ctx context.Context, payload *capella.SilaExecutionPayload) (bool, error) {
	return true, nil
}

func (n NoOpExecutionEngine) BellatrixNotifySilaNewPayload(ctx context.Context, executionPayload *bellatrix.SilaExecutionPayload) (valid bool, err error) {
	return true, nil
}

func (n NoOpExecutionEngine) BellatrixIsValidBlockHash(ctx context.Context, payload *bellatrix.SilaExecutionPayload) (bool, error) {
	return true, nil
}

var _ bellatrix.ExecutionEngine = (*NoOpExecutionEngine)(nil)
var _ capella.ExecutionEngine = (*NoOpExecutionEngine)(nil)
var _ deneb.ExecutionEngine = (*NoOpExecutionEngine)(nil)

var _ common.ExecutionEngine = (*NoOpExecutionEngine)(nil)
