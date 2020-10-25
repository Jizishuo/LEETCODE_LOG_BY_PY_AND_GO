select
    name, population, area
from
    world
where
    area > 3000000 or population > 25000000;

-- select
--     name, population, area
-- from
--     world
-- where
--     area > 3000000
-- union
-- select
--     name, population, area
-- from
--     world
-- where
--     population > 25000000