CREATE TABLE "video_by_admin" (
  "id" uuid DEFAULT uuid_generate_v4 () PRIMARY KEY,
  "title" varchar NOT NULL,
  "file_address" varchar NOT NULL,
  "uploaded_by" uuid NOT NULL,
  "status" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "video_by_admin" ("uploaded_by");
ALTER TABLE "video_by_admin" ADD FOREIGN KEY ("uploaded_by") REFERENCES "users" ("id");
ALTER TABLE video_by_admin 
ADD CONSTRAINT check_file_address CHECK (file_address <> ''),
ADD CONSTRAINT check_title CHECK (title <> '');
