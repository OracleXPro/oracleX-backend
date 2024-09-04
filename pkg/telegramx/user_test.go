package telegramx

import (
	"testing"
)

func TestGetUser(t *testing.T) {
	testCases := []struct{ token, botToken string }{
		{
			token:    `query_id=AAGcqlFKAAAAAJyqUUp6-Y62&user=%7B%22id%22%3A1246866076%2C%22first_name%22%3A%22Dante%22%2C%22last_name%22%3A%22%22%2C%22username%22%3A%22S_User%22%2C%22language_code%22%3A%22en%22%7D&auth_date=1651689536&hash=de7f6b26aadbd667a36d76d91969ecf6ffec70ffaa40b3e98d20555e2406bfbb`,
			botToken: "5139539316:AAGVhDje2A3mB9yA_7l8-TV8xikC7KcudNk",
		},
		{
			token:    `user=%7B%22id%22%3A5279501263%2C%22first_name%22%3A%22R%22%2C%22last_name%22%3A%22%22%2C%22username%22%3A%22ptrnil%22%2C%22language_code%22%3A%22zh-hans%22%2C%22allows_write_to_pm%22%3Atrue%7D&chat_instance=-7588080706109704348&chat_type=sender&auth_date=1708368597&hash=9c9efe28d81294fff9414ec1d4652438f8dabcfbf5ef71ec531f802c391e40c5`,
			botToken: "6594273839:AAG07ADcbLp3GzRiR-WaHm4At12o_BegQw0",
		},
		{
			token:    `user=%7B%22id%22%3A5279501263%2C%22first_name%22%3A%22R%22%2C%22last_name%22%3A%22%22%2C%22username%22%3A%22ptrnil%22%2C%22language_code%22%3A%22zh-hans%22%2C%22allows_write_to_pm%22%3Atrue%7D&chat_instance=-7588080706109704348&chat_type=sender&auth_date=1708368597&hash=9c9efe28d81294fff9414ec1d4652438f8dabcfbf5ef71ec531f802c391e40c5`,
			botToken: "6594273839:AAG07ADcbLp3GzRiR-WaHm4At12o_BegQw0",
		},
	}

	for _, kase := range testCases {
		u, err := GetUser(kase.token, kase.botToken)
		if err != nil {
			t.Fatal(err)
		} else {
			t.Logf("%+v\n", u)
		}
	}
}
