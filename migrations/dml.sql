USE scout_store_db;

INSERT INTO Products (ProductName, Price, Stock)
VALUES ('Topi Baret', 20000, 50),
       ('Dasi Kacu', 5000, 60),
       ('Pin Pramuka', 5000, 40),
       ('Sabuk Pramuka', 15000, 30);

INSERT INTO Roles (RoleName)
VALUES ('Admin'),
       ('Kasir');

INSERT INTO Staff (StaffName, RoleID)
VALUES ('John', 1),
       ('Doe', 2),
       ('Steve', 2);

INSERT INTO Sales (Date, ProductID, Quantity, TotalAmount, StaffID)
VALUES ('2024-07-01', 1, 5, 100000, 2),
       ('2024-07-02', 2, 10, 50000, 2),
       ('2024-07-03', 3, 15, 150000, 3);