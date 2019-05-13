package controllers

import (
	"bus-booking/models"
	"bus-booking/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AllBuses(c *gin.Context) {
	buses := make([]models.Bus, 0)
	err := models.AllBuses(&buses)
	if err != nil {
		util.BadRequest(c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"bus":  buses,
	})
}

func OneBus(c *gin.Context) {
	bus := models.Bus{BusID: c.Param("busID")}
	err := models.OneBus(&bus)
	if err != nil {
		util.BadRequest(c)
		return
	}
	session, _ := c.Cookie("session")
	if session != "" {
		var user models.User
		err := models.NowUser(&user, &session)
		util.Report(err)
		favorited := models.Favorited(&user.UserID, &bus.BusID)
		c.JSON(http.StatusOK, gin.H{
			"code":      http.StatusOK,
			"bus":       bus,
			"favorited": favorited,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"bus":  bus,
		})
	}
}
func InsertBus(c *gin.Context) {
	session, err := c.Cookie("session")
	if err != nil {
		util.Unauthorized(c)
		return
	}
	user:=models.User{}
	License := c.PostForm("license")
	//TotalSeats:= c.PostForm("totalSeats")
	TotalSeats,err := strconv.Atoi(c.PostForm("totalSeats"))
	Departure  := c.PostForm("departure")
	Destination:= c.PostForm("destination")
	BeginAt    := c.PostForm("startAt")
	EndAt      := c.PostForm("endAt")
	Price ,err := strconv.ParseFloat(c.PostForm("price"),64)
	Info      := c.PostForm("info")
	week ,err := strconv.Atoi(c.PostForm("weekly"))
	var weekly=uint8(week)
	sta ,err := strconv.Atoi(c.PostForm("status"))
	var status=uint8(sta)
	bus := models.Bus{License: License,TotalSeats: TotalSeats,Departure: Departure,Destination: Destination,BeginAt: BeginAt,EndAt: EndAt,Price: Price,Info: Info,Weekly: weekly,Status: status}
	err = models.InsertBus(&bus,&user,&session)
	if err != nil {
		util.BadRequest(c)
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK})
}
func UpdateBus(c *gin.Context) {
	//b *models.Bus
	session, err := c.Cookie("session")
	if err != nil {
		util.Unauthorized(c)
		return
	}
	user:=models.User{}
	bus := models.Bus{BusID: c.Param("busID")}
	err = models.OneBus(&bus)
	if err != nil {
		util.BadRequest(c)
		return
	}
	nLicense := c.PostForm("license")
	if nLicense!=""{
		//bus = models.Bus{License: nLicense}
		bus.License=nLicense
	}
	nTotalSeats:= c.PostForm("totalSeats")
	if nTotalSeats!=""{
		TotalSeats,err:=strconv.Atoi(nTotalSeats)
		if err != nil {
			util.BadRequest(c)
			return
		}
		//bus = models.Bus{TotalSeats: TotalSeats,BusID: busid}
		bus.TotalSeats=TotalSeats
	}
	nDeparture:= c.PostForm("departure")
	if nDeparture!=""{
		//bus = models.Bus{Departure: nDeparture,BusID: busid}
		bus.Departure=nDeparture
	}
	nDestination:= c.PostForm("destination")
	if nDestination!=""{
		//bus = models.Bus{Destination: nDestination,BusID: busid}
		bus.Destination=nDestination
	}
	nBeginAt    := c.PostForm("startAt")
	if nBeginAt!=""{
		//bus = models.Bus{BeginAt: nBeginAt,BusID: busid}
		bus.BeginAt=nBeginAt
	}
	nEndAt      := c.PostForm("endAt")
	if nEndAt!=""{
		//bus = models.Bus{EndAt: nEndAt,BusID: busid}
		bus.EndAt=nEndAt
	}
	nPrice     := c.PostForm("price")
	if nPrice!=""{
		Price,err := strconv.ParseFloat(nPrice,64)
		if err != nil {
			util.BadRequest(c)
			return
		}
		//bus = models.Bus{Price: Price,BusID: busid}
		bus.Price=Price
	}
	nInfo      := c.PostForm("info")
	if nInfo!=""{
		//bus = models.Bus{Info: nInfo,BusID: busid}
		bus.Info=nInfo
	}
	nweek  := c.PostForm("weekly")
	if nweek!=""{
		week,err:=strconv.Atoi(nweek)
		if err != nil {
			util.BadRequest(c)
			return
		}
		weekly:=uint8(week)
		//bus = models.Bus{Weekly: weekly,BusID: busid}
		bus.Weekly=weekly
	}
	nsta  := c.PostForm("status")
	if nsta!=""{
		sta,err:=strconv.Atoi(nsta)
		if err != nil {
			util.BadRequest(c)
			return
		}
		status:=uint8(sta)
		//bus = models.Bus{Status: status,BusID: busid}
		bus.Status=status
	}
	err = models.UpdateBus(&user,&bus,&session)
	if err != nil {
		util.BadRequest(c)
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK})
}