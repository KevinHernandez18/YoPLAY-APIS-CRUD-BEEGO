package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Clasificacion struct {
	Id                int       `orm:"column(id_clasificacion);pk:auto"`
	IdEncuentro       int       `orm:"column(id_encuentro)"`
	PartidoJugado     int       `orm:"column(partido_jugado);null"`
	PartidoGanado     int       `orm:"column(partido_ganado);null"`
	PartidoEmpatado   int       `orm:"column(partido_empatado);null"`
	PartidoDerrota    int       `orm:"column(partido_derrota);null"`
	PuntosPartidos    int       `orm:"column(puntos_partidos);null"`
	IdTorneo          int 		`orm:"column(id_torneo)"`
	IdEquipo          int  		`orm:"column(id_equipo)"`
	Activo            bool      `orm:"column(activo)"`
	FechaCreacion     time.Time `orm:"column(fecha_creacion);type(timestamp without time zone);null;auto_now_add"`
	FechaModificacion time.Time `orm:"column(fecha_modificacion);type(timestamp without time zone);null;auto_now"`
}

func (t *Clasificacion) TableName() string {
	return "clasificacion"
}

func init() {
	orm.RegisterModel(new(Clasificacion))
}

// AddClasificacion insert a new Clasificacion into database and returns
// last inserted Id on success.
func AddClasificacion(m *Clasificacion) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetClasificacionById retrieves Clasificacion by Id. Returns error if
// Id doesn't exist
func GetClasificacionById(id int) (v *Clasificacion, err error) {
	o := orm.NewOrm()
	v = &Clasificacion{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllClasificacion retrieves all Clasificacion matches certain condition. Returns empty list if
// no records exist
func GetAllClasificacion(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Clasificacion))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Clasificacion
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateClasificacion updates Clasificacion by Id and returns error if
// the record to be updated doesn't exist
func UpdateClasificacionById(m *Clasificacion) (err error) {
	o := orm.NewOrm()
	v := Clasificacion{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteClasificacion deletes Clasificacion by Id and returns error if
// the record to be deleted doesn't exist
func DeleteClasificacion(id int) (err error) {
	o := orm.NewOrm()
	v := Clasificacion{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Clasificacion{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
