package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Db_20260520_161614 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Db_20260520_161614{}
	m.Created = "20260520_161614"

	migration.Register("Db_20260520_161614", m)
}

// Run the migrations
func (m *Db_20260520_161614) Up() {
	file, err := ioutil.ReadFile("../scripts/20260520_161614_yoplay_up.sql")
	if err !=nil{
		fmt.Println(err)
	}

	requests:= strings.Split(string(file), ";")

	for _, request:= range requests{
		fmt.Println(request)
		m.SQL(request)
	}
}

// Reverse the migrations
func (m *Db_20260520_161614) Down() {
	file, err := ioutil.ReadFile("../scripts/20260520_161614_yoplay_down.sql")
	if err !=nil{
		fmt.Println(err)
	}

	requests:= strings.Split(string(file), ";")

	for _, request:= range requests{
		fmt.Println(request)
		m.SQL(request)
	}

}
