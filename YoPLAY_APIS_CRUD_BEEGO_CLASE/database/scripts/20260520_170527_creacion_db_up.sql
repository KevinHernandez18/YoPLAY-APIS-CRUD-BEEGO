--Creacion de Base de Datos
--CREATE DATABASE IF NOT EXISTS tienda_db;

--Creacion de Schema
CREATE SCHEMA IF NOT EXISTS clase;

--Creacion de Tabla Clientes
CREATE TABLE clase.reglas (
    id_reglas           SERIAL            PRIMARY KEY,
    cantidad_reglas     INTEGER      NOT NULL,
    descripcion         TEXT      NOT NULL,
    activo              BOOLEAN           NOT NULL DEFAULT TRUE,
    fecha_creacion      TIMESTAMP         DEFAULT CURRENT_TIMESTAMP,
    fecha_modificacion TIMESTAMP         DEFAULT CURRENT_TIMESTAMP
);

--Creacion de Tabla Productos
CREATE TABLE clase.tipo_deporte(
    id_tipo_deporte      SERIAL            PRIMARY KEY,
    nombre_deporte     VARCHAR(100)      NOT NULL,
    descripcion         TEXT,
    activo              BOOLEAN           NOT NULL DEFAULT TRUE,
    fecha_creacion      TIMESTAMP         DEFAULT CURRENT_TIMESTAMP,
    fecha_modificacion TIMESTAMP         DEFAULT CURRENT_TIMESTAMP
);

--Creacion de Tabla Pedidos
CREATE TABLE clase.tipo_distribucion (
    id_tipo_distribucion          SERIAL            PRIMARY KEY,
    equipos              VARCHAR(2)     NOT NULL,
    activo              BOOLEAN           NOT NULL DEFAULT TRUE,
    fecha_creacion      TIMESTAMP         DEFAULT CURRENT_TIMESTAMP,
    fecha_modificacion TIMESTAMP         DEFAULT CURRENT_TIMESTAMP

);

CREATE TABLE clase.tipo_reglamento(
    id_tipo_reglamento      SERIAL            PRIMARY KEY,
    id_tipo_deporte      INTEGER            NOT NULL,
    nombre_tipo         VARCHAR(30)    NOT NULL,
    descripcion         TEXT,
    activo              BOOLEAN           NOT NULL DEFAULT TRUE,
    fecha_creacion      TIMESTAMP         DEFAULT CURRENT_TIMESTAMP,
    fecha_modificacion TIMESTAMP         DEFAULT CURRENT_TIMESTAMP

);