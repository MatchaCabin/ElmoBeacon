package handler

import (
	"ElmoBeacon/db"
	"ElmoBeacon/model"
	"ElmoBeacon/service"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"slices"
)

//
//func (a *App) GetPoolTypeList(userId int64) (poolTypeList []int64, err error) {
//	var user model.User
//	_, err = db.Engine.ID(userId).Get(&user)
//	if err != nil {
//		return nil, err
//	}
//
//	gachaPoolTypeList, err := service.GetGachaPoolTypeList(user.GameDataDir, user.GameServer)
//	if err != nil {
//		log.Error().Err(err).Msg("")
//		return nil, errors.New("error occurred when get gacha pool type list")
//	}
//
//	return gachaPoolTypeList, nil
//}

type DisplayRecord struct {
	Name      string
	Icon      string
	Count     int64
	Timestamp int64
	IsMissing bool
}

type PoolInfo struct {
	StoredCount  int64           `json:"storedCount"`
	RecordList   []DisplayRecord `json:"recordList"`
	TotalCount   int64           `json:"totalCount"`
	Rank5Count   int64           `json:"rank5Count"`
	Rank4Count   int64           `json:"rank4Count"`
	Rank3Count   int64           `json:"rank3Count"`
	Rank5Rate    float64         `json:"rank5Rate"`
	Rank4Rate    float64         `json:"rank4Rate"`
	Rank3Rate    float64         `json:"rank3Rate"`
	Rank5Avg     float64         `json:"rank5Avg"`
	Rank5UpAvg   float64         `json:"rank5UpAvg"`
	MissingCount int64           `json:"missingCount"`
	MissingRate  float64         `json:"missingRate"`
}

func (a *App) GetPoolInfo(userId, poolType int64) (poolInfo PoolInfo, err error) {
	var user model.User
	_, err = db.Engine.ID(userId).Get(&user)
	if err != nil {
		return PoolInfo{}, err
	}
	lang := model.Setting{Key: "lang"}
	_, err = db.Engine.Get(&lang)
	if err != nil {
		return PoolInfo{}, err
	}

	langDataMap, err := service.GetLangDataMap(user.GameDataDir, user.GameServer, lang.Value)
	if err != nil {
		log.Error().Err(err).Msg("")
		return PoolInfo{}, errors.New("error occurred when get lang map")
	}

	itemDataMap, err := service.GetItemDataMap(user.GameDataDir, user.GameServer)
	if err != nil {
		log.Error().Err(err).Msg("")
		return PoolInfo{}, errors.New("error occurred when get item map")
	}

	dollIconMap, err := service.GetDollIconMap(user.GameDataDir, user.GameServer)
	if err != nil {
		log.Error().Err(err).Msg("")
		return PoolInfo{}, errors.New("error occurred when get doll icon map")
	}

	weaponIconMap, err := service.GetWeaponIconMap(user.GameDataDir, user.GameServer)
	if err != nil {
		log.Error().Err(err).Msg("")
		return PoolInfo{}, errors.New("error occurred when get weapon icon map")
	}

	gachaPoolMap, err := service.GetGachaPoolMap(user.GameDataDir, user.GameServer)
	if err != nil {
		log.Error().Err(err).Msg("")
		return PoolInfo{}, errors.New("error occurred when get gacha pool map")
	}

	var recordList []model.Record
	err = db.Engine.OrderBy("id").Find(&recordList, &model.Record{UserId: userId, PoolType: poolType})
	if err != nil {
		log.Error().Err(err).Msg("")
		return PoolInfo{}, errors.New("error occurred when get user record list by poolType from db")
	}
	if len(recordList) > 0 {
		var displayRecordList []DisplayRecord
		var count, judgeCount int64
		var isPreMissing bool
		for _, record := range recordList {
			gachaPoolInfo := gachaPoolMap[record.PoolId]
			if item, hasItemData := itemDataMap[record.ItemId]; hasItemData {
				poolInfo.TotalCount++
				count++
				if _, isRank5 := gachaPoolInfo.Rank5Item[item.Id]; isRank5 {
					if text, hasLangData := langDataMap[item.Name.Id]; hasLangData {
						poolInfo.Rank5Count++
						if !isPreMissing {
							judgeCount++
						}
						//todo handle special pool
						var isMissing bool
						if _, isUp := gachaPoolInfo.UpItem[item.Id]; len(gachaPoolInfo.UpItem) > 0 && !isUp {
							isMissing = true
							poolInfo.MissingCount++
						}

						icon := "Avatar_Head_Unknown_Spine.png"
						if item.Type == 10 {
							if i, hasIcon := dollIconMap[item.Id]; hasIcon {
								icon = i
							} else {
								return PoolInfo{}, errors.Errorf("error occurred when get doll icon by id %d", item.Id)
							}
						} else if item.Type == 20 {
							if i, hasIcon := weaponIconMap[item.Id]; hasIcon {
								icon = i
							} else {
								return PoolInfo{}, errors.Errorf("error occurred when get weapon icon by id %d", item.Id)
							}
						}

						displayRecordList = append(displayRecordList, DisplayRecord{
							Name:      text,
							Icon:      icon,
							Count:     count,
							Timestamp: record.Timestamp,
							IsMissing: isMissing,
						})
						isPreMissing = isMissing
					} else {
						return PoolInfo{}, errors.Errorf("error occurred when get item name by id:%d", record.ItemId)
					}
					count = 0
				} else if _, isRank4 := gachaPoolInfo.Rank4Item[item.Id]; isRank4 {
					poolInfo.Rank4Count++
				} else if _, isRank3 := gachaPoolInfo.Rank3Item[item.Id]; isRank3 {
					poolInfo.Rank3Count++
				} else {
					log.Warn().Int64("poolId", record.PoolId).Int64("itemId", item.Id).Msg("unknown rank")
				}
			} else {
				return PoolInfo{}, errors.Errorf("error occurred when get item data by id:%d", record.ItemId)
			}
		}
		poolInfo.StoredCount = count
		if poolInfo.MissingCount > 0 {
			poolInfo.MissingRate = float64(poolInfo.MissingCount) / float64(judgeCount)
		}
		if poolInfo.TotalCount > 0 {
			poolInfo.Rank5Rate = float64(poolInfo.Rank5Count) / float64(poolInfo.TotalCount)
			poolInfo.Rank4Rate = float64(poolInfo.Rank4Count) / float64(poolInfo.TotalCount)
			poolInfo.Rank3Rate = float64(poolInfo.Rank3Count) / float64(poolInfo.TotalCount)
		}

		if poolInfo.Rank5Count > 0 {
			poolInfo.Rank5Avg = float64(poolInfo.TotalCount) / float64(poolInfo.Rank5Count)
		}
		if poolInfo.Rank5Count-poolInfo.MissingCount > 0 {
			poolInfo.Rank5UpAvg = float64(poolInfo.TotalCount) / float64(poolInfo.Rank5Count-poolInfo.MissingCount)
		}

		slices.Reverse(displayRecordList)
		poolInfo.RecordList = displayRecordList
	}

	return poolInfo, nil
}
