package pkg

import (
	"encoding/json"
	"github.com/Baidu-AIP/golang-sdk/aip/censor"
)

type TextCensorResp struct {
	LogId          uint   `json:"log_id"`
	ErrorCode      uint   `json:"error_code"`
	ErrorMsg       string `json:"error_msg"`
	Conclusion     string `json:"conclusion"`
	ConclusionType uint   `json:"conclusionType"`
}

func TextCensor(content string) bool {
	client := censor.NewClient("4cbtVOEaoIWhZWKIFQdmS9gE", "tix1wSQIzBBkfGp7eqciA8RehHX84Rka")
	//如果是百度云ak sk,使用下面的客户端

	res := client.TextCensor(content)
	var resp TextCensorResp
	json.Unmarshal([]byte(res), &resp)
	if resp.ConclusionType == 1 {
		return true
	}
	return false

}
