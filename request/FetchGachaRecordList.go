package request

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type GachaResponseBody struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data"`
}

type GachaRecordListData struct {
	RecordList []struct {
		PoolId         int64 `json:"pool_id"`
		ItemId         int64 `json:"item"`
		GachaTimestamp int64 `json:"time"`
	} `json:"list"`
	Next string `json:"next"`
}

func FetchGachaRecordList(gachaUrl, accessToken, next string, poolType int64) (data GachaRecordListData, err error) {
	values := url.Values{}
	if next != "" {
		values.Set("next", next)
	}
	values.Set("type_id", strconv.FormatInt(poolType, 10))

	req, err := http.NewRequest("POST", gachaUrl, strings.NewReader(values.Encode()))
	if err != nil {
		return GachaRecordListData{}, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", accessToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return GachaRecordListData{}, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return GachaRecordListData{}, err
	}

	var body GachaResponseBody
	err = json.Unmarshal(bodyBytes, &body)
	if err != nil {
		return GachaRecordListData{}, err
	}

	if body.Code == 0 {
		err = json.Unmarshal(body.Data, &data)
		if err != nil {
			return GachaRecordListData{}, err
		}

		return data, nil
	} else {
		return GachaRecordListData{}, errors.Errorf("%s(Code:%d)", body.Message, body.Code)
	}
}
