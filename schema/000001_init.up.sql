CREATE TABLE users
(
    id SERIAL not null unique,
    username varchar(255) not null unique,
    email varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE category
(
    id serial not null unique,
    slug varchar(255) not null,
    title varchar(255) not null,
    description text,
    image_url varchar(512) not null,
    parent_id int default null
);

CREATE TABLE course
(
    id serial not null unique,
    slug varchar(255) not null,
    title varchar(255) not null,
    description text,
    language varchar(255) default 'ru',
    likes int default 0 not null,
    dislikes int default 0 not null,
    is_active boolean not null default false,
    image_url varchar(512) default '',
    archive_url varchar(512) default '',
    materials_url varchar(512) default '',
    category_id int references category (id) on delete cascade not null,
    created_at timestamp DEFAULT now(),
    updated_at timestamp DEFAULT now()
);

CREATE TABLE users_courses
(
    id serial not null unique,
    user_id int references users (id) on delete cascade not null,
    course_id int references course (id) on delete cascade not null
);

CREATE TABLE video
(
    id serial not null unique,
    video_order serial not null,
    video_url varchar(512) not null,
    file_name varchar(255) not null,
    preview_image_url varchar(512),
    course_id int references course (id) on delete cascade not null
);

CREATE TABLE comments
(
    id serial not null unique,
    body text,
    course_id int references course (id) on delete cascade not null,
    user_id int references users (id) on delete cascade not null,
    created_at timestamp,
    updated_at timestamp
);

CREATE TABLE course_rating
(
    id serial not null unique,
    user_id int references users (id) on delete cascade not null,
    course_id int references course (id) on delete cascade not null,
    is_like boolean
);
