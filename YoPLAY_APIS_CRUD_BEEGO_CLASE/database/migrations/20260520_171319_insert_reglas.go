package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertReglas_20260520_171319 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertReglas_20260520_171319{}
	m.Created = "20260520_171319"

	migration.Register("InsertReglas_20260520_171319", m)
}

// Run the migrations
func (m *InsertReglas_20260520_171319) Up() {
m.SQL("INSERT INTO clase.reglas (cantidad_reglas, descripcion) VALUES ('1','no se haran remates')")

}

// Reverse the migrations
func (m *InsertReglas_20260520_171319) Down() {
m.SQL("DELETE FROM clase.reglas WHERE id_reglas='1'")

}
