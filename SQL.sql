SET search_path TO public;

select *from pg_available_extensions;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

SELECT uuid_generate_v4();


create Table users(
    id uuid primary key not null default uuid_generate_v4(),
    login varchar not null unique,
    password varchar not null,
    name varchar not null,
    access_token text not null,
    active bool,
    created_at timestamptz not null default current_timestamp,
    updated_at timestamptz
);

create table tokens (
--     id uuid primary key not null default uuid_generate_v4(),
    user_id uuid not null references users,
    token text not null,
    platform text not null,
    created_at timestamptz not null default current_timestamp
);

create table organization (
    id uuid primary key not null default uuid_generate_v4(),
    name varchar not null,
    created_by uuid references users,
    active bool not null,
    created_at timestamptz not null default current_timestamp,
    updated_at timestamptz
    );

-- organization_users

create table projects(
    id uuid primary key not null default uuid_generate_v4(),
    name varchar not null,
    remove_right bool not null,
    created_by uuid references users,
    organization_id uuid references organization,
    active bool not null,
    created_at timestamptz not null default current_timestamp,
    updated_at timestamptz

);

create table keys (
    id uuid primary key not null default uuid_generate_v4(),
    name varchar not null,
    rule uuid references rules,
    project_id uuid references projects
);

create table rules(
    Type RuleType
);

Enum RuleType('KeyType')

Keys Pasport_fio
    Rule_1





--
--     //// -- LEVEL 1
-- //// -- Schemas, Tables and References
--
-- // Creating tables
-- // You can define the tables with full schema names
-- Table users{
--   id uuid
--   name varchar
--   password varchar
--   active bool
--   created_at timestamptz
--   update_at timestamptz
-- }
--
-- Table tokens{
--   id uuid
--   user_id uuid [ref: > users.id]
--   token varchar
-- }
--
-- Table organization{
--   id uuid
--   name varchar
--   created_by uuid [ref: > users.id]
--   active bool
--   created_at timestamptz
--   update_at timestamptz
-- }
--
-- table projects{
--   id uuid
--   name varchar
--   remove_right bool // can use like api key
--   created_by uuid [ref: > users.id]
--   organization_id uuid [ref: > organization.id]
--   active bool
--   created_at timestamptz
--   update_at timestamptz
-- }
--
-- table keys{
--   id uuid
--   name varchar unique
--   rule uuid [ref: > rules.id]
--   project_id uuid [ref: > projects.id]
--
-- }
--
-- table rules{
--   id uuid
--   name varchar unique
--   type varchar
--   stringTest int
--   range int
--   slice int
--   bool int
-- }

