package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["gestion_deportiva/controllers:ClasificacionController"] = append(beego.GlobalControllerRouter["gestion_deportiva/controllers:ClasificacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gestion_deportiva/controllers:ClasificacionController"] = append(beego.GlobalControllerRouter["gestion_deportiva/controllers:ClasificacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gestion_deportiva/controllers:ClasificacionController"] = append(beego.GlobalControllerRouter["gestion_deportiva/controllers:ClasificacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gestion_deportiva/controllers:ClasificacionController"] = append(beego.GlobalControllerRouter["gestion_deportiva/controllers:ClasificacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gestion_deportiva/controllers:ClasificacionController"] = append(beego.GlobalControllerRouter["gestion_deportiva/controllers:ClasificacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gestion_deportiva/controllers:DistribucionController"] = append(beego.GlobalControllerRouter["gestion_deportiva/controllers:DistribucionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gestion_deportiva/controllers:DistribucionController"] = append(beego.GlobalControllerRouter["gestion_deportiva/controllers:DistribucionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gestion_deportiva/controllers:DistribucionController"] = append(beego.GlobalControllerRouter["gestion_deportiva/controllers:DistribucionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gestion_deportiva/controllers:DistribucionController"] = append(beego.GlobalControllerRouter["gestion_deportiva/controllers:DistribucionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gestion_deportiva/controllers:DistribucionController"] = append(beego.GlobalControllerRouter["gestion_deportiva/controllers:DistribucionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gestion_deportiva/controllers:EquipoController"] = append(beego.GlobalControllerRouter["gestion_deportiva/controllers:EquipoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gestion_deportiva/controllers:EquipoController"] = append(beego.GlobalControllerRouter["gestion_deportiva/controllers:EquipoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gestion_deportiva/controllers:EquipoController"] = append(beego.GlobalControllerRouter["gestion_deportiva/controllers:EquipoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gestion_deportiva/controllers:EquipoController"] = append(beego.GlobalControllerRouter["gestion_deportiva/controllers:EquipoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gestion_deportiva/controllers:EquipoController"] = append(beego.GlobalControllerRouter["gestion_deportiva/controllers:EquipoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gestion_deportiva/controllers:ImagenController"] = append(beego.GlobalControllerRouter["gestion_deportiva/controllers:ImagenController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gestion_deportiva/controllers:ImagenController"] = append(beego.GlobalControllerRouter["gestion_deportiva/controllers:ImagenController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gestion_deportiva/controllers:ImagenController"] = append(beego.GlobalControllerRouter["gestion_deportiva/controllers:ImagenController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gestion_deportiva/controllers:ImagenController"] = append(beego.GlobalControllerRouter["gestion_deportiva/controllers:ImagenController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gestion_deportiva/controllers:ImagenController"] = append(beego.GlobalControllerRouter["gestion_deportiva/controllers:ImagenController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gestion_deportiva/controllers:IntegrantesController"] = append(beego.GlobalControllerRouter["gestion_deportiva/controllers:IntegrantesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gestion_deportiva/controllers:IntegrantesController"] = append(beego.GlobalControllerRouter["gestion_deportiva/controllers:IntegrantesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gestion_deportiva/controllers:IntegrantesController"] = append(beego.GlobalControllerRouter["gestion_deportiva/controllers:IntegrantesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gestion_deportiva/controllers:IntegrantesController"] = append(beego.GlobalControllerRouter["gestion_deportiva/controllers:IntegrantesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gestion_deportiva/controllers:IntegrantesController"] = append(beego.GlobalControllerRouter["gestion_deportiva/controllers:IntegrantesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gestion_deportiva/controllers:PremiacionController"] = append(beego.GlobalControllerRouter["gestion_deportiva/controllers:PremiacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gestion_deportiva/controllers:PremiacionController"] = append(beego.GlobalControllerRouter["gestion_deportiva/controllers:PremiacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gestion_deportiva/controllers:PremiacionController"] = append(beego.GlobalControllerRouter["gestion_deportiva/controllers:PremiacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gestion_deportiva/controllers:PremiacionController"] = append(beego.GlobalControllerRouter["gestion_deportiva/controllers:PremiacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gestion_deportiva/controllers:PremiacionController"] = append(beego.GlobalControllerRouter["gestion_deportiva/controllers:PremiacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gestion_deportiva/controllers:ReglamentoController"] = append(beego.GlobalControllerRouter["gestion_deportiva/controllers:ReglamentoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gestion_deportiva/controllers:ReglamentoController"] = append(beego.GlobalControllerRouter["gestion_deportiva/controllers:ReglamentoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gestion_deportiva/controllers:ReglamentoController"] = append(beego.GlobalControllerRouter["gestion_deportiva/controllers:ReglamentoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gestion_deportiva/controllers:ReglamentoController"] = append(beego.GlobalControllerRouter["gestion_deportiva/controllers:ReglamentoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gestion_deportiva/controllers:ReglamentoController"] = append(beego.GlobalControllerRouter["gestion_deportiva/controllers:ReglamentoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gestion_deportiva/controllers:TorneoController"] = append(beego.GlobalControllerRouter["gestion_deportiva/controllers:TorneoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gestion_deportiva/controllers:TorneoController"] = append(beego.GlobalControllerRouter["gestion_deportiva/controllers:TorneoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gestion_deportiva/controllers:TorneoController"] = append(beego.GlobalControllerRouter["gestion_deportiva/controllers:TorneoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gestion_deportiva/controllers:TorneoController"] = append(beego.GlobalControllerRouter["gestion_deportiva/controllers:TorneoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gestion_deportiva/controllers:TorneoController"] = append(beego.GlobalControllerRouter["gestion_deportiva/controllers:TorneoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
