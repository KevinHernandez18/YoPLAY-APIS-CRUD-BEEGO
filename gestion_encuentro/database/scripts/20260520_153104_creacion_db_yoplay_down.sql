-- DROP TABLES in reverse creation order
DROP TABLE IF EXISTS gestion_deportiva.historialtorneos;
DROP TABLE IF EXISTS gestion_deportiva.clasificacion;
DROP TABLE IF EXISTS gestion_encuentro.grupo_encuentro;
DROP TABLE IF EXISTS gestion_encuentro.grupo_equipo;
DROP TABLE IF EXISTS gestion_encuentro.encuentro;
DROP TABLE IF EXISTS gestion_deportiva.integrantes;
DROP TABLE IF EXISTS gestion_deportiva.reglamento;
DROP TABLE IF EXISTS gestion_deportiva.distribucion;
DROP TABLE IF EXISTS gestion_deportiva.imagen;
DROP TABLE IF EXISTS gestion_deportiva.torneo;
DROP TABLE IF EXISTS usuarios.historial_acceso;
DROP TABLE IF EXISTS usuarios.contrasena;
DROP TABLE IF EXISTS usuarios.usuario;
DROP TABLE IF EXISTS clase.tipo_reglamento;
DROP TABLE IF EXISTS clase.tipo_distribucion;
DROP TABLE IF EXISTS gestion_deportiva.premiacion;
DROP TABLE IF EXISTS gestion_deportiva.equipo;
DROP TABLE IF EXISTS usuarios.documento;
DROP TABLE IF EXISTS clase.reglas;
DROP TABLE IF EXISTS clase.tipo_deporte;
DROP TABLE IF EXISTS gestion_encuentro.grupo;
DROP TABLE IF EXISTS tutoriales;

-- DROP TYPES
DROP TYPE IF EXISTS fase_torneo;
DROP TYPE IF EXISTS resultado_tipo;
DROP TYPE IF EXISTS estado_torneo;

-- DROP SCHEMAS
DROP SCHEMA IF EXISTS gestion_encuentro CASCADE;
DROP SCHEMA IF EXISTS gestion_deportiva CASCADE;
DROP SCHEMA IF EXISTS usuarios CASCADE;
DROP SCHEMA IF EXISTS clase CASCADE;
