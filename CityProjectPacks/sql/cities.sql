create table cities (
    id serial constraint pk_cities Primary Key,
    name varchar(200) not null,
    pop int
);