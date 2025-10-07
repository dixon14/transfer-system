-- Create accounts table
CREATE TABLE IF NOT EXISTS accounts (
    account_id SERIAL PRIMARY KEY,
    balance DECIMAL(20, 5) NOT NULL DEFAULT 0.00000,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Create transactions table
CREATE TABLE IF NOT EXISTS transactions (
    id SERIAL PRIMARY KEY,
    source_account_id SERIAL NOT NULL,
    destination_account_id SERIAL NOT NULL,
    amount DECIMAL(20, 5) NOT NULL,
    status VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_source_account FOREIGN KEY (source_account_id) REFERENCES accounts(account_id),
    CONSTRAINT fk_destination_account FOREIGN KEY (destination_account_id) REFERENCES accounts(account_id)
);

-- Create indexes for better query performance
CREATE INDEX idx_transactions_source_account ON transactions(source_account_id);
CREATE INDEX idx_transactions_destination_account ON transactions(destination_account_id);
CREATE INDEX idx_transactions_created_at ON transactions(created_at);
