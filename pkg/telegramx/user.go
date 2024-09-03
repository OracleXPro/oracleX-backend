package telegramx

import (
	"encoding/json"
	"fmt"
	"net/url"
	"sort"
	"strings"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type User struct {
	Id              int64  `json:"id"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Username        string `json:"username"`
	LanguageCode    string `json:"language_code"`
	AllowsWriteToPm bool   `json:"allows_write_to_pm"`
	AuthDate        string `json:"auth_date"`
}

type kv struct {
	k, v string
}

func GetUser(token, botToken string) (*User, error) {
	if token == "" {
		return nil, errors.New("empty token")
	}

	if botToken == "" {
		return nil, errors.New("empty botToken")
	}

	unquote, err := url.QueryUnescape(token)
	if err != nil {
		return nil, errors.Wrap(err, "fail to unescape token")
	}

	params := getParams(unquote)
	hash2 := params["hash"]
	initData := getInitData(params)
	userStr := params["user"]
	authData := params["auth_date"]

	//now := time.Now().Unix()
	//delta := now - int64(authData)
	//if delta > 86400 || delta < 0 {
	//	return nil, errx.InvalidTimestamp
	//}

	logx.Debugf("%+v\n", params)

	u := &User{}
	ok, err := Verify(botToken, hash2, initData)
	if ok && err == nil {
		err := json.Unmarshal([]byte(userStr), u)
		if err != nil {
			logx.Error(err)
			// NOTICE:
		}
		u.AuthDate = authData
		return u, nil
	}
	return nil, errors.Wrap(err, "fail to verify user")
}

func getParams(token string) map[string]string {
	params := strings.Split(token, "&")
	mp := make(map[string]string)
	for _, p := range params {
		k0v1 := strings.Split(p, "=")
		mp[k0v1[0]] = k0v1[1]
	}
	return mp
}

func getInitData(params map[string]string) string {
	var list []*kv
	for k, v := range params {
		if k != "hash" {
			list = append(list, &kv{
				k: k,
				v: v,
			})
		}
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].k < list[j].k
	})

	var kvs []string
	for _, x := range list {
		kvs = append(kvs, fmt.Sprintf("%s=%s", x.k, x.v))
	}
	return strings.Join(kvs, "\n")
}
