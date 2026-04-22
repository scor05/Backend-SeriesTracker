/*
    Script para init de las tablas necesarias para proyecto #1 de web
    Autor: Santiago Cordero
*/

CREATE TABLE IF NOT EXISTS series(
    id_serie INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    description TEXT,
    current_episode INT NOT NULL,
    total_episodes INT NOT NULL,
    img_src TEXT
);
