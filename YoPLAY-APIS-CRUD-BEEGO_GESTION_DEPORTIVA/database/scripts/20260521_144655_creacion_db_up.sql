CREATE SCHEMA IF NOT EXISTS gestion_deportiva;



CREATE TABLE gestion_deportiva.premiacion (
    id_premiacion          SERIAL PRIMARY KEY,
    cantidad_premiacion    INTEGER NOT NULL,
    descripcion            TEXT NOT NULL,
    activo                 BOOLEAN NOT NULL DEFAULT TRUE,
    fecha_creacion         TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    fecha_modificacion     TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);



CREATE TABLE gestion_deportiva.equipo (
    id_equipo              SERIAL PRIMARY KEY,
    nombre_equipo          VARCHAR(100) NOT NULL,
    categoria              VARCHAR(100),
    cantidad_jugadores     INTEGER,
    region                 VARCHAR(100) NOT NULL,
    activo                 BOOLEAN DEFAULT TRUE,
    fecha_creacion         TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    fecha_modificacion     TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);



CREATE TABLE gestion_deportiva.torneo (
    id_torneo              SERIAL PRIMARY KEY,
    id_tipo_deporte        INTEGER NOT NULL,
    nombre_torneo          VARCHAR(100) NOT NULL,
    fecha_inicio           DATE NOT NULL,
    fecha_fin              DATE NOT NULL,
    ubicacion              VARCHAR(200) NOT NULL,
    objetivo               TEXT NOT NULL,
    id_imagen              INTEGER,
    fecha_fase             VARCHAR(100),
    cantidad_equipo        INTEGER NOT NULL,
    id_tipo_distribucion   INTEGER NOT NULL,
    id_premiacion          INTEGER,
    activo                 BOOLEAN DEFAULT TRUE,
    fecha_creacion         TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    fecha_modificacion     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_torneo_premiacion
        FOREIGN KEY (id_premiacion)
        REFERENCES gestion_deportiva.premiacion(id_premiacion)
);



CREATE TABLE gestion_deportiva.imagen (
    id_imagen              SERIAL PRIMARY KEY,
    id_torneo              INTEGER,
    url_imagen             TEXT NOT NULL,
    tipo_imagen            VARCHAR(100) NOT NULL,
    activo                 BOOLEAN DEFAULT TRUE,
    fecha_creacion         TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    fecha_modificacion     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_imagen_torneo
        FOREIGN KEY (id_torneo)
        REFERENCES gestion_deportiva.torneo(id_torneo)
);


CREATE TABLE gestion_deportiva.distribucion (
    id_distribucion        SERIAL PRIMARY KEY,
    id_torneo              INTEGER,
    id_tipo_distribucion   INTEGER NOT NULL,
    confirmacion           BOOLEAN NOT NULL DEFAULT FALSE,
    activo                 BOOLEAN NOT NULL DEFAULT TRUE,
    fecha_creacion         TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    fecha_modificacion     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_distribucion_torneo
        FOREIGN KEY (id_torneo)
        REFERENCES gestion_deportiva.torneo(id_torneo)
);



CREATE TABLE gestion_deportiva.reglamento (
    id_reglamento          SERIAL PRIMARY KEY,
    id_torneo              INTEGER,
    id_tipo_reglamento     INTEGER NOT NULL,
    id_reglas              INTEGER NOT NULL,
    id_tipo_deporte        INTEGER NOT NULL,
    idioma                 VARCHAR(50) NOT NULL,
    activo                 BOOLEAN DEFAULT TRUE,
    fecha_creacion         TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    fecha_modificacion     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_reglamento_torneo
        FOREIGN KEY (id_torneo)
        REFERENCES gestion_deportiva.torneo(id_torneo)
);



CREATE TABLE gestion_deportiva.integrantes (
    id_integrante          SERIAL PRIMARY KEY,

    id_equipo              INTEGER,
    nombre                 VARCHAR(100) NOT NULL,
    posicion               VARCHAR(100) NOT NULL,
    correo                 VARCHAR(100) NOT NULL,
    telefono               VARCHAR(20),
    activo                 BOOLEAN DEFAULT TRUE,
    fecha_creacion         TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    fecha_modificacion     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_integrante_equipo
        FOREIGN KEY (id_equipo)
        REFERENCES gestion_deportiva.equipo(id_equipo)
);



CREATE TABLE gestion_deportiva.clasificacion (
    id_clasificacion       SERIAL PRIMARY KEY,
    id_encuentro           INTEGER NOT NULL,
    partido_jugado         INTEGER,
    partido_ganado         INTEGER,
    partido_empatado       INTEGER,
    partido_derrota        INTEGER,
    puntos_partidos        INTEGER,
    id_torneo              INTEGER,
    id_equipo              INTEGER,
    activo                 BOOLEAN NOT NULL DEFAULT TRUE,
    fecha_creacion         TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    fecha_modificacion     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_clasificacion_torneo
        FOREIGN KEY (id_torneo)
        REFERENCES gestion_deportiva.torneo(id_torneo),

    CONSTRAINT fk_clasificacion_equipo
        FOREIGN KEY (id_equipo)
        REFERENCES gestion_deportiva.equipo(id_equipo)
);



CREATE TABLE gestion_deportiva.historialtorneos (
    id_historial_torneo    SERIAL PRIMARY KEY,
    nombre_torneo          VARCHAR(100) NOT NULL,
    id_torneo              INTEGER,
    id_usuario             INTEGER NOT NULL,
    deporte                VARCHAR(100) NOT NULL,
    reglamento             TEXT,
    equipos                INTEGER NOT NULL,
    id_tipo_distribucion   INTEGER NOT NULL,
    distribucion           VARCHAR(100),
    activo                 BOOLEAN NOT NULL DEFAULT TRUE,
    fecha_creacion         TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    fecha_modificacion     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_historial_torneo
        FOREIGN KEY (id_torneo)
        REFERENCES gestion_deportiva.torneo(id_torneo)
);