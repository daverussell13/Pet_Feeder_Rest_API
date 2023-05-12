ALTER TABLE "devices" DROP CONSTRAINT IF EXISTS "unique_device_name";

DROP TABLE IF EXISTS "devices";

DROP EXTENSION IF EXISTS "uuid-ossp";