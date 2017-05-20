INSERT INTO users (id, name, email, role, encrypted_password)
    VALUES (1, 'anton', 'anton.grabko@synesis.ru', 0, 'password');
INSERT INTO users (id, name, email, role, encrypted_password)
    VALUES (2, 'kirill', 'kirill.kotlarov@synesis.ru', 0, 'password');

INSERT INTO categories (id, name, title, visible) VALUES (1, 'sport', 'This is sport', 1);
INSERT INTO categories (id, name, title, visible) VALUES (2, 'IT', 'This is IT', 1);

INSERT INTO tags (id, name) VALUES (1, 'sql');
INSERT INTO tags (id, name) VALUES (2, 'oracle');
INSERT INTO tags (id, name) VALUES (3, 'golang');

INSERT INTO content_formats (id, background_image) VALUES (1, 'http::/127.0.0.1/background.jpeg');

INSERT INTO contents (id, title, content, published, published_at, category_id, content_format_id)
    VALUES (1, 'first', 'first SPORT content', 1, SYSDATE, 1, 1);
INSERT INTO contents (id, title, content, published, published_at, category_id, content_format_id)
    VALUES (2, 'second', 'second IT content', 1, SYSDATE, 2, 1);

INSERT INTO taggings (id, tag_id, content_id) VALUES (1, 1, 1);
INSERT INTO taggings (id, tag_id, content_id) VALUES (2, 2, 1);
INSERT INTO taggings (id, tag_id, content_id) VALUES (3, 2, 2);
INSERT INTO taggings (id, tag_id, content_id) VALUES (4, 3, 2);

INSERT INTO personalities (id, content_id, user_id) VALUES (1, 1, 1);
INSERT INTO personalities (id, content_id, user_id) VALUES (2, 1, 2);
INSERT INTO personalities (id, content_id, user_id) VALUES (3, 2, 1);

INSERT INTO images (id, file_name, content_type, photographer, location, content_id)
    VALUES (1, 'file', 'image/jpeg', 'photographer', 'Minsk Belarus', 1);
INSERT INTO infografics (id, html, content_id) VALUES (1, 'html', 1);
INSERT INTO places (id, place_type, content_type, content_id) VALUES (1, 1, 'text/json', 1);
INSERT INTO pages (id, title, content_id) VALUES (1, 'page', 1);
INSERT INTO banners (id, name, image, link, height, content_id) VALUES (1, 'banner', 'image', 'link to banner', 320, 1);
INSERT INTO news (id, title, body, photo, content_id) VALUES (1, 'new', 'new body', 'photo', 1);

COMMIT;
