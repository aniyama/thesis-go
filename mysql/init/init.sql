CREATE DATABASE IF NOT EXISTS thesis_db;
USE thesis_db;

DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS modules;
DROP TABLE IF EXISTS thesis;
DROP TABLE IF EXISTS tags;
DROP TABLE IF EXISTS category;

CREATE TABLE users (
    id int not null auto_increment primary key,
    name varchar(32) not null,
    password varchar(128) not null
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE categories (
    id int not null auto_increment primary key,
    title varchar(32) not null
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE tags (
    id int not null auto_increment primary key,
    title varchar(32) not null,
    user_id int not null,
    category_id int not null,
    constraint fk_user_id foreign key (user_id) references users (id) on delete cascade,
    constraint fk_category_id foreign key (category_id) references categories (id) on delete cascade
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE modules (
    id int not null auto_increment primary key,
    title varchar(64) not null,
    content varchar(512) not null,
    user_id int not null,
    tag_id int ,
    created_at datetime,
    updated_at datetime,
    constraint fk_modules_user_id foreign key (user_id) references users (id) on delete cascade,
    constraint fk_tag_id foreign key (tag_id) references tags (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE theses (
    id int not null auto_increment primary key,
    thesis_title varchar(64) not null,
    thesis_content varchar(4096) not null,
    user_id int not null,
    tag_id int ,
    created_at datetime,
    updated_at datetime,
    constraint fk_thesis_user_id foreign key (user_id) references users (id) on delete cascade,
    constraint fk_thesis_tag_id foreign key (tag_id) references tags (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO categories (id, title) VALUES (1, 'THESIS');
INSERT INTO categories (id, title) VALUES (2, 'MODULE');
