package dbutils

// UsersInsert Insert an initial user set in the database
var UsersInsert = `INSERT INTO USERS (UUID, NAME, EMAIL, AGE, CREATED_AT, UPDATED_AT)
VALUES ("6fa1b5af-f751-422e-a578-09fff549fe48", "KALU RAM", "kaluram@saathindustani.com", 34, '2021-01-01 15:54:20.954', '2021-01-01 15:54:20.954'),
("244607ae-24df-48d1-9f1e-de752277d1f0", "SALKI AUSSIE", "saussie@saathindustani.com", 34, '2021-01-01 15:54:20.954', '2021-01-01 15:54:20.954'),
("b999b167-e91b-48f9-9182-53e6dc65f2de", "MOTA", "motoa@saathindustani.com", 34, '2021-01-01 15:54:20.954', '2021-01-01 15:54:20.954'),
("c0931805-6349-4c37-8691-920cbbf555d0", "chhotu", "chhotu@saathindustani.com", 34, '2021-01-01 15:54:20.954', '2021-01-01 15:54:20.954'),
("cbb04455-6c20-4cf6-9d0a-837d14304d71", "ran darlo", "randarlo@saathindustani.com", 34, '2021-01-01 15:54:20.954', '2021-01-01 15:54:20.954');`