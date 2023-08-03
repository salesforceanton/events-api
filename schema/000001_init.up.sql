CREATE TABLE users 
(
    id serial not null unique,
    email varchar(255) not null,
    username varchar(255) not null,
    password_hash varchar(255) not null
);

CREATE TABLE events 
(
    id serial not null unique,
    title varchar(255) not null,
    timezoneId varchar(255) not null,
    startDatetime timestamp not null,
    organizerId int references users(id) on delete cascade not null,
    description varchar(255)
);
