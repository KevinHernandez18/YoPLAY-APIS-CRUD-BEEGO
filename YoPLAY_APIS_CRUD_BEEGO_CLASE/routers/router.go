// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"YoPLAY_APIS_CRUD_BEEGO_CLASE/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/tipo_deporte",
			beego.NSInclude(
				&controllers.TipoDeporteController{},
			),
		),

		beego.NSNamespace("/tipo_reglamento",
			beego.NSInclude(
				&controllers.TipoReglamentoController{},
			),
		),

		beego.NSNamespace("/tipo_distribucion",
			beego.NSInclude(
				&controllers.TipoDistribucionController{},
			),
		),

		beego.NSNamespace("/reglas",
			beego.NSInclude(
				&controllers.ReglasController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
