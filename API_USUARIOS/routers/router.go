// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"YoPLAY-APIS-CRUD-BEEGO/API_USUARIOS/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/documento",
			beego.NSInclude(
				&controllers.DocumentoController{},
			),
		),

		beego.NSNamespace("/usuario",
			beego.NSInclude(
				&controllers.UsuarioController{},
			),
		),

		beego.NSNamespace("/contrasena",
			beego.NSInclude(
				&controllers.ContrasenaController{},
			),
		),

		beego.NSNamespace("/historial_acceso",
			beego.NSInclude(
				&controllers.HistorialAccesoController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
