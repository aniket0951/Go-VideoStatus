CREATE TABLE "verify_videos" (
  "id" uuid DEFAULT uuid_generate_v4 () PRIMARY KEY,
  "video_id" uuid UNIQUE NOT NULL,
  "verify_by" uuid NOT NULL,
  "status" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "verify_videos" ADD FOREIGN KEY ("video_id") REFERENCES "video_by_admin" ("id");

ALTER TABLE "verify_videos" ADD FOREIGN KEY ("verify_by") REFERENCES "users" ("id");
