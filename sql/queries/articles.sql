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