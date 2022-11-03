INSERT INTO devbook.usuarios (nome, nick, email, senha)
VALUES ("Adalberto", "Dadá", "adalberto@email.com", "superSenhaSecreta1"),
("Giovane", "Gigio", "giovane@email.com", "superSenhaSecreta2"),
("Jessy", "Jess", "damasceno@email.com", "superSenhaSecreta3"),
("Rafael", "Primo", "rafael@email.com", "superSenhaSecreta4"),
("Maristela", "Mari", "maristela@email.com", "superSenhaSecreta5"),
("Ana", "Aninha", "anahelena@email.com", "superSenhaSecreta6"),
("Clair", "JojoDuck", "clair@email.com", "superSenhaSecreta7");

INSERT INTO devbook.seguidores (usuario_id, seguidor_id)
VALUES (1, 3), (1, 4), (3, 1), (3, 4), (5, 6), (6, 5);