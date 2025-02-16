package service

import (
	"ElmoBeacon/pb"
	"ElmoBeacon/util"
)

func GetLangDataMap(gameDataDir, gameServer, lang string) (map[int64]string, error) {
	langMap := make(map[int64]string)
	switch {
	case gameServer == string(GameServerCN):
		var langData pb.LangPackageTableCnData
		err := util.GetTableData(gameDataDir, "", &langData)
		if err != nil {
			return nil, err
		}
		for _, unit := range langData.Units {
			langMap[unit.Id] = unit.Content
		}
	case lang == "zh-TW", lang == "zh-CN":
		var langData pb.LangPackageTableZhtcData
		err := util.GetTableData(gameDataDir, "", &langData)
		if err != nil {
			return nil, err
		}
		for _, unit := range langData.Units {
			langMap[unit.Id] = unit.Content
		}
	case lang == "ja":
		var langData pb.LangPackageTableJajpData
		err := util.GetTableData(gameDataDir, "", &langData)
		if err != nil {
			return nil, err
		}
		for _, unit := range langData.Units {
			langMap[unit.Id] = unit.Content
		}
	case lang == "kr":
		var langData pb.LangPackageTableKokrData
		err := util.GetTableData(gameDataDir, "", &langData)
		if err != nil {
			return nil, err
		}
		for _, unit := range langData.Units {
			langMap[unit.Id] = unit.Content
		}
	default:
		var langData pb.LangPackageTableEnusData
		err := util.GetTableData(gameDataDir, "", &langData)
		if err != nil {
			return nil, err
		}
		for _, unit := range langData.Units {
			langMap[unit.Id] = unit.Content
		}
	}

	return langMap, nil
}
