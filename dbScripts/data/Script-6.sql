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

create table battles (
	id serial not null primary key,
	id_winner int not null references account(id),
	id_figter2 int not null references account(id),
	id_competition int not null references competitions(id),
	score_winner int not null,
	score_figher2 int not null
);

CREATE table users (
	email varchar(100) not null unique references account(email) primary key,
	password varchar(1000) not null,
	role varchar(10) not null check(role = 'admin' or role = 'user')
);

drop table users;

create table account (
	id bigserial not null primary key,
	name varchar not null,
	birthday date not null,
	sex varchar not null check (sex = 'female' or sex = 'male'),
	email varchar not null unique 
);

create table athlets (
	idffr bigserial not null primary key,
	id_account int not null references account(id),
	hand varchar check(hand = 'left' or hand = 'right'),
	insurance bool default false,
	license bool default false,
	weapon_type varchar not null,
	rank varchar
);

drop table athlets ;

select id from competitions c where age_category = 'взрослые' and sex = 'female';

create table AthletComp (
	id bigserial not null primary key,
	email varchar(100) not null references users(email),
	id_competition int not null references competitions(id),
	constraint uniquePair unique (email, id_competition)	
); 

drop table athletcomp;

create or replace function get_age_category(b date) returns varchar(10) as $$
	begin 
		if (current_date - b) < 12 * 365 then 
			return 'дети';
		end if;
		if (current_date - b) < 16 * 365 then
			return 'кадеты';
		end if;
		if (current_date - b) < 23 * 365 then
			return 'юниоры';
		end if;
		return 'взрослые';
	end;
$$ language plpgsql;

select get_age_category(current_date);

create trigger updateNum before insert on AthletComp
for each row execute function updateNum();

drop trigger updateNum on athletcomp ;

create or replace function updateNum() returns trigger as $$
	declare 
		bd date;
	begin
		bd := (select birthday from account a where a.email = new.email);
		if (select sex from account a where a.email = new.email) = 
		   (select sex from competitions c where c.id = new.id_competition) and 
		   (get_age_category(bd) = (select age_category from competitions c where c.id = new.id_competition)) then 
			update competitions
			set numOfAthlets = (select numOfAthlets from competitions where id=new.id_competition) + 1
			where id=new.id_competition;
			return new;
		else
			raise exception 'Вы не можете подать заявку на данный турнир';
--			return null;
		end if;
	end;
$$ language plpgsql;

create trigger decNum after delete on AthletComp
for each row execute function decNum();

drop trigger decNum on athletcomp ;

create or replace function decNum() returns trigger as $$
	begin
		update competitions
		set numOfAthlets = (select numOfAthlets from competitions where id=old.id_competition) - 1
		where id=old.id_competition;
		return old;
	end
$$ language plpgsql;

insert into account (name, birthday, sex, email)
values ('LOL', current_date, 'female', 'lol@example.com');

insert into users(email, "password", "role")
values ('ololo@example.com', 'koko', 'user');

delete from athletcomp 
where id_competition = 11;

select * from account a ;
select * from users u;

insert into athletcomp (email, id_competition)
values ('lol@example.com', 12);

delete from athletcomp where id_competition = 12;

select * from athletcomp a ;

select *
from competitions
where id = 12;

update competitions 
set numofathlets = 0
where id = 12;

drop role myapp;

create role myapp
login
password 'myapp';

grant all
on athletcomp
to myapp;

revoke all on users from myapp;

create role guest login password '1234';
grant select on table competitions to guest;
grant insert on table users to guest;

create role use login password '1234';
grant insert on table  athletcomp to use;
grant delete on table  athletcomp to use;
grant select on table  athletcomp to use;

grant insert on table  users to use;
grant select on table  competitions to use;
grant select on table  account to use;

create role administrator login password '1234';
grant all privileges on table users to administrator;
grant all privileges on table competitions to administrator;
grant all privileges on table battles to administrator;
grant all privileges on table athletcomp to administrator;
grant all privileges on table athlets to administrator;








