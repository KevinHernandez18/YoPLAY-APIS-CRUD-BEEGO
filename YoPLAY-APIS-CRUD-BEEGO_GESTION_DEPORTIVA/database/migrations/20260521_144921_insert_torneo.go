package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertTorneo_20260521_144921 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTorneo_20260521_144921{}
	m.Created = "20260521_144921"

	migration.Register("InsertTorneo_20260521_144921", m)
}

// Run the migrations
func (m *InsertTorneo_20260521_144921) Up() {
m.SQL("INSERT INTO gestion_deportiva.torneo (id_tipo_deporte,nombre_torneo,fecha_inicio,fecha_fin,ubicacion,objetivo,fecha_fase,cantidad_equipo,id_tipo_distribucion) VALUES ('1','cascada','12/06/2026','12/08/2026','yopal','somentar deporte','12/10/2027','8','1')")

}

// Reverse the migrations
func (m *InsertTorneo_20260521_144921) Down() {
m.SQL("DELETE FROM gestion_deportiva.torneo WHERE id_tipo_deporte='1'")

}
