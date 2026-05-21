--SCRIPT PARA LA INSERCION DE DATOS 

-- 1. clase.tipo_deporte
INSERT INTO clase.tipo_deporte (nombre_deporte, activo) VALUES
    ('Fútbol',      TRUE),
    ('Voleibol',    TRUE),
    ('Baloncesto',  TRUE),
    ('Microfútbol', TRUE),
    ('Patinaje',    TRUE),
    ('Ping-Pong',   TRUE);

 
-- 2. clase.reglas
INSERT INTO clase.reglas (cantidad_reglas, descripcion, activo) VALUES
    (17, 'Reglamento oficial FIFA para torneos de fútbol 11.',              TRUE),
    (12, 'Reglamento FIVB para torneos de voleibol indoor.',                TRUE),
    (12, 'Reglamento FIBA para competencias de baloncesto.',                TRUE),
    (10, 'Normativa AMF para torneos de microfútbol.',                      TRUE),
    (8,  'Reglamento World Skate para competencias de patinaje artístico.', TRUE),
    (6,  'Reglamento ITTF para torneos de ping-pong.',                      TRUE);
 
-- 3. usuarios.documento
INSERT INTO usuarios.documento (tipo_documento, activo) VALUES
    ('Cédula de Ciudadanía', TRUE),
    ('Pasaporte',            TRUE),
    ('DNI',                  TRUE),
    ('Tarjeta de Identidad', TRUE),
    ('Cédula Extranjería',   TRUE);
 

-- 4. clase.tipo_distribucion
INSERT INTO clase.tipo_distribucion (equipos, activo) VALUES
    ('16', TRUE),
    ('8',  TRUE),
    ('8',  TRUE),
    ('12', TRUE),
    ('6',  TRUE),
    ('8',  TRUE);
 

-- 5. gestion_deportiva.equipo
INSERT INTO gestion_deportiva.equipo (nombre_equipo, categoria, cantidad_jugadores, region, activo) VALUES
    ('Leones FC',          'Sub-20',   18, 'Centro',    TRUE),
    ('Rayos Voleibol',     'Femenino', 14, 'Sur',       TRUE),
    ('Tigres Baloncesto',  'Abierta',  12, 'Norte',     TRUE),
    ('Cóndores Micro',     'Sub-17',   10, 'Oriente',   TRUE),
    ('Estrellas Patinaje', 'Juvenil',   8, 'Occidente', TRUE),
    ('Dragones Ping-Pong', 'Abierta',   6, 'Centro',    TRUE);
 

-- 6. gestion_deportiva.premiacion
-- Copa Regional de Fútbol 2024       → id 1, 2, 3  (3 premios)
-- Liga Voleibol Femenino 2024        → id 4, 5     (2 premios)
-- Torneo Baloncesto Norte 2024       → id 6        (1 premio)
-- Torneo Microfútbol Interbarrios    → id 7        (1 premio)
-- Copa Patinaje Artístico 2024       → id 8        (1 premio)
-- Torneo Ping-Pong Regional 2024     → id 9        (1 premio)

INSERT INTO gestion_deportiva.premiacion (cantidad_premiacion, descripcion) VALUES
    (1, '1er Puesto - Copa Dorada | Copa Regional de Fútbol 2024');

INSERT INTO gestion_deportiva.premiacion (cantidad_premiacion, descripcion) VALUES
    (1, '2do Puesto - Medalla de Plata | Copa Regional de Fútbol 2024');

INSERT INTO gestion_deportiva.premiacion (cantidad_premiacion, descripcion) VALUES
    (1, '3er Puesto - Reconocimiento de Bronce | Copa Regional de Fútbol 2024');

INSERT INTO gestion_deportiva.premiacion (cantidad_premiacion, descripcion) VALUES
    (1, '1er Puesto - Trofeo Campeón | Liga Voleibol Femenino 2024');

INSERT INTO gestion_deportiva.premiacion (cantidad_premiacion, descripcion) VALUES
    (1, '2do Puesto - Medalla Subcampeón | Liga Voleibol Femenino 2024');

INSERT INTO gestion_deportiva.premiacion (cantidad_premiacion, descripcion) VALUES
    (1, '1er Puesto - Copa Campeón | Torneo Baloncesto Norte 2024');

INSERT INTO gestion_deportiva.premiacion (cantidad_premiacion, descripcion) VALUES
    (1, '1er Puesto - Trofeo de Oro | Torneo Microfútbol Interbarrios 2024');

INSERT INTO gestion_deportiva.premiacion (cantidad_premiacion, descripcion) VALUES
    (1, '1er Puesto - Medalla Dorada | Copa Patinaje Artístico 2024');

INSERT INTO gestion_deportiva.premiacion (cantidad_premiacion, descripcion) VALUES
    (1, '1er Puesto - Trofeo de Cristal | Torneo Ping-Pong Regional 2024');

 
-- 7. gestion_encuentro.grupo
INSERT INTO gestion_encuentro.grupo (grupo, activo) VALUES
    ('Grupo A', TRUE),
    ('Grupo B', TRUE),
    ('Grupo C', TRUE),
    ('Grupo D', TRUE),
    ('Grupo E', TRUE),
    ('Grupo F', TRUE);
 

-- 8. configuracion.tutoriales
INSERT INTO tutoriales (nombre_tutorial, link_tutorial, fecha_publicacion, descripcion, activo) VALUES
    ('Crear un torneo',       'https://tuto.io/1', '2024-01-10', 'Guía paso a paso para crear y publicar un torneo en la plataforma.',      TRUE),
    ('Registrar equipos',     'https://tuto.io/2', '2024-02-05', 'Cómo registrar equipos, jugadores y asignar categorías.',                 TRUE),
    ('Configurar reglamento', 'https://tuto.io/3', '2024-03-12', 'Pasos para adjuntar y personalizar reglamentos por tipo de deporte.',     TRUE),
    ('Ver historial',         'https://tuto.io/4', '2024-04-18', 'Accede al historial completo de torneos finalizados y sus resultados.',   TRUE),
    ('Subir imágenes',        'https://tuto.io/5', '2024-05-22', 'Tutorial para cargar imágenes de portada y galería de cada torneo.',      TRUE),
    ('Gestionar encuentros',  'https://tuto.io/6', '2024-06-01', 'Cómo programar, editar y registrar resultados de encuentros deportivos.', TRUE);
 

-- 9. clase.tipo_reglamento
INSERT INTO clase.tipo_reglamento (id_tipo_deporte, nombre_tipo, descripcion, activo) VALUES
    (1, 'FIFA Oficial',    'Reglamento internacional para partidos de fútbol 11.',         TRUE),
    (2, 'FIVB Indoor',     'Normativa internacional de voleibol indoor.',                  TRUE),
    (3, 'FIBA Oficial',    'Reglamento de baloncesto para torneos regionales.',            TRUE),
    (4, 'AMF Microfútbol', 'Reglas de microfútbol para torneos aficionados.',              TRUE),
    (5, 'World Skate Art', 'Reglamento de patinaje artístico para categorías juveniles.', TRUE),
    (6, 'ITTF Regional',   'Reglas de ping-pong para torneos no profesionales.',          TRUE);

 
-- 10. usuarios.usuario
INSERT INTO usuarios.usuario (nombre, apellido, numero_documento, email, telefono, fecha_nacimiento, id_documento, activo) VALUES
    ('Carlos',    'Ramírez',  '1012345678', 'carlos.ramirez@mail.com',  '3101234567', '1990-05-14', 1, TRUE),
    ('Luisa',     'Mendoza',  '1023456789', 'luisa.mendoza@mail.com',   '3202345678', '1995-08-22', 1, TRUE),
    ('Andrés',    'Torres',   'AB123456',   'andres.torres@mail.com',   '3153456789', '1988-11-03', 2, TRUE),
    ('Valentina', 'Gómez',    '1034567890', 'valentina.gomez@mail.com', '3004567890', '2000-02-17', 4, TRUE),
    ('Miguel',    'Herrera',  '987654321',  'miguel.herrera@mail.com',  '3115678901', '1993-07-30', 3, TRUE),
    ('Sara',      'Pineda',   '1045678901', 'sara.pineda@mail.com',     '3126789012', '1998-03-11', 1, TRUE);

 
-- 11. usuarios.contrasena
INSERT INTO usuarios.contrasena (contrasena, hash_contrasena, id_usuario, activo) VALUES
    ('Pass@2024!',   '$2b$10$hashCarlos001',  1, TRUE),
    ('Segura#456',   '$2b$10$hashLuisa002',   2, TRUE),
    ('Torneo$789',   '$2b$10$hashAndres003',  3, TRUE),
    ('Deporte@321',  '$2b$10$hashValenti004', 4, TRUE),
    ('Club%2024',    '$2b$10$hashMiguel005',  5, TRUE),
    ('Sara@Dep2024', '$2b$10$hashSara006',    6, TRUE);

 
-- 12. usuarios.historial_acceso
INSERT INTO usuarios.historial_acceso (exitoso, ip_origen, fallo_motivo, id_usuario, id_contrasena, activo) VALUES
    (TRUE,  '192.168.1.10', NULL,                    1, 1, TRUE),
    (FALSE, '192.168.1.21', 'Contraseña incorrecta', 2, 2, TRUE),
    (TRUE,  '10.0.0.5',     NULL,                    3, 3, TRUE),
    (FALSE, '172.16.0.33',  'Usuario no encontrado', 4, 4, TRUE),
    (TRUE,  '192.168.2.44', NULL,                    5, 5, TRUE),
    (TRUE,  '10.0.1.12',    NULL,                    6, 6, TRUE);

 
-- 13. gestion_deportiva.torneo  (id_imagen = NULL por ahora)
--insert de gestión_deportiva
INSERT INTO gestion_deportiva.torneo
    (id_tipo_deporte, nombre_torneo, fecha_inicio, fecha_fin, ubicacion, objetivo,
     id_imagen, fecha_fase, cantidad_equipo, id_tipo_distribucion, id_premiacion, activo)
VALUES
    (1, 'Copa Regional de Fútbol 2024',
     '2024-03-01', '2024-03-20',
     'Estadio El Campín, Bogotá',
     'Promover el fútbol juvenil y descubrir nuevos talentos a nivel regional.',
     NULL,
     'Grupos: 01-05/03 | Cuartos: 08/03 | Semifinal: 15/03 | Final: 20/03',
     16, 1, 1, TRUE),

    (2, 'Liga Voleibol Femenino 2024',
     '2024-04-05', '2024-04-25',
     'Coliseo Cubierto, Cali',
     'Impulsar el voleibol femenino en la región suroccidental del país.',
     NULL,
     'Grupos: 05-12/04 | Semifinal: 20/04 | Final: 25/04',
     8, 2, 4, TRUE),

    (3, 'Torneo Baloncesto Norte 2024',
     '2024-05-10', '2024-05-28',
     'Coliseo Norte, Barranquilla',
     'Fomentar el baloncesto competitivo en la región Caribe.',
     NULL,
     'Grupos: 10-18/05 | Semifinal: 24/05 | Final: 28/05',
     8, 3, 6, TRUE),

    (4, 'Torneo Microfútbol Interbarrios 2024',
     '2024-06-01', '2024-06-15',
     'Cancha Sintética La Victoria, Bogotá',
     'Integrar comunidades barriales a través del microfútbol.',
     NULL,
     'Grupos: 01-08/06 | Semifinal: 12/06 | Final: 15/06',
     12, 4, 7, TRUE),

    (5, 'Copa Patinaje Artístico 2024',
     '2024-07-20', '2024-07-28',
     'Pista de Patinaje, Medellín',
     'Identificar y premiar talentos de patinaje artístico en categoría juvenil.',
     NULL,
     'Eliminatoria: 20-24/07 | Final: 28/07',
     6, 5, 8, TRUE),

    (6, 'Torneo Ping-Pong Regional 2024',
     '2024-08-05', '2024-08-12',
     'Club Deportivo Oriente, Tunja',
     'Desarrollar habilidades y fomentar la competencia en ping-pong regional.',
     NULL,
     'Grupos: 05-08/08 | Semifinal: 10/08 | Final: 12/08',
     8, 6, 9, TRUE);

 
-- 14. gestion_deportiva.imagen  (ahora que los torneos ya existen)
INSERT INTO gestion_deportiva.imagen (id_torneo, url_imagen, tipo_imagen, activo) VALUES
    (1, 'https://example.com/torneo1.jpg', 'JPG',  TRUE),
    (2, 'https://example.com/torneo2.png', 'PNG',  TRUE),
    (3, 'https://example.com/torneo3.webp', 'WEBP', TRUE),
    (4, 'https://example.com/torneo4.jpg', 'JPG',  TRUE),
    (5, 'https://example.com/torneo5.png', 'PNG',  TRUE),
    (6, 'https://example.com/torneo6.jpg', 'JPG',  TRUE); 

-- 15. UPDATE: asignar id_imagen a cada torneo
UPDATE gestion_deportiva.torneo SET id_imagen = 1 WHERE id_torneo = 1;
UPDATE gestion_deportiva.torneo SET id_imagen = 2 WHERE id_torneo = 2;
UPDATE gestion_deportiva.torneo SET id_imagen = 3 WHERE id_torneo = 3;
UPDATE gestion_deportiva.torneo SET id_imagen = 4 WHERE id_torneo = 4;
UPDATE gestion_deportiva.torneo SET id_imagen = 5 WHERE id_torneo = 5;
UPDATE gestion_deportiva.torneo SET id_imagen = 6 WHERE id_torneo = 6;


-- 16. gestion_deportiva.distribucion
INSERT INTO gestion_deportiva.distribucion (id_torneo, id_tipo_distribucion, confirmacion, activo) VALUES
    (1, 1, TRUE,  TRUE),
    (2, 2, TRUE,  TRUE),
    (3, 3, TRUE,  TRUE),
    (4, 4, TRUE,  TRUE),
    (5, 5, TRUE,  TRUE),
    (6, 6, FALSE, TRUE);

 
-- 17. gestion_deportiva.reglamento
INSERT INTO gestion_deportiva.reglamento (id_torneo, id_tipo_reglamento, id_reglas, id_tipo_deporte, idioma, activo) VALUES
    (1, 1, 1, 1, 'Español', TRUE),
    (2, 2, 2, 2, 'Español', TRUE),
    (3, 3, 3, 3, 'Español', TRUE),
    (4, 4, 4, 4, 'Español', TRUE),
    (5, 5, 5, 5, 'Español', TRUE),
    (6, 6, 6, 6, 'Español', TRUE);

 
-- 18. gestion_deportiva.integrantes
INSERT INTO gestion_deportiva.integrantes (id_equipo, nombre, posicion, correo, telefono, activo) VALUES
    (1, 'Jorge Medina',    'Delantero',       'j.medina@leones.com',     '3101112233', TRUE),
    (2, 'Diana Salcedo',   'Central',         'd.salcedo@rayos.com',     '3202223344', TRUE),
    (3, 'Pablo Ríos',      'Alero',           'p.rios@tigres.com',       '3153334455', TRUE),
    (4, 'Felipe Castillo', 'Portero',         'f.castillo@condores.com', '3004445566', TRUE),
    (5, 'Sofía Morales',   'Artístico Libre', 's.morales@estrellas.com', '3115556677', TRUE),
    (6, 'Camilo Vega',     'Ataque Rápido',   'c.vega@dragones.com',     '3226667788', TRUE);
 

-- 19. gestion_encuentro.encuentro
INSERT INTO gestion_encuentro.encuentro (fecha, id_torneo, id_tipo_distribucion, fase_torneo, id_equipo1, resultado_equipo1, id_equipo2, resultado_equipo2, activo) VALUES
    ('2024-03-20 15:00:00', 1, 1, 'Final', 1, 'Ganador',  2, 'Perdedor', TRUE),
    ('2024-04-25 14:00:00', 2, 2, 'Final', 2, 'Ganador',  3, 'Perdedor', TRUE),
    ('2024-05-28 16:00:00', 3, 3, 'Final', 3, 'Ganador',  4, 'Perdedor', TRUE),
    ('2024-06-15 10:00:00', 4, 4, 'Final', 4, 'Ganador',  5, 'Perdedor', TRUE),
    ('2024-07-28 09:00:00', 5, 5, 'Final', 5, 'Ganador',  6, 'Perdedor', TRUE),
    ('2024-08-12 11:00:00', 6, 6, 'Final', 6, 'Ganador',  1, 'Perdedor', TRUE);

 
-- 20. gestion_encuentro.grupo_equipo
INSERT INTO gestion_encuentro.grupo_equipo (id_grupo, id_equipo, partidos_jugados, empates, victorias, derrotas, activo) VALUES
    (1, 1, 5, 1, 3, 1, TRUE),
    (2, 2, 5, 0, 4, 1, TRUE),
    (3, 3, 4, 2, 2, 0, TRUE),
    (4, 4, 6, 1, 4, 1, TRUE),
    (5, 5, 5, 0, 5, 0, TRUE),
    (6, 6, 4, 1, 2, 1, TRUE);
 

-- 21. gestion_encuentro.grupo_encuentro
INSERT INTO gestion_encuentro.grupo_encuentro (id_grupo, id_encuentro, activo) VALUES
    (1, 1, TRUE),
    (2, 2, TRUE),
    (3, 3, TRUE),
    (4, 4, TRUE),
    (5, 5, TRUE),
    (6, 6, TRUE);
 

-- 22. gestion_deportiva.clasificacion
INSERT INTO gestion_deportiva.clasificacion (id_encuentro, partido_jugado, partido_ganado, partido_empatado, partido_derrota, puntos_partidos, id_torneo, id_equipo, activo) VALUES
    (1, 5, 4, 1, 0, 13, 1, 1, TRUE),
    (2, 5, 4, 0, 1, 12, 2, 2, TRUE),
    (3, 4, 3, 1, 0, 10, 3, 3, TRUE),
    (4, 6, 5, 1, 0, 16, 4, 4, TRUE),
    (5, 5, 5, 0, 0, 15, 5, 5, TRUE),
    (6, 4, 3, 0, 1,  9, 6, 6, TRUE);

 
--23. gestión_deportiva.historialtorneos
INSERT INTO gestion_deportiva.historialtorneos
    (nombre_torneo, id_torneo, id_usuario, deporte, reglamento, equipos, id_tipo_distribucion, distribucion, estado, finalizacion, activo)
VALUES
    ('Copa Regional de Fútbol 2024',        1, 1, 'Fútbol',      'FIFA Oficial',    16, 1, 'Liga fase grupos',      'Finalizado', 'Ganador', TRUE),
    ('Liga Voleibol Femenino 2024',          2, 2, 'Voleibol',    'FIVB Indoor',      8, 2, 'Eliminatoria directa',  'Finalizado', 'Ganador', TRUE),
    ('Torneo Baloncesto Norte 2024',         3, 3, 'Baloncesto',  'FIBA Oficial',     8, 3, 'Grupos y llaves',       'Finalizado', 'Ganador', TRUE),
    ('Torneo Microfútbol Interbarrios 2024', 4, 4, 'Microfútbol', 'AMF Microfútbol', 12, 4, 'Round Robin',           'Finalizado', 'Ganador', TRUE),
    ('Copa Patinaje Artístico 2024',         5, 5, 'Patinaje',    'World Skate Art',  6, 5, 'Series eliminatorias',  'Finalizado', 'Ganador', TRUE),
    ('Torneo Ping-Pong Regional 2024',       6, 6, 'Ping-Pong',   'ITTF Regional',    8, 6, 'Grupos y semifinal',    'Finalizado', 'Ganador', TRUE);