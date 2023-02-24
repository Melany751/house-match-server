create
database housematch;

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
    module_id   uuid                           not null
        constraint views_modules_id_fk references domain.modules,
    "name"      varchar,
    description varchar,
    url         varchar,
    icon        varchar
);

create table domain.roles_views
(
    role_id       uuid not null
        constraint roles_views_roles_id_fk references domain.roles,
    view_id       uuid not null
        constraint roles_views_views_id_fk references domain.views,
    view_order    int,
    view_position varchar
);

create table domain.modules
(
    id          uuid default gen_random_uuid() not null primary key unique,
    "name"      varchar                        not null,
    description varchar,
    icon        varchar                        not null,
    "order"     int                            not null
);

create table domain.users_roles
(
    user_id uuid not null
        constraint users_roles_users_id_fk references domain.users,
    role_id uuid not null
        constraint users_roles_roles_id_fk references domain.roles
);

create table domain.properties
(
    id               uuid default gen_random_uuid() not null primary key unique,
    user_id          uuid                           not null
        constraint users_roles_users_id_fk references domain.users,
    description      varchar,
    "type"           varchar                        not null,
    "length"         float,
    width            float,
    area             float,
    floor            int                            not null,
    number_of_floors int
);

CREATE TABLE domain.location_person
(
    id       uuid default gen_random_uuid() not null primary key unique,
    country  varchar                        not null,
    city     varchar                        not null,
    province varchar,
    district varchar
);