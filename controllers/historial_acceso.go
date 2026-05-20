package controllers

import (
	"YoPLAY-APIS-CRUD-BEEGO/models"
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	beego "github.com/beego/beego/v2/server/web"
)

// HistorialAccesoController operations for HistorialAcceso
type HistorialAccesoController struct {
	beego.Controller
}

// URLMapping ...
func (c *HistorialAccesoController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create HistorialAcceso
// @Param	body		body 	models.HistorialAcceso	true		"body for HistorialAcceso content"
// @Success 201 {int} models.HistorialAcceso
// @Failure 403 body is empty
// @router / [post]
func (c *HistorialAccesoController) Post() {
	var v models.HistorialAcceso
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddHistorialAcceso(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = map[string]interface{}{"succes":true, "status":201, "Messaje": "SE CREA CORRECTAMENTE LOS DATOS", "Data": v}
		} else {
			c.Data["json"] = map[string]interface{}{"succes":false, "status":400, "Messaje":"NO SE ENCONTRO NINGUN PARAMETRO PARA CREAR"}
		}
	} else {
		c.Data["json"] = map[string]interface{}{"succes": false, "status": 400, "Messaje": "ERROR JSON VACIO O PARAMETROS INVALIDOS"}
}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get HistorialAcceso by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.HistorialAcceso
// @Failure 403 :id is empty
// @router /:id [get]
func (c *HistorialAccesoController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetHistorialAccesoById(id)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"succes":false, "status":400, "Messsaje": "ERROR EN EL SERVICIO: DATOS INVALIDOS O NO SE ENCUTRAN REGISTROS"}
	} else {
		c.Data["json"] = map[string]interface{}{"succes":true, "status":200, "Messaje": "PETICION EXITOSA", "Data": v}
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get HistorialAcceso
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.HistorialAcceso
// @Failure 403
// @router / [get]
func (c *HistorialAccesoController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllHistorialAcceso(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		if l== nil{
			c.Data["json"] = map[string]interface{}{"succes":false,"status":400, "Messaje":"ERROR EN EL SERVICIO DE GetAll: LA SOLICITUD CONTIENE UN PARAMETRO INCORRECTO O NO EXISTE NINGUN REGISTRO"}
		}else{
		c.Data["json"] = map[string]interface{}{"succes":true, "status":200, "Messaje":"EXITOSO", "Data": l}
	}
}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the HistorialAcceso
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.HistorialAcceso	true		"body for HistorialAcceso content"
// @Success 200 {object} models.HistorialAcceso
// @Failure 403 :id is not int
// @router /:id [put]
func (c *HistorialAccesoController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.HistorialAcceso{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateHistorialAccesoById(&v); err == nil {
			c.Data["json"] = map[string]interface{}{
				"succes":true, "status": 200, "Messaje": "SE REALIZO ACTUALIZACION CON EXITO", "Data" :v}
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = map[string]interface{}{"succes":true, "status": 400, "Messaje":"ERROR EN EL SERVICIO PUT: LA SOLICITUD CONTIENE UN PARAMETRO INCORRECTO O NO EXISTE NINGUN REGISTRO"}
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the HistorialAcceso
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *HistorialAccesoController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteHistorialAcceso(id); err == nil {
		c.Data["json"] = map[string]interface{}{"succes": true, "status":200, "Messaje": "SE ELIMINO EL REGISTRO" , "Data":id}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
