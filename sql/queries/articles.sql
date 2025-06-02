-- name: CreateArticle :one
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

-- name: GetArticles :many
SELECT * FROM articles WHERE title ILIKE $1;

-- name: GetArtidleByID :one
SELECT * FROM articles WHERE id = $1; 

-- name: UpdateArticle :one
UPDATE articles SET
title= $1,
content=$2,
updated_at = NOW()
WHERE id = $3
RETURNING *;

-- name: DeleteArticle :exec
DELETE FROM articles WHERE id = $1;