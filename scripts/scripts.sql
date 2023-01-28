create database housematch;

create schema domain;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create table domain.users
(
    id       uuid default uuid_generate_v4() not null primary key unique,
    user     varchar                         not null unique,
    password varchar                         not null,
    email    varchar                         not null unique,
    theme    varchar
);
