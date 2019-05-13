package models

import (
	"bus-booking/util"
	"errors"
)

type Codes struct {
	CodeID     string
	Codecode   string
	CodeStatus uint8
	UseAt      string
}

func AllCodes(b *[]Codes, session *string) error {
	var user User
	err := NowUser(&user, session)
	if err != nil || user.IsAdmin != true {
		return errors.New("admin error")
	}
	stmt, err := util.DB.Prepare("SELECT * FROM codes")
	util.Report(err)
	var codes Codes
	rows, err := stmt.Query()
	util.Report(err)
	for rows.Next() {
		err := rows.Scan(&codes.CodeID, &codes.Codecode, &codes.CodeStatus, &user.UserID, &codes.UseAt)
		util.Report(err)
		*b = append(*b, codes)
	}
	return nil
}
func FAcode(b *Codes, session *string) error {
	var user User
	err := NowUser(&user, session)
	if err != nil || user.IsAdmin != true {
		return errors.New("admin error")
	}
	go func(b *Codes) {
		stmt, err := util.DB.Prepare("UPDATE codes SET status = 0 WHERE id = ?")
		util.Report(err)
		_, err = stmt.Query(b.CodeID)
		util.Report(err)
	}(b)
	return nil
}
