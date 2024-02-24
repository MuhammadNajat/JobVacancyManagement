1. GraphQL services for example:
   
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

mutation DeleteJobDescription($id: ID!) {
  deleteJobDescription(id: $id) {
		deletedJobDescriptionID
  }
}



2. Query and Mutation calls to perfrom from GraphQL playground:
   
A. To perfrom Create, use any of the followings as input
{
  "input": {
  	"positionName" : "Programmer",
  	"yearsOfExperience" : 3,
    	"skills": ["C++", "JavaScript", "Node.js", "React"],
    	"minSalary": 80000,
    	"active": true
  }
}

{
  "input": {
  	"positionName" : "Mechanical Engineer",
  	"yearsOfExperience" : 13,
    	"skills": ["VLSI", "Robotics", "Microcontroller"],
    	"minSalary": 800000,
    	"active": true
  }
}

{
  "input": {
  	"positionName" : "Science Teacher",
  	"yearsOfExperience" : 7,
    	"skills": ["Teaching"],
    	"minSalary": 45000,
    	"active": true
  }
}


B. To Update, you may use (make sure to give a real `id`):
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
