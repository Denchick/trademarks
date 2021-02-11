CREATE INDEX idx_trademark_names ON trademarks(name);
CREATE EXTENSION IF NOT EXISTS pg_trgm;
CREATE INDEX trgm_idx_trademarks_name ON trademarks USING GIST (name gist_trgm_ops);
