# About model

### How to migrate struct model?

If you have created a model, the structs you define will not be migrated immeditely.

You must register your pointer struct (ex: ```&Student```) in ```Migrate()``` function in ```migrate.go``` so that the ```Migrate()``` function will automatically migrate when you run the application again.

### How to create an enum type (for postgresql migration)?

You can create SQl query variable in ```CreatePostgresEnums()``` function in ```enums.go``` then execute the SQL query variable you created with ```db.Exec(yourEnumSQL)```. The enums you define will be migrated when you run the application again. After that, create an enum type below in the same file if you want to use that type for other models.