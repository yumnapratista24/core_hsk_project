-- 1. Tabel Sumber HSK
CREATE TABLE hsk_sources (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL, 
    level INTEGER NOT NULL,      
    year_released INTEGER,
    hsk_version VARCHAR(100),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

-- 2. Tabel Kata (Words)
CREATE TABLE words (
    id SERIAL PRIMARY KEY,
    hanzi VARCHAR(100) NOT NULL,
    pinyin VARCHAR(100) NOT NULL,
    hsk_source_id INTEGER REFERENCES hsk_sources(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

-- 3. Tabel Contoh Kalimat (Examples)
CREATE TABLE examples (
    id SERIAL PRIMARY KEY,
    word_id INTEGER REFERENCES words(id) ON DELETE CASCADE,
    hanzi TEXT NOT NULL,
    pinyin TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

-- 4. Tabel Terjemahan Kata (Word Translations)
CREATE TABLE word_translations (
    id SERIAL PRIMARY KEY,
    word_id INTEGER NOT NULL REFERENCES words(id) ON DELETE CASCADE,
    language VARCHAR(10) NOT NULL,      -- 'id' atau 'en'
    value TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

-- 5. Tabel Terjemahan Contoh Kalimat (Example Translations)
CREATE TABLE example_translations (
    id SERIAL PRIMARY KEY,
    example_id INTEGER NOT NULL REFERENCES examples(id) ON DELETE CASCADE,
    language VARCHAR(10) NOT NULL,      -- 'id' atau 'en'
    value TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);