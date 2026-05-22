package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertTipoDistribucion_20260520_171332 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTipoDistribucion_20260520_171332{}
	m.Created = "20260520_171332"

	migration.Register("InsertTipoDistribucion_20260520_171332", m)
}

// Run the migrations
func (m *InsertTipoDistribucion_20260520_171332) Up() {
	m.SQL("INSERT INTO clase.tipo_distribucion (equipos) VALUES ('10')") 

}

// Reverse the migrations
func (m *InsertTipoDistribucion_20260520_171332) Down() {
			m.SQL("DELETE FROM clase.tipo_distribucion WHERE equipos='10'")

}
