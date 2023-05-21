--
--
-- name: CreateCategory :exec
INSERT INTO "categories" (
        "id",
        "name",
        "description",
        "enterprise_id",
        "updated_at"
    )
VALUES($1, $2, $3, $4, Now());