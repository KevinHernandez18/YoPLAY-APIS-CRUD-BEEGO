package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["YoPLAY_APIS_CRUD_BEEGO_CLASE/controllers:ReglasController"] = append(beego.GlobalControllerRouter["YoPLAY_APIS_CRUD_BEEGO_CLASE/controllers:ReglasController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YoPLAY_APIS_CRUD_BEEGO_CLASE/controllers:ReglasController"] = append(beego.GlobalControllerRouter["YoPLAY_APIS_CRUD_BEEGO_CLASE/controllers:ReglasController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YoPLAY_APIS_CRUD_BEEGO_CLASE/controllers:ReglasController"] = append(beego.GlobalControllerRouter["YoPLAY_APIS_CRUD_BEEGO_CLASE/controllers:ReglasController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YoPLAY_APIS_CRUD_BEEGO_CLASE/controllers:ReglasController"] = append(beego.GlobalControllerRouter["YoPLAY_APIS_CRUD_BEEGO_CLASE/controllers:ReglasController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YoPLAY_APIS_CRUD_BEEGO_CLASE/controllers:ReglasController"] = append(beego.GlobalControllerRouter["YoPLAY_APIS_CRUD_BEEGO_CLASE/controllers:ReglasController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YoPLAY_APIS_CRUD_BEEGO_CLASE/controllers:TipoDeporteController"] = append(beego.GlobalControllerRouter["YoPLAY_APIS_CRUD_BEEGO_CLASE/controllers:TipoDeporteController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YoPLAY_APIS_CRUD_BEEGO_CLASE/controllers:TipoDeporteController"] = append(beego.GlobalControllerRouter["YoPLAY_APIS_CRUD_BEEGO_CLASE/controllers:TipoDeporteController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YoPLAY_APIS_CRUD_BEEGO_CLASE/controllers:TipoDeporteController"] = append(beego.GlobalControllerRouter["YoPLAY_APIS_CRUD_BEEGO_CLASE/controllers:TipoDeporteController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YoPLAY_APIS_CRUD_BEEGO_CLASE/controllers:TipoDeporteController"] = append(beego.GlobalControllerRouter["YoPLAY_APIS_CRUD_BEEGO_CLASE/controllers:TipoDeporteController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YoPLAY_APIS_CRUD_BEEGO_CLASE/controllers:TipoDeporteController"] = append(beego.GlobalControllerRouter["YoPLAY_APIS_CRUD_BEEGO_CLASE/controllers:TipoDeporteController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YoPLAY_APIS_CRUD_BEEGO_CLASE/controllers:TipoDistribucionController"] = append(beego.GlobalControllerRouter["YoPLAY_APIS_CRUD_BEEGO_CLASE/controllers:TipoDistribucionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YoPLAY_APIS_CRUD_BEEGO_CLASE/controllers:TipoDistribucionController"] = append(beego.GlobalControllerRouter["YoPLAY_APIS_CRUD_BEEGO_CLASE/controllers:TipoDistribucionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YoPLAY_APIS_CRUD_BEEGO_CLASE/controllers:TipoDistribucionController"] = append(beego.GlobalControllerRouter["YoPLAY_APIS_CRUD_BEEGO_CLASE/controllers:TipoDistribucionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YoPLAY_APIS_CRUD_BEEGO_CLASE/controllers:TipoDistribucionController"] = append(beego.GlobalControllerRouter["YoPLAY_APIS_CRUD_BEEGO_CLASE/controllers:TipoDistribucionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YoPLAY_APIS_CRUD_BEEGO_CLASE/controllers:TipoDistribucionController"] = append(beego.GlobalControllerRouter["YoPLAY_APIS_CRUD_BEEGO_CLASE/controllers:TipoDistribucionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YoPLAY_APIS_CRUD_BEEGO_CLASE/controllers:TipoReglamentoController"] = append(beego.GlobalControllerRouter["YoPLAY_APIS_CRUD_BEEGO_CLASE/controllers:TipoReglamentoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YoPLAY_APIS_CRUD_BEEGO_CLASE/controllers:TipoReglamentoController"] = append(beego.GlobalControllerRouter["YoPLAY_APIS_CRUD_BEEGO_CLASE/controllers:TipoReglamentoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YoPLAY_APIS_CRUD_BEEGO_CLASE/controllers:TipoReglamentoController"] = append(beego.GlobalControllerRouter["YoPLAY_APIS_CRUD_BEEGO_CLASE/controllers:TipoReglamentoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YoPLAY_APIS_CRUD_BEEGO_CLASE/controllers:TipoReglamentoController"] = append(beego.GlobalControllerRouter["YoPLAY_APIS_CRUD_BEEGO_CLASE/controllers:TipoReglamentoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YoPLAY_APIS_CRUD_BEEGO_CLASE/controllers:TipoReglamentoController"] = append(beego.GlobalControllerRouter["YoPLAY_APIS_CRUD_BEEGO_CLASE/controllers:TipoReglamentoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
