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

create table domain.views
(
    id          uuid default gen_random_uuid() not null primary key unique,
    module_id   uuid                           not null constraint views_modules_id_fk references domain.modules,
    "name"      varchar,
    description varchar,
    url         varchar,
    icon        varchar
);

create table domain.roles_views
(
    role_id       uuid not null constraint roles_views_roles_id_fk references domain.roles,
    view_id       uuid not null constraint roles_views_views_id_fk references domain.views,
    view_order    int,
    view_position varchar
);