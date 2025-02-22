package handler

import (
	"ElmoBeacon/db"
	"ElmoBeacon/model"
	"ElmoBeacon/request"
	"ElmoBeacon/service"
	"fmt"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"golang.org/x/exp/slices"
	"strconv"
)

type SyncResult struct {
	Server   string
	Uid      uint64
	PoolType int64
	Count    int
}

func (a *App) SyncRecordsFromServer() (syncResultList []SyncResult, err error) {
	// Step.1 extract user info from the logs
	gameUserInfoList := service.GetUserInfoList()

	if len(gameUserInfoList) == 0 {
		return nil, errors.New("Failed to find any game user info, please check whether a valid login has been performed in this machine recently.")
	}

	for _, gameUserInfo := range gameUserInfoList {
		// Step.2 check if the user exists
		cond := model.User{GameServer: string(gameUserInfo.GameServer), Uid: gameUserInfo.Uid} //double conditions prevent users from having the same UID on different servers.It seems highly improbable, but it does exist among the users registered at the start of different servers.
		hasUser, err := db.Engine.Get(&cond)
		if err != nil {
			log.Error().Err(err).Msg("")
			return nil, errors.New("Failed to query user")
		}
		var userId int64
		user := model.User{
			Uid:         gameUserInfo.Uid,
			GameServer:  string(gameUserInfo.GameServer),
			GameDataDir: gameUserInfo.GameDataDir,
		}
		if !hasUser {
			_, err = db.Engine.Insert(&user)
			if err != nil {
				log.Error().Err(err).Msg("")
				return nil, errors.New("Failed to insert user")
			}
			userId = user.Id
		} else {
			userId = cond.Id
			_, err = db.Engine.ID(userId).Update(user)
			if err != nil {
				log.Error().Err(err).Msg("")
				return nil, errors.New("Failed to update user")
			}
		}

		// Step.3 fetch gacha records from official server until it matches the latest local record
		gachaPoolTypeList, err := service.GetGachaPoolTypeList(gameUserInfo.GameDataDir, string(gameUserInfo.GameServer))
		if err != nil {
			log.Error().Err(err).Msg("")
			return nil, errors.New("Failed to get gacha pool type list")
		}

		for _, poolType := range gachaPoolTypeList {
			latestLocalRecord := model.Record{
				UserId:   userId,
				PoolType: poolType,
			}
			_, err = db.Engine.Desc("id").Get(&latestLocalRecord)
			if err != nil {
				log.Error().Err(err).Msg("")
				return nil, errors.New("Failed to query latest local record")
			}

			var incrementalRecordList []model.Record
			var next string
		loopFetchingRemoteRecord:
			for {
				remoteRecordList, err := request.FetchGachaRecordList(gameUserInfo.GachaRecordUrl, gameUserInfo.GameAccessToken, next, poolType)
				if err != nil {
					log.Error().Err(err).Msg("")
					return nil, errors.New("Failed to fetch gacha record list")
				}
				for _, remoteRecord := range remoteRecordList.RecordList {
					if remoteRecord.GachaTimestamp == latestLocalRecord.Timestamp && remoteRecord.ItemId == latestLocalRecord.ItemId {
						break loopFetchingRemoteRecord
					} else {
						incrementalRecordList = append(incrementalRecordList, model.Record{
							UserId:    userId,
							PoolType:  poolType,
							PoolId:    remoteRecord.PoolId,
							ItemId:    remoteRecord.ItemId,
							Timestamp: remoteRecord.GachaTimestamp,
						})
					}
				}
				if remoteRecordList.Next != "" {
					next = remoteRecordList.Next
				} else {
					break
				}
			}
			log.Info().Str("server", string(gameUserInfo.GameServer)).Uint64("uid", gameUserInfo.Uid).Int64("poolType", poolType).Int("count", len(incrementalRecordList)).Msg("")

			// Step.4 merge the incremental gacha records into the database
			if len(incrementalRecordList) > 0 {
				slices.Reverse(incrementalRecordList)
				var lastTimestamp, order int64
				for i, record := range incrementalRecordList {
					if record.Timestamp != lastTimestamp {
						order = 0
					} else {
						order++
					}
					virtualId, _ := strconv.ParseUint(fmt.Sprintf("%d%03d", record.Timestamp, order), 10, 64)
					incrementalRecordList[i].Id = virtualId
					lastTimestamp = record.Timestamp
				}

				_, err = db.Engine.Insert(&incrementalRecordList)
				if err != nil {
					log.Error().Err(err).Msg("")
					return nil, errors.New("Failed to insert incremental record list")
				}

				syncResultList = append(syncResultList, SyncResult{
					Server:   string(gameUserInfo.GameServer),
					Uid:      gameUserInfo.Uid,
					PoolType: poolType,
					Count:    len(incrementalRecordList),
				})
			}
		}
	}

	return syncResultList, nil
}
