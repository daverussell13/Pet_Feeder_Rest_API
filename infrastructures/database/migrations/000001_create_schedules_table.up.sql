CREATE TABLE "schedules" (
  "id" SERIAL PRIMARY KEY,
  "day_of_week" VARCHAR(10) NOT NULL,
  "feed_time" TIME NOT NULL
);

ALTER TABLE "schedules" ADD CONSTRAINT "unique_day_time" UNIQUE ("day_of_week", "feed_time");

CREATE INDEX "idx_day_of_week_feed_time" ON "schedules" ("day_of_week", "feed_time");