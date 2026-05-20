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
	m.SQL("INSERT INTO usuarios.contrasena (contrasena, hash_contrasena, id_usuario, activo) VALUES ('Pass@2024!',   '$2b$10$hashCarlos001',  1, TRUE), ('Segura#456',   '$2b$10$hashLuisa002',   2, TRUE), ('Torneo$789',   '$2b$10$hashAndres003',  3, TRUE) ('Deporte@321',  '$2b$10$hashValenti004', 4, TRUE), ('Club%2024',    '$2b$10$hashMiguel005',  5, TRUE), ('Sara@Dep2024', '$2b$10$hashSara006',    6, TRUE);") 

}

// Reverse the migrations
func (m *InsertContrasena_20260520_162547) Down() {
	m.SQL("DELETE FROM usuarios.contrasena") 
}
