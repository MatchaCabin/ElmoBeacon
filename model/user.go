package model

type User struct {
	Id           int64  `json:"id"`
	Uid          uint64 `json:"uid"`
	GameServer   string `json:"server"`
	GameDataDir  string `json:"gameDataDir"`
	LastBBSToken string `json:"lastBBSToken"`
}
