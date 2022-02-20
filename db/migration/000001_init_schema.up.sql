CREATE TABLE "users" (
  "id" bigint PRIMARY KEY AUTO_INCREMENT,
  "username" varchar(255),
  "email" varchar(255),
  "password" varchar(255)
);

CREATE TABLE "members" (
  "id" bigint PRIMARY KEY,
  "user_id" bigint,
  "room_id" bigint
);

CREATE TABLE "rooms" (
  "id" bigint PRIMARY KEY
);

CREATE TABLE "messsages" (
  "id" bigint PRIMARY KEY,
  "room_id" bigint,
  "sent" timestamptz,
  "sender_id" bigint,
  "seen_user_id" bigint,
  "messsages" varchar(255)
);
