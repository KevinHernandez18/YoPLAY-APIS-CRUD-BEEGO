package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type HistorialAcceso struct {
	Id                int         `orm:"column(id_historial_acceso);pk;auto"`
	FechaIntento      time.Time   `orm:"column(fecha_intento);type(timestamp without time zone);null"`
	Exitoso           bool        `orm:"column(exitoso)"`
	IpOrigen          string      `orm:"column(ip_origen);null"`
	FalloMotivo       string      `orm:"column(fallo_motivo);null"`
	IdUsuario         *Usuario    `orm:"column(id_usuario);rel(fk)"`
	IdContrasena      *Contrasena `orm:"column(id_contrasena);rel(fk)"`
	Activo            bool        `orm:"column(activo)"`
	FechaCreacion     time.Time   `orm:"column(fecha_creacion);type(timestamp without time zone);null;auto_now_add"`
	FechaModificacion time.Time   `orm:"column(fecha_modificacion);type(timestamp without time zone);null;auto_now"`
}

func (t *HistorialAcceso) TableName() string {
	return "historial_acceso"
}

func init() {
	orm.RegisterModel(new(HistorialAcceso))
}

// AddHistorialAcceso insert a new HistorialAcceso into database and returns
// last inserted Id on success.
func AddHistorialAcceso(m *HistorialAcceso) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetHistorialAccesoById retrieves HistorialAcceso by Id. Returns error if
// Id doesn't exist
func GetHistorialAccesoById(id int) (v *HistorialAcceso, err error) {
	o := orm.NewOrm()
	v = &HistorialAcceso{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllHistorialAcceso retrieves all HistorialAcceso matches certain condition. Returns empty list if
// no records exist
func GetAllHistorialAcceso(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(HistorialAcceso))
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

	var l []HistorialAcceso
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

// UpdateHistorialAcceso updates HistorialAcceso by Id and returns error if
// the record to be updated doesn't exist
func UpdateHistorialAccesoById(m *HistorialAcceso) (err error) {
	o := orm.NewOrm()
	v := HistorialAcceso{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteHistorialAcceso deletes HistorialAcceso by Id and returns error if
// the record to be deleted doesn't exist
func DeleteHistorialAcceso(id int) (err error) {
	o := orm.NewOrm()
	v := HistorialAcceso{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&HistorialAcceso{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
