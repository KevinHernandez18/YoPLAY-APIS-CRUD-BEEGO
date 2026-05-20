package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["YoPLAY-APIS-CRUD-BEEGO/YoPLAY-APIS-CRUD-BEEGO/controllers:ContrasenaController"] = append(beego.GlobalControllerRouter["YoPLAY-APIS-CRUD-BEEGO/YoPLAY-APIS-CRUD-BEEGO/controllers:ContrasenaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YoPLAY-APIS-CRUD-BEEGO/YoPLAY-APIS-CRUD-BEEGO/controllers:ContrasenaController"] = append(beego.GlobalControllerRouter["YoPLAY-APIS-CRUD-BEEGO/YoPLAY-APIS-CRUD-BEEGO/controllers:ContrasenaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YoPLAY-APIS-CRUD-BEEGO/YoPLAY-APIS-CRUD-BEEGO/controllers:ContrasenaController"] = append(beego.GlobalControllerRouter["YoPLAY-APIS-CRUD-BEEGO/YoPLAY-APIS-CRUD-BEEGO/controllers:ContrasenaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YoPLAY-APIS-CRUD-BEEGO/YoPLAY-APIS-CRUD-BEEGO/controllers:ContrasenaController"] = append(beego.GlobalControllerRouter["YoPLAY-APIS-CRUD-BEEGO/YoPLAY-APIS-CRUD-BEEGO/controllers:ContrasenaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YoPLAY-APIS-CRUD-BEEGO/YoPLAY-APIS-CRUD-BEEGO/controllers:ContrasenaController"] = append(beego.GlobalControllerRouter["YoPLAY-APIS-CRUD-BEEGO/YoPLAY-APIS-CRUD-BEEGO/controllers:ContrasenaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YoPLAY-APIS-CRUD-BEEGO/YoPLAY-APIS-CRUD-BEEGO/controllers:DocumentoController"] = append(beego.GlobalControllerRouter["YoPLAY-APIS-CRUD-BEEGO/YoPLAY-APIS-CRUD-BEEGO/controllers:DocumentoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YoPLAY-APIS-CRUD-BEEGO/YoPLAY-APIS-CRUD-BEEGO/controllers:DocumentoController"] = append(beego.GlobalControllerRouter["YoPLAY-APIS-CRUD-BEEGO/YoPLAY-APIS-CRUD-BEEGO/controllers:DocumentoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YoPLAY-APIS-CRUD-BEEGO/YoPLAY-APIS-CRUD-BEEGO/controllers:DocumentoController"] = append(beego.GlobalControllerRouter["YoPLAY-APIS-CRUD-BEEGO/YoPLAY-APIS-CRUD-BEEGO/controllers:DocumentoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YoPLAY-APIS-CRUD-BEEGO/YoPLAY-APIS-CRUD-BEEGO/controllers:DocumentoController"] = append(beego.GlobalControllerRouter["YoPLAY-APIS-CRUD-BEEGO/YoPLAY-APIS-CRUD-BEEGO/controllers:DocumentoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YoPLAY-APIS-CRUD-BEEGO/YoPLAY-APIS-CRUD-BEEGO/controllers:DocumentoController"] = append(beego.GlobalControllerRouter["YoPLAY-APIS-CRUD-BEEGO/YoPLAY-APIS-CRUD-BEEGO/controllers:DocumentoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YoPLAY-APIS-CRUD-BEEGO/YoPLAY-APIS-CRUD-BEEGO/controllers:HistorialAccesoController"] = append(beego.GlobalControllerRouter["YoPLAY-APIS-CRUD-BEEGO/YoPLAY-APIS-CRUD-BEEGO/controllers:HistorialAccesoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YoPLAY-APIS-CRUD-BEEGO/YoPLAY-APIS-CRUD-BEEGO/controllers:HistorialAccesoController"] = append(beego.GlobalControllerRouter["YoPLAY-APIS-CRUD-BEEGO/YoPLAY-APIS-CRUD-BEEGO/controllers:HistorialAccesoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YoPLAY-APIS-CRUD-BEEGO/YoPLAY-APIS-CRUD-BEEGO/controllers:HistorialAccesoController"] = append(beego.GlobalControllerRouter["YoPLAY-APIS-CRUD-BEEGO/YoPLAY-APIS-CRUD-BEEGO/controllers:HistorialAccesoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YoPLAY-APIS-CRUD-BEEGO/YoPLAY-APIS-CRUD-BEEGO/controllers:HistorialAccesoController"] = append(beego.GlobalControllerRouter["YoPLAY-APIS-CRUD-BEEGO/YoPLAY-APIS-CRUD-BEEGO/controllers:HistorialAccesoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YoPLAY-APIS-CRUD-BEEGO/YoPLAY-APIS-CRUD-BEEGO/controllers:HistorialAccesoController"] = append(beego.GlobalControllerRouter["YoPLAY-APIS-CRUD-BEEGO/YoPLAY-APIS-CRUD-BEEGO/controllers:HistorialAccesoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YoPLAY-APIS-CRUD-BEEGO/YoPLAY-APIS-CRUD-BEEGO/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["YoPLAY-APIS-CRUD-BEEGO/YoPLAY-APIS-CRUD-BEEGO/controllers:UsuarioController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YoPLAY-APIS-CRUD-BEEGO/YoPLAY-APIS-CRUD-BEEGO/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["YoPLAY-APIS-CRUD-BEEGO/YoPLAY-APIS-CRUD-BEEGO/controllers:UsuarioController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YoPLAY-APIS-CRUD-BEEGO/YoPLAY-APIS-CRUD-BEEGO/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["YoPLAY-APIS-CRUD-BEEGO/YoPLAY-APIS-CRUD-BEEGO/controllers:UsuarioController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YoPLAY-APIS-CRUD-BEEGO/YoPLAY-APIS-CRUD-BEEGO/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["YoPLAY-APIS-CRUD-BEEGO/YoPLAY-APIS-CRUD-BEEGO/controllers:UsuarioController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["YoPLAY-APIS-CRUD-BEEGO/YoPLAY-APIS-CRUD-BEEGO/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["YoPLAY-APIS-CRUD-BEEGO/YoPLAY-APIS-CRUD-BEEGO/controllers:UsuarioController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
