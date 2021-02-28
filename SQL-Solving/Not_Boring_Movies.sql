SELECT id, movie, description, rating 
FROM cinema 
WHERE mod(id, 2) = 1 and  description <> 'boring'
ORDER BY rating DESC