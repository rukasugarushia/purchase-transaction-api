create table purchase_transactions(
    id serial primary key,
    description varchar(50) NOT NULL,
    transaction_date TIMESTAMP NOT NULL,
    purchase_amount NUMERIC(10, 2) NOT NULL
);

INSERT INTO purchase_transactions (description, transaction_date, purchase_amount)
VALUES
    ('Sample Purchase 1', '2023-10-26', 199.99),
    ('Sample Purchase 2', '2023-10-25', 299.50),
    ('Sample Purchase 3', '2023-10-24', 99.99),
    ('Sample Purchase 4', '2023-10-23', 149.25);