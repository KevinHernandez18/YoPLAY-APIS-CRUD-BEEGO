package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsercionDatosYoplay_20260520_153235 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsercionDatosYoplay_20260520_153235{}
	m.Created = "20260520_153235"

	migration.Register("InsercionDatosYoplay_20260520_153235", m)
}

// Run the migrations
func (m *InsercionDatosYoplay_20260520_153235) Up() {
	file, err := ioutil.ReadFile("../scripts/20260520_153235_insercion_datos_yoplay_up.sql")

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
func (m *InsercionDatosYoplay_20260520_153235) Down() {
	file, err := ioutil.ReadFile("../scripts/20260520_153235_insercion_datos_yoplay_down.sql")

	if err != nil{
		fmt.Println(err)
	}
	
	requests := strings.Split(string(file),";")

	for _, request := range requests{
		fmt.Println(request)
		m.SQL(request)
	}
}
