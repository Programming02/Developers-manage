CREATE TABLE users (
    id UUID PRIMARY KEY ,
    full_name TEXT NOT NULL ,
    avatar TEXT,
    role TEXT NOT NULL ,
    birth_day DATE NOT NULL ,
    phone TEXT NOT NULL ,
    position TEXT NOT NULL
);


CREATE TABLE project (
    name TEXT NOT NULL ,
    start_date DATE NOT NULL ,
    end_date DATE NOT NULL ,
    status TEXT NOT NULL ,
    teamlead_id NOT NULL REFERENCES user(id) ,
    attachments TEXT
);


CREATE TABLE task (
    title TEXT NOT NULL ,
    description TEXT NOT NULL ,
    start_at TIME NOT NULL ,
    finish_at TIME NOT NULL ,
    status TEXT NOT NULL
    started_at TIME NOT NULL ,
    finished_at TIME NOT NULL ,
    programmer_id UUID NOT NULL REFERENCES users(id) ,
    attachments TEXT
);