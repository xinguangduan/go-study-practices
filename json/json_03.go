package main

import (
	"encoding/json"
	"fmt"
)

type Token struct {
	Code    int    `json:"code"`
	Data    Data   `json:"data"`
	Message string `json:"message"`
}
type DeviceInfo struct {
	HwModel     string `json:"hwModel"`
	DeviceID    string `json:"deviceId"`
	ProductId   string `json:"productId"`
	CsType      string `json:"csType"`
	Vendor      string `json:"vendor"`
	ClientSwVer string `json:"clientSwVer"`
	Sn          string `json:"sn"`
	Mac         string `json:"mac"`
}
type Data struct {
	UID           string     `json:"uid"`
	DeviceInfo    DeviceInfo `json:"deviceInfo"`
	DeviceID      string     `json:"deviceId"`
	UserID        string     `json:"userId"`
	ProductID     string     `json:"productId"`
	ExpireAt      int64      `json:"expireAt"`
	Category      string     `json:"category"`
	ApplicationID string     `json:"applicationId"`
	CsType        string     `json:"csType"`
}

func main() {
	di := DeviceInfo{
		HwModel:     "sksd",
		DeviceID:    "12334",
		ProductId:   "sbbbb",
		CsType:      "sc",
		Vendor:      "ssd",
		ClientSwVer: "dd",
		Sn:          "fff",
		Mac:         "ss-3223-4-3-2-3-4",
	}

	token := Token{}
	token.Code = 0
	token.Data = Data{
		UID:           "1234",
		DeviceInfo:    di,
		DeviceID:      "weee",
		UserID:        "aaaa",
		ProductID:     "dddd",
		ExpireAt:      0,
		Category:      "pc",
		ApplicationID: "addk",
		CsType:        "pc",
	}
	token.Message = "ok"

	jsonBytes, _ := json.Marshal(token)

	fmt.Println(string(jsonBytes))
}
