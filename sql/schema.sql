CREATE TABLE users (
    id SERIAL NOT NULL PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    is_active BOOLEAN NOT NULL,
    is_staff BOOLEAN NOT NULL,
    is_superuser BOOLEAN NOT NULL,
    last_login TIMESTAMP NOT NULL,
    date_joined TIMESTAMP NOT NULL
);
CREATE TABLE category (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) UNIQUE NOT NULL
);
CREATE TABLE subcategory (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) UNIQUE NOT NULL
);
CREATE TABLE objectives (
    id SERIAL NOT NULL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users (id) DEFERRABLE INITIALLY DEFERRED,
    category_id INTEGER REFERENCES category(id),
    subcategory_id INTEGER REFERENCES subcategory(id),
    how_am_i TEXT NOT NULL,
    how_do_i_want_to_be TEXT NOT NULL,
    what_will_i_do TEXT NOT NULL,
    when_will_i_do_it TEXT NOT NULL
);
