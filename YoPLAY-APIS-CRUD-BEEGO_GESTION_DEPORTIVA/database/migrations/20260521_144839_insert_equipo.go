package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertEquipo_20260521_144839 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertEquipo_20260521_144839{}
	m.Created = "20260521_144839"

	migration.Register("InsertEquipo_20260521_144839", m)
}

// Run the migrations
func (m *InsertEquipo_20260521_144839) Up() {
m.SQL("INSERT INTO gestion_deportiva.equipo (nombre_equipo,categoria,cantidad_jugadores,region) VALUES ('dorados','sub-15','10','casanare')")

}

// Reverse the migrations
func (m *InsertEquipo_20260521_144839) Down() {
m.SQL("DELETE FROM gestion_deportiva.equipo WHERE categoria='sub-15'")

}
