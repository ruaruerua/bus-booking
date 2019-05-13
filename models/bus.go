package models

import (
	"bus-booking/util"
	"errors"
)

type Bus struct {
	BusID       string  `json:"busID"`
	License     string  `json:"license"`
	TotalSeats  int     `json:"totalSeats"`
	EmptySeats  int     `json:"emptySeats"`
	Departure   string  `json:"departure"`
	Destination string  `json:"destination"`
	BeginAt     string  `json:"beginAt"`
	EndAt       string  `json:"endAt"`
	Price       float64 `json:"price"`
	Info        string  `json:"info"`
	Weekly      uint8   `json:"weekly"`
	Status      uint8   `json:"status"`
}

func AllBuses(b *[]Bus) error {
	stmt, err := util.DB.Prepare("SELECT * FROM buses")
	util.Report(err)
	var bus Bus
	rows, err := stmt.Query()
	util.Report(err)
	for rows.Next() {
		err := rows.Scan(&bus.BusID, &bus.License, &bus.TotalSeats, &bus.EmptySeats, &bus.Departure,
			&bus.Destination, &bus.BeginAt, &bus.EndAt, &bus.Price, &bus.Info, &bus.Weekly, &bus.Status)
		util.Report(err)
		*b = append(*b, bus)
	}
	return nil
}

func OneBus(b *Bus) error {
	stmt, err := util.DB.Prepare("SELECT * FROM buses WHERE id = ?")
	util.Report(err)
	rows, err := stmt.Query(b.BusID)
	util.Report(err)
	if !rows.Next() {
		return errors.New("bus: error")
	}
	err = rows.Scan(&b.BusID, &b.License, &b.TotalSeats, &b.EmptySeats, &b.Departure, &b.Destination, &b.BeginAt,
		&b.EndAt, &b.Price, &b.Info, &b.Weekly, &b.Status)
	util.Report(err)
	return nil
}

func UpdateBus(b *Bus, session *string) error {
	var user User
	err := NowUser(&user, session)
	if err != nil || user.IsAdmin != true {
		return errors.New("update bus: error")
	}
	go func(b *Bus) {
		stmt, err := util.DB.Prepare("UPDATE buses SET license = ?, total_seats = ? , empty_seats = ?, departure = ?, destination = ?, start_at = ?, end_at = ?, price = ?, info = ?, weekly = ?, status = ? WHERE id = ?")
		util.Report(err)
		_, err = stmt.Query(b.License, b.TotalSeats, b.TotalSeats, b.Departure, b.Destination, b.BeginAt, b.EndAt, b.Price, b.Info, b.Weekly, b.Status, b.BusID)
		util.Report(err)
	}(b)
	return nil
}
func InsertBus(b *Bus, session *string) error {
	var user User
	err := NowUser(&user, session)
	if err != nil || user.IsAdmin != true {
		return errors.New("update bus: error")
	}
	go func(b *Bus) {
		stmt, err := util.DB.Prepare("INSERT INTO buses (license, total_seats,empty_seats,departure,destination,start_at,end_at,price,info,weekly,status) VALUES ( ?, ?,?,?,?,?,?,?,?,?,?)")
		util.Report(err)
		_, err = stmt.Query(b.License, b.TotalSeats, b.TotalSeats, b.Departure, b.Destination, b.BeginAt, b.EndAt, b.Price, b.Info, b.Weekly, b.Status)
		util.Report(err)
	}(b)
	return nil
}
