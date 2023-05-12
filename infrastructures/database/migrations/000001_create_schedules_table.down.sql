ALTER TABLE "schedules"
    DROP CONSTRAINT IF EXISTS "unique_day_time";

DROP TABLE IF EXISTS "schedules";
