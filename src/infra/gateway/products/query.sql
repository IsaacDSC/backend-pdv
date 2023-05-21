--
--
-- name: CreateProduct :exec
INSERT INTO "products" (
        "id",
        "category_id",
        "enterprise_id",
        "name",
        "description",
        "image",
        "price",
        "updated_at"
    )
VALUES($1, $2, $3, $4, $5, $6, $7, Now());