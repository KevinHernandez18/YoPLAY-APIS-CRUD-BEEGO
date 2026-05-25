package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertImagen_20260521_144854 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertImagen_20260521_144854{}
	m.Created = "20260521_144854"

	migration.Register("InsertImagen_20260521_144854", m)
}

// Run the migrations
func (m *InsertImagen_20260521_144854) Up() {
m.SQL("INSERT INTO gestion_deportiva.imagen (id_torneo, url_imagen,tipo_imagen) VALUES ('1','3sadsadsada','png')")

}

// Reverse the migrations
func (m *InsertImagen_20260521_144854) Down() {
m.SQL("DELETE FROM gestion_deportiva.imagen WHERE id_torneo='1'")

}
