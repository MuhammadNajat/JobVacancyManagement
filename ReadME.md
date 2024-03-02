1. To run the project, run the server written / configured in go.
Please run the following command from the root directory:

`go run server.go`

The application is expected to be running on `localhost:8080` for testing with GraphQL playground. However, the GraphQL services can be consumed at `localhost:8080/graphql` in case you want to test them using cURL or Postman.


2. GraphQL services for example:

i.
`
mutation CreateJobDescription($input: JobDescriptionCreationInput!) {
  createJobDescription(input: $input) {
    _id
    positionName
    skills
    yearsOfExperience
    minSalary
    active
  }
}
`

ii.
`
query GetJobDescriptions {
 jobDescriptions {
    _id
    positionName
    skills
    yearsOfExperience
    minSalary
    active
 }
}
`

iii.
`
query GetJobDescription($id : ID!) {
 jobDescription(id : $id) {
    _id
    positionName
    skills
    yearsOfExperience
    minSalary
    active
 }
}
`

iv.
`
mutation UpdateJobDescription($id: ID!, $input: JobDescriptionUpdateInput!) {
  updateJobDescription(id: $id, input: $input) {
    _id
    positionName
    skills
    yearsOfExperience
    minSalary
    active
  }
}
`

v.
`
mutation DeleteJobDescription($id: ID!) {
  deleteJobDescription(id: $id) {
		deletedJobDescriptionID
  }
}
`


3. Query and Mutation calls to perfrom from GraphQL playground:
   
A. To perfrom `Create`, use any of the followings as input:

i.
`
{
  "input": {
  	"positionName" : "Programmer",
  	"yearsOfExperience" : 3,
    	"skills": ["C++", "JavaScript", "Node.js", "React"],
    	"minSalary": 80000,
    	"active": true
  }
}
`

ii.
`
{
  "input": {
  	"positionName" : "Mechanical Engineer",
  	"yearsOfExperience" : 13,
    	"skills": ["VLSI", "Robotics", "Microcontroller"],
    	"minSalary": 800000,
    	"active": true
  }
}
`

iii.
`
{
  "input": {
  	"positionName" : "Science Teacher",
  	"yearsOfExperience" : 7,
    	"skills": ["Teaching"],
    	"minSalary": 45000,
    	"active": true
  }
}
`

B. To `Update`, you may use (make sure to give a real `id`):

`
{
  "id": "123abc",
  "input": {
  	"positionName" : null,
  	"yearsOfExperience" : null,
    	"skills": ["Teaching", "Physics"],
    	"minSalary": null,
    	"active": null
  }
}
`

C. To `Read` all the Job Descriptions, no variables are needed to set. To get a Job Description by its ID:

`
{
  "id": "123abc"
}
`

D. To `Delete` by ID:

`
{
  "id": "123abc"
}
`
