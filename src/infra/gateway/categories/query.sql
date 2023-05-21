--
--
-- name: CreateCategory :exec
INSERT INTO "categories" (
        "name",
        "description",
        "enterprise_id",
        "updated_at"
    )
VALUES($1, $2, $3, Now());