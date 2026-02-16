-- Migration: insert_hsk_1_2026_data
-- Created at: 2026-02-14 11:53:00

-- Delete in reverse dependency order
DELETE FROM example_translations WHERE example_id IN (SELECT id FROM examples WHERE word_id IN (SELECT id FROM words WHERE hsk_source_id = 1));
DELETE FROM word_translations WHERE word_id IN (SELECT id FROM words WHERE hsk_source_id = 1);
DELETE FROM examples WHERE word_id IN (SELECT id FROM words WHERE hsk_source_id = 1);
DELETE FROM words WHERE hsk_source_id = 1;
DELETE FROM hsk_sources WHERE id = 1;
