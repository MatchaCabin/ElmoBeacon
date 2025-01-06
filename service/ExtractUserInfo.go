package service

import (
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

const (
	gameLogPathCN      = "AppData/LocalLow/SunBorn/少女前线2：追放/Player.log"
	gameLogPathOversea = "AppData/LocalLow/SunBorn/EXILIUM/Player.log"
)

const (
	exprGameDataDir           = `\[Subsystems] Discovering subsystems at path (.+)/UnitySubsystems`
	exprGachaRecordUrlCN      = `"gacha_record_url":"(.*?)"`
	exprGachaRecordUrlOversea = `"k":"gacha_record_url","v":"(.*?)"`
	exprLoginInfoCN           = `"access_token":"(.+?)".+"uid":(\d+)`
	exprLoginInfoOversea      = `"uid":(\d+).+"access_token":"(.+?)"`
)

type gameServer string

const (
	GameServerCN     gameServer = "cn"   //DarkWinter China
	GameServerUS     gameServer = "us"   //DarkWinter USA
	GameServerGlobal gameServer = "intl" //HaoPlay Global
	GameServerJP     gameServer = "jp"   //HaoPlay Japan
	GameServerKR     gameServer = "kr"   //HaoPlay Korea
	GameServerAsia   gameServer = "tw"   //HaoPlay Asia
)

type GameUserInfo struct {
	Uid             uint64
	GameServer      gameServer
	GameDataDir     string
	GameAccessToken string
	GachaRecordUrl  string
}

func ExtractUserInfoCN() (*GameUserInfo, error) {
	userHome, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	logDataBytes, err := os.ReadFile(filepath.Join(userHome, gameLogPathCN))
	if err != nil {
		log.Error().Msg("game log file read failed(CN)")
		return nil, nil
	}

	//extract GameDataDir
	regexpGameDataDir, err := regexp.Compile(exprGameDataDir)
	if err != nil {
		return nil, err
	}
	resultGameDataDir := regexpGameDataDir.FindSubmatch(logDataBytes)
	if resultGameDataDir == nil {
		log.Error().Msg("game data directory path not found(CN)")
		return nil, nil
	}
	gameDataDir := filepath.Join(string(resultGameDataDir[1]), "LocalCache/Data")

	//extract GachaRecordUrl
	regexpGachaRecordUrl, err := regexp.Compile(exprGachaRecordUrlCN)
	if err != nil {
		return nil, err
	}
	resultGachaRecordUrlList := regexpGachaRecordUrl.FindAllSubmatch(logDataBytes, -1)
	if resultGachaRecordUrlList == nil {
		log.Error().Msg("gacha record url not found(CN)")
		return nil, nil
	}
	latestResultGachaRecordUrl := resultGachaRecordUrlList[len(resultGachaRecordUrlList)-1]
	gachaRecordUrl := string(latestResultGachaRecordUrl[1])

	//extract uid and accessToken
	regexpLoginInfo, err := regexp.Compile(exprLoginInfoCN)
	if err != nil {
		return nil, err
	}
	resultLoginInfoList := regexpLoginInfo.FindAllSubmatch(logDataBytes, -1)
	if resultLoginInfoList == nil {
		log.Error().Msg("game login information not found(CN)")
		return nil, nil
	}
	latestResultLoginInfo := resultLoginInfoList[len(resultLoginInfoList)-1]
	gameAccessToken := string(latestResultLoginInfo[1])
	uid, err := strconv.ParseUint(string(latestResultLoginInfo[2]), 10, 64)
	if err != nil {
		return nil, errors.WithMessage(err, "uid format error(CN)")
	}

	return &GameUserInfo{
		Uid:             uid,
		GameServer:      GameServerCN,
		GameDataDir:     gameDataDir,
		GameAccessToken: gameAccessToken,
		GachaRecordUrl:  gachaRecordUrl,
	}, nil
}

func ExtractUserInfoOversea() (*GameUserInfo, error) {
	userHome, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	logDataBytes, err := os.ReadFile(filepath.Join(userHome, gameLogPathOversea))
	if err != nil {
		log.Error().Msg("game log file read failed(Oversea)")
		return nil, nil
	}

	//extract GameDataDir
	regexpGameDataDir, err := regexp.Compile(exprGameDataDir)
	if err != nil {
		return nil, err
	}
	resultGameDataDir := regexpGameDataDir.FindSubmatch(logDataBytes)
	if resultGameDataDir == nil {
		log.Error().Msg("game data directory path not found(Oversea)")
		return nil, nil
	}
	gameDataDir := filepath.Join(string(resultGameDataDir[1]), "LocalCache/Data")

	//extract GachaRecordUrl
	regexpGachaRecordUrl, err := regexp.Compile(exprGachaRecordUrlOversea)
	if err != nil {
		return nil, err
	}
	resultGachaRecordUrlList := regexpGachaRecordUrl.FindAllSubmatch(logDataBytes, -1)
	if resultGachaRecordUrlList == nil {
		log.Error().Msg("gacha record url not found(Oversea)")
		return nil, nil
	}
	latestResultGachaRecordUrl := resultGachaRecordUrlList[len(resultGachaRecordUrlList)-1]
	gachaRecordUrl := string(latestResultGachaRecordUrl[1])

	//extract uid and accessToken
	regexpLoginInfo, err := regexp.Compile(exprLoginInfoOversea)
	if err != nil {
		return nil, err
	}
	resultLoginInfoList := regexpLoginInfo.FindAllSubmatch(logDataBytes, -1)
	if resultLoginInfoList == nil {
		log.Error().Msg("game login information not found(Oversea)")
		return nil, nil
	}
	latestResultLoginInfo := resultLoginInfoList[len(resultLoginInfoList)-1]
	gameAccessToken := string(latestResultLoginInfo[2])
	uid, err := strconv.ParseUint(string(latestResultLoginInfo[1]), 10, 64)
	if err != nil {
		return nil, errors.WithMessage(err, "uid format error(Oversea)")
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
		return nil, errors.Errorf("unknown server:%s", gachaRecordUrl)
	}

	return &GameUserInfo{
		Uid:             uid,
		GameServer:      server,
		GameDataDir:     gameDataDir,
		GameAccessToken: gameAccessToken,
		GachaRecordUrl:  gachaRecordUrl,
	}, nil
}
