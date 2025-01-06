package service

import (
	"ElmoBeacon/pb"
	"ElmoBeacon/util"
)

func GetItemDataMap(gameDataDir string, gameServer string) (map[int64]*pb.Item, error) {
	var itemData pb.ItemData
	if gameServer == string(GameServerCN) {
		err := util.GetTableData(gameDataDir, "", &itemData)
		if err != nil {
			return nil, err
		}
	} else {
		err := util.GetTableData(gameDataDir, gameServer, &itemData)
		if err != nil {
			return nil, err
		}
	}

	itemMap := make(map[int64]*pb.Item)
	for i, unit := range itemData.Units {
		itemMap[unit.Id] = itemData.Units[i]
	}

	return itemMap, nil
}
