package model

import "os"

func GetQueryString() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	dat, err := os.ReadFile(wd + "/model/helper/migration/migration.sql")
	if err != nil {
		return "", err
	}
	return string(dat), nil
}
