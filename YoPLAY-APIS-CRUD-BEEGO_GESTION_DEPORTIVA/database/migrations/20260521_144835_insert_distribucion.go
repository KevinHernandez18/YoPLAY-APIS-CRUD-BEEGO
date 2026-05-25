package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertDistribucion_20260521_144835 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertDistribucion_20260521_144835{}
	m.Created = "20260521_144835"

	migration.Register("InsertDistribucion_20260521_144835", m)
}

// Run the migrations
func (m *InsertDistribucion_20260521_144835) Up() {
m.SQL("INSERT INTO gestion_deportiva.distribucion (id_tipo_distribucion,confirmacion) VALUES ('1','true')")

}

// Reverse the migrations
func (m *InsertDistribucion_20260521_144835) Down() {
m.SQL("DELETE FROM gestion_deportiva.distribucion WHERE id_tipo_distribucon='1'")

}
