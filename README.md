# YoPLAY-APIS-CRUD-BEEGO

Este repositorio contiene el servicio `gestion_encuentro`, una API REST construida con Go y Beego para gestionar grupos, equipos, encuentros y la relación entre grupos y encuentros.

## Estructura principal

- `gestion_encuentro/`
  - `main.go` - Punto de entrada de la aplicación Beego.
  - `conf/app.conf` - Configuración de Beego y variables de entorno para PostgreSQL.
  - `controllers/` - Controladores REST para los recursos.
  - `models/` - Modelos y lógica de acceso a datos.
  - `routers/` - Rutas y namespace de la API.
  - `database/migrations/` - Migraciones de la base de datos.
  - `database/scripts/` - Scripts SQL de creación y reversión de esquema y datos.
  - `swagger/` - Documentación Swagger UI y archivos de esquema.

## Requisitos

- Go 1.20+ (o compatible con Beego v2)
- PostgreSQL
- Variables de entorno para la base de datos
- `github.com/beego/beego/v2` y dependencias gestionadas por `go.mod`

## Configuración

La aplicación carga variables de entorno mediante `godotenv` y luego usa `conf/app.conf` con referencias a estas variables.

Variables de entorno necesarias:

- `API_CRUD_GESTION_ENCUENTRO_HTTP_PORT` - Puerto HTTP donde escuchará la API.
- `API_CRUD_GESTION_ENCUENTRO_RUN_MODE` - Modo de ejecución (por ejemplo, `dev`).
- `API_CRUD_GESTION_ENCUENTRO_PGUSER` - Usuario de PostgreSQL.
- `API_CRUD_GESTION_ENCUENTRO_PGPASS` - Contraseña de PostgreSQL.
- `API_CRUD_GESTION_ENCUENTRO_PGHOST` - Host de PostgreSQL.
- `API_CRUD_GESTION_ENCUENTRO_PGPORT` - Puerto de PostgreSQL.
- `API_CRUD_GESTION_ENCUENTRO_PGDB` - Nombre de la base de datos.
- `API_CRUD_GESTION_ENCUENTRO_PGSCHEMA` - Esquema PostgreSQL.

## Ejecución

Desde la carpeta `gestion_encuentro`:

```bash
go run .
```

La aplicación inicializa la conexión a PostgreSQL y arranca el servidor Beego. En modo `dev`, la carpeta `swagger/` se expone automáticamente en `/swagger`.

## Endpoints disponibles

Todos los endpoints están bajo el prefijo base `/v1`.

### Recursos

- `/v1/grupo`
- `/v1/grupo_equipo`
- `/v1/grupo_encuentro`
- `/v1/encuentro`

### Operaciones CRUD estándar

Para cada recurso:

- `POST /v1/{recurso}/` - Crear un nuevo registro.
- `GET /v1/{recurso}/` - Obtener todos los registros.
- `GET /v1/{recurso}/:id` - Obtener un registro por ID.
- `PUT /v1/{recurso}/:id` - Actualizar un registro por ID.
- `DELETE /v1/{recurso}/:id` - Eliminar un registro por ID.

### Query params para `GET /v1/{recurso}/`

- `fields` - Campos a devolver (`col1,col2`).
- `query` - Filtros (`col1:v1,col2:v2`).
- `sortby` - Campos de ordenamiento (`col1,col2`).
- `order` - Orden (`asc,desc`).
- `limit` - Límite de resultados.
- `offset` - Offset de resultados.

## Base de datos

La conexión PostgreSQL se configura en `main.go` usando las variables de entorno y `orm.RegisterDataBase`.

Las migraciones y scripts SQL están en:

- `gestion_encuentro/database/migrations/`
- `gestion_encuentro/database/scripts/`

## Documentación Swagger

Si la aplicación se ejecuta en modo `dev`, Swagger UI está disponible en:

- `http://localhost:<PUERTO>/swagger`

## Notas

- El proyecto está diseñado como un CRUD para el dominio de encuentros y grupos.
- Los controladores devuelven respuestas JSON con campos `success`, `status`, `Message` o `Messaje`, y `Data`.
- `main.go` imprime la cadena de conexión de PostgreSQL en el inicio para facilitar la depuración.

