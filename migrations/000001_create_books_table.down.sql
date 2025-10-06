-- Drop books table and related functions
DROP FUNCTION IF EXISTS fn_books_delete(UUID);
DROP FUNCTION IF EXISTS fn_books_update(UUID, VARCHAR(255), DECIMAL(10,2));
DROP FUNCTION IF EXISTS fn_books_get_by_id(UUID);
DROP FUNCTION IF EXISTS sp_books_create(UUID, VARCHAR(255), DECIMAL(10,2));
DROP TABLE IF EXISTS books;
