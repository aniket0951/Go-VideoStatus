CREATE TABLE "published_videos" (
  "id" uuid DEFAULT uuid_generate_v4 () PRIMARY KEY,
  "video_id" uuid UNIQUE NOT NULL,
  "published_by" uuid NOT NULL,
  "status" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "published_videos" ADD FOREIGN KEY ("video_id") REFERENCES "video_by_admin" ("id");

ALTER TABLE "published_videos" ADD FOREIGN KEY ("published_by") REFERENCES "users" ("id");