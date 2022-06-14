CREATE TABLE wallets
(
    id SERIAL PRIMARY KEY,
    amount NUMERIC(15,2)
);

INSERT INTO wallets
VALUES
    (1, 100),
    (2, 9999999999999.99),
    (3, 300),
    (4, 400),
    (5, 500),
    (6, 1);

DELETE FROM wallets WHERE id = 6;