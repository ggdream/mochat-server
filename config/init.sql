use `test`;

create table if not exists `user` (
  id INT(5) PRIMARY KEY AUTO_INCREMENT NOT NULL,
  username VARCHAR(10) not null,
  password VARCHAR(10) not null
)ENGINE=InnoDB DEFAULT CHARACTER SET utf8mb4;

create table if not exists `msg` (
 id INT(10) PRIMARY KEY AUTO_INCREMENT NOT NULL,
 me VARCHAR(10) not null,
 you VARCHAR(10) not null,
 time INT(13) not null
)ENGINE=InnoDB DEFAULT CHARACTER SET utf8mb4;