package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertDocumento_20260520_162529 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertDocumento_20260520_162529{}
	m.Created = "20260520_162529"

	migration.Register("InsertDocumento_20260520_162529", m)
}

// Run the migrations
func (m *InsertDocumento_20260520_162529) Up() {
	m.SQL("INSERT INTO usuarios.documento (tipo_documento, activo) VALUES	('Cédula de Ciudadanía', TRUE),('Pasaporte',            TRUE), ('DNI',                  TRUE), ('Tarjeta de Identidad', TRUE), ('Cédula Extranjería',   TRUE);")

}

// Reverse the migrations
func (m *InsertDocumento_20260520_162529) Down() {
	m.SQL("DELETE FROM usuarios.documento") 

}
