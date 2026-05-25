package main

import (
	"fmt"
	"strings"
	"io/ioutil"

	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type CreacionDbYoplay_20260520_153104 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreacionDbYoplay_20260520_153104{}
	m.Created = "20260520_153104"

	migration.Register("CreacionDbYoplay_20260520_153104", m)
}

// Run the migrations
func (m *CreacionDbYoplay_20260520_153104) Up() {
	file, err := ioutil.ReadFile("../scripts/20260520_153104_creacion_db_yoplay_up.sql")

	if err != nil{
		fmt.Println(err)
	}
	
	requests := strings.Split(string(file),";")

	for _, request := range requests{
		fmt.Println(request)
		m.SQL(request)
	}
}

// Reverse the migrations
func (m *CreacionDbYoplay_20260520_153104) Down() {
	file, err := ioutil.ReadFile("../scripts/20260520_153104_creacion_db_yoplay_down.sql")

	if err != nil{
		fmt.Println(err)
	}
	
	requests := strings.Split(string(file),";")

	for _, request := range requests{
		fmt.Println(request)
		m.SQL(request)
	}
}
