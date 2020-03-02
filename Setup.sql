
-- 記事テーブル
create table articles (
	 	id int not null AUTO_INCREMENT,
	 	title varchar(100) not null,
	 	body varchar(5000) not null,
        create_user_id int not null ,
	 	created_at timestamp not null DEFAULT CURRENT_TIMESTAMP,
	 	primary key (id));

-- ユーザテーブル
create table users (
  user_id         int AUTO_INCREMENT primary key,
  name       varchar(255),
  email      varchar(255) not null unique,
  password   varchar(255) not null,
  created_at timestamp not null  DEFAULT CURRENT_TIMESTAMP
);

-- セッションテーブル
create table sessions (
  id         serial primary key,
  email      varchar(255),
  user_id    integer references users(id),
  created_at timestamp not null   
);
