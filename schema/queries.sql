-- name: GetPageViews :one
SELECT page_views FROM stats LIMIT 1;

-- name: IncrementPageViews :exec
UPDATE stats SET page_views = page_views + 1 WHERE TRUE;
