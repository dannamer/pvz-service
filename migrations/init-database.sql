-- Active: 1744228605253@@127.0.0.1@5432@postgres
CREATE TABLE users (
    id UUID PRIMARY KEY,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    role TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL
);

CREATE TABLE pvz (
    id UUID PRIMARY KEY,
    created_by UUID,
    city TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    CONSTRAINT fk_created_by FOREIGN KEY (created_by) REFERENCES users(id) ON DELETE SET NULL
);

CREATE TABLE receptions (
    id UUID PRIMARY KEY,
    pvz_id UUID NOT NULL,
    created_by UUID NOT NULL,
    status VARCHAR(20) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    CONSTRAINT fk_pvz FOREIGN KEY (pvz_id) REFERENCES pvz(id) ON DELETE CASCADE,
    CONSTRAINT fk_created_by FOREIGN KEY (created_by) REFERENCES users(id) ON DELETE CASCADE
);
CREATE TABLE products (
    id UUID PRIMARY KEY,
    pvz_id UUID NOT NULL,
    reception_id UUID NOT NULL,
    type TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,

    CONSTRAINT fk_pvz FOREIGN KEY (pvz_id) REFERENCES pvz(id) ON DELETE CASCADE,
    CONSTRAINT fk_reception FOREIGN KEY (reception_id) REFERENCES receptions(id) ON DELETE CASCADE
);