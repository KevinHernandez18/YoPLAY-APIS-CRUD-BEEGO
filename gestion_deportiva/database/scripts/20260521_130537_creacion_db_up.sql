-- 1. SCHEMAS
CREATE SCHEMA IF NOT EXISTS clase;
CREATE SCHEMA IF NOT EXISTS usuarios;
CREATE SCHEMA IF NOT EXISTS gestion_deportiva;
CREATE SCHEMA IF NOT EXISTS gestion_encuentro;
 
-- 2. ENUMs
DO $$ BEGIN
    CREATE TYPE estado_torneo AS ENUM ('Próximo', 'Activo', 'Finalizado', 'Cancelado');
EXCEPTION WHEN duplicate_object THEN NULL;
END $$;

DO $$ BEGIN
    CREATE TYPE resultado_tipo AS ENUM ('Ganador', 'Perdedor');
EXCEPTION WHEN duplicate_object THEN NULL;
END $$;

DO $$ BEGIN
    CREATE TYPE fase_torneo AS ENUM ('Grupos', 'Octavos', 'Cuartos', 'Semifinal', 'Final');
EXCEPTION WHEN duplicate_object THEN NULL;
END $$;

--CREACION DE TABLAS
CREATE TABLE IF NOT EXISTS clase.tipo_deporte (
    id_tipo_deporte    SERIAL PRIMARY KEY,
    nombre_deporte     VARCHAR(30) NOT NULL,
    activo             BOOLEAN     DEFAULT TRUE,
    fecha_creacion     TIMESTAMP   NOT NULL DEFAULT NOW(),
    fecha_modificacion TIMESTAMP   NOT NULL DEFAULT NOW()
);

--TABLA DE REGLAS
CREATE TABLE IF NOT EXISTS clase.reglas (
    id_reglas          SERIAL PRIMARY KEY,
    cantidad_reglas    INT,
    descripcion        TEXT,
    activo             BOOLEAN   NOT NULL DEFAULT TRUE,
    fecha_creacion     TIMESTAMP NOT NULL DEFAULT NOW(),
    fecha_modificacion TIMESTAMP NOT NULL DEFAULT NOW()
);
--TABLA DE DOCUMENTO DE USUARIOS
CREATE TABLE IF NOT EXISTS usuarios.documento (
    id_documento       SERIAL PRIMARY KEY,
    tipo_documento     VARCHAR(30) NOT NULL,
    activo             BOOLEAN   NOT NULL DEFAULT TRUE,
    fecha_creacion     TIMESTAMP NOT NULL DEFAULT NOW(),
    fecha_modificacion TIMESTAMP NOT NULL DEFAULT NOW()
);
--TABLA DE GESTIOS DEPORTIVA
CREATE TABLE IF NOT EXISTS gestion_deportiva.equipo (
    id_equipo          SERIAL PRIMARY KEY,
    nombre_equipo      VARCHAR(50) NOT NULL,
    categoria          VARCHAR(40),
    cantidad_jugadores INT,
    region             VARCHAR(20) NOT NULL,
    activo             BOOLEAN     DEFAULT TRUE,
    fecha_creacion     TIMESTAMP   NOT NULL DEFAULT NOW(),
    fecha_modificacion TIMESTAMP   NOT NULL DEFAULT NOW()
);
--TABLA DE PREMIACION
CREATE TABLE IF NOT EXISTS gestion_deportiva.premiacion (
    id_premiacion       SERIAL PRIMARY KEY,
    cantidad_premiacion INT       NOT NULL,
    descripcion         TEXT      NOT NULL,
    activo             BOOLEAN     NOT NULL DEFAULT TRUE,
    fecha_creacion      TIMESTAMP NOT NULL DEFAULT NOW(),
    fecha_modificacion  TIMESTAMP NOT NULL DEFAULT NOW()
);
--TABLA DE GRUPO
CREATE TABLE IF NOT EXISTS gestion_encuentro.grupo (
    id_grupo           SERIAL PRIMARY KEY,
    grupo              VARCHAR(100) NOT NULL,
    activo             BOOLEAN     NOT NULL DEFAULT TRUE,
    fecha_creacion     TIMESTAMP   NOT NULL DEFAULT NOW(),
    fecha_modificacion TIMESTAMP   NOT NULL DEFAULT NOW()
);

--TABLA DE TUTORIALES DE LA PAGINA
CREATE TABLE IF NOT EXISTS tutoriales (
    id_tutoriales      SERIAL PRIMARY KEY,
    nombre_tutorial    VARCHAR(30) NOT NULL,
    link_tutorial      VARCHAR(30) NOT NULL,
    fecha_publicacion  DATE        NOT NULL,
    descripcion        TEXT        NOT NULL,
    activo             BOOLEAN     NOT NULL DEFAULT TRUE,
    fecha_creacion     TIMESTAMP   NOT NULL DEFAULT NOW(),
    fecha_modificacion TIMESTAMP   NOT NULL DEFAULT NOW()
);
--TABLA DE TIPO DE DISTRIBUCION
CREATE TABLE IF NOT EXISTS clase.tipo_distribucion (
    id_tipo_distribucion SERIAL PRIMARY KEY,
    equipos              VARCHAR(2) NOT NULL,
    activo               BOOLEAN    NOT NULL DEFAULT TRUE,
    fecha_creacion       TIMESTAMP  NOT NULL DEFAULT NOW(),
    fecha_modificacion   TIMESTAMP  NOT NULL DEFAULT NOW()
);
--TABLA DE TIPO DE REGLAMENTO
CREATE TABLE IF NOT EXISTS clase.tipo_reglamento (
    id_tipo_reglamento SERIAL PRIMARY KEY,
    id_tipo_deporte    INT         NOT NULL REFERENCES clase.tipo_deporte(id_tipo_deporte),
    nombre_tipo        VARCHAR(30) NOT NULL,
    descripcion        TEXT        NOT NULL,
    activo             BOOLEAN     DEFAULT TRUE,
    fecha_creacion     TIMESTAMP   NOT NULL DEFAULT NOW(),
    fecha_modificacion TIMESTAMP   NOT NULL DEFAULT NOW()
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

--TABLA DE TORNEO
CREATE TABLE IF NOT EXISTS gestion_deportiva.torneo (
    id_torneo            SERIAL PRIMARY KEY,
    id_tipo_deporte      INT          NOT NULL REFERENCES clase.tipo_deporte(id_tipo_deporte),
    nombre_torneo        VARCHAR(50)  NOT NULL,
    fecha_inicio         DATE         NOT NULL,
    fecha_fin            DATE         NOT NULL,
    ubicacion            VARCHAR(150) NOT NULL,
    objetivo             VARCHAR(200) NOT NULL,
    id_imagen            INT          NULL,       -- Se actualiza después de insertar imágenes
    fecha_fase           VARCHAR(150),
    cantidad_equipo      INT          NOT NULL,
    id_tipo_distribucion INT          NOT NULL REFERENCES clase.tipo_distribucion(id_tipo_distribucion),
    id_premiacion        INT          NOT NULL REFERENCES gestion_deportiva.premiacion(id_premiacion),
    activo               BOOLEAN      DEFAULT TRUE,
    fecha_creacion       TIMESTAMP    NOT NULL DEFAULT NOW(),
    fecha_modificacion   TIMESTAMP    NOT NULL DEFAULT NOW()
);

--TABLA DE IMÁGENES DEL TORNEO
CREATE TABLE IF NOT EXISTS gestion_deportiva.imagen (
    id_imagen          SERIAL PRIMARY KEY,
    id_torneo          INT         NOT NULL REFERENCES gestion_deportiva.torneo(id_torneo),
url_imagen 	VARCHAR (100) NOT NULL,
    tipo_imagen        VARCHAR(10) NOT NULL,
    activo             BOOLEAN     DEFAULT TRUE,
    fecha_creacion     DATE        NOT NULL DEFAULT NOW(),
    fecha_modificacion DATE        NOT NULL DEFAULT NOW()
);

-- FK de torneo hacia imagen (ahora que imagen ya existe)
DO $$ BEGIN
    ALTER TABLE gestion_deportiva.torneo
        ADD CONSTRAINT fk_torneo_imagen
        FOREIGN KEY (id_imagen) REFERENCES gestion_deportiva.imagen(id_imagen);
EXCEPTION WHEN duplicate_object THEN NULL;
END $$;

--TABLA DE DISTRIBUCION PARA TORNEO
CREATE TABLE IF NOT EXISTS gestion_deportiva.distribucion (
    id_distribucion      SERIAL PRIMARY KEY,
    id_torneo            INT       NOT NULL REFERENCES gestion_deportiva.torneo(id_torneo),
    id_tipo_distribucion INT       NOT NULL REFERENCES clase.tipo_distribucion(id_tipo_distribucion),
    confirmacion         BOOLEAN   NOT NULL DEFAULT TRUE,
    activo               BOOLEAN   NOT NULL DEFAULT TRUE,
    fecha_creacion       TIMESTAMP NOT NULL DEFAULT NOW(),
    fecha_modificacion   TIMESTAMP NOT NULL DEFAULT NOW()
);
--TABLA DE REGLAMENTO
CREATE TABLE IF NOT EXISTS gestion_deportiva.reglamento (
    id_reglamento      SERIAL PRIMARY KEY,
    id_torneo          INT         NOT NULL REFERENCES gestion_deportiva.torneo(id_torneo),
    id_tipo_reglamento INT         NOT NULL REFERENCES clase.tipo_reglamento(id_tipo_reglamento),
    id_reglas          INT         NOT NULL REFERENCES clase.reglas(id_reglas),
    id_tipo_deporte    INT         NOT NULL REFERENCES clase.tipo_deporte(id_tipo_deporte),
    idioma             VARCHAR(20) NOT NULL,
    activo             BOOLEAN     DEFAULT TRUE,
    fecha_creacion     TIMESTAMP   NOT NULL DEFAULT NOW(),
    fecha_modificacion TIMESTAMP   NOT NULL DEFAULT NOW()
);
--TABLA DE EQUIPO
CREATE TABLE IF NOT EXISTS gestion_deportiva.integrantes (
    id_integrante      SERIAL PRIMARY KEY,
    id_equipo          INT          NOT NULL REFERENCES gestion_deportiva.equipo(id_equipo),
    nombre             VARCHAR(50)  NOT NULL,
    posicion           VARCHAR(30)  NOT NULL,
    correo             VARCHAR(100) NOT NULL,
    telefono           VARCHAR(15),
    activo             BOOLEAN      DEFAULT TRUE,
    fecha_creacion     TIMESTAMP    NOT NULL DEFAULT NOW(),
    fecha_modificacion TIMESTAMP    NOT NULL DEFAULT NOW()
);

--TABLA DE ENCUENTROS 
CREATE TABLE IF NOT EXISTS gestion_encuentro.encuentro (
    id_encuentro         SERIAL PRIMARY KEY,
    fecha                TIMESTAMP      NOT NULL,
    id_torneo            INT            NOT NULL REFERENCES gestion_deportiva.torneo(id_torneo),
    id_tipo_distribucion INT            NOT NULL REFERENCES clase.tipo_distribucion(id_tipo_distribucion),
    fase_torneo          fase_torneo,
    id_equipo1           INT            NOT NULL REFERENCES gestion_deportiva.equipo(id_equipo),
    resultado_equipo1    resultado_tipo,
    id_equipo2           INT            NOT NULL REFERENCES gestion_deportiva.equipo(id_equipo),
    resultado_equipo2    resultado_tipo,
    activo               BOOLEAN        NOT NULL DEFAULT TRUE,
    fecha_creacion       TIMESTAMP      NOT NULL DEFAULT NOW(),
    fecha_modificacion   TIMESTAMP      NOT NULL DEFAULT NOW()
);
--TABLAS DE GRUPOS_EQUIPOS
CREATE TABLE IF NOT EXISTS gestion_encuentro.grupo_equipo (
    id_grupo_equipo    SERIAL PRIMARY KEY,
    id_grupo           INT       NOT NULL REFERENCES gestion_encuentro.grupo(id_grupo),
    id_equipo          INT       NOT NULL REFERENCES gestion_deportiva.equipo(id_equipo),
    partidos_jugados   INT       DEFAULT 0,
    empates            INT       DEFAULT 0,
    victorias          INT       DEFAULT 0,
    derrotas           INT       DEFAULT 0,
    activo             BOOLEAN   NOT NULL DEFAULT TRUE,
    fecha_creacion     TIMESTAMP NOT NULL DEFAULT NOW(),
    fecha_modificacion TIMESTAMP NOT NULL DEFAULT NOW()
);
--TABLA DE ENCUENTROS GRUPOS
CREATE TABLE IF NOT EXISTS gestion_encuentro.grupo_encuentro (
    id_grupo_encuentro SERIAL PRIMARY KEY,
    id_grupo           INT       NOT NULL REFERENCES gestion_encuentro.grupo(id_grupo),
    id_encuentro       INT       NOT NULL REFERENCES gestion_encuentro.encuentro(id_encuentro),
    activo             BOOLEAN   NOT NULL DEFAULT TRUE,
    fecha_creacion     TIMESTAMP NOT NULL DEFAULT NOW(),
    fecha_modificacion TIMESTAMP NOT NULL DEFAULT NOW()
);
--TABLA DE CLASIFICACION
CREATE TABLE IF NOT EXISTS gestion_deportiva.clasificacion (
    id_clasificacion   SERIAL PRIMARY KEY,
    id_encuentro       INT       NOT NULL REFERENCES gestion_encuentro.encuentro(id_encuentro),
    partido_jugado     INT       DEFAULT 0,
    partido_ganado     INT       DEFAULT 0,
    partido_empatado   INT       DEFAULT 0,
    partido_derrota    INT       DEFAULT 0,
    puntos_partidos    INT       DEFAULT 0,
    id_torneo          INT       NOT NULL REFERENCES gestion_deportiva.torneo(id_torneo),
    id_equipo          INT       NOT NULL REFERENCES gestion_deportiva.equipo(id_equipo),
    activo             BOOLEAN   NOT NULL DEFAULT TRUE,
    fecha_creacion     TIMESTAMP NOT NULL DEFAULT NOW(),
    fecha_modificacion TIMESTAMP NOT NULL DEFAULT NOW()
);

--TABLA PARA HISTORIAL DE TORNEOS
CREATE TABLE IF NOT EXISTS gestion_deportiva.historialtorneos (
    id_historial_torneo  SERIAL PRIMARY KEY,
    nombre_torneo        VARCHAR(50)    NOT NULL,
    id_torneo            INT            NOT NULL REFERENCES gestion_deportiva.torneo(id_torneo),
    id_usuario           INT            NOT NULL REFERENCES usuarios.usuario(id_usuario),
    deporte              VARCHAR(50)    NOT NULL,
    reglamento           VARCHAR(60),
    equipos              INT            NOT NULL,
    id_tipo_distribucion INT            NOT NULL REFERENCES clase.tipo_distribucion(id_tipo_distribucion),
    distribucion         VARCHAR(50),
    estado               estado_torneo  NOT NULL DEFAULT 'Finalizado',
    finalizacion         resultado_tipo NOT NULL,
    activo               BOOLEAN        NOT NULL DEFAULT TRUE,
    fecha_creacion       TIMESTAMP      NOT NULL DEFAULT NOW(),
    fecha_modificacion   TIMESTAMP      NOT NULL DEFAULT NOW()
);
