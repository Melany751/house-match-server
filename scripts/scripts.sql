create database housematch;

create schema domain;

create table domain.users
(
    id       uuid default gen_random_uuid() not null primary key unique,
    "user"   varchar                        not null unique,
    password varchar                        not null,
    email    varchar                        not null unique,
    theme    varchar
);

create table domain.roles
(
    id          uuid default gen_random_uuid() not null primary key unique,
    "name"      varchar                        not null,
    description varchar                        not null,
    "order"     int                            not null
);
