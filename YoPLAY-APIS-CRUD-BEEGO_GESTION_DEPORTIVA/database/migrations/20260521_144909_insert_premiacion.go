package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertPremiacion_20260521_144909 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertPremiacion_20260521_144909{}
	m.Created = "20260521_144909"

	migration.Register("InsertPremiacion_20260521_144909", m)
}

// Run the migrations
func (m *InsertPremiacion_20260521_144909) Up() {
m.SQL("INSERT INTO gestion_deportiva.premiacion (cantidad_premiacion,descripcion) VALUES ('1','se ganara un millon de pesos')")

}

// Reverse the migrations
func (m *InsertPremiacion_20260521_144909) Down() {
m.SQL("DELETE FROM gestion_deportiva.premiacion WHERE id_premiacion='1'")

}
