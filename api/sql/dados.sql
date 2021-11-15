insert into usuarios (nome, nick, email, senha)
values
("usuário 1", "usuario_a", "a@gmail.com", "$2a$10$DH9ixdHKuyuclIc5VywcH.RGGywlW3eXMKukjlaeZ.yVo/5mO5FfK"), -- usuario 1
("usuário 2", "usuario_b", "b@gmail.com", "$2a$10$k/C9KPGttg3L.koyFZ5YIOLb18hamCxoEY1J.DCjXieqLZ020xRcG"), -- usuario 2
("usuário 3", "usuario_c", "c@gmail.com", "$2a$10$0oYlpXGO105ZQou/k3zG/u6UoMP6mVQLlvtSPj/W0e7XJOhv838RG"); -- usuario 3

insert into seguidores(usuario_id, seguidor_id)
values
(1, 2),
(3, 1),
(1, 3);

insert into publicacoes(titulo, conteudo, autor_id)
values
("Publicação do usuário 1", "Essa é a publicação do usuário 1!", 1),
("Publicação do usuário 2", "Essa é a publicação do usuário 2!", 2),
("Publicação do usuário 3", "Essa é a publicação do usuário 3!", 3);
