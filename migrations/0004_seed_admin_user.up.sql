INSERT INTO users (name, email, password_hash, role_id)
VALUES (
    'System Admin',
    'admin@movie.com',
    '$2a$10$REPLACE_WITH_YOUR_HASH',
    (SELECT id FROM roles WHERE name = 'ADMIN')
);
