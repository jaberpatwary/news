CREATE TABLE users (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(150) UNIQUE NOT NULL,
    phone VARCHAR(20),
    password_hash TEXT NOT NULL, 
    avatar_url TEXT,
    
    status VARCHAR(20) NOT NULL 
        DEFAULT 'active'
        CHECK (status IN ('active', 'inactive', 'banned')),
    
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);
