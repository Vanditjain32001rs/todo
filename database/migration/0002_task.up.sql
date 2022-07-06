DROP TABLE IF EXISTS tasks;
CREATE TABLE tasks
(
    task_id          UUID PRIMARY KEY         DEFAULT gen_random_uuid(),
    id               UUID,
    task_description TEXT NOT NULL,
    is_complete      BOOLEAN NOT NULL         DEFAULT false,
    archived_at     TIMESTAMP with time zone,
    FOREIGN KEY (id)
        references todo (id)
);