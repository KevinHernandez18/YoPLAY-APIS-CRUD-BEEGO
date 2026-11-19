package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertContrasena_20260520_162547 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertContrasena_20260520_162547{}
	m.Created = "20260520_162547"

	migration.Register("InsertContrasena_20260520_162547", m)
}

// Run the migrations
func (m *InsertContrasena_20260520_162547) Up() {
	m.SQL("INSERT INTO usuarios.contrasena (hash_contrasena, id_usuario, activo) VALUES ( '$2b$10$hashCarlos001',  1, TRUE), (   '$2b$10$hashLuisa002',   2, TRUE), ('$2b$10$hashAndres003',  3, TRUE),(  '$2b$10$hashValenti004', 4, TRUE), ('$2b$10$hashMiguel005',  5, TRUE), ('$2b$10$hashSara006',    6, TRUE);")

}

// Reverse the migrations
func (m *InsertContrasena_20260520_162547) Down() {
	m.SQL("DELETE FROM usuarios.contrasena") 
}
