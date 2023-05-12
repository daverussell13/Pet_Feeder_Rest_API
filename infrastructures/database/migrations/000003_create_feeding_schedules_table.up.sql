CREATE TABLE "feeding_schedules" (
  "id" SERIAL PRIMARY KEY,
  "device_id" UUID NOT NULL,
  "schedule_id" INTEGER NOT NULL,
  "feed_amount" INTEGER NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT (CURRENT_TIMESTAMP AT TIME ZONE 'Asia/Jakarta')
);

ALTER TABLE "feeding_schedules" ADD CONSTRAINT "fk_device_id" FOREIGN KEY ("device_id") REFERENCES "devices" ("id");
ALTER TABLE "feeding_schedules" ADD CONSTRAINT "fk_schedule_id" FOREIGN KEY ("schedule_id") REFERENCES "schedules" ("id");

ALTER TABLE "feeding_schedules" ADD CONSTRAINT "unique_device_id_schedule_id" UNIQUE ("device_id", "schedule_id");

CREATE INDEX "idx_device_id_schedule_id" ON "feeding_schedules" ("device_id", "schedule_id");
