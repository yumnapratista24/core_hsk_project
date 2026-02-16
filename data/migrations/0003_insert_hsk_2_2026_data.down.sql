-- Migration: insert_hsk_2_2026_data
-- Created at: 2026-02-14 12:41:14

-- Delete in reverse dependency order
DELETE FROM example_translations WHERE example_id IN (SELECT id FROM examples WHERE word_id IN (SELECT id FROM words WHERE hsk_source_id = 2));
DELETE FROM word_translations WHERE word_id IN (SELECT id FROM words WHERE hsk_source_id = 2);
DELETE FROM examples WHERE word_id IN (SELECT id FROM words WHERE hsk_source_id = 2);
DELETE FROM words WHERE hsk_source_id = 2;
DELETE FROM hsk_sources WHERE id = 2;
