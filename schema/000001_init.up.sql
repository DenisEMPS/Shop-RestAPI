CREATE TABLE adress
(
    adress_id SERIAL PRIMARY KEY,
    country VARCHAR(50) NOT NULL,
    city VARCHAR(50) NOT NULL,
    street VARCHAR(50) NOT NULL
);

CREATE TABLE images
(
    image_id UUID NOT NULL PRIMARY KEY,
    image bytea
);

CREATE TABLE client
(   
    client_id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    surname VARCHAR(50) NOT NULL,
    birthday DATE,
    gender BOOLEAN,
    registration_date TIMESTAMP NOT NULL,
    adress_id INT,
    FOREIGN KEY (adress_id) REFERENCES adress (adress_id) ON DELETE RESTRICT
);

CREATE TABLE supplier
(
    supplier_id SERIAL PRIMARY KEY,
    supplier_name VARCHAR(50) NOT NULL,
    supplier_adress_id INT NOT NULL,
    supplier_phone_number VARCHAR(15),
    FOREIGN KEY (adress_id) REFERENCES adress (adress_id) ON DELETE RESTRICT
);


CREATE TABLE product
(
    product_id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    category VARCHAR(50) NOT NULL,
    price NUMERIC (8,2) NOT NULL,
    available_stock INT NOT NULL,
    last_update_date DATE NOT NULL,
    supplier_id INT NOT NULL,
    image_id UUID,
    FOREIGN KEY (supplier_id) REFERENCES supplier (supplier_id) ON DELETE RESTRICT,
    FOREIGN KEY (image_id) REFERENCES images (image_id) ON DELETE SET NULL
);
