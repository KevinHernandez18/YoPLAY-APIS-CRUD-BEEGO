-- se elimina primero la que tiene la relacion 
DROP TABLE IF EXISTS usuarios.documento;
-- sigue secuencial 
DROP TABLE IF EXISTS usuarios.usuario;
DROP TABLE IF EXISTS usuarios.contrasena;
DROP TABLE IF EXISTS usuarios.historial_acceso

-- se elimina el esquema 
DROP SCHEMA IF EXISTS usuarios;