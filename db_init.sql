drop table if exists tr_favorites;
drop table if exists tr_users;

create table if not exists tr_users(
	pk_user serial primary key,
	username varchar(20) unique,
	names varchar(50),
	surnames varchar(50),
	email varchar(50) unique,
	password text
);

create table if not exists tr_favorites(
	pk_favorite serial primary key,
	video_playlist_id varchar unique not null,
	item JSONB not null,
	fk_user int4 references tr_users
);
