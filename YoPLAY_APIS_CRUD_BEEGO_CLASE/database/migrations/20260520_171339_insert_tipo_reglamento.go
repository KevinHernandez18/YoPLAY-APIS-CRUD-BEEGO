package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertTipoReglamento_20260520_171339 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTipoReglamento_20260520_171339{}
	m.Created = "20260520_171339"

	migration.Register("InsertTipoReglamento_20260520_171339", m)
}

// Run the migrations
func (m *InsertTipoReglamento_20260520_171339) Up() {
 m.SQL(("insert into clase.tipo_reglamento (id_tipo_deporte,nombre_tipo,descripcion)values ('1','todos contra todos','sera un partido amistoso de todo contra todos ')")) 

}

// Reverse the migrations
func (m *InsertTipoReglamento_20260520_171339) Down() {
	m.SQL("DELETE FROM clase.tipo_reglamento WHERE id_tipo_deporte='1'")

}
