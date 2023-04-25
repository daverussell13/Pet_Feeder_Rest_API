CREATE TABLE "schedules" (
  "id" SERIAL PRIMARY KEY,
  "day_of_week" VARCHAR(10) NOT NULL,
  "feed_time" TIME NOT NULL
);

ALTER TABLE "schedules" ADD CONSTRAINT "unique_day_time" UNIQUE ("day_of_week", "feed_time");