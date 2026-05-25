-- 1. SCHEMAS
CREATE SCHEMA IF NOT EXISTS gestion_encuentro;

DO $$ BEGIN
    CREATE TYPE resultado_tipo AS ENUM ('Ganador', 'Perdedor');
EXCEPTION WHEN duplicate_object THEN NULL;
END $$;

DO $$ BEGIN
    CREATE TYPE fase_torneo AS ENUM ('Grupos', 'Octavos', 'Cuartos', 'Semifinal', 'Final');
EXCEPTION WHEN duplicate_object THEN NULL;
END $$;

--TABLA DE GRUPO
CREATE TABLE IF NOT EXISTS gestion_encuentro.grupo (
    id_grupo           SERIAL PRIMARY KEY,
    grupo              VARCHAR(100) NOT NULL,
    activo             BOOLEAN     NOT NULL DEFAULT TRUE,
    fecha_creacion     TIMESTAMP   DEFAULT NOW(),
    fecha_modificacion TIMESTAMP   NOT NULL DEFAULT NOW()
);

--TABLA DE ENCUENTROS 
CREATE TABLE IF NOT EXISTS gestion_encuentro.encuentro (
    id_encuentro         SERIAL PRIMARY KEY,
    fecha                TIMESTAMP      NOT NULL,
    id_torneo            INT            NOT NULL,
    id_tipo_distribucion INT            NOT NULL,
    fase_torneo          fase_torneo,
    id_equipo1           INT            NOT NULL,
    resultado_equipo1    resultado_tipo,
    id_equipo2           INT            NOT NULL,
    resultado_equipo2    resultado_tipo,
    activo               BOOLEAN        NOT NULL DEFAULT TRUE,
    fecha_creacion       TIMESTAMP      DEFAULT NOW(),
    fecha_modificacion   TIMESTAMP      NOT NULL DEFAULT NOW()
);
--TABLAS DE GRUPOS_EQUIPOS
CREATE TABLE IF NOT EXISTS gestion_encuentro.grupo_equipo (
    id_grupo_equipo    SERIAL PRIMARY KEY,
    id_grupo           INT       NOT NULL REFERENCES gestion_encuentro.grupo(id_grupo) ON DELETE CASCADE,
    id_equipo          INT       NOT NULL,
    partidos_jugados   INT       DEFAULT 0,
    empates            INT       DEFAULT 0,
    victorias          INT       DEFAULT 0,
    derrotas           INT       DEFAULT 0,
    activo             BOOLEAN   NOT NULL DEFAULT TRUE,
    fecha_creacion     TIMESTAMP DEFAULT NOW(),
    fecha_modificacion TIMESTAMP NOT NULL DEFAULT NOW()
);
--TABLA DE ENCUENTROS GRUPOS
CREATE TABLE IF NOT EXISTS gestion_encuentro.grupo_encuentro (
    id_grupo_encuentro SERIAL PRIMARY KEY,
    id_grupo           INT       NOT NULL REFERENCES gestion_encuentro.grupo(id_grupo) ON DELETE CASCADE,
    id_encuentro       INT       NOT NULL REFERENCES gestion_encuentro.encuentro(id_encuentro) ON DELETE CASCADE,
    activo             BOOLEAN   NOT NULL DEFAULT TRUE,
    fecha_creacion     TIMESTAMP DEFAULT NOW(),
    fecha_modificacion TIMESTAMP NOT NULL DEFAULT NOW()
);