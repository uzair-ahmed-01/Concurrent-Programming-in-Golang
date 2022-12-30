###This will create a table called "numbers" with two columns: "id" and "number". The "id" column is an auto-incrementing integer that serves as the primary key for the table, and the "number" column is an integer that will store the numbers produced by the producer.

CREATE TABLE numbers (
  id INT AUTO_INCREMENT PRIMARY KEY,
  number INT
);

###You can then use the INSERT INTO statement in the Go code to insert the numbers produced by the producer into the "numbers" table.
