package db

import (
  "labix.org/v2/mgo"
  "labix.org/v@/mgo/bson"
)

// This is might not work.

type DBRecord interface {
  Save(DBRecord) (DBRecord, error)
  Find(DBRecord) []DBRecord
  FindOrCreate(DBRecord) DBRecord
}










func open_connection() (session, err) {
  // This is here so we don't have to change it anywhere
  // else.
  server_address = "localhost"
  
  // Todo: connection pooling?
  return mgo.Dial(server_address)
}
