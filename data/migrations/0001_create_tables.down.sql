-- Down migration: Drop all tables in reverse order to handle foreign key constraints

-- 1. Drop translation tables first (they reference words/examples)
DROP TABLE IF EXISTS example_translations;
DROP TABLE IF EXISTS word_translations;

-- 2. Drop examples (references words)
DROP TABLE IF EXISTS examples;

-- 3. Drop words (references hsk_sources)
DROP TABLE IF EXISTS words;

-- 4. Drop hsk_sources last (referenced by words)
DROP TABLE IF EXISTS hsk_sources;