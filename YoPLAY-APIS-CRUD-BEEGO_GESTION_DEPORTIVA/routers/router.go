// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/torneo",
			beego.NSInclude(
				&controllers.TorneoController{},
			),
		),

		beego.NSNamespace("/premiacion",
			beego.NSInclude(
				&controllers.PremiacionController{},
			),
		),

		beego.NSNamespace("/imagen",
			beego.NSInclude(
				&controllers.ImagenController{},
			),
		),

		beego.NSNamespace("/distribucion",
			beego.NSInclude(
				&controllers.DistribucionController{},
			),
		),

		beego.NSNamespace("/reglamento",
			beego.NSInclude(
				&controllers.ReglamentoController{},
			),
		),

		beego.NSNamespace("/equipo",
			beego.NSInclude(
				&controllers.EquipoController{},
			),
		),

		beego.NSNamespace("/integrantes",
			beego.NSInclude(
				&controllers.IntegrantesController{},
			),
		),

		beego.NSNamespace("/clasificacion",
			beego.NSInclude(
				&controllers.ClasificacionController{},
			),
		),

		beego.NSNamespace("/historialtorneos",
			beego.NSInclude(
				&controllers.HistorialtorneosController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
