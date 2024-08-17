package worker

import (
	"goro/model"
	"goro/service"
)

func SyncAssets() {
	conf := service.GetContext()
	assets, err := service.GetAssets()
	if err != nil {
		error.Error(err)
	}

	dbAssets := make([]*model.Asset, len(assets))

	for i, a := range assets {
		dbAssets[i] = &model.Asset{
			ExternalId:   a.ExternalId,
			Name:         a.Name,
			Status:       a.Status,
			Type:         a.Details.Type,
			MinSize:      a.MinSize,
			MaxPrecision: a.MaxPrecision,
		}

		if len(a.SupportedNetworks) == 0 {
			continue
		}

		//defaultNetwork := a.SupportedNetworks[0]
		//
		//dbAssets[i].DefaultNetwork = model.Network{
		//	ExternalId:          defaultNetwork.ExternalId,
		//	Name:                defaultNetwork.Name,
		//	Status:              defaultNetwork.Status,
		//	MinWithdrawalAmount: defaultNetwork.MinWithdrawalAmount,
		//	MaxWithdrawalAmount: defaultNetwork.MaxWithdrawalAmount,
		//}
	}

	conf.Db.Create(dbAssets)

}
