
create schema IF NOT EXISTS tutoriales;


CREATE TABLE IF NOT EXISTS tutoriales.tutoriales (
    id_tutoriales      SERIAL PRIMARY KEY,
    nombre_tutorial    VARCHAR(100) NOT NULL,
    link_tutorial      VARCHAR(300) NOT NULL,
    fecha_publicacion  DATE        NOT NULL,
    descripcion        TEXT        NOT NULL,
    activo             BOOLEAN     NOT NULL DEFAULT TRUE,
    fecha_creacion     TIMESTAMP   NOT NULL DEFAULT NOW(),
    fecha_modificacion TIMESTAMP   NOT NULL DEFAULT NOW()
);