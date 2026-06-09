-- Write your query below
select name from customers as c
left join orders as o on o.customer_id = c.id
where o is null 