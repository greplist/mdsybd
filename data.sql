INSERT INTO users (id, name, email, role, encrypted_password)
    VALUES (1, 'anton', 'anton.grabko@synesis.ru', 0, 'password');
INSERT INTO users (id, name, email, role, encrypted_password)
    VALUES (2, 'kirill', 'kirill.kotlarov@synesis.ru', 0, 'password');

INSERT INTO categories (id, name, title, visible) VALUES (1, 'sport', 'This is sport', 1);
INSERT INTO categories (id, name, title, visible) VALUES (2, 'IT', 'This is IT', 1);

INSERT INTO tags (id, name) VALUES (1, 'sql');
INSERT INTO tags (id, name) VALUES (2, 'golang');

INSERT INTO content_formats (id, background_image) VALUES (1, 'http::/127.0.0.1/background.jpeg');

INSERT INTO contents (id, title, content, published, published_at, category_id, content_format_id)
    VALUES (1, 'first', 'first SPORT content', 1, SYSDATE, 1, 1);
INSERT INTO contents (id, title, content, published, published_at, category_id, content_format_id)
    VALUES (2, 'second', 'second IT content', 1, SYSDATE, 2, 1);
