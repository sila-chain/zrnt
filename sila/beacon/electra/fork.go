package electra

import (
	"errors"

	"github.com/sila-chain/zrnt/sila/beacon/common"
	"github.com/sila-chain/zrnt/sila/beacon/deneb"
)

func UpgradeToElectra(spec *common.Spec, epc *common.EpochsContext, pre *deneb.BeaconStateView) (*BeaconStateView, error) {
	return nil, errors.New("upgrade of deneb state to electra state is not supported")
}
