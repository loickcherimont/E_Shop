DROP DATABASE IF EXISTS e_shop;
CREATE DATABASE e_shop;
USE e_shop;

DROP TABLE IF EXISTS products;
CREATE TABLE products (
    id INT AUTO_INCREMENT NOT NULL,
    title VARCHAR(128) NOT NULL,
    img VARBINARY(128) NOT NULL,
    price DECIMAL(5,2) NOT NULL,
    quantity INT,
    PRIMARY KEY (`id`)
);

INSERT INTO products
    (title, img, price, quantity)
VALUES
    ("Laptop", "laptop.jpg", 500.99, 5),
    ("Smartphone", "smartphone.jpg", 300.50, 10),
    ("Tablet", "tablet.jpg", 100.50, 3);