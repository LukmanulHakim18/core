# Http doc

## ClientV2

This lib was integration with APM Http for reporting performance and can trace error.

```go
import (
	"context"
	"net/http"
	"testing"

	commHttp "github.com/LukmanulHakim18/core/http"
)
func RefreshToken(){

   ep := commHttp.NewEndpoint("/token/7364612d31363834", nil, "GET")
   cli := commHttp.NewClient("http://localhost:1407")

   h := http.Header{}
   h.Add("App-Version", "6.2.0")

   httpResp, err := cli.Exec(context.Background(), ep, h, nil)
   if err != nil {
      return err
   }

   var result map[string]interface{}

   if err := commHttp.MappingResponse(httpResp, &result, &result); err != nil {
      return err
   }
   return result
}
```
