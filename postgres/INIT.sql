CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TYPE trx_type AS ENUM ('in', 'out');

CREATE TABLE "user" (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(225) NOT NULL,
    email VARCHAR(225) NOT NULL,
    password VARCHAR(225) NOT NULL,
    no_telp VARCHAR(225),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE toko (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    name VARCHAR(225) NOT NULL,
    address VARCHAR(225),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_toko_user FOREIGN KEY (user_id) REFERENCES "user"(id) ON DELETE CASCADE
);

CREATE TABLE gudang (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    toko_id UUID NOT NULL,
    name VARCHAR(225) NOT NULL,
    address VARCHAR(225),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_gudang_toko FOREIGN KEY (toko_id) REFERENCES toko(id) ON DELETE CASCADE
);

CREATE TABLE category (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(225) NOT NULL,
    description VARCHAR(225),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE barang (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    gudang_id UUID NOT NULL,
    category_id UUID NOT NULL,
    name VARCHAR(225) NOT NULL,
    sku VARCHAR(225) NOT NULL,
    image_url VARCHAR(225) NOT NULL,
    stock INT DEFAULT 0,
    safety_stock INT DEFAULT 0,
    need_restock BOOLEAN DEFAULT FALSE,
    lead_time_days INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_barang_gudang FOREIGN KEY (gudang_id) REFERENCES gudang(id) ON DELETE CASCADE,
    CONSTRAINT fk_barang_category FOREIGN KEY (category_id) REFERENCES category(id) ON DELETE SET NULL
);

CREATE TABLE trx_log (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    barang_id UUID NOT NULL,
    gudang_id UUID NOT NULL,
    qty INT NOT NULL,
    type trx_type NOT NULL, 
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_trx_barang FOREIGN KEY (barang_id) REFERENCES barang(id) ON DELETE CASCADE,
    CONSTRAINT fk_trx_gudang FOREIGN KEY (gudang_id) REFERENCES gudang(id) ON DELETE CASCADE
);

CREATE TABLE smart_log (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    barang_id UUID NOT NULL,
    gudang_id UUID NOT NULL,
    period_month INT,
    period_year INT,
    eoq_calculation_result INT,
    rop_value INT,
    ai_insight TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_smart_barang FOREIGN KEY (barang_id) REFERENCES barang(id) ON DELETE CASCADE,
    CONSTRAINT fk_smart_gudang FOREIGN KEY (gudang_id) REFERENCES gudang(id)  ON DELETE CASCADE  
);

CREATE INDEX idx_barang_sku ON barang(sku);
CREATE INDEX idx_user_email ON "user"(email);