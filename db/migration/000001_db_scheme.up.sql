CREATE TABLE IF NOT EXISTS account (
    account_id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS deck (
    deck_id SERIAL PRIMARY KEY,
    account_id INT REFERENCES account(account_id) ON DELETE CASCADE NOT NULL,
    title VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS flashcard (
    flashcard_id SERIAL PRIMARY KEY,
    deck_id INT REFERENCES deck(deck_id) ON DELETE CASCADE NOT NULL,
    question TEXT NOT NULL,
    answer TEXT NOT NULL,
    next_review_date TIMESTAMP,
    interval INT,
    repetitions INT,
    easiness_factor DOUBLE PRECISION,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    is_archived BOOLEAN DEFAULT FALSE
);

CREATE INDEX IF NOT EXISTS idx_deck_account_id ON deck (account_id);
CREATE INDEX IF NOT EXISTS idx_flashcard_deck_id ON flashcard (deck_id);
CREATE INDEX IF NOT EXISTS idx_flashcard_next_review_date ON flashcard (next_review_date);
-- TO DO INDEX OF USERNAME