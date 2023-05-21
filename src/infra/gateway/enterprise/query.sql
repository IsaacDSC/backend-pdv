-- name: CreateEnterprise :exec
INSERT INTO "enterprise" (
        id,
        corporate_name,
        fantasy_name,
        logo,
        cnpj,
        token,
        updated_at
    )
VALUES ($1, $2, $3, $4, $5, $6, Now());