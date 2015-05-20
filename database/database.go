package database

import (
    "fmt"
    "log"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

/*
type Database struct {
    session *Session
}
*/

func GetSession(url string) (s, error) {
    session, err := mgo.Dial(url)
    if err != nil {
        return nil, err
    }
    defer session.Close()

    session.SetMode(mgo.Monotonic, true)
    return session, nil
}

func GetDB(s *Session, dbName string) (*Database) {
    db := s.session.DB(dbName)
    return db
}

func GetCollection(db *Database, colName string) (*Collection) {
    col := db(colName)
    return col
}

func Insert(col *Collection, elem *) (error) {
    err := col.Insert(elem)
    if err != nil {
        return err
    }
    return nil
}
