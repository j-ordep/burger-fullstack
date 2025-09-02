-- Active: 1756584010900@@localhost@5433@burger-gorm
ALTER TABLE burger_opcionais
ADD CONSTRAINT fk_burger_opcionais_pedidos
FOREIGN KEY (pedidos_id)
REFERENCES pedidos(id)
ON DELETE CASCADE;

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