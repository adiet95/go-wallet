package libs

import "time"

func TimeNow() (string, error) {
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		return "", err
	}
	now := time.Now().In(loc).Format("2006-01-02 15:04:05")
	return now, nil
}
