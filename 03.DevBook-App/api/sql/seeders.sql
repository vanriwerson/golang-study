INSERT INTO devbook.usuarios (nome, nick, email, senha)
VALUES ("Adalberto", "Dadá", "adalberto@email.com", "$2a$10$jSgHypvvykWC4jJcFzKKV.J5TFSBFGnQGI2xcfnLWTgfbIpEYNlxa"), #senha: 123456
("Giovane", "Gigio", "giovane@email.com", "$2a$10$jSgHypvvykWC4jJcFzKKV.J5TFSBFGnQGI2xcfnLWTgfbIpEYNlxa"), #senha: 123456
("Jessy", "Jess", "damasceno@email.com", "$2a$10$jSgHypvvykWC4jJcFzKKV.J5TFSBFGnQGI2xcfnLWTgfbIpEYNlxa"), #senha: 123456
("Rafael", "Primo", "rafael@email.com", "$2a$10$jSgHypvvykWC4jJcFzKKV.J5TFSBFGnQGI2xcfnLWTgfbIpEYNlxa"), #senha: 123456
("Maristela", "Mari", "maristela@email.com", "$2a$10$jSgHypvvykWC4jJcFzKKV.J5TFSBFGnQGI2xcfnLWTgfbIpEYNlxa"), #senha: 123456
("Ana", "Aninha", "anahelena@email.com", "$2a$10$jSgHypvvykWC4jJcFzKKV.J5TFSBFGnQGI2xcfnLWTgfbIpEYNlxa"), #senha: 123456
("Lilian", "Lily", "lilian@email.com", "$2a$10$jSgHypvvykWC4jJcFzKKV.J5TFSBFGnQGI2xcfnLWTgfbIpEYNlxa"), #senha: 123456
("Gabriel", "Leite", "leite@email.com", "$2a$10$jSgHypvvykWC4jJcFzKKV.J5TFSBFGnQGI2xcfnLWTgfbIpEYNlxa"), #senha: 123456
("Bruno", "Riwerson", "bruno@email.com", "$2a$10$tU2.Ka7nlOAafo1u/gpLiOpEIDhCNnPqoH2AM0KbuNF7Ys4kmaT2K"); # senha: admin

INSERT INTO devbook.seguidores (usuario_id, seguidor_id)
VALUES (1, 3), (1, 4), (3, 1), (3, 4), (5, 6), (6, 5);