CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TYPE trx_type AS ENUM ("in", "out")

CREATE TABLE "user" (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(225) NOT NULL,
    email VARCHAR(225) NOT NULL UNIQUE,
    password VARCHAR(225) NOT NULL,
    no_telp VARCHAR(225),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE toko (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    name VARCHAR(225) NOT NULL,
    address VARCHAR(225),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT fk_toko_user FOREIGN KEY (user_id) REFERENCES "user"(id)
);

CREATE TABLE gudang (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    toko_id UUID NOT NULL,
    name VARCHAR(225) NOT NULL,
    address VARCHAR(225),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT fk_gudang_toko FOREIGN KEY (toko_id) REFERENCES toko(id)
);

CREATE TABLE category (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(225) NOT NULL,
    description VARCHAR(225),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE barang (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    gudang_id UUID NOT NULL,
    category_id UUID NOT NULL,
    name VARCHAR(225) NOT NULL,
    sku VARCHAR(225) NOT NULL,
    stock INT DEFAULT 0,
    image_url VARCHAR(225) NOT NULL,
    need_restock BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT fk_barang_gudang FOREIGN KEY (gudang_id) REFERENCES gudang(id),
    CONSTRAINT fk_barang_category FOREIGN KEY (category_id) REFERENCES category(id)
);

CREATE TABLE trx_log (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    barang_id UUID NOT NULL,
    qty INT NOT NULL,
    type trx_type NOT NULL, 
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT fk_trx_barang FOREIGN KEY (barang_id) REFERENCES barang(id)
);

CREATE TABLE smart_log (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    barang_id UUID NOT NULL,
    gudang_id UUID NOT NULL,
    category_id UUID NOT NULL,
    eoq_calculation_result INT,
    decision VARCHAR(225)
    ai_insight VARCHAR(225),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT fk_smart_barang FOREIGN KEY (barang_id) REFERENCES barang(id),
    CONSTRAINT fk_smart_gudang FOREIGN KEY (gudang_id) REFERENCES gudang(id),
    CONSTRAINT fk_smart_category FOREIGN KEY (category_id) REFERENCES category(id)
);

CREATE INDEX idx_barang_sku ON barang(sku);
CREATE INDEX idx_user_email ON "user"(email);