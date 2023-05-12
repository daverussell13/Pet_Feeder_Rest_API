CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "devices" (
  "id" UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
  "device_name" VARCHAR(255) NOT NULL,
  "device_type" VARCHAR(255) NOT NULL
);

ALTER TABLE "devices" ADD CONSTRAINT "unique_device_name" UNIQUE ("device_name");