package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertHistorialtorneos_20260521_144848 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertHistorialtorneos_20260521_144848{}
	m.Created = "20260521_144848"

	migration.Register("InsertHistorialtorneos_20260521_144848", m)
}

// Run the migrations
func (m *InsertHistorialtorneos_20260521_144848) Up() {
m.SQL("INSERT INTO gestion_deportiva.historialtorneos (nombre_torneo,id_torneo,id_usuario,deporte,reglamento,equipos,id_tipo_distribucion,distribucion) VALUES ('fifa','1','1','fultbol','fifa','8','1','ramas')")

}

// Reverse the migrations
func (m *InsertHistorialtorneos_20260521_144848) Down() {
m.SQL("DELETE FROM gestion_deportiva.historialtorneos WHERE nombre_torneo='fifa'")

}
