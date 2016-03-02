package server

import (
	"github.com/julienschmidt/httprouter"
	"net/http"

	"github.com/huyjack178/izi-guest-manager-ver0.1/cmd/util/logs"
	"github.com/huyjack178/izi-guest-manager-ver0.1/cmd/util/mysql"
	"github.com/huyjack178/izi-guest-manager-ver0.1/cmd/util/xhttp"
	"github.com/huyjack178/izi-guest-manager-ver0.1/cmd/store"
	"github.com/huyjack178/izi-guest-manager-ver0.1/cmd/handler"
	"github.com/gorilla/context"
)

var log = logs.New("server")

type setupStruct struct {
	Config

	//MgoDB *mongodb.Instance
	MySqlDB *mysql.Instance
	Handler http.Handler
}

func setup(cfg Config) *setupStruct {
	s := &setupStruct{Config: cfg}
	//s.setupMongo()
	s.setUpMySql()
	s.setupRoutes()
	return s
}
//
//func (s *setupStruct) setupMongo() {
//	cfg := s.Config
//
//	mgoIns, err := mongodb.NewInstance(mongodb.ConnectOpt{
//		Address: cfg.Mongo.Addr,
//		Database: cfg.Mongo.DBName,
//		Collections: cfg.Mongo.Collections,
//	})
//
//	if err != nil {
//		log.Println("Cannot connect DB: error ", err)
//	}
//
//	s.MgoDB = mgoIns
//}

func (s *setupStruct) setUpMySql() {
	cfg := s.Config

	mySqlIns, err := mysql.NewInstance(mysql.ConnectOpt{
		UserName: cfg.MySQL.UserName,
		Password:cfg.MySQL.Password,
		Host:cfg.MySQL.Host,
		Database:cfg.MySQL.Database,
	})

	if err != nil {
		log.Println("Cannot connect mySQL DB: error ", err)
	}

	s.MySqlDB = mySqlIns
}

func (s *setupStruct) setupRoutes() {

	normal := func(h http.HandlerFunc) httprouter.Handle {
		return xhttp.Adapt(h)
	}

	router := httprouter.New()

	guestStore := store.NewGuestStore(s.MySqlDB)

	{
		guestCtrl := handler.NewGuestCtrl(guestStore)
		router.GET("/guests", normal(guestCtrl.List))
		router.GET("/guest/:id", normal(guestCtrl.Get))
		router.POST("/guests", normal(guestCtrl.Create))
		router.PUT("/guest/:id", normal(guestCtrl.Update))
		router.DELETE("/guest/:id", normal(guestCtrl.Delete))

	}

	s.Handler = context.ClearHandler(router)
}

