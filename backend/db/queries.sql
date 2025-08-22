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

CREATE TABLE burger_opcionais (
    burger_id INTEGER REFERENCES burgers(id) ON DELETE CASCADE,
    opcional_id INTEGER REFERENCES opcionais(id),
    PRIMARY KEY (burger_id, opcional_id)
);

CREATE TABLE burgers (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(100) NOT NULL,
    pao VARCHAR(50) NOT NULL,
    carne VARCHAR(50) NOT NULL,
    status VARCHAR(50) NOT NULL
);