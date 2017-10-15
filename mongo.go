package dataface

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MongoDB struct {
	Address        string `json:"address"`
	DatabaseName   string
	CollectionName string
	client         *mgo.Session
}

func NewMongoDB(host string, username string, password string) (*MongoDB, error) {
	var m MongoDB

	var err error
	m.client, err = mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{host},
		Username: username,
		Password: password,
	})
	err = m.client.Ping()

	m.DatabaseName = "defaultDB"
	m.CollectionName = "defaultCollection"
	return &m, err
}

func (m MongoDB) Put(key string, value []byte) error {
	c := m.client.DB(m.DatabaseName).C(m.CollectionName)
	err := c.Insert(bson.M{key: value})
	return err
}

func (m MongoDB) Get(key string) ([]byte, error) {
	c := m.client.DB(m.DatabaseName).C(m.CollectionName)
	var result []byte
	err := c.Find(bson.M{}).One(&result)
	return result, err
}

func (m MongoDB) Close() error {
	m.client.Close()
	return nil
}
