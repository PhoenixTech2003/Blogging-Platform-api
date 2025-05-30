-- name: createArticle :one
INSERT INTO articles (id, title, content, created_at, updated_at)
VALUES
(
    gen_random_uuid(),
    $1,
    $2,
    NOW(),
    NOW()
)

RETURNING *;