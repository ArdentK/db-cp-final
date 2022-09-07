CREATE TABLE competitions (
    id serial not null primary key,
    name varchar(240) not null,
    dt date not null,
    age_category varchar(20) not null,
    weapon_type varchar(20) not null,
    is_team bool not null,
    status varchar(100),
    sex varchar(10) not null check (sex = 'female' or sex = 'male'),
    type varchar(1000)
);

alter table competitions 
add column numOfAthlets int not null default 0;

alter table competitions 
drop column type;

CREATE table users (
	id serial not null primary key,
	email varchar(100) not null unique,
	password varchar(1000) not null,
	role varchar(10) not null check(role = 'admin' or role = 'user')
);

create table account (
	id bigserial not null primary key,
	name varchar not null,
	birthday date not null,
	sex varchar not null check (sex = 'female' or sex = 'male'),
	email varchar not null references users(email)
);

create table AthletComp (
	id bigserial not null primary key,
	id_athlet int not null references account(id),
	id_competition int not null references competitions(id),
	constraint uniquePair unique (id_athlet, id_competition)	
); 

drop table athletcomp;

create trigger updateNum after insert on AthletComp
for each row execute function updateNum();

create or replace function updateNum() returns trigger as $$
	begin 
		update competitions
		set numOfAthlets = (select numOfAthlets from competitions where id=new.id_competition) + 1
		where id=new.id_competition;
		return new;
	end;
$$ language plpgsql;

create trigger decNum after delete on AthletComp
for each row execute function decNum();

create or replace function decNum() returns trigger as $$
	begin
		update competitions
		set numOfAthlets = (select numOfAthlets from competitions where id=new.id_competition) - 1
		where id=new.id_competition;
		return new;
	end
$$ language plpgsql;

insert into account (name, birthday, sex, email)
values ('Ololo', current_date, 'female', 'ololo@example.com');

insert into users(email, "password", "role")
values ('ololo@example.com', 'koko', 'user');

delete from users 
where id = 2;

select * from users u;

insert into athletcomp (id_athlet, id_competition)
values (1, 2);

select *
from competitions
where id = 2

drop role myapp;

create role myapp
login
password 'myapp';

grant all
on users
to myapp;

revoke all on users from myapp;


