# go-what-did-i-do

[![Build Application](https://github.com/joukojo/go-what-did-i-do/actions/workflows/build.yml/badge.svg?branch=develop)](https://github.com/joukojo/go-what-did-i-do/actions/workflows/build.yml)

Simple command line tool for handling project tasks

# Data relations 

```mermaid
erDiagram
    CUSTOMER ||--o{ PROJECT : has
    PROJECT ||--o{ TASK : contains
    TASK ||--o{ WORK_EVENT : has

    CUSTOMER {
        int id
        string name
    }
    PROJECT {
        int id
        string name
        int customer_id
    }
    TASK {
        int id
        string title
        int project_id
    }
    WORK_EVENT {
        int id
        string description
        int task_id
    }



# Data storage

The data is stored under user's home directory 

```bash
$HOME/.what-did-i-do  
```

| File    | Description |
| -------- | ------- |
| customers.json  |  customer-data in json-format    |

# Usage 

Help page is available for every command by: 

```bash 
what-did-i-do --help 
```


## Customers 

The customers command group allows you to manage your customer list. A customer can have multiple projects associated with them.
🔧 Usage

what-did-i-do customers [command] [flags]

### Available Commands

| Command    | Description |
| -------- | ------- |
| add  | Adds a new customer to the system    |
| delete | Deletes an existing customer by ID     |

### Flags

    -h, --help Show help for the customers command

### Examples 

#### Add a Customer

```bash
what-did-i-do customers add --name "Acme Corporation"
```

#### Delete a Customer

```bash
what-did-i-do customers delete --id 123445 
```

#### List Customers 

```bash
what-did-i-do customers 
```

## Projects 

This comamnd handles projects under the customer. One customer could have non or many projects. 

#### Add a project to customer


```bash
what-did-i-do projects add --customerId 12345 --name "Acme Corporation"
```

#### Delete a project

```bash
what-did-i-do projects delete --id 12346
```

#### List projects 

```bash
what-did-i-do projects 
```