package store

import (
	"github.com/huyjack178/izi-guest-manager-ver0.1/cmd/util/logs"
	"github.com/huyjack178/izi-guest-manager-ver0.1/cmd/util/mysql"
	"github.com/huyjack178/izi-guest-manager-ver0.1/cmd/model"
)

var log = logs.New("guestStore")

type GuestStore struct {
	sqlDB *mysql.Instance
}

func NewGuestStore(db *mysql.Instance) *GuestStore {
	return &GuestStore{
		sqlDB: db,
	}
}

func (this *GuestStore) List() (guests []model.Guest, err error) {
	//err = this.mgoDB.Collection(mongodb.GuestCollection).Find(nil).All(&guests)
	//
	//return guests, err
	return nil, nil
}

func (this *GuestStore) Create(guest *model.Guest) error {
	//return this.mgoDB.Collection(mongodb.GuestCollection).Insert(guest)
	return nil
}

func (this *GuestStore) Get(guestId string) (guest *model.Guest, err error) {
	//guestObjId, err := this.mgoDB.CheckObjectId(guestId)
	//if err != nil {
	//	return nil, err
	//}
	//
	//err = this.mgoDB.Collection(mongodb.GuestCollection).FindId(guestObjId).One(&guest)
	//
	//return guest, err
	return nil, nil
}

func (this *GuestStore) Update(guestId string, guest *model.Guest) error {

	//guestServer, err := this.Get(guestId)
	//if (err != nil) {
	//	return err
	//}
	//
	//guestServer.FullName = guest.FullName
	//guestServer.Address = guest.Address
	//guestServer.Email = guest.Email
	//guestServer.HomeTown = guest.HomeTown
	//guestServer.Phone = guest.Phone
	//
	//err = this.mgoDB.Collection(mongodb.GuestCollection).UpdateId(bson.ObjectIdHex(guestId), guestServer)
	//
	//return err
	return nil
}

func (this *GuestStore) Delete(guestId string) error {
	//_, err := this.Get(guestId)
	//if (err != nil) {
	//	return err
	//}
	//
	//err = this.mgoDB.Collection(mongodb.GuestCollection).RemoveId(bson.ObjectIdHex(guestId))
	//
	//return err
	return nil
}