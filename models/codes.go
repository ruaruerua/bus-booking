package models

import (
	"bus-booking/util"
	"errors"
	"log"
	"strings"
	"time"
)

type Codes struct {
	CodeID     string
	Codecode   string
	CodeStatus int
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
	if !rows.Next() {
		return errors.New("error")
	}
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
func Billing(b *Codes, session *string, u *User) error {
	err := NowUser(u, session)
	util.Report(err)
	c := make(chan bool, 1)
	go checkCode(b, c)
	for i := 0; i < cap(c); i++ {
		if !<-c {
			return errors.New("error")
		}
	}
	stmta, err := util.DB.Prepare(`UPDATE users SET users.balance = ? WHERE users.id = ?`)
	util.Report(err)
	stmtb, err := util.DB.Prepare(`UPDATE codes SET status = 0 ,useAT=?,user_id=? WHERE code = ?`)
	util.Report(err)
	t := time.Now()
	b.UseAt = t.Format("2006-01-02 15:04:05")
	u.Balance = u.Balance + 100
	_, err = stmta.Query(u.Balance, u.UserID)
	util.Report(err)
	_, err = stmtb.Query(b.UseAt, u.UserID, b.Codecode)
	util.Report(err)
	log.Print(u)
	if check(u) {
		str := strings.Replace(*session, "session:", "", -1)
		updateUserCache(u, &str)
	}
	return nil
}
func checkCode(b *Codes, c chan bool) {
	var code Codes
	stmt, err := util.DB.Prepare("SELECT code,status FROM codes WHERE code = ?")
	util.Report(err)
	rows, err := stmt.Query(b.Codecode)
	util.Report(err)
	for rows.Next() {
		err := rows.Scan(&code.Codecode, &code.CodeStatus)
		util.Report(err)
	}
	if code.CodeStatus < 0 && b.Codecode == code.Codecode {
		*b = code
		c <- true
	} else {
		c <- false
	}
}
