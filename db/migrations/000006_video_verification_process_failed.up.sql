CREATE TABLE "video_verification_process_failed" (
  "Id" uuid DEFAULT uuid_generate_v4 () PRIMARY KEY,
  "video_id" uuid UNIQUE NOT NULL,
  "verification_failed_by" uuid,
  "unpublished_by" uuid,
  "status" varchar NOT NULL,
  "reason" varchar NOT NULL,
  "is_verification_failed" bool,
  "is_unpublished" bool,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "video_verification_process_failed" ("video_id");

ALTER TABLE "video_verification_process_failed" ADD FOREIGN KEY ("video_id") REFERENCES "video_by_admin" ("id");

ALTER TABLE "video_verification_process_failed" ADD FOREIGN KEY ("verification_failed_by") REFERENCES "users" ("id");

ALTER TABLE "video_verification_process_failed" ADD FOREIGN KEY ("unpublished_by") REFERENCES "users" ("id");