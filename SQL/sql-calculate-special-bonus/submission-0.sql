WITH regional_sales AS (
    select employee_id, salary from employees 
    where name not like 'M%'
    and employee_id % 2 != 0
)
SELECT 
    e.employee_id, -- Четко указываем, откуда брать ID (из r)
    CASE 
        WHEN r.salary IS NULL THEN 0 -- Четко указываем, откуда брать salary
        ELSE r.salary 
    END AS bonus 
FROM regional_sales as r
right join employees as e on e.employee_id = r.employee_id;