package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertTipoDeporte_20260520_171326 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTipoDeporte_20260520_171326{}
	m.Created = "20260520_171326"

	migration.Register("InsertTipoDeporte_20260520_171326", m)
}

// Run the migrations
func (m *InsertTipoDeporte_20260520_171326) Up() {
	m.SQL("INSERT INTO clase.tipo_deporte (nombre_deporte) VALUES ('futbol')")
}



// Reverse the migrations
func (m *InsertTipoDeporte_20260520_171326) Down() {
	m.SQL("DELETE FROM clase.tipo_deporte WHERE nombre_deporte='futbol'")

}
