# API Documentation

This is the API documentation for Billy's Pokemon API. It allows users to create, read, update, and delete Pokemon from the database, search for Pokemon by name. The API is built using Golang and the Gin framework.

## Setup

Create a `.env` file in the root of the server directory with the following:

```
PORT=8080
MONGODB_URI=mongodb://localhost:27017/your-database-name
```

To run the server locally, you need to have Golang installed on your machine. Run the following commands:

```bash
go mod download
go run main.go
```

The API will start running on port 8080. You can access the API by visiting the following URL in your browser:

```bash
http://localhost:8080/api/
```

## Endpoints

The API has the following endpoints:

| Method | Endpoint                      | Description                  |
| ------ | ----------------------------- | ---------------------------- |
| GET    | `/api/pokemon`                | Get all Pokemon              |
| GET    | `/api/pokemon/:id`            | Get a Pokemon by ID          |
| GET    | `/api/pokemon/search?q=bulba` | Search for a Pokemon by name |
| POST   | `/api/pokemon`                | Create a new Pokemon         |
| PUT    | `/api/pokemon/:id`            | Update a Pokemon by ID       |
| DELETE | `/api/pokemon/:id`            | Delete a Pokemon by ID       |

### Data model

The Pokémon data model is defined as follows:

```go
type Pokemon struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name            string             `bson:"name" json:"name"`
	Number          int                `bson:"number" json:"number"`
	ImageURL        string             `bson:"imageUrl" json:"imageUrl"`
	Description     string             `bson:"description" json:"description"`
	BackgroundColor string             `bson:"backgroundColor" json:"backgroundColor"`
}
```

The `ID` field is an auto-generated unique identifier for each Pokémon. The `Name`, `Number`, `ImageURL`, `Description`, and `BackgroundColor` fields are used to store information about each Pokémon.
