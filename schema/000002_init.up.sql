INSERT INTO category (title, description, slug, image_url, parent_id)
VALUES
    ('Frontend', '', 'frontend', '', 0),
    ('Backend', '', 'backend', '', 0),
    ('System programming', 'System programming involves designing and writing computer programs that allow the computer hardware to interface with the programmer and the user, leading to the effective execution of application software on the computer system.', 'systemprogramming', '', 0),
    ('HTML', 'The HyperText Markup Language, or HTML is the standard markup language for documents designed to be displayed in a web browser.', 'html', '', 1),
    ('CSS', 'Cascading Style Sheets is a style sheet language used for describing the presentation of a document written in a markup language such as HTML.', 'css', '', 1),
    ('Javascript', 'JavaScript, often abbreviated as JS, is a programming language that conforms to the ECMAScript specification. JavaScript is high-level, often just-in-time compiled, and multi-paradigm. It has curly-bracket syntax, dynamic typing, prototype-based object-orientation, and first-class functions.', 'javascript', '', 1),
    ('PHP', 'PHP is a general-purpose scripting language especially suited to web development. It was originally created by Danish-Canadian programmer Rasmus Lerdorf in 1994.', 'php', '', 2),
    ('Rust', 'Rust is a multi-paradigm programming language designed for performance and safety, especially safe concurrency. Rust is syntactically similar to C++, but can guarantee memory safety by using a borrow checker to validate references.', 'rust', '', 3);