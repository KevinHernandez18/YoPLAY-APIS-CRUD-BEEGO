package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["gestion_encuentro/controllers:EncuentroController"] = append(beego.GlobalControllerRouter["gestion_encuentro/controllers:EncuentroController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gestion_encuentro/controllers:EncuentroController"] = append(beego.GlobalControllerRouter["gestion_encuentro/controllers:EncuentroController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gestion_encuentro/controllers:EncuentroController"] = append(beego.GlobalControllerRouter["gestion_encuentro/controllers:EncuentroController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gestion_encuentro/controllers:EncuentroController"] = append(beego.GlobalControllerRouter["gestion_encuentro/controllers:EncuentroController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gestion_encuentro/controllers:EncuentroController"] = append(beego.GlobalControllerRouter["gestion_encuentro/controllers:EncuentroController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gestion_encuentro/controllers:GrupoController"] = append(beego.GlobalControllerRouter["gestion_encuentro/controllers:GrupoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gestion_encuentro/controllers:GrupoController"] = append(beego.GlobalControllerRouter["gestion_encuentro/controllers:GrupoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gestion_encuentro/controllers:GrupoController"] = append(beego.GlobalControllerRouter["gestion_encuentro/controllers:GrupoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gestion_encuentro/controllers:GrupoController"] = append(beego.GlobalControllerRouter["gestion_encuentro/controllers:GrupoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gestion_encuentro/controllers:GrupoController"] = append(beego.GlobalControllerRouter["gestion_encuentro/controllers:GrupoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gestion_encuentro/controllers:GrupoEncuentroController"] = append(beego.GlobalControllerRouter["gestion_encuentro/controllers:GrupoEncuentroController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gestion_encuentro/controllers:GrupoEncuentroController"] = append(beego.GlobalControllerRouter["gestion_encuentro/controllers:GrupoEncuentroController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gestion_encuentro/controllers:GrupoEncuentroController"] = append(beego.GlobalControllerRouter["gestion_encuentro/controllers:GrupoEncuentroController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gestion_encuentro/controllers:GrupoEncuentroController"] = append(beego.GlobalControllerRouter["gestion_encuentro/controllers:GrupoEncuentroController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gestion_encuentro/controllers:GrupoEncuentroController"] = append(beego.GlobalControllerRouter["gestion_encuentro/controllers:GrupoEncuentroController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gestion_encuentro/controllers:GrupoEquipoController"] = append(beego.GlobalControllerRouter["gestion_encuentro/controllers:GrupoEquipoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gestion_encuentro/controllers:GrupoEquipoController"] = append(beego.GlobalControllerRouter["gestion_encuentro/controllers:GrupoEquipoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gestion_encuentro/controllers:GrupoEquipoController"] = append(beego.GlobalControllerRouter["gestion_encuentro/controllers:GrupoEquipoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gestion_encuentro/controllers:GrupoEquipoController"] = append(beego.GlobalControllerRouter["gestion_encuentro/controllers:GrupoEquipoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gestion_encuentro/controllers:GrupoEquipoController"] = append(beego.GlobalControllerRouter["gestion_encuentro/controllers:GrupoEquipoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
