package modules

import (
	"encoding/json"
	"net/http"
	"net/url"

	conf "github.com/ShiinaAiiko/meow-whisper-core-sfu/config"
	"github.com/cherrai/nyanyago-utils/nresponse"
	"github.com/go-resty/resty/v2"
)

var (
	request = resty.New()
)

type CustomData struct {
	AppId  string
	Uid    string
	RoomId string
}

func WSAuth(r *http.Request) bool {

	u, _ := url.Parse(r.URL.String())
	values := u.Query()

	log.Info("------------------开始校验------------------")

	token := values.Get("token")

	log.Info(token)
	var customData CustomData
	err := json.Unmarshal([]byte(values.Get("customData")), &customData)
	if err != nil {
		log.Error("err: ", err)
		return false
	}
	// nlogger.Info(token, uid, customData)
	if conf.Config.MeowWhisperCore.AppId != customData.AppId {
		log.Error("Connection failed, appid is incorrect.")
		return false
	}

	res, err := request.R().SetFormData(map[string]string{
		"appId":  conf.Config.MeowWhisperCore.AppId,
		"appKey": conf.Config.MeowWhisperCore.AppKey,
		"uid":    customData.Uid,
		"roomId": customData.RoomId,
		"token":  token,
	}).Post(conf.Config.MeowWhisperCore.Url + "/api/v1/call/token/verify")
	if err != nil {
		log.Error(err)
		return false
	}

	var m nresponse.NResponse
	err = json.Unmarshal([]byte(res.Body()), &m)
	// log.Info(conf.Config.MeowWhisperCore.Url, m)
	if err != nil {
		log.Error("Unmarshal with error: %+v\n", err)
		return false
	}
	if m.Code != 200 {
		log.Error("Connection failed", m)
		log.Error("Parameter:", token, customData)
		return false
	}
	log.Info("Connection succeeded")
	return true
}
