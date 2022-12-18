DROP DATABASE IF EXISTS devbook;
CREATE DATABASE devbook;
USE devbook;

DROP TABLE IF EXISTS publicacoes;
DROP TABLE IF EXISTS usuarios;
DROP TABLE IF EXISTS seguidores;


CREATE TABLE usuarios
(
    id       int auto_increment primary key,
    nome     varchar(50)  not null,
    nick     varchar(50)  not null unique,
    email    varchar(50)  not null unique,
    senha    varchar(200) not null,
    criadoEm timestamp default current_timestamp()
) ENGINE = INNODB;


CREATE TABLE seguidores
(
    usuario_id  int not null,
    FOREIGN KEY (usuario_id) references usuarios (id) on delete cascade,
    seguidor_id int not null,
    FOREIGN KEY (seguidor_id) references usuarios (id) on delete cascade,
    PRIMARY KEY (usuario_id, seguidor_id)
) ENGINE = INNODB;

CREATE TABLE publicacoes
(
    id       int auto_increment primary key,
    titulo   varchar(50)  not null,
    conteudo varchar(300) not null,
    autor_id int          not null,
    foreign key (autor_id) references usuarios (id) on delete cascade,

    curtidas int       default 0,
    criadaEm timestamp default current_timestamp
) ENGINE = INNODB;