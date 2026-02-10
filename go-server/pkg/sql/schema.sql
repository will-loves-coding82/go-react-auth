create database if not exists blueprint;

create table if not exists users (
    id serial PRIMARY KEY,
    google_id varchar(255) NOT NULL UNIQUE,
    email varchar(255) NOT NULL UNIQUE,
    picture_url varchar(255)
);