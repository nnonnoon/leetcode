SELECT Email 
FROM
(
    SELECT Email, count(Email) as num
    FROM Person
    GROUP BY Email
) as statistic 

WHERE NUM > 1

-- way-2 --

SELECT Email
FROM Person
GROUP BY Email
HAVING count(Email) > 1