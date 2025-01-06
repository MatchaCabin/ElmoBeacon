package service

import (
	"ElmoBeacon/pb"
	"ElmoBeacon/util"
)

func GetGachaPoolTypeList(gameDataDir string, gameServer string) (gachaPoolTypeList []int64, err error) {
	var gachaTypeListData pb.GachaTypeListData
	if gameServer == string(GameServerCN) {
		err = util.GetTableData(gameDataDir, "", &gachaTypeListData)
		if err != nil {
			return nil, err
		}
	} else {
		err = util.GetTableData(gameDataDir, gameServer, &gachaTypeListData)
		if err != nil {
			return nil, err
		}
	}

	for _, unit := range gachaTypeListData.Units {
		gachaPoolTypeList = append(gachaPoolTypeList, unit.Id)
	}

	return gachaPoolTypeList, nil
}
