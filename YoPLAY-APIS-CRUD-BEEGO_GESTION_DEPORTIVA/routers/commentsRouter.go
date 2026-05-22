package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:ClasificacionController"] = append(beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:ClasificacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:ClasificacionController"] = append(beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:ClasificacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:ClasificacionController"] = append(beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:ClasificacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:ClasificacionController"] = append(beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:ClasificacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:ClasificacionController"] = append(beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:ClasificacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:DistribucionController"] = append(beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:DistribucionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:DistribucionController"] = append(beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:DistribucionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:DistribucionController"] = append(beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:DistribucionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:DistribucionController"] = append(beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:DistribucionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:DistribucionController"] = append(beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:DistribucionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:EquipoController"] = append(beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:EquipoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:EquipoController"] = append(beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:EquipoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:EquipoController"] = append(beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:EquipoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:EquipoController"] = append(beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:EquipoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:EquipoController"] = append(beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:EquipoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:HistorialtorneosController"] = append(beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:HistorialtorneosController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:HistorialtorneosController"] = append(beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:HistorialtorneosController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:HistorialtorneosController"] = append(beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:HistorialtorneosController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:HistorialtorneosController"] = append(beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:HistorialtorneosController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:HistorialtorneosController"] = append(beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:HistorialtorneosController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:ImagenController"] = append(beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:ImagenController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:ImagenController"] = append(beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:ImagenController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:ImagenController"] = append(beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:ImagenController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:ImagenController"] = append(beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:ImagenController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:ImagenController"] = append(beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:ImagenController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:IntegrantesController"] = append(beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:IntegrantesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:IntegrantesController"] = append(beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:IntegrantesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:IntegrantesController"] = append(beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:IntegrantesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:IntegrantesController"] = append(beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:IntegrantesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:IntegrantesController"] = append(beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:IntegrantesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:PremiacionController"] = append(beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:PremiacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:PremiacionController"] = append(beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:PremiacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:PremiacionController"] = append(beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:PremiacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:PremiacionController"] = append(beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:PremiacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:PremiacionController"] = append(beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:PremiacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:ReglamentoController"] = append(beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:ReglamentoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:ReglamentoController"] = append(beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:ReglamentoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:ReglamentoController"] = append(beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:ReglamentoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:ReglamentoController"] = append(beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:ReglamentoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:ReglamentoController"] = append(beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:ReglamentoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:TorneoController"] = append(beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:TorneoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:TorneoController"] = append(beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:TorneoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:TorneoController"] = append(beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:TorneoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:TorneoController"] = append(beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:TorneoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:TorneoController"] = append(beego.GlobalControllerRouter["YOPLAY-APIS-CRUD-BEEGO_GESTION_DEPORTIVA/controllers:TorneoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
