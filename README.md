# Todo lists    
Simple Todo application.  

Todos have certain attributes. Lists contain a bunch of Todos  
## App description
1. Create lists (POST /lists)
2. Create Todo (POST /todo)
3. Delete Todo (DELETE /todo/:todoID)
4. Update Todo (PUT /todo/:todoID)
5. Get Todo (GET /todo)

Example Todo Create Payload:  
{  
	"name": "New Todo",  
	"notes": "Description about Todo",  
	"list_id": 1,  
	"completed": true,  
	"due_date": "2018-09-07T00:00:00Z"  
}  

Note: The format for due_date field should be in RFC3339 format.

Update Todo includes:  
* Setting whether it's completed or not  
* Setting due date, notes and other fields  

Get Todo:  
In conjunction with Ajax, this can support auto-complete of results while typing  
* Retrieving results based on whether they are completed.  
* Retrieving results based on name  
The two search params of "completed" and "name" can be used individually or together.  

Eg: /todo?completed=true  
Eg: /todo?completed=true&name=hello  
Eg: /todo?name=world  

### Model description  
Lists contains just two fields:  
* ID : Primary Key  
* Name : String (Name of List)  

Todos contains 6 fields:  
* ID : Primary Key  
* ListID: Foreign Key to row on Lists table  
* Name: String (Name of Todo)  
* Notes: Text (Long description of the Todo)  
* DueDate: Time (Due date of Todo)  
* Completed: Boolean (Denotes whether the Todo has been completed or not)  


Prerequisites: Install docker  

## Run application
```
make dev
```

## Run tests
```
make test
```
