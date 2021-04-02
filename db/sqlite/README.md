Library that used 
 - https://github.com/mattn/go-sqlite3

 Scenario :

 Database : user
 Tables :

# Create table
```sql
user
 ------
 user_id | int
 name | string
 sex | string
 age | int
 ------
```
# Alter
```sql
user
 ------
 user_id | int | primary | auto increment
 name | string | not null
 sex | string | not null
 age | int | not null
 hobby | string | nullable
 occupation | string | nullable
 ------
```
# Insert
```sql
------
Fakih | Male   | 17 | Swimming | Swimmer
Hikaf | Female | 16 | - | Office Assistant
Kafhi | Male   | 11 | Jogging | -
Fhika | Female | 20 | - | -
Ikfha | Female | 55 | - | -
------
```
# Update
Update Female Hobby to be all Swimming
# Delete
Delete Ikfha
# Select
List all user that have age < 20 
List all user that have age > 20
