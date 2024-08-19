-- Some queries that we can use to extend service functionalities

-- name: HasRoles :one
SELECT id, phone,
  CASE
    WHEN '{"DBA"}' <@ users.roles THEN 'Yes'
    ELSE 'No'
  END 
  AS has_roles
FROM users WHERE id = $1;

-- name: CreateUserOrUpdateIfExists :one
INSERT INTO users (name, password, roles) 
VALUES ($1, $3, $4::roles[])
ON CONFLICT("name") 
DO UPDATE SET 
    roles = (select array_agg(distinct e) from unnest(array_append(users.roles, $5::roles)) e)
RETURNING *;