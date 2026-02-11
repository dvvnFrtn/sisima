# About DTO

### DTO (```dto.go```)
You can define request and response contracts at here.

*You will must interact with this file and create a structs in it.*

### Pagination (```pagination.go```)
Actually using pagination to find all resources is better than not using it. You can use ```NewPagination()``` for it. Run two queries in the service before using this, namely: count all tables and select all. Example:
```
config.DB.Model(&model.Student{}).Count(&total)
config.DB.Limit(limit).Offset((page - 1) * limit).Find(&students)
```

*You will probably change this code very rarely.*

### Validation (```validation.go```)
Using for validation error that returning HTTP status 422.

*You will probably change this code very rarely.*