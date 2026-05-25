package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertClasificacion_20260521_144828 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertClasificacion_20260521_144828{}
	m.Created = "20260521_144828"

	migration.Register("InsertClasificacion_20260521_144828", m)
}

// Run the migrations
func (m *InsertClasificacion_20260521_144828) Up() {
m.SQL("INSERT INTO gestion_deportiva.clasificacion (id_encuentro,partido_jugado,partido_ganado,partido_empatado,partido_derrota,puntos_partidos,id_torneo,id_equipo) VALUES ('1','10','8','1','1','20','1','1')")

}

// Reverse the migrations
func (m *InsertClasificacion_20260521_144828) Down() {
m.SQL("DELETE FROM gestion_deportiva.clasificacion WHERE id_encuentro='1'")

}
