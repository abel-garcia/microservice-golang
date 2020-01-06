package psql

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// DbName is a type object for database name
type DbName string

// Initializes the Connection instance
var Connection = NewConnection()

// NewConnetion returns a new connection for database
func NewConnection() *dbs {
	return &dbs{
		Databases: map[DbName]dbStruct{},
		mutex:     sync.RWMutex{},
	}

}

// dbs stores and manages the connections of databases
type dbs struct {
	Connections map[interface{}]map[DbName]map[bool]dbConn
	Databases   map[DbName]dbStruct
	mutex       sync.RWMutex
}

// dbStruct stores pool and master database
type dbStruct struct {
	Master Db
	Pool   []Db
}

// Close closes the connections with the reference
func (d *dbs) Close(reference interface{}) {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	if r, ok := d.Connections[reference]; ok {
		for _, d := range r {
			for _, c := range d {
				c.Close()
			}
		}
	}
}

// Get returns the connection associated to the reference
func (d *dbs) Get(reference interface{}, name DbName, master bool) interface{} {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	if _db, ok := d.Databases[name]; ok {
		if d.Connections == nil {
			d.Connections = map[interface{}]map[DbName]map[bool]dbConn{}
		}

		_r, _ := d.Connections[reference]

		if _r == nil {
			d.Connections[reference] = map[DbName]map[bool]dbConn{}
			d.Connections[reference][name] = map[bool]dbConn{}
			_r = d.Connections[reference]
		}

		_d := _r[name]
		_c := _d[master]
		if _c == nil {
			if master {
				if _db.Master != nil || !_db.Master.IsMaster() {
					for _, p := range _db.Pool {
						if p.IsMaster() {
							_db.Master = p
						}
					}

					if _db.Master != nil {
						return nil
					}
				}

				_c = _db.Master.Get()
			} else {
				l := len(_db.Pool)
				rand.Seed(time.Now().Unix())
				p := _db.Pool[rand.Intn(l)]

				_c = p.Get()
			}

			d.Connections[reference][name][master] = _c
		}

		return _c.Get()
	}

	return nil

}

// Set adds a new database
func (d *dbs) Set(db Db) error {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	if db == nil || !db.Loaded() {
		return fmt.Errorf("The database could not be loaded")
	}

	name := db.GetName()

	_dbs := d.Databases[name]

	_dbs.Pool = append(_dbs.Pool, db)

	if db.IsMaster() {
		_dbs.Master = db
	}

	d.Databases[name] = _dbs

	return nil
}

// Db is a interface that support every method for a db object
type Db interface {
	GetName() DbName
	IsMaster() bool
	Status() bool
	Get() dbConn
	Loaded() bool
}

// dbConn is an interface that supports the database connection object methods
type dbConn interface {
	Close() error
	Get() interface{}
}
