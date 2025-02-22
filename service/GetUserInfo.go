package service

import (
	"github.com/rs/zerolog/log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func GetUserInfoCN() *GameUserInfo {
	userHome, err := os.UserHomeDir()
	if err != nil {
		log.Err(err).Msg("Failed to get user home dir")
		return nil
	}

	logDataBytes, err := os.ReadFile(filepath.Join(userHome, gameLogPathCN))
	if err != nil {
		log.Err(err).Msg("Failed to read game log file(CN)")
		return nil
	}

	//extract GameDataDir
	regexpGameDataDir, err := regexp.Compile(exprGameDataDir)
	if err != nil {
		log.Err(err).Msg("Failed to compile exprGameDataDir(CN)")
		return nil
	}
	resultGameDataDir := regexpGameDataDir.FindSubmatch(logDataBytes)
	if resultGameDataDir == nil {
		log.Err(err).Msg("Failed to find game data directory(CN)")
		return nil
	}
	gameDataDir := filepath.Join(string(resultGameDataDir[1]), "LocalCache/Data")

	//extract GachaRecordUrl
	regexpGachaRecordUrl, err := regexp.Compile(exprGachaRecordUrlCN)
	if err != nil {
		log.Err(err).Msg("Failed to compile exprGachaRecordUrl(CN)")
		return nil
	}
	resultGachaRecordUrlList := regexpGachaRecordUrl.FindAllSubmatch(logDataBytes, -1)
	if resultGachaRecordUrlList == nil {
		log.Err(err).Msg("Failed to find gacha record url(CN)")
		return nil
	}
	latestResultGachaRecordUrl := resultGachaRecordUrlList[len(resultGachaRecordUrlList)-1]
	gachaRecordUrl := string(latestResultGachaRecordUrl[1])

	//extract uid and accessToken
	regexpLoginInfo, err := regexp.Compile(exprLoginInfoCN)
	if err != nil {
		log.Err(err).Msg("Failed to compile exprLoginInfo(CN)")
		return nil
	}
	resultLoginInfoList := regexpLoginInfo.FindAllSubmatch(logDataBytes, -1)
	if resultLoginInfoList == nil {
		log.Err(err).Msg("Failed to find game login information(CN)")
		return nil
	}
	latestResultLoginInfo := resultLoginInfoList[len(resultLoginInfoList)-1]
	gameAccessToken := string(latestResultLoginInfo[1])
	uid, err := strconv.ParseUint(string(latestResultLoginInfo[2]), 10, 64)
	if err != nil {
		log.Err(err).Msgf("Failed to parse uid:'%s'(CN)", latestResultLoginInfo[2])
		return nil
	}

	return &GameUserInfo{
		Uid:             uid,
		GameServer:      GameServerCN,
		GameDataDir:     gameDataDir,
		GameAccessToken: gameAccessToken,
		GachaRecordUrl:  gachaRecordUrl,
	}
}

func GetUserInfoOS() *GameUserInfo {
	userHome, err := os.UserHomeDir()
	if err != nil {
		log.Err(err).Msg("Failed to get user home dir")
		return nil
	}

	logDataBytes, err := os.ReadFile(filepath.Join(userHome, gameLogPathOversea))
	if err != nil {
		log.Err(err).Msg("Failed to read game log file(Oversea)")
		return nil
	}

	//extract GameDataDir
	regexpGameDataDir, err := regexp.Compile(exprGameDataDir)
	if err != nil {
		log.Err(err).Msg("Failed to compile exprGameDataDir(Oversea)")
		return nil
	}
	resultGameDataDir := regexpGameDataDir.FindSubmatch(logDataBytes)
	if resultGameDataDir == nil {
		log.Err(err).Msg("Failed to find game data directory(Oversea)")
		return nil
	}
	gameDataDir := filepath.Join(string(resultGameDataDir[1]), "LocalCache/Data")

	//extract GachaRecordUrl
	regexpGachaRecordUrl, err := regexp.Compile(exprGachaRecordUrlOversea)
	if err != nil {
		log.Err(err).Msg("Failed to compile exprGachaRecordUrl(Oversea)")
		return nil
	}
	resultGachaRecordUrlList := regexpGachaRecordUrl.FindAllSubmatch(logDataBytes, -1)
	if resultGachaRecordUrlList == nil {
		log.Err(err).Msg("Failed to find gacha record url(Oversea)")
		return nil
	}
	latestResultGachaRecordUrl := resultGachaRecordUrlList[len(resultGachaRecordUrlList)-1]
	gachaRecordUrl := string(latestResultGachaRecordUrl[1])

	//extract uid and accessToken
	regexpLoginInfo, err := regexp.Compile(exprLoginInfoOversea)
	if err != nil {
		log.Err(err).Msg("Failed to compile exprLoginInfo(Oversea)")
		return nil
	}
	resultLoginInfoList := regexpLoginInfo.FindAllSubmatch(logDataBytes, -1)
	if resultLoginInfoList == nil {
		log.Err(err).Msg("Failed to find game login information(Oversea)")
		return nil
	}
	latestResultLoginInfo := resultLoginInfoList[len(resultLoginInfoList)-1]
	gameAccessToken := string(latestResultLoginInfo[2])
	uid, err := strconv.ParseUint(string(latestResultLoginInfo[1]), 10, 64)
	if err != nil {
		log.Err(err).Msgf("Failed to parse uid:'%s'(Oversea)", latestResultLoginInfo[1])
		return nil
	}

	//determine the server
	var server gameServer
	switch {
	case strings.Contains(gachaRecordUrl, "gf2-gacha-record-us"):
		server = GameServerUS
	case strings.Contains(gachaRecordUrl, "gf2-gacha-record-intl"):
		server = GameServerGlobal
	case strings.Contains(gachaRecordUrl, "gf2-gacha-record-jp"):
		server = GameServerJP
	case strings.Contains(gachaRecordUrl, "gf2-gacha-record-kr"):
		server = GameServerKR
	case strings.Contains(gachaRecordUrl, "gf2-gacha-record-asia"):
		server = GameServerAsia
	default:
		log.Error().Str("gachaRecordUrl", gachaRecordUrl).Msgf("Failed to determine server(Oversea)")
		return nil
	}

	return &GameUserInfo{
		Uid:             uid,
		GameServer:      server,
		GameDataDir:     gameDataDir,
		GameAccessToken: gameAccessToken,
		GachaRecordUrl:  gachaRecordUrl,
	}
}

func GetUserInfoList() (userInfoList []*GameUserInfo) {
	userInfoCN := GetUserInfoCN()
	if userInfoCN != nil {
		userInfoList = append(userInfoList, userInfoCN)
	}
	userInfoOS := GetUserInfoOS()
	if userInfoOS != nil {
		userInfoList = append(userInfoList, userInfoOS)
	}
	return userInfoList
}
