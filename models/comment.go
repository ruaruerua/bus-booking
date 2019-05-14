package models

import (
	"bus-booking/util"
	"errors"
	"log"
	"time"
)

type Comment struct {
	CommentID      string
	Content        string
	CommentAt      string
	Stars          uint8
	Status         uint8
	IsReplied      uint8
	ContentReplied string
}

func Commentlist(b *[]Comment, busID *string) error {
	var user User
	var bus Bus
	var comment Comment
	stmt, err := util.DB.Prepare("SELECT * FROM comments WHERE bus_id = ?")
	util.Report(err)
	rows, err := stmt.Query(busID)
	util.Report(err)
	if !rows.Next() {
		return errors.New("error")
	}
	for rows.Next() {
		err = rows.Scan(&comment.CommentID, &comment.Content, &comment.CommentAt, &comment.Status, &comment.Stars, &comment.IsReplied, &comment.ContentReplied,
			&user.UserID, &bus.BusID)
		*b = append(*b, comment)
	}
	util.Report(err)
	return nil
}
func InsertComment(b *Comment, c *Bus, session *string) error {
	var user User
	err := NowUser(&user, session)
	if err != nil {
		return errors.New("error")
	}
	go func(b *Comment) {
		t := time.Now()
		b.CommentAt = t.Format("2006-01-02 15:04:05")
		stmt, err := util.DB.Prepare("INSERT INTO comments (content,comment_at,stars,user_id,bus_id) VALUES (?,?,?,?,?)")
		util.Report(err)
		_, err = stmt.Query(b.Content, b.CommentAt, b.Stars, user.UserID, c.BusID)
		util.Report(err)
	}(b)
	return nil
}
func Deletecomment(session *string, CommentID *string) error {
	var bus Bus
	var user User
	var ouser User
	err := NowUser(&user, session)
	util.Report(err)
	var comment Comment
	stmt, err := util.DB.Prepare("SELECT * FROM comments WHERE id = ?")
	util.Report(err)
	rows, err := stmt.Query(CommentID)
	util.Report(err)
	if !rows.Next() {
		return errors.New("error")
	}
	err = rows.Scan(&comment.CommentID, &comment.Content, &comment.CommentAt, &comment.Status, &comment.Stars, &comment.IsReplied, &comment.ContentReplied,
		&ouser.UserID, &bus.BusID)
	util.Report(err)
	log.Print(*CommentID)
	if user.UserID == ouser.UserID || user.IsAdmin == true {
		go func() {
			stmt, err := util.DB.Prepare("DELETE FROM comments WHERE id = ? ")
			util.Report(err)
			_, err = stmt.Query(CommentID)
			util.Report(err)
		}()
		return nil
	}
	return errors.New("没权限,error")
}
func ReplyComment(b *Comment, session *string) error {
	var user User
	err := NowUser(&user, session)
	if err != nil || user.IsAdmin != true {
		return errors.New("admin error")
	}
	go func(b *Comment) {
		stmt, err := util.DB.Prepare("UPDATE comments SET content_replied = ?,is_replied=? WHERE id = ?")
		util.Report(err)
		_, err = stmt.Query(b.ContentReplied, 1, b.CommentID)
		util.Report(err)
	}(b)
	return nil
}
