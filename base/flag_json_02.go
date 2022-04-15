package main

import (
	"encoding/json"
	"fmt"
)

type AuthToken struct {
	Code int `json:"code"`
	Data struct {
		UID        string `json:"uid"`
		DeviceInfo struct {
			HwModel     string `json:"hwModel"`
			DeviceID    string `json:"deviceId"`
			ProductId   string `json:"productId"`
			CsType      string `json:"csType"`
			Vendor      string `json:"vendor"`
			ClientSwVer string `json:"clientSwVer"`
			Sn          string `json:"sn"`
			DeviceId    string `json:"deviceId"`
			Mac         string `json:"mac"`
		} `json:"deviceInfo"`
		DeviceID      string `json:"deviceId"`
		UserID        string `json:"userId"`
		ProductID     string `json:"productId"`
		ExpireAt      int64  `json:"expireAt"`
		Category      string `json:"category"`
		ApplicationID string `json:"applicationId"`
		CsType        string `json:"csType"`
	} `json:"data"`
	Message string `json:"message"`
}

func main() {

	token := "{\n\t\"code\":0,\n    \"data\":{\n\t\t\"uid\":\"baad1917d207d42f6506a66b5118d943\",\n\t\t\"deviceInfo\":{\n\t\t\t\"hwModel\":\"ThinkPad X1 Fold Gen 1\",\n\t\t\t\"deviceId\":\"TEST-APP-1dac859b60bcdd045dcc604f32e1597e\",\n\t\t\t\"productId\":\"\",\n\t\t\t\"csType\":\"xiaoxin\",\n\t\t\t\"vendor\":\"lva team\",\n\t\t\t\"clientSwVer\":\"v2.34\",\n\t\t\t\"sn\":\"\",\n\t\t\t\"deviceId\":\"TEST-APP-1dac859b60bcdd045dcc604f32e1597e\",\n\t\t\t\"mac\":\"56B53E87-314D-46B9-9D0E-49D49D93641E\"\n\t\t},\n\t\t\"deviceId\":\"TEST-APP-1dac859b60bcdd045dcc604f32e1597e\",\n\t\t\"userId\":\"baad1917d207d42f6506a66b5118d943\",\n\t\t\"productId\":\"PC_LVA\",\n\t\t\"expireAt\":1645761722441,\"category\":\"pc\",\n\t\t\"applicationId\":\"5ef9a71e94df8bcf53512d7e\",\n\t\t\"csType\":\"xiaoxin\"\n\t},\n\t\"message\":\"ok\"\n}"

	authToken := &AuthToken{}
	json.Unmarshal([]byte(token), authToken)
	fmt.Println(authToken.Code)

}
