-- Basic database setup for Book Service
-- This migration sets up the basic database structure

-- Create database extensions if they don't exist
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create a simple users table for basic authentication (if needed)
-- This is kept minimal for the book service demo
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Insert a default admin user for testing (password: admin123)
INSERT INTO users (id, username, email, password_hash) VALUES 
    ('2c1dd240-e6b1-4b1c-9f03-b96fb81def94', 'admin', 'admin@bookservice.com', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi')
ON CONFLICT (username) DO NOTHING;