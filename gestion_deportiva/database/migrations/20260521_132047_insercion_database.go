package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsercionDatabase_20260521_132047 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsercionDatabase_20260521_132047{}
	m.Created = "20260521_132047"

	migration.Register("InsercionDatabase_20260521_132047", m)
}

// Run the migrations
func (m *InsercionDatabase_20260521_132047) Up() {
	file, err := ioutil.ReadFile("../scripts/20260521_130537_creacion_db_up.sql")

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
func (m *InsercionDatabase_20260521_132047) Down() {
	file, err := ioutil.ReadFile("../scripts/20260521_130537_creacion_db_down.sql")

	if err != nil{
		fmt.Println(err)
	}
	
	requests := strings.Split(string(file),";")

	for _, request := range requests{
		fmt.Println(request)
		m.SQL(request)
	}
}
