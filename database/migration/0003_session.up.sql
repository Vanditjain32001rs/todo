CREATE TABLE IF NOT EXISTS sessions
(
    session_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    id         uuid,
    archived_at timestamp with time zone,
    expires_at timestamp with time zone,
    Foreign key (id)
        references todo (id)
);
