-- create table Games
CREATE TABLE Games (
    GameId INT PRIMARY KEY AUTO_INCREMENT,
    Name VARCHAR(255),
    Description VARCHAR(255),
    Genre VARCHAR(255),
    SalePrice DECIMAL(10, 2),
    RentalPrice DECIMAL(10, 2),
    Studio VARCHAR(255),
    Stock INT
);

-- create table Users
CREATE TABLE Users (
    UserId INT PRIMARY KEY AUTO_INCREMENT,
    Name VARCHAR(255),
    Role VARCHAR(255),
    Email VARCHAR(255) UNIQUE,
    PhoneNumber VARCHAR(255),
    PasswordHash VARCHAR(255)
);

-- create table Sales
CREATE TABLE Sales (
    SaleId INT PRIMARY KEY AUTO_INCREMENT,
    GameId INT,
    UserId INT,
    SaleDate DATE,
    PurchasedPrice DECIMAL(10, 2),
    Quantity INT,
    FOREIGN KEY (GameId) REFERENCES Games(GameId),
    FOREIGN KEY (UserId) REFERENCES Users(UserId)
);

-- create table Reviews
CREATE TABLE Reviews (
    ReviewId INT PRIMARY KEY AUTO_INCREMENT,
    UserId INT,
    GameId INT,
    Rating DECIMAL(3, 2),
    ReviewMsg VARCHAR(255),
    FOREIGN KEY (UserId) REFERENCES Users(UserId),
    FOREIGN KEY (GameId) REFERENCES Games(GameId),
    CHECK (Rating <= 5)
);

-- create table Rentals
CREATE TABLE Rentals (
    RentalId INT PRIMARY KEY AUTO_INCREMENT,
    UserId INT,
    GameId INT,
    StartDate DATE,
    EndDate DATE,
    Status VARCHAR(255),
    FOREIGN KEY (UserId) REFERENCES Users(UserId),
    FOREIGN KEY (GameId) REFERENCES Games(GameId)
);

