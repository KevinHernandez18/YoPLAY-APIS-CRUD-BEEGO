package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertReglamento_20260521_144915 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertReglamento_20260521_144915{}
	m.Created = "20260521_144915"

	migration.Register("InsertReglamento_20260521_144915", m)
}

// Run the migrations
func (m *InsertReglamento_20260521_144915) Up() {
m.SQL("INSERT INTO gestion_deportiva.reglamento (id_torneo,id_tipo_reglamento,id_reglas,id_tipo_deporte,idioma) VALUES ('1','1','1','1','español')")


}

// Reverse the migrations
func (m *InsertReglamento_20260521_144915) Down() {
m.SQL("DELETE FROM gestion_deportiva.reglamento WHERE id_reglamento='1'")

}
