package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertTutoriales_20260521_153916 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTutoriales_20260521_153916{}
	m.Created = "20260521_153916"

	migration.Register("InsertTutoriales_20260521_153916", m)
}

// Run the migrations
func (m *InsertTutoriales_20260521_153916) Up() {
	m.SQL("INSERT INTO tutoriales.tutoriales (nombre_tutorial, link_tutorial, fecha_publicacion, descripcion, activo) VALUES ('Crear un torneo',       'https://tuto.io/1', '2024-01-10', 'Guía paso a paso para crear y publicar un torneo en la plataforma.',      TRUE),('Registrar equipos',     'https://tuto.io/2', '2024-02-05', 'Cómo registrar equipos, jugadores y asignar categorías.',                 TRUE), ('Configurar reglamento', 'https://tuto.io/3', '2024-03-12', 'Pasos para adjuntar y personalizar reglamentos por tipo de deporte.',     TRUE), ('Ver historial',         'https://tuto.io/4', '2024-04-18', 'Accede al historial completo de torneos finalizados y sus resultados.',   TRUE), ('Subir imágenes',        'https://tuto.io/5', '2024-05-22', 'Tutorial para cargar imágenes de portada y galería de cada torneo.',      TRUE), ('Gestionar encuentros',  'https://tuto.io/6', '2024-06-01', 'Cómo programar, editar y registrar resultados de encuentros deportivos.', TRUE);") 

}

// Reverse the migrations
func (m *InsertTutoriales_20260521_153916) Down() {
	m.SQL("DELETE FROM tutoriales") 

}
