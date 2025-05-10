# Package Name: pubsub

## Usage

Below are examples of how to use the functions available in the <b>pubsub</b> package:

1. Function NewPubSub

   ```go
   import (
       "context"
       "fmt"
       "log"

       "github.com/LukmanulHakim18/core/pubsub"
   )

   func main() {
       projectID := "your-project-id"
       client, err := pubsub.NewPubSub(projectID)
       if err != nil {
           log.Fatalf(err.Error())
       }
       defer client.Close()
   }
   ```

2. Function Parse

   ```go
   import (
       "context"
       "fmt"
       "log"

       commonPubsub "github.com/LukmanulHakim18/core/pubsub"
       "cloud.google.com/go/pubsub"
   )

   func main() {
       projectID := "your-project-id"
       client, err := commonPubsub.NewPubSub(projectID)
       if err != nil {
           log.Fatalf(err.Error())
       }
       defer client.Close()

       sub := client.Conn.Subscription("your-subscription-id")
       subExist, err := sub.Exists(context.Background())
       if err != nil {
           log.Fatalf(err.Error())
       }

       if !subExist {
           return
       }

       err = sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
           var reqBase = commonPubsub.NewBaseMsg()
           var req = YourRequestStruct{}
           ctx, err := reqBase.Parse(ctx, msg.Data, &req)
           if err != nil {
               fmt.Printf("parse message error %v\n", err)
               msg.Ack()
               return
           }

           fmt.Printf("request: %v \nmetadata: %v\n", req, reqBase.Metadata)

           md, _ := metadata.FromIncomingContext(ctx)
           fmt.Printf("metadata from context: %v\n", md)
       })

       if err != nil {
           log.Fatalf(err.Error())
       }
   }
   ```
