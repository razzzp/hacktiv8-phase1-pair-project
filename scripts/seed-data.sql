-- Insert mock data into Games
INSERT INTO Games (Name, Description, Genre, SalePrice, RentalPrice, Studio, Stock)
VALUES 
('The Legend of Zelda: Breath of the Wild', 'An action-adventure game set in a vast open world.', 'Action-Adventure', 59.99, 4.99, 'Nintendo', 1),
('Minecraft', 'A sandbox game where players can build and explore infinite worlds.', 'Sandbox', 26.95, 2.99, 'Mojang Studios', 3),
('The Witcher 3: Wild Hunt', 'An RPG where players control Geralt of Rivia, a monster hunter.', 'RPG', 39.99, 3.99, 'CD Projekt Red', 4),
('SimCity', 'A city-building simulation game where players develop and manage a city.', 'Simulation', 29.99, 2.49, 'Maxis', 2),
('Civilization VI', 'A turn-based strategy game where players build and expand their own empire.', 'Strategy', 49.99, 3.49, 'Firaxis Games', 3),
('FIFA 21', 'A football simulation game featuring real-world players and teams.', 'Sports', 59.99, 4.99, 'EA Sports', 5),
('Tetris', 'A puzzle game where players arrange falling blocks to clear lines.', 'Puzzle', 9.99, 0.99, 'Alexey Pajitnov', 4),
('Resident Evil Village', 'A horror game where players explore a village filled with terrifying creatures.', 'Horror', 59.99, 5.99, 'Capcom', 1),
('Mario Kart 8 Deluxe', 'A racing game featuring characters from the Mario franchise.', 'Racing', 59.99, 4.99, 'Nintendo', 5),
('Street Fighter V', 'A fighting game where players control characters in one-on-one combat.', 'Fighting', 39.99, 3.49, 'Capcom', 3);


-- Insert mock data into Users
INSERT INTO Users (Name, Role, Email, PhoneNumber, PasswordHash)
VALUES 
('Alice Johnson', 'Admin', 'alice.johnson@example.com', '555-123-4567', SHA2(CONCAT('randomSalt1', '12345'), 256)),
('Bob Smith', 'User', 'bob.smith@example.com', '555-234-5678',  SHA2(CONCAT('randomSalt2', '12345'), 256)),
('Charlie Brown', 'User', 'charlie.brown@example.com', '555-345-6789',  SHA2(CONCAT('randomSalt3', '12345'), 256)),
('David Wilson', 'User', 'david.wilson@example.com', '555-456-7890',  SHA2(CONCAT('randomSalt4', '12345'), 256)),
('Eve Davis', 'User', 'eve.davis@example.com', '555-567-8901',  SHA2(CONCAT('randomSalt5', '12345'), 256)),
('Frank Miller', 'User', 'frank.miller@example.com', '555-678-9012', SHA2(CONCAT('randomSalt6', '12345'), 256)),
('Grace Lee', 'User', 'grace.lee@example.com', '555-789-0123',  SHA2(CONCAT('randomSalt7', '12345'), 256)),
('Hannah White', 'User', 'hannah.white@example.com', '555-890-1234', SHA2(CONCAT('randomSalt8', '12345'), 256)),
('Ivan Green', 'User', 'ivan.green@example.com', '555-901-2345',  SHA2(CONCAT('randomSalt9', '12345'), 256)),
('Julia Adams', 'Admin', 'julia.adams@example.com', '555-012-3456',  SHA2(CONCAT('randomSalt10', '12345'), 256));


-- Insert mock data into Sales
INSERT INTO Sales (GameId, UserId, SaleDate, PurchasedPrice)
VALUES 
(1, 1, '2024-01-01', 59.99),
(2, 2, '2024-02-01', 49.99),
(3, 3, '2024-03-01', 39.99),
(4, 4, '2024-04-01', 29.99),
(5, 5, '2024-05-01', 19.99),
(6, 6, '2024-06-01', 69.99),
(7, 7, '2024-07-01', 9.99),
(8, 8, '2024-08-01', 79.99),
(9, 9, '2024-09-01', 49.99),
(10, 10, '2024-10-01', 39.99);

-- Insert mock data into Reviews
INSERT INTO Reviews (UserId, GameId, Rating, ReviewMsg)
VALUES 
(1, 1, 4.5, 'Great game!'),
(2, 2, 4.0, 'Really enjoyed it.'),
(3, 3, 3.5, 'It was okay.'),
(4, 4, 5.0, 'Amazing!'),
(5, 5, 2.0, 'Not my type.'),
(6, 6, 4.8, 'Highly recommend!'),
(7, 7, 3.8, 'Good game.'),
(8, 8, 2.5, 'Scary but fun.'),
(9, 9, 4.2, 'Fast and fun.'),
(10, 10, 3.0, 'Decent.');

-- Insert mock data into Rentals
INSERT INTO Rentals (UserId, GameId, StartDate, EndDate, Status)
VALUES 
(1, 1, '2024-01-01', '2024-01-10', 'Returned'),
(2, 2, '2024-02-01', '2024-02-10', 'Returned'),
(3, 3, '2024-03-01', '2024-03-10', 'Returned'),
(4, 4, '2024-04-01', '2024-04-10', 'Returned'),
(5, 5, '2024-05-01', '2024-05-10', 'Returned'),
(6, 6, '2024-06-01', '2024-06-10', 'Returned'),
(7, 7, '2024-07-01', '2024-07-10', 'Returned'),
(8, 8, '2024-08-01', '2024-08-10', 'Returned'),
(9, 9, '2024-09-01', '2024-09-10', 'Returned'),
(10, 10, '2024-10-01', '2024-10-10', 'Returned');
