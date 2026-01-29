CREATE TABLE IF NOT EXISTS attendees (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    event_id INTEGER NOT NULL,
    CONSTRAINT fk_attendees_user
        FOREIGN KEY (user_id)
        REFERENCES users (id)
        ON DELETE CASCADE,
    CONSTRAINT fk_attendees_event
        FOREIGN KEY (event_id)
        REFERENCES events (id)
        ON DELETE CASCADE
);
