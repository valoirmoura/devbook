insert into usuarios (nome, nick, email, senha)
values ('Usuario 1', 'usuario_1', 'usuario1@gmail.com', '$2a$10$tJIozmJMYEfLrIy9JI.j0eKa3REQTOm32VC4hjo8UjxFCcZqVTnSy'),
       ('Usuario 2', 'usuario_2', 'usuario2@gmail.com', '$2a$10$tJIozmJMYEfLrIy9JI.j0eKa3REQTOm32VC4hjo8UjxFCcZqVTnSy'),
       ('Usuario 3', 'usuario_3', 'usuario3@gmail.com', '$2a$10$tJIozmJMYEfLrIy9JI.j0eKa3REQTOm32VC4hjo8UjxFCcZqVTnSy');

insert into seguidores (usuario_id, seguidor_id) values (1, 2);
insert into seguidores (usuario_id, seguidor_id) values (3, 1);
insert into seguidores (usuario_id, seguidor_id) values (1, 3);

insert into publicacoes(titulo, conteudo, autor_id) values ('Publicacao do Usuario 1', 'Esta é a Publicacao do Usuário 1', 1 );
insert into publicacoes(titulo, conteudo, autor_id) values ('Publicacao do Usuario 2', 'Esta é a Publicacao do Usuário 2', 2 );
insert into publicacoes(titulo, conteudo, autor_id) values ('Publicacao do Usuario 3', 'Esta é a Publicacao do Usuário 3', 3 );