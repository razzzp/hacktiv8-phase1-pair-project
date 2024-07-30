## ROC Gameshop - Entity Analysis

## Entities

### Game

Attributes:

- GameId (PK AI) INT
- Name VARCHAR
- Description VARCHAR  
- Genre VARCHAR
- SalePrice (to Buy) DECIMAL
- RentalPrice (per Day) DECIMAL
- Studio VARCHAR
- Stock INT

### User

Attributes:

- UserId (PK AI) INT
- Name VARCHAR
- Role VARCHAR
- Email VARCHAR
- PhoneNumber VARCHAR
- Salt VARCHAR
- PasswordHash VARCHAR

### Sale

Attributes:

- SaleId (PK AI) INT
- GameId (FK) INT
- UserId (FK) INT
- SaleDate DATE
- PurchasedPrice DECIMAL

### Review

Attributes:

- ReviewId (PK AI) INT
- UserId (FK) INT
- GameId (FK) INT
- Rating DECIMAL
- ReviewMsg VARCHAR

### Rental

Attributes:

- RentalId (PK AI) INT
- UserId (FK) INT
- GameId (FK) INT
- StartDate DATE
- EndDate DATE
- Status VARCHAR

## Relationships 


- `Games`- `Users` : Many to Many

- `Sales`, `Reviews`, and `Rentals` are the junction tables for MtoM relations between `Games` - `Users` .
 
 - One `game` can be bought/rented by many `users`, and one `user` can buy/rent many `games`
    