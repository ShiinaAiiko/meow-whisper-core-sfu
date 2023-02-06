package modules

func WSAuth() bool {

	log.Warn("开始校验ws")
	return true
}

// var (
// 	usersMap = map[string][]byte{}
// )

// func longTermCredentials(username string, sharedSecret string) (string, error) {
// 	mac := hmac.New(sha1.New, []byte(sharedSecret))
// 	_, err := mac.Write([]byte(username))
// 	if err != nil {
// 		return "", err // Not sure if this will ever happen
// 	}
// 	password := mac.Sum(nil)
// 	return base64.StdEncoding.EncodeToString(password), nil
// }

// func GetTurnAuth(credentials, sharedSecret, realm string) func(username string, realm string, srcAddr net.Addr) ([]byte, bool) {
// 	if len(usersMap) == 0 && sharedSecret == "" {
// 		log.Error("No turn auth provided", "Got err")
// 		return nil
// 	}

// 	// config.SfuConfig.Turn.Auth.Credentials
// 	for _, kv := range regexp.MustCompile(`(\w+)=(\w+)`).FindAllStringSubmatch(credentials, -1) {
// 		// config.SfuConfig.Turn.Realm
// 		usersMap[kv[1]] = turn.GenerateAuthKey(kv[1], realm, kv[2])

// 		log.Info(kv, usersMap[kv[1]], string(usersMap[kv[1]]))
// 	}

// 	u, p, err := turn.GenerateLongTermCredentials(sharedSecret, 60*time.Second)
// 	log.Error("err", u, p, err)
// 	turnAuth := func(username string, realm string, srcAddr net.Addr) ([]byte, bool) {
// 		log.Warn("开始校验turn")
// 		log.Info(1, "TurnAuth", username, realm, srcAddr, usersMap[username])

// 		if sharedSecret != "" {
// 			t, err := strconv.Atoi(username)
// 			if err != nil {
// 				log.Error("Invalid time-windowed username %q", username)
// 				return nil, false
// 			}
// 			if int64(t) < time.Now().Unix() {
// 				log.Error("Expired time-windowed username %q", username)
// 				return nil, false
// 			}
// 			password, err := longTermCredentials(username, sharedSecret)
// 			log.Info("这里开始校验密码", sharedSecret)
// 			log.Info("password", password)
// 			// 校验用户名有没有问题
// 			// u, _, err := turn.GenerateLongTermCredentials(sharedSecret, 60*time.Second)
// 			// log.Error("err", err)
// 			return turn.GenerateAuthKey(username, realm, ""), true
// 			// return nil, false
// 		}
// 		username = "meowWhisper"
// 		log.Info(usersMap[username])
// 		if key, ok := usersMap[username]; ok {
// 			log.Info("ok", ok, "key", key)
// 			return key, true
// 		}
// 		return nil, false
// 	}
// 	return turnAuth
// }
