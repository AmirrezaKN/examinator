CREATE SCHEMA IF NOT EXISTS examinator;
DO $$ BEGIN CREATE TYPE "examinator"."question_type" AS ENUM (
  'essay',
  'multiple_choice',
  'true_false',
  'short_answer'
);
EXCEPTION
WHEN duplicate_object THEN null;
END $$;
CREATE TABLE IF NOT EXISTS "examinator"."question" (
  "id" SERIAL PRIMARY KEY,
  "title" text,
  "major" text,
  "type" question_type,
  "description" text
);
CREATE TABLE IF NOT EXISTS "examinator"."choice" (
  "id" SERIAL PRIMARY KEY,
  "text" text,
  "is_answer" boolean
);
CREATE TABLE IF NOT EXISTS "examinator"."question_choice" (
  "question_id" bigserial,
  "choice_id" bigserial,
  PRIMARY KEY ("question_id", "choice_id")
);
CREATE TABLE IF NOT EXISTS "examinator"."exam" (
  "id" SERIAL PRIMARY KEY,
  "title" text,
  "date" datetime,
  "major" text
);
CREATE TABLE IF NOT EXISTS "examinator"."exam_question" (
  "exam_id" bigserial,
  "question_id" bigserial,
  PRIMARY KEY ("exam_id", "question_id")
);
ALTER TABLE "examinator"."question_choice"
ADD FOREIGN KEY ("question_id") REFERENCES "examinator"."question" ("id");
ALTER TABLE "examinator"."question_choice"
ADD FOREIGN KEY ("choice_id") REFERENCES "examinator"."choice" ("id");
ALTER TABLE "examinator"."exam_question"
ADD FOREIGN KEY ("exam_id") REFERENCES "examinator"."exam" ("id");
ALTER TABLE "examinator"."exam_question"
ADD FOREIGN KEY ("question_id") REFERENCES "examinator"."question" ("id");