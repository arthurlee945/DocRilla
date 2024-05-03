-- CreateEnum
CREATE TYPE "user_role" AS ENUM ('ADMIN', 'MODERATOR', 'USER', 'MOCK');

-- CreateEnum
CREATE TYPE "project_type" AS ENUM ('TEXT', 'IMAGE', 'NUMBER');

-- CreateTable
CREATE TABLE "account" (
    "id" SERIAL NOT NULL,
    "user_id" INTEGER NOT NULL,
    "type" TEXT NOT NULL,
    "provider" TEXT NOT NULL,
    "provider_account_id" TEXT,
    "refresh_token" TEXT,
    "access_token" TEXT,
    "expires_at" INTEGER,
    "token_type" TEXT,
    "scope" TEXT,
    "id_token" TEXT,
    "session_state" TEXT,

    CONSTRAINT "account_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "usr" (
    "id" SERIAL NOT NULL,
    "name" TEXT,
    "email" TEXT NOT NULL,
    "email_verified" BOOLEAN DEFAULT false,
    "email_verification_token" TEXT,
    "password" TEXT,
    "role" "user_role" NOT NULL DEFAULT 'USER',
    "password_changed_at" TIMESTAMP(3),
    "reset_password_token" TEXT,
    "reset_password_expires" TIMESTAMP(3),
    "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "active" BOOLEAN NOT NULL DEFAULT true,

    CONSTRAINT "usr_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "session" (
    "id" SERIAL NOT NULL,
    "session_token" TEXT NOT NULL,
    "user_id" INTEGER NOT NULL,
    "expires" TIMESTAMP(3) NOT NULL,

    CONSTRAINT "session_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "verification_token" (
    "identifier" TEXT NOT NULL,
    "token" TEXT NOT NULL,
    "expires" TIMESTAMP(3) NOT NULL
);

-- CreateTable
CREATE TABLE "project" (
    "id" SERIAL NOT NULL,
    "user_id" INTEGER NOT NULL,
    "uuid" TEXT NOT NULL DEFAULT gen_random_uuid(),
    "route" TEXT NOT NULL,
    "token" TEXT,
    "title" TEXT NOT NULL,
    "description" TEXT,
    "document_url" TEXT NOT NULL,
    "archived" BOOLEAN NOT NULL DEFAULT false,
    "visited_at" TIMESTAMP(3),
    "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT "project_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "field" (
    "id" SERIAL NOT NULL,
    "uuid" TEXT NOT NULL DEFAULT gen_random_uuid(),
    "project_id" TEXT NOT NULL,
    "x" DOUBLE PRECISION NOT NULL,
    "y" DOUBLE PRECISION NOT NULL,
    "width" DOUBLE PRECISION NOT NULL,
    "height" DOUBLE PRECISION NOT NULL,
    "page" INTEGER NOT NULL,
    "type" "project_type" NOT NULL,

    CONSTRAINT "field_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "submission" (
    "id" SERIAL NOT NULL,
    "uuid" TEXT NOT NULL DEFAULT gen_random_uuid(),
    "project_id" TEXT NOT NULL,
    "submitted_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT "submission_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "submitted_field" (
    "id" SERIAL NOT NULL,
    "field_id" TEXT NOT NULL,
    "submission_id" TEXT NOT NULL,
    "value" BYTEA NOT NULL,

    CONSTRAINT "submitted_field_pkey" PRIMARY KEY ("id")
);

-- CreateIndex
CREATE UNIQUE INDEX "account_provider_provider_account_id_key" ON "account"("provider", "provider_account_id");

-- CreateIndex
CREATE UNIQUE INDEX "usr_email_key" ON "usr"("email");

-- CreateIndex
CREATE UNIQUE INDEX "session_session_token_key" ON "session"("session_token");

-- CreateIndex
CREATE UNIQUE INDEX "verification_token_token_key" ON "verification_token"("token");

-- CreateIndex
CREATE UNIQUE INDEX "verification_token_identifier_token_key" ON "verification_token"("identifier", "token");

-- CreateIndex
CREATE UNIQUE INDEX "project_uuid_key" ON "project"("uuid");

-- CreateIndex
CREATE UNIQUE INDEX "project_route_key" ON "project"("route");

-- CreateIndex
CREATE INDEX "project_uuid_idx" ON "project"("uuid");

-- CreateIndex
CREATE UNIQUE INDEX "field_uuid_key" ON "field"("uuid");

-- CreateIndex
CREATE INDEX "field_uuid_idx" ON "field"("uuid");

-- CreateIndex
CREATE UNIQUE INDEX "submission_uuid_key" ON "submission"("uuid");

-- CreateIndex
CREATE INDEX "submission_uuid_idx" ON "submission"("uuid");

-- AddForeignKey
ALTER TABLE "account" ADD CONSTRAINT "account_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "usr"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "session" ADD CONSTRAINT "session_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "usr"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "project" ADD CONSTRAINT "project_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "usr"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "field" ADD CONSTRAINT "field_project_id_fkey" FOREIGN KEY ("project_id") REFERENCES "project"("uuid") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "submission" ADD CONSTRAINT "submission_project_id_fkey" FOREIGN KEY ("project_id") REFERENCES "project"("uuid") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "submitted_field" ADD CONSTRAINT "submitted_field_submission_id_fkey" FOREIGN KEY ("submission_id") REFERENCES "submission"("uuid") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "submitted_field" ADD CONSTRAINT "submitted_field_field_id_fkey" FOREIGN KEY ("field_id") REFERENCES "field"("uuid") ON DELETE CASCADE ON UPDATE CASCADE;
