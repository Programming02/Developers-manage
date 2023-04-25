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
    id TEXT NOT NULL PRIMARY KEY ,
    name TEXT NOT NULL ,
    start_date DATE NOT NULL ,
    end_date DATE NOT NULL ,
    status TEXT NOT NULL ,
    teamlead_id UUID NOT NULL REFERENCES user(id),
    attachments TEXT
);


CREATE TABLE task (
    id TEXT NOT NULL PRIMARY KEY ,
    title TEXT NOT NULL ,
    description TEXT NOT NULL ,
    start_at TIMESTAMP NOT NULL ,
    finish_at TIMESTAMP NOT NULL ,
    status TEXT NOT NULL,
    started_at TIMESTAMP NOT NULL ,
    finished_at TIMESTAMP NOT NULL ,
    programmer_id UUID NOT NULL REFERENCES users(id) ,
    attachments TEXT,
    project_id TEXT REFERENCES project(id)
);

CREATE TABLE attendance (
    type TEXT NOT NULL ,
    user_id UUID NOT NULL REFERENCES users(id),
    created_at TIMESTAMP
);

CREATE TABLE comments (
    task_id TEXT NOT NULL REFERENCES task(id),
    programmer_id UUID NOT NULL REFERENCES users(id),
    created_at TIMESTAMP NOT NULL
);