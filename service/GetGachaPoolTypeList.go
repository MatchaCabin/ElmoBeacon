package service

import (
	"ElmoBeacon/pb"
	"ElmoBeacon/util"
)

func GetGachaPoolTypeList(gameDataDir string) (gachaPoolTypeList []int64, err error) {
	var gachaTypeListData pb.GachaTypeListData

	err = util.GetTableData(gameDataDir, "", &gachaTypeListData)
	if err != nil {
		return nil, err
	}

	for _, unit := range gachaTypeListData.Units {
		gachaPoolTypeList = append(gachaPoolTypeList, unit.Id)
	}

	return gachaPoolTypeList, nil
}
