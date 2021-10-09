package main
import (
    "net/http"

    "github.com/gin-gonic/gin"
"context"
    "fmt"
    "time"
 
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/mongo/readpref"
)
func close(client *mongo.Client, ctx context.Context,
           cancel context.CancelFunc){
            
    // CancelFunc to cancel to context
    defer cancel()
     
    // client provides a method to close
    // a mongoDB connection.
    defer func(){
     
        // client.Disconnect method also has deadline.
        // returns error if any,
        if err := client.Disconnect(ctx); err != nil{
            panic(err)
        }
    }()
}
func connect(uri string)(*mongo.Client, context.Context,
                          context.CancelFunc, error) {
                           
    // ctx will be used to set deadline for process, here
    // deadline will of 30 seconds.
    ctx, cancel := context.WithTimeout(context.Background(),
                                       30 * time.Second)
     
    // mongo.Connect return mongo.Client method
    client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
    return client, ctx, cancel, err
}
 
// This is a user defined method that accepts
// mongo.Client and context.Context
// This method used to ping the mongoDB, return error if any.
func ping(client *mongo.Client, ctx context.Context) error{
 
    // mongo.Client has Ping to ping mongoDB, deadline of
    // the Ping method will be determined by cxt
    // Ping method return error if any occored, then
    // the error can be handled.
    if err := client.Ping(ctx, readpref.Primary()); err != nil {
        return err
    }
    fmt.Println("connected successfully")
    return nil
}
// album represents data about a record album.
type album struct {
        ID     string  `json:"id"`
        NAME  string  `json:"Name"`
        EMAIL string  `json:"Email"`
        PASSWORD  float64 `json:"Password"`
}
// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
        c.IndentedJSON(http.StatusOK, albums)
}
// albums slice to seed record album data.
var albums = []album{
        {ID: "1", NAME: "ABC", EMAIL: "abc@gmail.com", PASSWORD: 2334},
        {ID: "2", NAME: "DEF", EMAIL: "def@gamil.com", PASSWORD: 4456},
        {ID: "3", NAME: "HIJ", EMAIL: "hij@gmail.com", PASSWORD: 3456},
}
func main() {
 client, ctx, cancel, err := connect("mongodb://localhost:27017")
    if err != nil
    {
        panic(err)
    }
     
    // Release resource when the main
    // function is returned.
    defer close(client, ctx, cancel)
     
    // Ping mongoDB with Ping method
    ping(client, ctx)
        router := gin.Default()
        router.GET("/albums", getAlbums)

        router.Run("localhost:8080")
}

