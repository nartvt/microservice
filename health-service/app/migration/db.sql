CREATE TYPE newsfeed_sections_type_enum AS ENUM (
    'top_page',
    'personal_record',
    'about_recommendations'
    );

CREATE TABLE "newsfeed_sections"
(
    "id"         SERIAL PRIMARY KEY,
    "name"       varchar(50) NOT NULL,
    "active"     BOOLEAN,
    "type"       newsfeed_sections_type_enum,
    "created_at" timestamp   NOT NULL DEFAULT (CURRENT_TIMESTAMP),
    "updated_at" timestamp   NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "dishes"
(
    "id"         SERIAL PRIMARY KEY,
    "name"       VARCHAR(50) NOT NULL,
    "image"      TEXT,
    "active"     BOOLEAN,
    "created_at" timestamp   NOT NULL DEFAULT (CURRENT_TIMESTAMP),
    "updated_at" timestamp   NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "section_dishes"
(
    "section_id" INT,
    "dish_id"    INT,
    PRIMARY KEY (section_id, dish_id)
);
CREATE INDEX CONCURRENTLY section_dishes_dish_id_idx ON section_dishes (dish_id);
CREATE INDEX CONCURRENTLY section_dishes_section_id_idx ON section_dishes (section_id);


CREATE TABLE "users"
(
    "id"           SERIAL PRIMARY KEY,
    "username"     VARCHAR(255) UNIQUE NOT NULL,
    "password"     VARCHAR(255)        NOT NULL,
    "email"        VARCHAR(255) UNIQUE NOT NULL,
    "phone_number" VARCHAR(11) UNIQUE  NOT NULL,
    "full_name"    VARCHAR(255),
    "created_at"   timestamp           NOT NULL DEFAULT (CURRENT_TIMESTAMP),
    "updated_at"   timestamp           NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);

CREATE INDEX CONCURRENTLY users_phone_number_idx ON users (phone_number);
CREATE INDEX CONCURRENTLY users_email_idx ON users (email);


CREATE TABLE "user_body_records"
(
    "id"         SERIAL PRIMARY KEY,
    "user_id"    INT           NOT NULL,
    "weight"     DECIMAL(5, 2) NOT NULL,
    "height"     INT           NOT NULL,
    "percentage" DECIMAL(4, 2) NOT NULL,
    "created_at" timestamp     NOT NULL DEFAULT (CURRENT_TIMESTAMP),
    "updated_at" timestamp     NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);
CREATE INDEX CONCURRENTLY user_body_records_user_id_idx ON user_body_records (user_id);

CREATE TABLE "user_exercises"
(
    "id"              SERIAL PRIMARY KEY,
    "user_id"         INT       NOT NULL,
    "at_time"         INT       NOT NULL,
    "description"     TEXT      NOT NULL,
    "calories_burned" INT       NOT NULL,
    "created_at"      timestamp NOT NULL DEFAULT (CURRENT_TIMESTAMP),
    "updated_at"      timestamp NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);
CREATE INDEX CONCURRENTLY user_exercises_user_id_idx ON user_exercises (user_id);

CREATE TABLE "user_diaries"
(
    "id"          SERIAL PRIMARY KEY,
    "user_id"     INT       NOT NULL,
    "at_time"     INT       NOT NULL,
    "description" TEXT      NOT NULL,
    "calories"    INT       NOT NULL,
    "created_at"  timestamp NOT NULL DEFAULT (CURRENT_TIMESTAMP),
    "updated_at"  timestamp NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);
CREATE INDEX CONCURRENTLY user_diaries_user_id_idx ON user_diaries (user_id);

CREATE TABLE "abouts"
(
    "id"         SERIAL PRIMARY KEY,
    "name"       VARCHAR(50) NOT NULL,
    "image"      TEXT,
    "active"     BOOLEAN,
    "created_at" timestamp   NOT NULL DEFAULT (CURRENT_TIMESTAMP),
    "updated_at" timestamp   NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "sections_about"
(
    "section_id" INT,
    "about_id"   INT,
    PRIMARY KEY (section_id, about_id)
);
CREATE INDEX CONCURRENTLY sections_about_about_id_idx ON sections_about (about_id);
CREATE INDEX CONCURRENTLY sections_about_section_id_idx ON sections_about (section_id);


ALTER TABLE "user_body_records" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "user_exercises" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "user_diaries" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

-- Insert data into "newsfeed_sections" table
INSERT INTO newsfeed_sections(name, active, type)
VALUES ('Morning', true, 'top_page'),
       ('Section 2', true, 'personal_record'),
       ('Section 3', true, 'about_recommendations'),
       ('Snack', true, 'top_page'),
       ('Dinner', true, 'personal_record'),
       ('Section 6', true, 'about_recommendations'),
       ('Lunch', true, 'top_page'),
       ('Section 8', true, 'personal_record'),
       ('Section 9', true, 'about_recommendations'),
       ('Section 10', true, 'top_page');

-- Insert data into "dishes" table
INSERT INTO dishes(name, image, active)
VALUES ('Dish 1', 'https://example.com/dish1.jpg', true),
       ('Dish 2', 'https://example.com/dish2.jpg', true),
       ('Dish 3', 'https://example.com/dish3.jpg', true),
       ('Dish 4', 'https://example.com/dish4.jpg', true),
       ('Dish 5', 'https://example.com/dish5.jpg', true),
       ('Dish 6', 'https://example.com/dish6.jpg', true),
       ('Dish 7', 'https://example.com/dish7.jpg', true),
       ('Dish 8', 'https://example.com/dish8.jpg', true),
       ('Dish 9', 'https://example.com/dish9.jpg', true),
       ('Dish 10', 'https://example.com/dish10.jpg', true);

-- Insert data into "sections_dish" table
INSERT INTO section_dishes(section_id, dish_id)
VALUES (1, 2),
       (1, 4),
       (1, 6),
       (2, 1),
       (2, 3),
       (2, 5),
       (3, 8),
       (3, 10),
       (4, 7),
       (4, 9);

INSERT INTO users (username, password, email, phone_number, full_name)
VALUES ('user1', 'password1', 'user1@example.com', '12345678901', 'User One'),
       ('user2', 'password2', 'user2@example.com', '23456789012', 'User Two'),
       ('user3', 'password3', 'user3@example.com', '34567890123', 'User Three'),
       ('user4', 'password4', 'user4@example.com', '45678901234', 'User Four'),
       ('user5', 'password5', 'user5@example.com', '56789012345', 'User Five'),
       ('user6', 'password6', 'user6@example.com', '67890123456', 'User Six'),
       ('user7', 'password7', 'user7@example.com', '78901234567', 'User Seven'),
       ('user8', 'password8', 'user8@example.com', '89012345678', 'User Eight'),
       ('user9', 'password9', 'user9@example.com', '90123456789', 'User Nine'),
       ('user10', 'password10', 'user10@example.com', '01234567890', 'User Ten');


INSERT INTO user_body_records (user_id, weight, height, percentage) VALUES (1, 70.5, 170, 20.5);
INSERT INTO user_body_records (user_id, weight, height, percentage) VALUES (2, 60.2, 165, 18.5);
INSERT INTO user_body_records (user_id, weight, height, percentage) VALUES (3, 80.1, 180, 25.0);
INSERT INTO user_body_records (user_id, weight, height, percentage) VALUES (4, 90.0, 185, 30.5);
INSERT INTO user_body_records (user_id, weight, height, percentage) VALUES (5, 75.5, 175, 22.5);

INSERT INTO user_exercises (user_id, at_time, description, calories_burned) VALUES (1, 8, 'Running', 200);
INSERT INTO user_exercises (user_id, at_time, description, calories_burned) VALUES (2, 9, 'Cycling', 250);
INSERT INTO user_exercises (user_id, at_time, description, calories_burned) VALUES (3, 10, 'Swimming', 300);
INSERT INTO user_exercises (user_id, at_time, description, calories_burned) VALUES (4, 11, 'Yoga', 150);
INSERT INTO user_exercises (user_id, at_time, description, calories_burned) VALUES (5, 12, 'Walking', 100);

INSERT INTO user_diaries (user_id, at_time, description, calories)VALUES (1, 8, 'Breakfast - Oatmeal', 200);
INSERT INTO user_diaries (user_id, at_time, description, calories) VALUES (2, 9, 'Lunch - Grilled Chicken Salad', 300);
INSERT INTO user_diaries (user_id, at_time, description, calories) VALUES (3, 10, 'Dinner - Salmon with vegetables', 400);
INSERT INTO user_diaries (user_id, at_time, description, calories) VALUES (4, 11, 'Snack - Apple and Peanut Butter', 150);
INSERT INTO user_diaries (user_id, at_time, description, calories) VALUES (5, 12, 'Breakfast - Scrambled Eggs with Toast', 250);

INSERT INTO abouts (name, image, active) VALUES ('Workouts', 'workouts.jpg', true);
INSERT INTO abouts (name, image, active) VALUES ('Diets', 'diets.jpg', true);
INSERT INTO abouts (name, image, active) VALUES ('Community', 'community.jpg', true);
INSERT INTO abouts (name, image, active) VALUES ('Settings', 'settings.jpg', true);
INSERT INTO abouts (name, image, active) VALUES ('Profile', 'profile.jpg', true);

INSERT INTO sections_about (section_id, about_id) VALUES (3, 1);
INSERT INTO sections_about (section_id, about_id) VALUES (6, 2);
INSERT INTO sections_about (section_id, about_id) VALUES (9, 3);