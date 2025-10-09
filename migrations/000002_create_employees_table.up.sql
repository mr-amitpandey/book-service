-- Create books table
CREATE TABLE IF NOT EXISTS books (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    price DECIMAL(10,2) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create stored procedure for creating books
CREATE OR REPLACE FUNCTION sp_books_create(
    p_id UUID,
    p_name VARCHAR(255),
    p_price DECIMAL(10,2)
) RETURNS UUID AS $$
BEGIN
    INSERT INTO books (id, name, price)
    VALUES (p_id, p_name, p_price);
    RETURN p_id;
END;
$$ LANGUAGE plpgsql;

-- Create function for getting book by ID
CREATE OR REPLACE FUNCTION fn_books_get_by_id(p_id UUID)
RETURNS TABLE(id UUID, name VARCHAR(255), price DECIMAL(10,2)) AS $$
BEGIN
    RETURN QUERY
    SELECT b.id, b.name, b.price
    FROM books b
    WHERE b.id = p_id;
END;
$$ LANGUAGE plpgsql;

-- Create function for updating books
CREATE OR REPLACE FUNCTION fn_books_update(
    p_id UUID,
    p_name VARCHAR(255),
    p_price DECIMAL(10,2)
) RETURNS VOID AS $$
BEGIN
    UPDATE books
    SET name = p_name,
        price = p_price,
        updated_at = CURRENT_TIMESTAMP
    WHERE id = p_id;
END;
$$ LANGUAGE plpgsql;

-- Create function for deleting books
CREATE OR REPLACE FUNCTION fn_books_delete(p_id UUID)
RETURNS VOID AS $$
BEGIN
    DELETE FROM books WHERE id = p_id;
END;
$$ LANGUAGE plpgsql;

-- Insert some sample data
INSERT INTO books (id, name, price) VALUES 
    ('00cd2c91-1bbd-46b8-90d3-c37bb6969fec', 'Clean Architecture', 499.99),
    ('ef97e391-2bf9-49c6-a840-b36fcaf42285', 'The Great Gatsby', 29.99),
    ('a4507552-5e03-40ae-85c0-aab5d2f3bf26', 'To Kill a Mockingbird', 19.99)
ON CONFLICT (id) DO NOTHING;
