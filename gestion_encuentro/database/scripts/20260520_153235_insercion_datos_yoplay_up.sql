
-- 7. gestion_encuentro.grupo
INSERT INTO gestion_encuentro.grupo (grupo, activo) VALUES
    ('Grupo A', TRUE),
    ('Grupo B', TRUE),
    ('Grupo C', TRUE),
    ('Grupo D', TRUE),
    ('Grupo E', TRUE),
    ('Grupo F', TRUE);
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