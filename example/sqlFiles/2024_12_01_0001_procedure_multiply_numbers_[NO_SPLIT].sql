DELIMITER //

CREATE PROCEDURE MultiplyNumbers(IN num1 INT, IN num2 INT, OUT result INT)
BEGIN
    SET result = num1 * num2;
END //

DELIMITER ;
