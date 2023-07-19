CREATE TABLE cars (
    id serial primary key,
    name varchar,
    history varchar
);

INSERT INTO cars (name, history) VALUES
('Fusca', 'The Volkswagen Beetle, also known as the Fusca in Brazil, is a car that was manufactured by the German automaker Volkswagen from 1938 to 2003. It is one of the most popular cars of all time, with over 21 million units sold worldwide.'),
('Chevette', 'The Chevrolet Chevette was a subcompact car that was produced by Chevrolet from 1976 to 1987. It was based on the Opel Kadett, and it was designed to compete with the Ford Pinto and the AMC Gremlin.')
;