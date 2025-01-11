## Pokédex Server

This directory contains RESTful API server for Pokédex clients. The server is implemented in Golang and uses MongoDB to store Pokémon data.

### Running the server

Create a `.env` file in the root of the server directory with the following:

```
PORT=8080
MONGODB_URI=mongodb://localhost:27017/your-database-name
```

You can run the following commands to start the server:

```bash
go install
go run main.go
```

The server will start on port 8080.

### API

The server provides the following endpoints:

- `GET /api/pokemon`: Returns a list of all Pokémon.
- `GET /api/pokemon/:id`: Returns a single Pokémon by ID.
- `GET /api/pokemon/search?q=...`: Returns a list of Pokémon that match the search query.
- `POST /api/pokemon`: Creates a new Pokémon.
- `PUT /api/pokemon/:id`: Updates a Pokémon by ID.
- `DELETE /api/pokemon/:id`: Deletes a Pokémon by ID.

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
