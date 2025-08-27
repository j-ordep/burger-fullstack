-- Active: 1755963471787@@localhost@5432@burger
CREATE TABLE paes (
    id SERIAL PRIMARY KEY,
    tipo VARCHAR(50) NOT NULL
);

CREATE TABLE carnes (
    id SERIAL PRIMARY KEY,
    tipo VARCHAR(50) NOT NULL
);

CREATE TABLE status (
    id SERIAL PRIMARY KEY,
    tipo VARCHAR(50) NOT NULL
);

CREATE TABLE opcionais (
    id SERIAL PRIMARY KEY,
    tipo VARCHAR(50) NOT NULL
);

CREATE TABLE pedidos (
    id SERIAL PRIMARY KEY,
    nome TEXT NOT NULL,
    pao TEXT NOT NULL,
    carne TEXT NOT NULL,
    status_id INT NOT NULL REFERENCES status(id)
);

CREATE TABLE burger_opcionais (
    id SERIAL PRIMARY KEY,
    pedidos_id INTEGER REFERENCES pedidos(id) ON DELETE CASCADE,
    opcional TEXT NOT NULL
);

INSERT INTO paes (id, tipo) VALUES
(1, 'Italiano Branco'),
(2, '3 Queijos'),
(3, 'Parmesão e Orégano'),
(4, 'Integral');

INSERT INTO carnes (id, tipo) VALUES
(1, 'Maminha'),
(2, 'Alcatra'),
(3, 'Picanha'),
(4, 'Veggie burger');

INSERT INTO opcionais (id, tipo) VALUES
(1, 'Bacon'),
(2, 'Cheddar'),
(3, 'Salame'),
(4, 'Tomate'),
(5, 'Cebola roxa'),
(6, 'Pepino');

INSERT INTO status (id, tipo) VALUES
(1, 'Solicitado'),
(2, 'Em produção'),
(3, 'Finalizado');