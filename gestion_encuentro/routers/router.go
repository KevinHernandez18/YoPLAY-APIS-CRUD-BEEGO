// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"gestion_encuentro/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/grupo",
			beego.NSInclude(
				&controllers.GrupoController{},
			),
		),

		beego.NSNamespace("/grupo_equipo",
			beego.NSInclude(
				&controllers.GrupoEquipoController{},
			),
		),

		beego.NSNamespace("/grupo_encuentro",
			beego.NSInclude(
				&controllers.GrupoEncuentroController{},
			),
		),

		beego.NSNamespace("/encuentro",
			beego.NSInclude(
				&controllers.EncuentroController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
