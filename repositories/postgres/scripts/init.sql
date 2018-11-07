CREATE TABLE IF NOT EXISTS profile(
  "id" BIGSERIAL
    CONSTRAINT "profile_id_primary_key" PRIMARY KEY,

  "username" VARCHAR(20)
    CONSTRAINT "profile_username_not_null" NOT NULL
    CONSTRAINT "profile_username_unique" UNIQUE
    CONSTRAINT "profile_username_check" CHECK("username" ~ '^\w+$'),

  "password_hash" VARCHAR(256)
    CONSTRAINT "profile_password_hash_not_null" NOT NULL,

  "password_salt" VARCHAR(256)
    CONSTRAINT "profile_password_salt_not_null" NOT NULL,

  "email" VARCHAR(100)
    CONSTRAINT "profile_email_not_null" NOT NULL
    CONSTRAINT "profile_email_unique" UNIQUE
    CONSTRAINT "profile_email_check" CHECK(email ~ '^.+@.+$'),

  "score" INTEGER
    DEFAULT 0
    CONSTRAINT "profile_score_not_null" NOT NULL
    CONSTRAINT "profile_score_check" CHECK("score" >= 0)
);

CREATE OR REPLACE FUNCTION create_profile(
  _username_ TEXT, _password_hash_ TEXT, _password_salt_ TEXT, _email_ TEXT)
RETURNS VOID
AS $$
BEGIN
  INSERT INTO "profile"("username", "password_hash", "password_salt", "email")
  VALUES(_username_, _password_hash_, _password_salt_, _email_);
END;
$$
LANGUAGE plpgsql;

CREATE TYPE "profile_info" AS (
  "id" BIGINT, "username" TEXT, "email" TEXT, "score" INTEGER
);

CREATE OR REPLACE FUNCTION get_profile(_profile_id_ BIGINT)
RETURNS "profile_info"
AS $$
DECLARE _result_ "profile_info";
BEGIN
  SELECT p."id", p."username", p."email", p."score"
  FROM "profile" p
  WHERE "id" = _profile_id_
  INTO _result_;

  IF _result_ IS NULL THEN
    RAISE 'profile not found';
  END IF;

  RETURN _result_;
END;
$$
LANGUAGE plpgsql;

CREATE TYPE "id_password_info" AS (
  "id" BIGINT, "password_hash" TEXT, "PASSWORD_SALT" TEXT
);

CREATE OR REPLACE FUNCTION get_profile(_username_ TEXT)
  RETURNS "id_password_info"
AS $$
DECLARE _result_ "id_password_info";
BEGIN
  SELECT p."id", p."password_hash", p."password_salt"
  FROM "profile" p
  WHERE "username" = _username_
  INTO _result_;

  IF _result_ IS NULL THEN
    RAISE 'profile not found';
  END IF;

  RETURN _result_;
END;
$$
LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION update_profile(_profile_id_ BIGINT,
  _username_ TEXT, _password_hash_ TEXT, _password_salt_ TEXT, _email_ TEXT)
RETURNS VOID
AS $$
DECLARE _profile_ "profile";
BEGIN
  SELECT * FROM "profile" WHERE "id" = _profile_id_
  INTO _profile_;

  IF _profile_ IS NULL THEN
    RAISE 'profile not found';
  END IF;

  IF _username_ != '' THEN
    _profile_."username" := _username_;
  END IF;
  IF _password_hash_ != '' THEN
    _profile_."password_hash" := _password_hash_;
  END IF;
  IF _password_salt_ != '' THEN
    _profile_."password_hash" := _password_hash_;
  END IF;
  IF _email_ != '' THEN
    _profile_."email" := _email_;
  END IF;

  UPDATE "profile" SET
    "username" = _profile_."username",
    "password_hash" = _profile_."password_hash",
    "password_salt" = _profile_."password_salt",
    "email" = _profile_."email"
  WHERE "id" = _profile_id_;
END;
$$
LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION delete_profile(_profile_id_ BIGINT)
RETURNS VOID
AS $$
DECLARE _profile_ "profile";
BEGIN
  DELETE FROM "profile" WHERE "id" = _profile_id_
  RETURNING * INTO _profile_;
  IF _profile_ IS NULL THEN
    RAISE 'profile not found';
  END IF;
END;
$$
LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION get_all_profiles(
  _page_index_ INTEGER, _page_size_ INTEGER)
RETURNS SETOF "profile_info"
AS $$
  SELECT p."id", p."username", p."email", p."score"
  FROM "profile" p
  ORDER BY p."score", p."username"
  LIMIT _page_size_ OFFSET _page_index_;
$$
LANGUAGE sql;

CREATE OR REPLACE FUNCTION update_profile_score(
  _profile_id_ BIGINT, _score_diff_ INTEGER)
RETURNS VOID
AS $$
DECLARE _score_ INTEGER;
BEGIN
  SELECT "score" FROM "profile" WHERE "id" = _profile_id_
  INTO _score_;

  IF _score_ IS NULL THEN
    RAISE 'profile not found';
  END IF;

  _score_ := _score_ + _score_diff_;
  IF _score_ < 0 THEN
    _score_ = 0;
  END IF;

  UPDATE "profile" SET
    "score" = _score_
  WHERE "id" = _profile_id_;
END;
$$
LANGUAGE plpgsql;
