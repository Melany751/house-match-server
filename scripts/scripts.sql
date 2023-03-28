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
    user_id uuid not null constraint users_roles_users_id_fk references domain.users,
    role_id uuid not null constraint users_roles_roles_id_fk references domain.roles
);

create table domain.medias
(
    id   uuid default gen_random_uuid() not null primary key,
    name varchar                        not null,
    url  varchar,
    size double precision,
    type varchar
);

create table domain.properties
(
    id               uuid default gen_random_uuid() not null primary key unique,
    user_id          uuid                           not null constraint users_roles_users_id_fk references domain.users,
    location_id      uuid constraint properties_location_id_fk references domain.locations,
    "description"    varchar,
    "type"           varchar                        not null,
    "length"         float,
    width            float,
    area             float,
    floor            int                            not null,
    number_of_floors int,
    rooms            int,
    bathrooms        int,
    yard             int,
    terrace          int,
    living_room      int,
    laundry_room     int,
    kitchen          int,
    garage           int
);

create table if not exists domain.properties_medias
(
    property_id uuid not null constraint properties_medias_properties_id_fk references domain.properties,
    media_id uuid not null constraint properties_medias_medias_id_fk references domain.medias
);

create table domain.locations
(
    id       uuid default gen_random_uuid() not null primary key,
    country  varchar                        not null,
    city     varchar                        not null,
    province varchar,
    district varchar,
    address  varchar,
    lat      double precision,
    long     double precision
);

create table domain.persons
(
    id             uuid default gen_random_uuid() not null primary key,
    document_type  varchar                        not null,
    document       varchar                        not null,
    "names"        varchar                        not null,
    lastname       varchar                        not null,
    m_lastname     varchar,
    phone          varchar,
    gender         varchar                        not null,
    marital_status varchar,
    date_birth     time,
    photo          uuid constraint persons_media_id_fk references domain.medias,
    location_id    uuid constraint persons_location_person_id_fk references domain.location_person
);