CREATE TABLE IF NOT EXISTS "enterprise" (
    "id" VARCHAR(100) PRIMARY KEY DEFAULT gen_random_uuid(),
    "corporate_name" VARCHAR(100),
    "fantasy_name" VARCHAR(255),
    "logo" VARCHAR(255),
    "cnpj" VARCHAR(255),
    "status" VARCHAR(60) DEFAULT 'TRIAL',
    "token" TEXT,
    "updated_at" TIMESTAMPTZ,
    "created_at" TIMESTAMPTZ DEFAULT Now()
);
CREATE TABLE IF NOT EXISTS "clients" (
    "id" VARCHAR(100) PRIMARY KEY DEFAULT gen_random_uuid(),
    "name" VARCHAR(100) NOT NULL,
    "email" VARCHAR(250) UNIQUE,
    "password" VARCHAR(255),
    "telephone" VARCHAR(20) NOT NULL UNIQUE,
    "city" VARCHAR(100) NOT NULL,
    "home_number" VARCHAR(250),
    "neighborhood" VARCHAR(100),
    "cep" VARCHAR(100),
    "address" VARCHAR(20),
    "observation" TEXT,
    "updated_at" TIMESTAMPTZ,
    "created_at" TIMESTAMPTZ DEFAULT Now()
);
CREATE TABLE IF NOT EXISTS "categories" (
    "id" VARCHAR(100) PRIMARY KEY DEFAULT gen_random_uuid(),
    "name" VARCHAR(100) NOT NULL,
    "description" TEXT,
    "enterprise_id" VARCHAR(100),
    "updated_at" TIMESTAMPTZ,
    "created_at" TIMESTAMPTZ DEFAULT Now()
);
CREATE TABLE IF NOT EXISTS "products" (
    "id" VARCHAR(100) PRIMARY KEY DEFAULT gen_random_uuid(),
    "category_id" VARCHAR(100) REFERENCES categories(id) NOT NULL,
    "enterprise_id" VARCHAR(100) REFERENCES enterprise(id) NOT NULL,
    "name" VARCHAR(150) NOT NULL,
    "description" TEXT,
    "image" VARCHAR(255),
    "price" REAL,
    "updated_at" TIMESTAMPTZ,
    "created_at" TIMESTAMPTZ DEFAULT Now()
);
CREATE TABLE IF NOT EXISTS "orders" (
    "id" VARCHAR(100) PRIMARY KEY DEFAULT gen_random_uuid(),
    "total_price" REAL,
    "price_delivery" REAL,
    "price_items" REAL,
    "address_id" VARCHAR(100) REFERENCES "address"(id),
    "client_id" VARCHAR(100) REFERENCES "clients"(id),
    "status" VARCHAR(100) DEFAULT 'PENDING',
    "updated_at" TIMESTAMPTZ,
    "created_at" TIMESTAMPTZ DEFAULT Now()
);
CREATE TABLE IF NOT EXISTS "order_items" (
    "id" VARCHAR(100) PRIMARY KEY DEFAULT gen_random_uuid(),
    "product_id" VARCHAR(100) REFERENCES products(id),
    "product_name" VARCHAR(255),
    "product_description" TEXT,
    "product_images" VARCHAR(255),
    "category_name" VARCHAR(100),
    "price_product" REAL,
    "updated_at" TIMESTAMPTZ,
    "created_at" TIMESTAMPTZ DEFAULT Now(),
    "order_id" VARCHAR(100) NOT NULL,
    CONSTRAINT fk_order_items FOREIGN KEY(order_id) REFERENCES "orders"(id)
);