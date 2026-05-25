package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertIntegrantes_20260521_144904 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertIntegrantes_20260521_144904{}
	m.Created = "20260521_144904"

	migration.Register("InsertIntegrantes_20260521_144904", m)
}

// Run the migrations
func (m *InsertIntegrantes_20260521_144904) Up() {
m.SQL("INSERT INTO gestion_deportiva.integrantes (id_equipo, nombre,posicion,correo,telefono) VALUES ('1','cristian','delantero','pera000890@gmail.com',3125183803)")


}

// Reverse the migrations
func (m *InsertIntegrantes_20260521_144904) Down() {
m.SQL("DELETE FROM gestion_deportiva.integrantes WHERE correo='pera000890@gmail.com'")

}
