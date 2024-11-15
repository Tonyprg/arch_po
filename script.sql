CREATE TABLE IF NOT EXISTS client (
    id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name VARCHAR(70) NOT NULL,
    last_name VARCHAR(70) NOT NULL,
    surname VARCHAR(70) NOT NULL,
    birthdate DATE NOT NULL,
    issue_plase VARCHAR(200) NOT NULL,
    phone_number BIGINT NOT NULL
);

CREATE TABLE IF NOT EXISTS agent (
    id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    serial_number INT NOT NULL
);

CREATE TABLE IF NOT EXISTS service (
    id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    cost INT NOT NULL,
    compensation INT NOT NULL
);

CREATE TABLE IF NOT EXISTS bid (
    id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    status VARCHAR(40) NOT NULL,
    client_id INT,
    agent_id INT,
    service_id INT,

    FOREIGN KEY (client_id)  REFERENCES client  (id),
    FOREIGN KEY (agent_id)   REFERENCES agent   (id),
    FOREIGN KEY (service_id) REFERENCES service (id)
);

DROP TABLE bid;
DROP TABLE service;
DROP TABLE agent;
DROP TABLE client;

SELECT * FROM client;
SELECT * FROM agent;
SELECT * FROM service;
SELECT * FROM bid;

\dt
