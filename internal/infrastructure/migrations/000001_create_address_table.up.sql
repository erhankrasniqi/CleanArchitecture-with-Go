CREATE TABLE Address (
    id INT IDENTITY(1,1) PRIMARY KEY,
    street NVARCHAR(255) NOT NULL,
    city NVARCHAR(100) NOT NULL,
    state NVARCHAR(100),
    country NVARCHAR(100) NOT NULL,
    user_id INT NOT NULL,
    created_at DATETIME DEFAULT GETDATE(),
    updated_at DATETIME DEFAULT GETDATE(),

    CONSTRAINT fk_address_user FOREIGN KEY (user_id) REFERENCES Users(id) ON DELETE CASCADE
);
