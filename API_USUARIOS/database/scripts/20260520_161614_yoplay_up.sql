
CREATE SCHEMA IF NOT EXISTS usuarios;

CREATE TABLE IF NOT EXISTS usuarios.documento (
    id_documento       SERIAL PRIMARY KEY,
    tipo_documento     VARCHAR(30) NOT NULL,
    activo             BOOLEAN NOT NULL DEFAULT TRUE,
    fecha_creacion     TIMESTAMP NOT NULL DEFAULT NOW(),
    fecha_modificacion TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS usuarios.usuario (
    id_usuario         SERIAL PRIMARY KEY,
    nombre             VARCHAR(30) NOT NULL,
    apellido           VARCHAR(30) NOT NULL,
    numero_documento   VARCHAR(30) NOT NULL,
    email              VARCHAR(50) NOT NULL UNIQUE,
    telefono           VARCHAR(12) NOT NULL,
    fecha_nacimiento   DATE NOT NULL,

    id_documento       INT NOT NULL,
    CONSTRAINT fk_usuario_documento
        FOREIGN KEY (id_documento)
        REFERENCES usuarios.documento(id_documento),

    fecha_registro     DATE DEFAULT NOW(),
    activo             BOOLEAN NOT NULL DEFAULT TRUE,
    fecha_creacion     TIMESTAMP NOT NULL DEFAULT NOW(),
    fecha_modificacion TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS usuarios.contrasena (
    id_contrasena      SERIAL PRIMARY KEY,
    contrasena         VARCHAR(255) NOT NULL,
    hash_contrasena    VARCHAR(255) NOT NULL,

    id_usuario         INT NOT NULL,
    CONSTRAINT fk_contrasena_usuario
        FOREIGN KEY (id_usuario)
        REFERENCES usuarios.usuario(id_usuario)
        ON DELETE CASCADE,

    activo             BOOLEAN NOT NULL DEFAULT TRUE,
    fecha_creacion     TIMESTAMP NOT NULL DEFAULT NOW(),
    fecha_modificacion TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS usuarios.historial_acceso (
    id_historial_acceso SERIAL PRIMARY KEY,
    fecha_intento       TIMESTAMP DEFAULT NOW(),
    exitoso             BOOLEAN NOT NULL,
    ip_origen           VARCHAR(20),
    fallo_motivo        VARCHAR(100),

    id_usuario          INT NOT NULL,
    CONSTRAINT fk_historial_usuario
        FOREIGN KEY (id_usuario)
        REFERENCES usuarios.usuario(id_usuario)
        ON DELETE CASCADE,

    id_contrasena       INT NOT NULL,
    CONSTRAINT fk_historial_contrasena
        FOREIGN KEY (id_contrasena)
        REFERENCES usuarios.contrasena(id_contrasena)
        ON DELETE CASCADE,

    activo              BOOLEAN NOT NULL DEFAULT TRUE,
    fecha_creacion      TIMESTAMP NOT NULL DEFAULT NOW(),
    fecha_modificacion  TIMESTAMP NOT NULL DEFAULT NOW()
);