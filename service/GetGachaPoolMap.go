package service

import (
	"ElmoBeacon/pb"
	"ElmoBeacon/util"
	"strconv"
	"strings"
)

type GachaPoolInfo struct {
	UpItem    map[int64]struct{}
	Rank5Item map[int64]struct{}
	Rank4Item map[int64]struct{}
	Rank3Item map[int64]struct{}
}

func GetGachaPoolMap(gameDataDir string, gameServer string) (map[int64]GachaPoolInfo, error) {
	var gachaData pb.GachaData
	if gameServer == string(GameServerCN) {
		err := util.GetTableData(gameDataDir, "", &gachaData)
		if err != nil {
			return nil, err
		}
	} else {
		err := util.GetTableData(gameDataDir, gameServer, &gachaData)
		if err != nil {
			return nil, err
		}
	}

	poolMap := make(map[int64]GachaPoolInfo)
	for _, unit := range gachaData.Units {
		gachaPoolInfo := GachaPoolInfo{
			UpItem:    make(map[int64]struct{}),
			Rank5Item: make(map[int64]struct{}),
			Rank4Item: make(map[int64]struct{}),
			Rank3Item: make(map[int64]struct{}),
		}

		//get up char
		if unit.GunUpCharacter != "" {
			upCharId, err := strconv.ParseInt(strings.ReplaceAll(unit.GunUpCharacter, "：", ":"), 10, 64)
			if err != nil {
				return nil, err
			}
			gachaPoolInfo.UpItem[upCharId] = struct{}{}
		}

		//get up weapon
		if unit.WeaponUpItem != "" {
			splitUpWeaponGroupStr := strings.Split(strings.ReplaceAll(unit.WeaponUpItem, "：", ":"), ",")
			for _, upWeaponGroupStr := range splitUpWeaponGroupStr {
				upWeaponGroupStr = strings.TrimPrefix(strings.TrimPrefix(strings.TrimPrefix(upWeaponGroupStr, "3:"), "4:"), "5:")
				splitUpWeaponIdStr := strings.Split(upWeaponGroupStr, ":")
				for _, upWeaponIdStr := range splitUpWeaponIdStr {
					upWeaponId, err := strconv.ParseInt(upWeaponIdStr, 10, 64)
					if err != nil {
						return nil, err
					}
					gachaPoolInfo.UpItem[upWeaponId] = struct{}{}
				}
			}
		}

		//get char list
		if unit.RateDesGun != "" {
			splitCharGroupStr := strings.Split(strings.ReplaceAll(unit.RateDesGun, "：", ":"), ",")
			for _, charGroupStr := range splitCharGroupStr {
				//get rank5 char
				if strings.HasPrefix(charGroupStr, "5:") {
					charGroupStr = strings.TrimPrefix(charGroupStr, "5:")
					splitCharIdStr := strings.Split(charGroupStr, ":")
					for _, charIdStr := range splitCharIdStr {
						charId, err := strconv.ParseInt(charIdStr, 10, 64)
						if err != nil {
							return nil, err
						}
						gachaPoolInfo.Rank5Item[charId] = struct{}{}
					}
				}
				//get rank4 char
				if strings.HasPrefix(charGroupStr, "4:") {
					charGroupStr = strings.TrimPrefix(charGroupStr, "4:")
					splitCharIdStr := strings.Split(charGroupStr, ":")
					for _, charIdStr := range splitCharIdStr {
						charId, err := strconv.ParseInt(charIdStr, 10, 64)
						if err != nil {
							return nil, err
						}
						gachaPoolInfo.Rank4Item[charId] = struct{}{}
					}
				}
				//get rank3 char
				if strings.HasPrefix(charGroupStr, "3:") {
					charGroupStr = strings.TrimPrefix(charGroupStr, "3:")
					splitCharIdStr := strings.Split(charGroupStr, ":")
					for _, charIdStr := range splitCharIdStr {
						charId, err := strconv.ParseInt(charIdStr, 10, 64)
						if err != nil {
							return nil, err
						}
						gachaPoolInfo.Rank3Item[charId] = struct{}{}
					}
				}
			}
		}

		//get weapon list
		if unit.RateDesWeapon != "" {
			splitWeaponGroupStr := strings.Split(strings.ReplaceAll(unit.RateDesWeapon, "：", ":"), ",")
			for _, weaponGroupStr := range splitWeaponGroupStr {
				//get rank5 weapon
				if strings.HasPrefix(weaponGroupStr, "5:") {
					weaponGroupStr = strings.TrimPrefix(weaponGroupStr, "5:")
					splitWeaponIdStr := strings.Split(weaponGroupStr, ":")
					for _, weaponIdStr := range splitWeaponIdStr {
						weaponId, err := strconv.ParseInt(weaponIdStr, 10, 64)
						if err != nil {
							return nil, err
						}
						gachaPoolInfo.Rank5Item[weaponId] = struct{}{}
					}
				}
				//get rank4 weapon
				if strings.HasPrefix(weaponGroupStr, "4:") {
					weaponGroupStr = strings.TrimPrefix(weaponGroupStr, "4:")
					splitWeaponIdStr := strings.Split(weaponGroupStr, ":")
					for _, weaponIdStr := range splitWeaponIdStr {
						weaponId, err := strconv.ParseInt(weaponIdStr, 10, 64)
						if err != nil {
							return nil, err
						}
						gachaPoolInfo.Rank4Item[weaponId] = struct{}{}
					}
				}
				//get rank3 weapon
				if strings.HasPrefix(weaponGroupStr, "3:") {
					weaponGroupStr = strings.TrimPrefix(weaponGroupStr, "3:")
					splitWeaponIdStr := strings.Split(weaponGroupStr, ":")
					for _, weaponIdStr := range splitWeaponIdStr {
						weaponId, err := strconv.ParseInt(weaponIdStr, 10, 64)
						if err != nil {
							return nil, err
						}
						gachaPoolInfo.Rank3Item[weaponId] = struct{}{}
					}
				}
			}
		}

		for _, itemId := range unit.OptionalItem {
			gachaPoolInfo.UpItem[itemId] = struct{}{}
			gachaPoolInfo.Rank5Item[itemId] = struct{}{}
		}

		poolMap[unit.Id] = gachaPoolInfo
	}

	return poolMap, nil
}
