package main

import (
	"fmt"
	"sync"
	"time"
)

type DataBase struct {
}

var db *DataBase
var lock sync.Mutex

func (DataBase) CrateSingleConnection() {
	fmt.Println("Creating singleton Database")
	time.Sleep(2 * time.Second)
	fmt.Println("Creation done")
}

func getDatabaseInstance() *DataBase {
	lock.Lock()
	defer lock.Unlock()
	if db == nil {
		fmt.Println("creando una conexion de base de datos")
		db = &DataBase{}
		db.CrateSingleConnection()
	} else {
		fmt.Println("DB already connected")
	}

	return db
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			getDatabaseInstance()
		}()
	}
	wg.Wait()
}
