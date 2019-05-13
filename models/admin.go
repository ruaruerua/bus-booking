package models

import (
	"bus-booking/util"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"math/rand"
)

func Alogin(u *User) (string, error) {
	var session string
	c := make(chan bool, 2)
	go func(session *string, c chan bool) {
		const charSet = "abcdefghijklmnopqrstuvwxyz0123456789"
		b := make([]byte, 26)
		for i := range b {
			b[i] = charSet[rand.Intn(len(charSet))]
		}
		*session = string(b)
		c <- true
	}(&session, c)
	go checkUserAdmin(u, c)
	for i := 0; i < cap(c); i++ {
		if !<-c {
			return "", errors.New("login: error")
		}
	}
	go updateUserCache(u, &session)
	return session, nil
}
func checkUserAdmin(u *User, c chan bool) {
	if u.Account == "" || u.Password == "" {
		c <- false
	} else {
		stmt, err := util.DB.Prepare("SELECT * FROM users WHERE account = ?")
		util.Report(err)
		var user User
		rows, err := stmt.Query(u.Account)
		util.Report(err)
		for rows.Next() {
			err := rows.Scan(&user.UserID, &user.Account, &user.Password, &user.Salt, &user.Balance,
				&user.IsAdmin)
			util.Report(err)
		}
		password := sha1.New()
		password.Write([]byte(u.Password + user.Salt))
		if user.Password == hex.EncodeToString(password.Sum(nil)) && user.IsAdmin == true {
			*u = user
			c <- true
		} else {
			c <- false
		}
	}
}
