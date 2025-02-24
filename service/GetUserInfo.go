package service

import (
	"github.com/pkg/errors"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func GetUserInfoCN() (*GameUserInfo, error) {
	userHome, err := os.UserHomeDir()
	if err != nil {
		return nil, errors.New("Failed to get user home dir")
	}

	logDataBytes, err := os.ReadFile(filepath.Join(userHome, gameLogPathCN))
	if err != nil {
		return nil, errors.New("Failed to read game log file(CN)")
	}

	//extract GameDataDir
	regexpGameDataDir, err := regexp.Compile(exprGameDataDir)
	if err != nil {
		return nil, errors.New("Failed to compile exprGameDataDir(CN)")
	}
	resultGameDataDir := regexpGameDataDir.FindSubmatch(logDataBytes)
	if resultGameDataDir == nil {
		return nil, errors.New("Failed to find game data directory(CN)")
	}
	gameDataDir := filepath.Join(string(resultGameDataDir[1]), "LocalCache/Data")

	//extract GachaRecordUrl
	regexpGachaRecordUrl, err := regexp.Compile(exprGachaRecordUrlCN)
	if err != nil {
		return nil, errors.New("Failed to compile exprGachaRecordUrl(CN)")
	}
	resultGachaRecordUrlList := regexpGachaRecordUrl.FindAllSubmatch(logDataBytes, -1)
	if resultGachaRecordUrlList == nil {
		return nil, errors.New("Failed to find gacha record url(CN)")
	}
	latestResultGachaRecordUrl := resultGachaRecordUrlList[len(resultGachaRecordUrlList)-1]
	gachaRecordUrl := string(latestResultGachaRecordUrl[1])

	//extract uid and accessToken
	regexpLoginInfo, err := regexp.Compile(exprLoginInfoCN)
	if err != nil {
		return nil, errors.New("Failed to compile exprLoginInfo(CN)")
	}
	resultLoginInfoList := regexpLoginInfo.FindAllSubmatch(logDataBytes, -1)
	if resultLoginInfoList == nil {
		return nil, errors.New("Failed to find game login information(CN)")
	}
	latestResultLoginInfo := resultLoginInfoList[len(resultLoginInfoList)-1]
	gameAccessToken := string(latestResultLoginInfo[1])
	uid, err := strconv.ParseUint(string(latestResultLoginInfo[2]), 10, 64)
	if err != nil {
		return nil, errors.Errorf("Failed to parse uid:'%s'(CN)", latestResultLoginInfo[2])
	}

	return &GameUserInfo{
		Uid:             uid,
		GameServer:      GameServerCN,
		GameDataDir:     gameDataDir,
		GameAccessToken: gameAccessToken,
		GachaRecordUrl:  gachaRecordUrl,
	}, nil
}

func GetUserInfoOS() (*GameUserInfo, error) {
	userHome, err := os.UserHomeDir()
	if err != nil {
		return nil, errors.New("Failed to get user home dir")
	}

	logDataBytes, err := os.ReadFile(filepath.Join(userHome, gameLogPathOversea))
	if err != nil {
		return nil, errors.New("Failed to read game log file(Oversea)")
	}

	//extract GameDataDir
	regexpGameDataDir, err := regexp.Compile(exprGameDataDir)
	if err != nil {
		return nil, errors.New("Failed to compile exprGameDataDir(Oversea)")
	}
	resultGameDataDir := regexpGameDataDir.FindSubmatch(logDataBytes)
	if resultGameDataDir == nil {
		return nil, errors.New("Failed to find game data directory(Oversea)")
	}
	gameDataDir := filepath.Join(string(resultGameDataDir[1]), "LocalCache/Data")

	//extract GachaRecordUrl
	regexpGachaRecordUrl, err := regexp.Compile(exprGachaRecordUrlOversea)
	if err != nil {
		return nil, errors.New("Failed to compile exprGachaRecordUrl(Oversea)")
	}
	resultGachaRecordUrlList := regexpGachaRecordUrl.FindAllSubmatch(logDataBytes, -1)
	if resultGachaRecordUrlList == nil {
		return nil, errors.New("Failed to find gacha record url(Oversea)")
	}
	latestResultGachaRecordUrl := resultGachaRecordUrlList[len(resultGachaRecordUrlList)-1]
	gachaRecordUrl := string(latestResultGachaRecordUrl[1])

	//extract uid and accessToken
	regexpLoginInfo, err := regexp.Compile(exprLoginInfoOversea)
	if err != nil {
		return nil, errors.New("Failed to compile exprLoginInfo(Oversea)")
	}
	resultLoginInfoList := regexpLoginInfo.FindAllSubmatch(logDataBytes, -1)
	if resultLoginInfoList == nil {
		return nil, errors.New("Failed to find game login information(Oversea)")
	}
	latestResultLoginInfo := resultLoginInfoList[len(resultLoginInfoList)-1]
	gameAccessToken := string(latestResultLoginInfo[2])
	uid, err := strconv.ParseUint(string(latestResultLoginInfo[1]), 10, 64)
	if err != nil {
		return nil, errors.Errorf("Failed to parse uid:'%s'(Oversea)", latestResultLoginInfo[1])
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
		return nil, errors.Errorf("Failed to determine server(Oversea),gacha url:%s", gachaRecordUrl)
	}

	return &GameUserInfo{
		Uid:             uid,
		GameServer:      server,
		GameDataDir:     gameDataDir,
		GameAccessToken: gameAccessToken,
		GachaRecordUrl:  gachaRecordUrl,
	}, nil
}
