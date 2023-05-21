--
--
--
-- name: FindClient :one
SELECT *
FROM "clients"
WHERE telephone = $1;
--
--
--
-- name: CreateCLient :exec
INSERT INTO "clients" (
        "name",
        "email",
        "password",
        "telephone",
        "home_number",
        city,
        neighborhood,
        cep,
        "address",
        observation,
        updated_at
    )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, Now());
--
--
--
-- name: UpdateCLient :exec
UPDATE "clients"
SET "name" = $1
WHERE telephone = $2;