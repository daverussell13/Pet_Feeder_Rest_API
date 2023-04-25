CREATE TABLE "feeding_schedules" (
  "id" SERIAL PRIMARY KEY,
  "device_id" UUID NOT NULL,
  "schedule_id" INTEGER NOT NULL,
  "feed_amount" INTEGER NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP DEFAULT NULL
);

ALTER TABLE "feeding_schedules" ADD CONSTRAINT "fk_device_id" FOREIGN KEY ("device_id") REFERENCES "devices" ("id");
ALTER TABLE "feeding_schedules" ADD CONSTRAINT "fk_schedule_id" FOREIGN KEY ("schedule_id") REFERENCES "schedules" ("id");
