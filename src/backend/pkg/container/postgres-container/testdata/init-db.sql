DROP TABLE IF EXISTS TICKET;
DROP TABLE IF EXISTS POSTER;
DROP TABLE IF EXISTS EVENT;
DROP TABLE IF EXISTS PLACE;
DROP TABLE IF EXISTS COMIC;
DROP TABLE IF EXISTS PERSON;
DROP TABLE IF EXISTS IMAGE;

CREATE TABLE Image(
                      id SERIAL PRIMARY KEY NOT NULL,
                      path TEXT UNIQUE NOT NULL
);

CREATE TABLE Person(
                       id SERIAL PRIMARY KEY NOT NULL,
                       login TEXT UNIQUE NOT NULL,
                       email TEXT UNIQUE NOT NULL,
                       salt TEXT NOT NULL DEFAULT 'ya4tr78gf78ert8',
                       hash TEXT NOT NULL,
                       role INTEGER NOT NULL DEFAULT 3
);

CREATE TABLE Comic(
                      id SERIAL PRIMARY KEY NOT NULL ,
                      image TEXT DEFAULT '',
                      name TEXT NOT NULL ,
                      city TEXT DEFAULT 'Столица юмора',
                      sentence TEXT DEFAULT '',
                      count_kek INTEGER DEFAULT 0 CHECK ( count_kek >= 0)
);

CREATE TABLE Place(
                      id SERIAL PRIMARY KEY NOT NULL,
                      name TEXT UNIQUE NOT NULL,
                      count_ticket INTEGER NOT NULL CHECK ( count_ticket >= 0 )
);

CREATE TABLE Event(
                      id SERIAL PRIMARY KEY NOT NULL,
                      name TEXT UNIQUE NOT NULL,
                      about TEXT DEFAULT ''
);

CREATE TABLE Poster(
                       id SERIAL PRIMARY KEY NOT NULL,
                       image TEXT DEFAULT '',
                       id_event INTEGER REFERENCES Event(id) NOT NULL,
                       id_plate INTEGER REFERENCES Place(id) NOT NULL,
                       date TIMESTAMP NOT NULL
);

CREATE TABLE Ticket(
                       id SERIAL PRIMARY KEY NOT NULL,
                       id_poster INTEGER REFERENCES Poster(id) NOT NULL,
                       id_person INTEGER REFERENCES Person(id) NOT NULL
);

