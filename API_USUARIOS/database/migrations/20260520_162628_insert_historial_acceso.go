package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertHistorialAcceso_20260520_162628 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertHistorialAcceso_20260520_162628{}
	m.Created = "20260520_162628"

	migration.Register("InsertHistorialAcceso_20260520_162628", m)
}

// Run the migrations
func (m *InsertHistorialAcceso_20260520_162628) Up() {
	m.SQL("INSERT INTO usuarios.historial_acceso (exitoso, ip_origen, fallo_motivo, id_usuario, id_contrasena, activo) VALUES (TRUE,  '192.168.1.10', NULL,                    1, 1, TRUE), (FALSE, '192.168.1.21', 'Contraseña incorrecta', 2, 2, TRUE), (TRUE,  '10.0.0.5',     NULL,                    3, 3, TRUE), (FALSE, '172.16.0.33',  'Usuario no encontrado', 4, 4, TRUE), (TRUE,  '192.168.2.44', NULL,                    5, 5, TRUE), (TRUE,  '10.0.1.12',    NULL,                    6, 6, TRUE);") 

}

// Reverse the migrations
func (m *InsertHistorialAcceso_20260520_162628) Down() {
	m.SQL("DELETE FROM usuarios.historial_acceso WHERE fallo_motivo = 'contraseña incorrecta'") 

}
