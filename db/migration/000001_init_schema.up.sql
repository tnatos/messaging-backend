CREATE TABLE IF NOT EXISTS "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "email" varchar NOT NULL,
  "password" varchar NOT NULL
);

CREATE TABLE "members" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "room_id" bigint NOT NULL
);

CREATE TABLE "rooms" (
  "id" bigserial PRIMARY KEY
);

CREATE TABLE "messsages" (
  "id" bigserial PRIMARY KEY,
  "room_id" bigint NOT NULL,
  "sent" timestamptz NOT NULL DEFAULT (now()),
  "sender_id" bigint NOT NULL,
  "seen_user_id" bigint NOT NULL,
  "messsages" varchar NOT NULL
);

AlTER TABLE "members" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

AlTER TABLE "members" ADD FOREIGN KEY ("room_id") REFERENCES "rooms" ("id");

AlTER TABLE "messsages" ADD FOREIGN KEY ("room_id") REFERENCES "rooms" ("id");

AlTER TABLE "messsages" ADD FOREIGN KEY ("sender_id") REFERENCES "users" ("id");

AlTER TABLE "messsages" ADD FOREIGN KEY ("seen_user_id") REFERENCES "users" ("id");
