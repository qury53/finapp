CREATE schema finapp;

CREATE TABLE finapp.users (
    id SERIAL PRIMARY KEY,
    version INT NOT NULL DEFAULT 1,
    username VARCHAR(32) UNIQUE NOT NULL CHECK(char_length(username) BETWEEN 3 AND 32),
    password VARCHAR(72) NOT NULL CHECK(char_length(password) BETWEEN 6 AND 72),
    email VARCHAR(255) NOT NULL CHECK(email ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$'),
    phone_number VARCHAR(15) CHECK (
        phone_number ~ '^\+[0-9]+$'
        AND char_length(phone_number) BETWEEN 10 AND 15
    ),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE finapp.user_providers (
    id SERIAL PRIMARY KEY,
    provider VARCHAR(100) UNIQUE NOT NULL,
    provider_id INTEGER UNIQUE NOT NULL,
    user_id INTEGER NOT NULL REFERENCES finapp.users(id)
);

CREATE TYPE si_status AS ENUM ('oncely', 'weekly', 'monthly');

CREATE TABLE finapp.spending (
    id SERIAL PRIMARY KEY,
    version INT NOT NULL DEFAULT 1,
    value NUMERIC(12, 2) NOT NULL,
    title VARCHAR(100) NOT NULL CHECK(char_length(title) BETWEEN 1 AND 100),
    description VARCHAR(500),
    status si_status NOT NULL DEFAULT 'oncely',
    category VARCHAR(100) NOT NULL DEFAULT 'Без категории',
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    user_id INTEGER NOT NULL REFERENCES finapp.users(id) ON DELETE CASCADE
);

CREATE TABLE finapp.income (
    id SERIAL PRIMARY KEY,
    version INT NOT NULL DEFAULT 1,
    value NUMERIC(12, 2) NOT NULL,
    title VARCHAR(100) NOT NULL CHECK(char_length(title) BETWEEN 1 AND 100),
    description VARCHAR(500),
    status si_status NOT NULL DEFAULT 'monthly',
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    user_id INTEGER NOT NULL REFERENCES finapp.users(id) ON DELETE CASCADE
);
