-- Создание таблицы "Национальности"
CREATE TABLE nationalities (
id SERIAL PRIMARY KEY,
name VARCHAR(100) NOT NULL
);

-- Создание таблицы "Актеры"
CREATE TABLE actors (
id SERIAL PRIMARY KEY,
first_name VARCHAR(100) NOT NULL,
last_name VARCHAR(100) NOT NULL,
nationality_id INTEGER REFERENCES nationalities(id),
film_count INTEGER NOT NULL,
total_earnings DECIMAL(10, 2) NOT NULL
);

-- Добавление данных в таблицу "Национальности"
INSERT INTO nationalities (name) VALUES
('Американец'),
('Британец'),
('Француз'),
('Русский'),
('Индиец'),
('Китаец'),
('Итальянец'),
('Немец'),
('Испанец'),
('Бразилец'),
('Канадец'),
('Казах');

-- Добавление данных в таблицу "Актеры"
INSERT INTO actors (first_name, last_name, nationality_id, film_count, total_earnings) VALUES
('Леонардо', 'ДиКаприо', 1, 40, 750.00),
('Том', 'Хэнкс', 1, 60, 800.00),
('Брэд', 'Питт', 1, 50, 700.00),
('Джонни', 'Депп', 1, 55, 650.00),
('Кейт', 'Уинслет', 2, 45, 600.00),
('Эмма', 'Уотсон', 2, 30, 500.00),
('Одри', 'Тоту', 3, 25, 450.00),
('Юрий', 'Борисов', 4, 35, 10.00),
('Аамир', 'Хан', 5, 40, 600.00),
('Джеки', 'Чан', 6, 50, 700.00),
('Роберт', 'Де Ниро', 1, 45, 750.00),
('Аль', 'Пачино', 1, 50, 700.00),
('Джулия', 'Робертс', 1, 40, 650.00),
('Родриго', 'Санторо', 10, 35, 15.00),
('Николас', 'Кейдж', 1, 55, 700.00);