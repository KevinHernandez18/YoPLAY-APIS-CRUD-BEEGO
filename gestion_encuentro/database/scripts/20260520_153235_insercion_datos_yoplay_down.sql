-- SCRIPT PARA ELIMINAR LAS INSERCIONES DE DATOS

-- 1. gestion_deportiva.historialtorneos
DELETE FROM gestion_deportiva.historialtorneos
WHERE nombre_torneo IN (
    'Copa Regional de Fútbol 2024',
    'Liga Voleibol Femenino 2024',
    'Torneo Baloncesto Norte 2024',
    'Torneo Microfútbol Interbarrios 2024',
    'Copa Patinaje Artístico 2024',
    'Torneo Ping-Pong Regional 2024'
);

-- 2. gestion_deportiva.clasificacion
DELETE FROM gestion_deportiva.clasificacion
WHERE id_encuentro IN (1, 2, 3, 4, 5, 6)
  AND id_torneo IN (1, 2, 3, 4, 5, 6)
  AND id_equipo IN (1, 2, 3, 4, 5, 6);

-- 3. gestion_encuentro.grupo_encuentro
DELETE FROM gestion_encuentro.grupo_encuentro
WHERE id_encuentro IN (1, 2, 3, 4, 5, 6);

-- 4. gestion_encuentro.grupo_equipo
DELETE FROM gestion_encuentro.grupo_equipo
WHERE id_grupo IN (1, 2, 3, 4, 5, 6)
  AND id_equipo IN (1, 2, 3, 4, 5, 6);

-- 5. gestion_encuentro.encuentro
DELETE FROM gestion_encuentro.encuentro
WHERE fecha IN (
    '2024-03-20 15:00:00',
    '2024-04-25 14:00:00',
    '2024-05-28 16:00:00',
    '2024-06-15 10:00:00',
    '2024-07-28 09:00:00',
    '2024-08-12 11:00:00'
)
  AND id_torneo IN (1, 2, 3, 4, 5, 6);

-- 6. gestion_deportiva.integrantes
DELETE FROM gestion_deportiva.integrantes
WHERE id_equipo IN (1, 2, 3, 4, 5, 6)
  AND nombre IN (
    'Jorge Medina',
    'Diana Salcedo',
    'Pablo Ríos',
    'Felipe Castillo',
    'Sofía Morales',
    'Camilo Vega'
);

-- 7. gestion_deportiva.reglamento
DELETE FROM gestion_deportiva.reglamento
WHERE id_torneo IN (1, 2, 3, 4, 5, 6)
  AND id_tipo_reglamento IN (1, 2, 3, 4, 5, 6)
  AND id_reglas IN (1, 2, 3, 4, 5, 6)
  AND id_tipo_deporte IN (1, 2, 3, 4, 5, 6);

-- 8. gestion_deportiva.distribucion
DELETE FROM gestion_deportiva.distribucion
WHERE id_torneo IN (1, 2, 3, 4, 5, 6)
  AND id_tipo_distribucion IN (1, 2, 3, 4, 5, 6);

-- 9. gestion_deportiva.imagen
DELETE FROM gestion_deportiva.imagen
WHERE id_torneo IN (1, 2, 3, 4, 5, 6)
  AND url_imagen IN (
    'https://example.com/torneo1.jpg',
    'https://example.com/torneo2.png',
    'https://example.com/torneo3.webp',
    'https://example.com/torneo4.jpg',
    'https://example.com/torneo5.png',
    'https://example.com/torneo6.jpg'
);

-- 10. Reiniciar referencias a imagen en torneo antes de borrar los torneos
UPDATE gestion_deportiva.torneo
SET id_imagen = NULL
WHERE id_torneo IN (1, 2, 3, 4, 5, 6);

-- 11. gestion_deportiva.torneo
DELETE FROM gestion_deportiva.torneo
WHERE id_torneo IN (1, 2, 3, 4, 5, 6)
  AND nombre_torneo IN (
    'Copa Regional de Fútbol 2024',
    'Liga Voleibol Femenino 2024',
    'Torneo Baloncesto Norte 2024',
    'Torneo Microfútbol Interbarrios 2024',
    'Copa Patinaje Artístico 2024',
    'Torneo Ping-Pong Regional 2024'
);

-- 12. usuarios.historial_acceso
DELETE FROM usuarios.historial_acceso
WHERE id_usuario IN (1, 2, 3, 4, 5, 6)
  AND id_contrasena IN (1, 2, 3, 4, 5, 6);

-- 13. usuarios.contrasena
DELETE FROM usuarios.contrasena
WHERE id_usuario IN (1, 2, 3, 4, 5, 6)
  AND hash_contrasena IN (
    '$2b$10$hashCarlos001',
    '$2b$10$hashLuisa002',
    '$2b$10$hashAndres003',
    '$2b$10$hashValenti004',
    '$2b$10$hashMiguel005',
    '$2b$10$hashSara006'
);

-- 14. usuarios.usuario
DELETE FROM usuarios.usuario
WHERE email IN (
    'carlos.ramirez@mail.com',
    'luisa.mendoza@mail.com',
    'andres.torres@mail.com',
    'valentina.gomez@mail.com',
    'miguel.herrera@mail.com',
    'sara.pineda@mail.com'
);

-- 15. clase.tipo_reglamento
DELETE FROM clase.tipo_reglamento
WHERE id_tipo_reglamento IN (1, 2, 3, 4, 5, 6)
  AND nombre_tipo IN (
    'FIFA Oficial',
    'FIVB Indoor',
    'FIBA Oficial',
    'AMF Microfútbol',
    'World Skate Art',
    'ITTF Regional'
);

-- 16. gestion_encuentro.grupo
DELETE FROM gestion_encuentro.grupo
WHERE grupo IN ('Grupo A', 'Grupo B', 'Grupo C', 'Grupo D', 'Grupo E', 'Grupo F');

-- 17. gestion_deportiva.equipo
DELETE FROM gestion_deportiva.equipo
WHERE nombre_equipo IN (
    'Leones FC',
    'Rayos Voleibol',
    'Tigres Baloncesto',
    'Cóndores Micro',
    'Estrellas Patinaje',
    'Dragones Ping-Pong'
);

-- 18. gestion_deportiva.premiacion
DELETE FROM gestion_deportiva.premiacion
WHERE descripcion IN (
    '1er Puesto - Copa Dorada | Copa Regional de Fútbol 2024',
    '2do Puesto - Medalla de Plata | Copa Regional de Fútbol 2024',
    '3er Puesto - Reconocimiento de Bronce | Copa Regional de Fútbol 2024',
    '1er Puesto - Trofeo Campeón | Liga Voleibol Femenino 2024',
    '2do Puesto - Medalla Subcampeón | Liga Voleibol Femenino 2024',
    '1er Puesto - Copa Campeón | Torneo Baloncesto Norte 2024',
    '1er Puesto - Trofeo de Oro | Torneo Microfútbol Interbarrios 2024',
    '1er Puesto - Medalla Dorada | Copa Patinaje Artístico 2024',
    '1er Puesto - Trofeo de Cristal | Torneo Ping-Pong Regional 2024'
);

-- 19. usuarios.documento
DELETE FROM usuarios.documento
WHERE tipo_documento IN (
    'Cédula de Ciudadanía',
    'Pasaporte',
    'DNI',
    'Tarjeta de Identidad',
    'Cédula Extranjería'
);

-- 20. clase.reglas
DELETE FROM clase.reglas
WHERE descripcion IN (
    'Reglamento oficial FIFA para torneos de fútbol 11.',
    'Reglamento FIVB para torneos de voleibol indoor.',
    'Reglamento FIBA para competencias de baloncesto.',
    'Normativa AMF para torneos de microfútbol.',
    'Reglamento World Skate para competencias de patinaje artístico.',
    'Reglamento ITTF para torneos de ping-pong.'
);

-- 21. clase.tipo_distribucion
DELETE FROM clase.tipo_distribucion
WHERE equipos IN ('16', '8', '12', '6');

-- 22. clase.tipo_deporte
DELETE FROM clase.tipo_deporte
WHERE nombre_deporte IN (
    'Fútbol',
    'Voleibol',
    'Baloncesto',
    'Microfútbol',
    'Patinaje',
    'Ping-Pong'
);

-- 23. tutoriales
DELETE FROM tutoriales
WHERE nombre_tutorial IN (
    'Crear un torneo',
    'Registrar equipos',
    'Configurar reglamento',
    'Ver historial',
    'Subir imágenes',
    'Gestionar encuentros'
);
