package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
	"fmt"
	"strings"
	"io/ioutil"
)

// DO NOT MODIFY
type CreacionDb_20260520_170527 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreacionDb_20260520_170527{}
	m.Created = "20260520_170527"

	migration.Register("CreacionDb_20260520_170527", m)
}

// Run the migrations
func (m *CreacionDb_20260520_170527) Up() {
	file, err := ioutil.ReadFile("../scripts/20260520_170527_creacion_db_up.sql")

	if err !=nil {
		fmt.Println(err)
		
	}
	
	requests:= strings.Split(string(file), ";")

	for _, request:=range requests{
		fmt.Println(request)
		m.SQL(request)
	}
}

// Reverse the migrations
func (m *CreacionDb_20260520_170527) Down() {
	file, err :=ioutil.ReadFile("../scripts/20260520_170527_creacion_db_down.sql")

	if err !=nil{
		fmt.Println(err)
	}
 
	// SPLIT sirve para segmentar cuando recorra todo el codigo 
	requests := strings.Split(string(file), ";")

	for _, request := range requests{
		fmt.Println(request)
		m.SQL(request)
	}
}