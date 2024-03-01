-- Create Customer Table
CREATE TABLE Customer (
    customer_id SERIAL PRIMARY KEY,
    customer_name VARCHAR(255) NOT NULL,
    phone_number VARCHAR(15) NOT NULL
);

-- Create Transaksi Table
CREATE TABLE Transaksi (
    transaction_id SERIAL PRIMARY KEY,
    entry_date DATE NOT NULL,
    completion_date DATE,
    received_by VARCHAR(255) NOT NULL,
    customer_id INT REFERENCES Customer(customer_id)
);

-- Create Detail_Transaksi Table
CREATE TABLE Detail_Transaksi (
    detail_id SERIAL PRIMARY KEY,
    transaction_id INT REFERENCES Transaksi(transaction_id),
    service_name VARCHAR(255) NOT NULL,
    quantity INT NOT NULL,
    unit VARCHAR(50) NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    total DECIMAL(10, 2) NOT NULL
);