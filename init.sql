DO $$
BEGIN
    IF NOT EXISTS (SELECT FROM pg_catalog.pg_roles WHERE rolname = 'shop_manager') THEN
        CREATE ROLE shop_manager WITH LOGIN PASSWORD 'MotDePasseCompliqueEtTropLong';
    END IF;
END
$$;

DO $$
BEGIN
    IF NOT EXISTS (SELECT FROM pg_catalog.pg_database WHERE datname = 'shops_management') THEN
        CREATE DATABASE shops_management;
    END IF;
END
$$;

GRANT ALL PRIVILEGES ON DATABASE shops_management TO shop_manager;
