package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertUsuario_20260520_162539 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertUsuario_20260520_162539{}
	m.Created = "20260520_162539"

	migration.Register("InsertUsuario_20260520_162539", m)
}

// Run the migrations
func (m *InsertUsuario_20260520_162539) Up() {
	m.SQL("INSERT INTO usuarios.usuario (nombre, apellido, numero_documento, email, telefono, fecha_nacimiento, id_documento, activo) VALUES ('Carlos',    'Ramírez',  '1012345678', 'carlos.ramirez@mail.com',  '3101234567', '1990-05-14', 1, TRUE), ('Luisa',     'Mendoza',  '1023456789', 'luisa.mendoza@mail.com',   '3202345678', '1995-08-22', 1, TRUE), ('Andrés',    'Torres',   'AB123456',   'andres.torres@mail.com',   '3153456789', '1988-11-03', 2, TRUE), ('Valentina', 'Gómez',    '1034567890', 'valentina.gomez@mail.com', '3004567890', '2000-02-17', 4, TRUE), ('Miguel',    'Herrera',  '987654321',  'miguel.herrera@mail.com',  '3115678901', '1993-07-30', 3, TRUE), ('Sara',      'Pineda',   '1045678901', 'sara.pineda@mail.com',     '3126789012', '1998-03-11', 1, TRUE);")

}

// Reverse the migrations
func (m *InsertUsuario_20260520_162539) Down() {
	m.SQL("DELETE FROM usuarios.usuario where nombre= carlos") 
}
