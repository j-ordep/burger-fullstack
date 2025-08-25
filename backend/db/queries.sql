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

CREATE TABLE burgers (
    id SERIAL PRIMARY KEY,
    nome_cliente TEXT NOT NULL,
    pao_id INT REFERENCES paes(id),
    carne_id INT REFERENCES carnes(id),
    status_id INT REFERENCES status(id),
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE burger_opcionais (
    burger_id INTEGER,
    opcional_id INTEGER,
    PRIMARY KEY (burger_id, opcional_id),
    FOREIGN KEY (burger_id) REFERENCES burgers(id) ON DELETE CASCADE,
    FOREIGN KEY (opcional_id) REFERENCES opcionais(id)
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
