# uPepperoni in GO
a theoretical small open reading frame finding tool

Database configuration requirement:
-
- Postgres
- Available schema named "upep"


Core tool Main.go parameter:
-
- driver: postgres
- dbName: chosen postgres db for storing table and schema
- user: postgres username
- pass: postgres password
- sslmode: postgres secure mode
- port: postgres access port
- initCodon: first time setup

First time setup:
-
- By running with initCodon, the program will attempt to the schema and migrate the necessary table.
- 
