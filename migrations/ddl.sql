CREATE DATABASE scout_store_db;

USE scout_store_db;

CREATE TABLE Products
(
    ProductID   INT PRIMARY KEY AUTO_INCREMENT,
    ProductName VARCHAR(255),
    Price       DECIMAL(10, 2),
    Stock       INT
);

CREATE TABLE Staff
(
    StaffID   INT PRIMARY KEY AUTO_INCREMENT,
    StaffName VARCHAR(255),
    RoleID    INT NOT NULL,
    FOREIGN KEY (RoleID) REFERENCES Roles (RoleID)
);

CREATE TABLE Roles
(
    RoleID   INT PRIMARY KEY AUTO_INCREMENT,
    RoleName VARCHAR(255)
);

CREATE TABLE Sales
(
    SaleID     INT PRIMARY KEY AUTO_INCREMENT,
    Date        DATE,
    ProductID   INT NOT NULL,
    Quantity    INT,
    TotalAmount DECIMAL(10, 2),
    StaffID     INT NOT NULL,
    FOREIGN KEY (ProductID) REFERENCES Products (ProductID),
    FOREIGN KEY (StaffID) REFERENCES Staff (StaffID)
);
