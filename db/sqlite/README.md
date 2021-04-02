Library that used 
 - https://github.com/mattn/go-sqlite3

 Scenario :

 Database : user
 Tables :

# Create table
```sql
user
 ------
 user_id | int | primary | auto increment/ alias for rowid
 name | text | not null | unique
 sex | text | not null
 age | int | not null
 phone | text | null | unique
 ------
```
# Alter (change)
```sql
user
 ------
 user_id | int | primary | auto increment/ alias for rowid
 name | text | not null | unique
 sex | text | not null
 age | int | not null
 phone | text | null | unique
 hobby | text | null
 occupation | text | null
 ------
```
# Alter - table and field name, drop phone column but delete column is not supported per march 2021, so just ignore this column
```sql
users
 ------
 user_id | int | primary | auto increment/ alias for rowid
 full_name | text | not null | unique
 sex | text | not null
 age | int | not null
 phone | text | null | unique
 hobby | text | null
 occupation | text | null
 ------
```
# Insert
with and without transaction
```sql
------
Fakih | Male   | 17 | Swimming | Swimmer
Hikaf | Female | 16 | - | Office Assistant
Kafhi | Male   | 11 | Jogging | -
Fhika | Female | 20 | - | -
Ikfha | Female | 55 | - | -
------
```
# Upsert
with and without transaction
```sql
Fakih 1 | Male   | 17 | Swimming | Swimmer
Hikaf 2 | Female | 16 | - | Office Assistant
Kafhi 3 | Male   | 11 | Jogging | -
Fhika 4 | Female | 20 | - | -
Ikfha 5 | Female | 55 | - | -
```
# Update
with and without transaction
Update for name like Ikfha to be all Swimming
# Delete
Delete Ikfha
Delete user with age 11
# Select
List all user that have age < 18 
List all user that have age > 18
