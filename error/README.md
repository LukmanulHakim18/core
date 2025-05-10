# Package Name: error

## Usage

Below are examples of how to use the functions available in the <b>error</b> package:

<b>Type Error</b>

```go
type Error struct {
	DeviceLang       constant.DeviceLang    `json:"-"`
	StatusCode       int                    `json:"-"`
	ErrorCode        string                 `json:"error_code"`
	ErrorMessage     string                 `json:"error_message"`
	ErrorField       string                 `json:"error_field,omitempty"`
	LocalizedMessage Message                `json:"localized_message"`
	Data             map[string]interface{} `json:"data,omitempty"`
	ErrorData        interface{}            `json:"error_data,omitempty"`
}
```

1. Function `NewError(code, message, english, indonesia string) *Error`

   ```go
   import (
       "fmt"

       commErrors "github.com/LukmanulHakim18/core/error"
   )

   func main() {
       err := commErrors.NewError(
           "your-error-code",
           "your-error-message",
           "your-error-message-in-english",
           "your-error-message-in-Indonesian")

       fmt.Printf("error: %+v", *err)
       // output
       // error: {
       //     DeviceLang:
       //     StatusCode: 0
       //     ErrorCode: your-error-code
       //     ErrorMessage: your-error-message
       //     ErrorField: LocalizedMessage:
       //         {
       //             English: your-error-message-in-english
       //             Indonesia: your-error-message-in-Indonesian
       //         }
       //     Data: map[]
       //     ErrorData:<nil>
       // }
   }
   ```

2. Function `NewErrorWithStatus(status int, code, message, english, indonesia string) *Error`

   ```go
   import (
       "fmt"

       commErrors "github.com/LukmanulHakim18/core/error"
   )

   func main() {
       yourErrorStatus := 404
       err := commErrors.NewErrorWithStatus(
           yourErrorStatus,
           "your-error-code",
           "your-error-message",
           "your-error-message-in-english",
           "your-error-message-in-Indonesian")

       fmt.Printf("error: %+v", *err)
       // output
       // error: {
       //     DeviceLang:
       //     StatusCode:1
       //     ErrorCode:your-error-code
       //     ErrorMessage:your-error-message
       //     ErrorField:
       //     LocalizedMessage:{
       //         English:your-error-message-in-english Indonesia:your-error-message-in-Indonesian
       //         }
       //     Data:map[]
       //     ErrorData:<nil>
       // }
   }
   ```

3. Function `GetInvalidParameterMessage(err *Error, parameter ...interface{}) *Error`

   ```go
   import (
       "fmt"

       commErrors "github.com/LukmanulHakim18/core/error"
   )

   func main() {
       err := commErrors.NewError(
           "your-error-code",
           "your-error-message",
           "your-error-message-in-english",
           "your-error-message-in-Indonesian")

       errParam := commErrors.GetInvalidParameterMessage(err, "your-parameter1", "your-parameter2")

       fmt.Printf("error: %+v", *errParam)
       // output
       // error: {
       //     DeviceLang:
       //     StatusCode:0
       //     ErrorCode:your-error-code
       //     ErrorMessage:your-error-message%!(EXTRA []interface {}=[your-parameter1 your-parameter2])
       //     ErrorField:
       //     LocalizedMessage:{
       //         English:your-error-message-in-english%!(EXTRA []interface {}=[your-parameter1 your-parameter2])
       //         Indonesia:your-error-message-in-Indonesian%!(EXTRA []interface {}=[your-parameter1 your-parameter2])
       //     }
       //     Data:map[]
       //     ErrorData:<nil>
       // }
   }
   ```

4. Function `GetCustomMessageWithParameters(err *Error, paramID, paramEN string) *Error`

   ```go
   import (
       "fmt"

       commErrors "github.com/LukmanulHakim18/core/error"
   )

   func main() {
       err := commErrors.NewError(
           "your-error-code",
           "your-error-message",
           "your-error-message-in-English",
           "your-error-message-in-Indonesian")

       errParam := commErrors.GetCustomMessageWithParameters(err, "your-parameter-in-Indonesian", "your-parameter-in-English")

       fmt.Printf("error: %+v", *errParam)
       // output
       // error: {
       //     DeviceLang:
       //     StatusCode:0
       //     ErrorCode:your-error-code
       //     ErrorMessage:your-error-message%!(EXTRA string=your-parameter-in-English)
       //     ErrorField:
       //     LocalizedMessage:{
       //         English:your-error-message-in-English%!(EXTRA string=your-parameter-in-English)
       //         Indonesia:your-error-message-in-Indonesian%!(EXTRA string=your-parameter-in-Indonesian)
       //     }
       //     Data:map[]
       //     ErrorData:<nil>
       // }
   }
   ```

5. function `GetCustomMessageWithArgs(err *Error, args ...interface{}) *Error`

   ```go
   import (
       "fmt"

       commErrors "github.com/LukmanulHakim18/core/error"
   )

   func main() {
       err := commErrors.NewError(
           "your-error-code",
           "your-error-message",
           "your-error-message-in-English",
           "your-error-message-in-Indonesian")

       errArgs := commErrors.GetCustomMessageWithArgs(err, "your-argument1", "your-argument2")

       fmt.Printf("error: %+v", *errArgs)
       // output
       // error: {
       //     DeviceLang:
       //     StatusCode:0
       //     ErrorCode:your-error-code
       //     ErrorMessage:your-error-message%!(EXTRA string=your-argument1, string=your-argument2)
       //     ErrorField:
       //     LocalizedMessage:{
       //         English:your-error-message-in-English%!(EXTRA string=your-argument1, string=your-argument2)
       //         Indonesia:your-error-message-in-Indonesian%!(EXTRA string=your-argument1, string=your-argument2)
       //     }
       //     Data:map[]
       //     ErrorData:<nil>
       // }
   }
   ```

6. Function `(e *Error) Error() string`

   ```go
   import (
       "fmt"

       "github.com/LukmanulHakim18/core/constant"
       commErrors "github.com/LukmanulHakim18/core/error"
   )

   func main() {
       err := commErrors.NewError(
           "your-error-code",
           "your-error-message",
           "your-error-message-in-English",
           "your-error-message-in-Indonesian")

       fmt.Printf("error: %s", err.Error())
       // output
       // error: your-error-message-in-English

       err.DeviceLang = constant.DEVICE_LANG_ID
       fmt.Printf("error: %s", err.Error())
       // output
       // error: your-error-message-in-Indonesian
   }
   ```

7. Function `(e *Error) WithData(data map[string]interface{}) *Error`

   ```go
   import (
       "fmt"

       commErrors "github.com/LukmanulHakim18/core/error"
   )

   func main() {
       err := commErrors.NewError(
           "your-error-code",
           "your-error-message",
           "your-error-message-in-English",
           "your-error-message-in-Indonesian")

       data := map[string]interface{}{
           "data1": "data1",
           "data2": "data2",
           "data3": "data3",
       }
       err = err.WithData(data)

       fmt.Printf("error: %+v", *err)
       // output
       // error: {
       //     DeviceLang:
       //     StatusCode:0
       //     ErrorCode:your-error-code
       //     ErrorMessage:your-error-message
       //     ErrorField:
       //     LocalizedMessage:{
       //         English:your-error-message-in-English
       //         Indonesia:your-error-message-in-Indonesian
       //         }
       //     Data:map[data1:data1 data2:data2 data3:data3]
       //     ErrorData:<nil>
       // }
   }
   ```

8. Function `(e *Error) WithErrorData(errCode string, errdata interface{}) *Error`

   ```go
   import (
       "fmt"

       commErrors "github.com/LukmanulHakim18/core/error"
   )

   func main() {
       err := commErrors.NewError(
           "your-error-code",
           "your-error-message",
           "your-error-message-in-English",
           "your-error-message-in-Indonesian")

       yourErrorCode := "404"
       errdata := "your error data"
       err = err.WithErrorData(yourErrorCode, errdata)

       fmt.Printf("error: %+v", *err)
       // output
       // error: {
       //     DeviceLang:
       //     StatusCode:0
       //     ErrorCode:404
       //     ErrorMessage:your-error-message
       //     ErrorField:
       //     LocalizedMessage:{
       //         English:your-error-message-in-English
       //         Indonesia:your-error-message-in-Indonesian
       //     }
       //     Data:map[]
       //     ErrorData:your error data
       // }
   }
   ```

9. Function `(e *Error) WithParameter(parameter ...interface{}) *Error`

   ```go
   import (
       "fmt"

       commErrors "github.com/LukmanulHakim18/core/error"
   )

   func main() {
       err := commErrors.NewError(
           "your-error-code",
           "your-error-message",
           "your-error-message-in-English",
           "your-error-message-in-Indonesian")

       err = err.WithParameter("your-parameter1", "your-parameter2")
       fmt.Printf("error: %+v", *err)
       // output
       // error: {
       //     DeviceLang:
       //     StatusCode:0
       //     ErrorCode:your-error-code
       //     ErrorMessage:your-error-message%!(EXTRA string=your-parameter1, string=your-parameter2)
       //     ErrorField:
       //     LocalizedMessage:{
       //         English:your-error-message-in-English%!(EXTRA string=your-parameter1, string=your-parameter2)
       //         Indonesia:your-error-message-in-Indonesian%!(EXTRA string=your-parameter1, string=your-parameter2)
       //     }
       //     Data:map[]
       //     ErrorData:<nil>
       // }
   }
   ```

10. Function `(e Error) WithField(field string) *Error`

    ```go
    import (
        "fmt"

        commErrors "github.com/LukmanulHakim18/core/error"
    )

    func main() {
        err := commErrors.NewError(
            "your-error-code",
            "your-error-message",
            "your-error-message-in-English",
            "your-error-message-in-Indonesian")

        err = err.WithField("your-field")
        fmt.Printf("error: %+v", *err)
        // output
        // error: {
        //     DeviceLang:
        //     StatusCode:0
        //     ErrorCode:your-error-code
        //     ErrorMessage:your-error-message
        //     ErrorField:your-field
        //     LocalizedMessage:{
        //         English:your-error-message-in-English
        //         Indonesia:your-error-message-in-Indonesian
        //     }
        //     Data:map[]
        //     ErrorData:<nil>
        // }
    }
    ```

11. Function `(e Error) GrpcCode() codes.Code`

    ```go
    import (
        "fmt"

        commErrors "github.com/LukmanulHakim18/core/error"
    )

    func main() {
        yourErrorStatus := 404
        err := commErrors.NewErrorWithStatus(
            yourErrorStatus,
            "your-error-code",
            "your-error-message",
            "your-error-message-in-english",
            "your-error-message-in-Indonesian")

        fmt.Printf("error code grpc: %d", err.GrpcCode())
        // output
        // error code grpc: 5
    }
    ```

12. Function `(e *Error) BuildError(ctx context.Context) error`

    ```go
    import (
        "context"
        "fmt"

        commErrors "github.com/LukmanulHakim18/core/error"
        "google.golang.org/grpc/metadata"
    )

    func main() {
        yourErrorStatus := 404
        commErr := commErrors.NewErrorWithStatus(
            yourErrorStatus,
            "your-error-code",
            "your-error-message",
            "your-error-message-in-english",
            "your-error-message-in-Indonesian")

        err := commErr.BuildError(context.Background())
        fmt.Printf("error : %+v\n", err)
        // output
        // error : rpc error: code = NotFound desc = your-error-message-in-english

        err = commErr.BuildError(
            metadata.NewIncomingContext(
                context.Background(),
                metadata.MD{"Accept-Language": {"ID"}}),
        )
        fmt.Printf("error : %+v", err)
        // output
        // error : rpc error: code = NotFound desc = your-error-message-in-Indonesian
    }
    ```
