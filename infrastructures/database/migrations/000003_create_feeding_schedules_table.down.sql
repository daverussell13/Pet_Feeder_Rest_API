ALTER TABLE "feeding_schedules" DROP CONSTRAINT IF EXISTS "fk_device_id";
ALTER TABLE "feeding_schedules" DROP CONSTRAINT IF EXISTS "fk_schedule_id";

DROP TABLE IF EXISTS "feeding_schedules";