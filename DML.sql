INSERT INTO Customer (customer_name, phone_number)
VALUES 
    ('angga', '081234567891'),
    ('bang vicko', '089876543212');


INSERT INTO Transaksi (entry_date, completion_date, received_by, customer_id)
VALUES 
    ('2024-02-29', '2024-03-05', 'Inne', 1),
    ('2024-02-28', '2024-03-03', 'Cabi', 2);


INSERT INTO Detail_Transaksi (transaction_id, service_name, quantity, unit, price, total)
VALUES 
    (1, 'Laundry', 2, 'kg', 5000.00, 10000.00),
    (1, 'Dry Cleaning', 1, 'piece', 3000.00, 3000.00),
    (2, 'Ironing', 3, 'item', 2000.00, 6000.00);
