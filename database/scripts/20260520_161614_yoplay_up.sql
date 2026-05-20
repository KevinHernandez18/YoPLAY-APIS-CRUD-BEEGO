CREATE DATABASE IF NOT EXISTS yoplay;

CREATE SCHEMA IF NOT EXISTS usuarios;

--TABLA DE DOCUMENTO DE USUARIOS
CREATE TABLE IF NOT EXISTS usuarios.documento (
    id_documento       SERIAL PRIMARY KEY,
    tipo_documento     VARCHAR(30) NOT NULL,
    activo             BOOLEAN   NOT NULL DEFAULT TRUE,
    fecha_creacion     TIMESTAMP NOT NULL DEFAULT NOW(),
    fecha_modificacion TIMESTAMP NOT NULL DEFAULT NOW()
);

--TABLA PARA USUARIO
CREATE TABLE IF NOT EXISTS usuarios.usuario (
    id_usuario         SERIAL PRIMARY KEY,
    nombre             VARCHAR(30) NOT NULL,
    apellido           VARCHAR(30) NOT NULL,
    numero_documento   VARCHAR(30) NOT NULL,
    email              VARCHAR(50) NOT NULL UNIQUE,
    telefono           VARCHAR(12) NOT NULL,
    fecha_nacimiento   DATE        NOT NULL,
    id_documento       INT         NOT NULL REFERENCES usuarios.documento(id_documento),
    fecha_registro     DATE        DEFAULT NOW(),
    activo             BOOLEAN     NOT NULL DEFAULT TRUE,
    fecha_creacion     TIMESTAMP   NOT NULL DEFAULT NOW(),
    fecha_modificacion TIMESTAMP   NOT NULL DEFAULT NOW()
);


--TABLA DE CONTRASEÑA DEL USUARIO
CREATE TABLE IF NOT EXISTS usuarios.contrasena (
    id_contrasena      SERIAL PRIMARY KEY,
    contrasena         VARCHAR(255) NOT NULL,
    hash_contrasena    VARCHAR(255) NOT NULL,
    id_usuario         INT          NOT NULL REFERENCES usuarios.usuario(id_usuario),
    activo             BOOLEAN      NOT NULL DEFAULT TRUE,
    fecha_creacion     TIMESTAMP    NOT NULL DEFAULT NOW(),
    fecha_modificacion TIMESTAMP    NOT NULL DEFAULT NOW()
);
--TABLA DE HISTORIAL DE ACCESO DEL USUARIO
CREATE TABLE IF NOT EXISTS usuarios.historial_acceso (
    id_historial_acceso SERIAL PRIMARY KEY,
    fecha_intento       TIMESTAMP    DEFAULT NOW(),
    exitoso             BOOLEAN      NOT NULL,
    ip_origen           VARCHAR(20),
    fallo_motivo        VARCHAR(100),
    id_usuario          INT          NOT NULL REFERENCES usuarios.usuario(id_usuario),
    id_contrasena       INT          NOT NULL REFERENCES usuarios.contrasena(id_contrasena),
    activo              BOOLEAN      NOT NULL DEFAULT TRUE,
    fecha_creacion      TIMESTAMP    NOT NULL DEFAULT NOW(),
    fecha_modificacion  TIMESTAMP    NOT NULL DEFAULT NOW()
);