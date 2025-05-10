# Package Name: metadata

## Usage

Below are examples of how to use the functions available in the <b>metadata</b> package:

<b>Type Metadata</b>

```go
type Metadata struct {
	DeviceLang      constant.DeviceLang
	AppVersion      *version.Version
	OSVersion       *version.Version
	DeviceType      string
	DeviceUUID      string
	RequestId       string
	UserId          string
	Token           string
	DeviceTime      *time.Time
	DeviceId        string
	Manufacturer    string
	DeviceModel     string
	OperatingSystem *constant.OperatingSystem
}
```

1. Function `MakeMetadataFromCtx(ctx context.Context) (meta Metadata, err error)`

   ```go
   import (
       "context"
       "fmt"
       "log"

       commMetadata "github.com/LukmanulHakim18/core/metadata"
       "google.golang.org/grpc/metadata"
   )

   func main() {
       ctx := metadata.NewIncomingContext(context.Background(), metadata.MD{
           "Accept-Language": {"EN"},
           "App-Version": {"v1.0.0"},
           "Device-Type": {"android"},
           "Device-Uuid": {"60e02f4d-3c9e-4a61-b265-7eac459e6fa4"},
           "Request-Id": {"12345"},
           "User-Id": {"1234"},
           "Token": {"token-example"},
       })

       md, err := commMetadata.MakeMetadataFromCtx(ctx)
       if err != nil {
           log.Fatalf(err.Error())
       }

       fmt.Printf("metadata: %+v", md)
       // output
       // metadata: {
       //     DeviceLang:EN
       //     AppVersion:1.0.0
       //     OSVersion:<nil>
       //     DeviceType:android DeviceUUID:60e02f4d-3c9e-4a61-b265-7eac459e6fa4
       //     RequestId:12345
       //     UserId:1234
       //     Token:token-example
       //     DeviceTime:<nil>
       //     DeviceId:
       //     Manufacturer:
       //     DeviceModel:
       //     OperatingSystem:<nil>
       // }
   }
   ```

2. Function `GetMetaDataFromContext(ctx context.Context) Metadata`

   ```go
   import (
       "context"
       "fmt"

       commMetadata "github.com/LukmanulHakim18/core/metadata"
       "google.golang.org/grpc/metadata"
   )

   func main() {
       ctx := metadata.NewIncomingContext(context.Background(), metadata.MD{
           "Accept-Language": {"EN"},
           "App-Version":     {"v1.0.0"},
           "Device-Type":     {"android"},
           "Device-Uuid":     {"60e02f4d-3c9e-4a61-b265-7eac459e6fa4"},
           "Request-Id":      {"12345"},
           "User-Id":         {"1234"},
           "Token":           {"token-example"},
       })

       md := commMetadata.GetMetaDataFromContext(ctx)
       fmt.Printf("metadata: %+v", md)
       // output
       // metadata: {
       // 	DeviceLang:EN
       // 	AppVersion:1.0.0
       // 	OSVersion:<nil>
       // 	DeviceType:android
       // 	DeviceUUID:60e02f4d-3c9e-4a61-b265-7eac459e6fa4
       // 	RequestId:
       // 	UserId:1234
       // 	Token:token-example
       // 	DeviceTime:<nil>
       // 	DeviceId:
       // 	Manufacturer:
       // 	DeviceModel:
       // 	OperatingSystem:<nil>
       // }
   }
   ```

3. Function `GetDeviceLanguageFromCtx(ctx context.Context) constant.DeviceLang`

   ```go
   import (
       "context"
       "fmt"

       commMetadata "github.com/LukmanulHakim18/core/metadata"
       "google.golang.org/grpc/metadata"
   )

   func main() {
       ctx := metadata.NewIncomingContext(context.Background(), metadata.MD{
           "Accept-Language": {"EN"},
           "App-Version":     {"v1.0.0"},
           "Device-Type":     {"android"},
           "Device-Uuid":     {"60e02f4d-3c9e-4a61-b265-7eac459e6fa4"},
           "Request-Id":      {"12345"},
           "User-Id":         {"1234"},
           "Token":           {"token-example"},
       })

       deviceLang := commMetadata.GetDeviceLanguageFromCtx(ctx)
       fmt.Printf("device language: %s", deviceLang)
       // output
       // device language: EN
   }
   ```

4. Function `GetDeviceVersionFromCtx(ctx context.Context) (ver *version.Version, err error)`

   ```go
   import (
       "context"
       "fmt"
       "log"

       commMetadata "github.com/LukmanulHakim18/core/metadata"
       "google.golang.org/grpc/metadata"
   )

   func main() {
       ctx := metadata.NewIncomingContext(context.Background(), metadata.MD{
           "Accept-Language": {"EN"},
           "App-Version":     {"v1.0.0"},
           "Device-Type":     {"android"},
           "Device-Uuid":     {"60e02f4d-3c9e-4a61-b265-7eac459e6fa4"},
           "Request-Id":      {"12345"},
           "User-Id":         {"1234"},
           "Token":           {"token-example"},
       })

       deviceVersion, err := commMetadata.GetDeviceVersionFromCtx(ctx)
       if err != nil {
           log.Fatalf(err.Error())
       }
       fmt.Printf("device version: %s", deviceVersion)
       // output
       // device version: 1.0.0
   }
   ```

5. Function `GetDeviceTypeFromCtx(ctx context.Context) (string, error)`

   ```go
   import (
       "context"
       "fmt"
       "log"

       commMetadata "github.com/LukmanulHakim18/core/metadata"
       "google.golang.org/grpc/metadata"
   )

   func main() {
       ctx := metadata.NewIncomingContext(context.Background(), metadata.MD{
           "Accept-Language": {"EN"},
           "App-Version":     {"v1.0.0"},
           "Device-Type":     {"android"},
           "Device-Uuid":     {"60e02f4d-3c9e-4a61-b265-7eac459e6fa4"},
           "Request-Id":      {"12345"},
           "User-Id":         {"1234"},
           "Token":           {"token-example"},
       })

       deviceType, err := commMetadata.GetDeviceTypeFromCtx(ctx)
       if err != nil {
           log.Fatalf(err.Error())
       }
       fmt.Printf("device type: %s", deviceType)
       // output
       // device type: android
   }
   ```

6. Function `GetDeviceUuid(ctx context.Context) (string, error)`

   ```go
   import (
       "context"
       "fmt"
       "log"

       commMetadata "github.com/LukmanulHakim18/core/metadata"
       "google.golang.org/grpc/metadata"
   )

   func main() {
       ctx := metadata.NewIncomingContext(context.Background(), metadata.MD{
           "Accept-Language": {"EN"},
           "App-Version":     {"v1.0.0"},
           "Device-Type":     {"android"},
           "Device-Uuid":     {"60e02f4d-3c9e-4a61-b265-7eac459e6fa4"},
           "Request-Id":      {"12345"},
           "User-Id":         {"1234"},
           "Token":           {"token-example"},
       })

       deviceId, err := commMetadata.GetDeviceUuid(ctx)
       if err != nil {
           log.Fatalf(err.Error())
       }
       fmt.Printf("device uuid: %s", deviceId)
       // output
       // device uuid: 60e02f4d-3c9e-4a61-b265-7eac459e6fa4
   }
   ```

7. Function `GetRequestId(ctx context.Context) (string, error)`

   ```go
   import (
       "context"
       "fmt"
       "log"

       commMetadata "github.com/LukmanulHakim18/core/metadata"
       "google.golang.org/grpc/metadata"
   )

   func main() {
       ctx := metadata.NewIncomingContext(context.Background(), metadata.MD{
           "Accept-Language": {"EN"},
           "App-Version":     {"v1.0.0"},
           "Device-Type":     {"android"},
           "Device-Uuid":     {"60e02f4d-3c9e-4a61-b265-7eac459e6fa4"},
           "Request-Id":      {"12345"},
           "User-Id":         {"1234"},
           "Token":           {"token-example"},
       })

       reqId, err := commMetadata.GetRequestId(ctx)
       if err != nil {
           log.Fatalf(err.Error())
       }
       fmt.Printf("request id: %s", reqId)
       // output
       // request id: 12345
   }
   ```

8. Function `GetUserId(ctx context.Context) (string, error)`

   ```go
   import (
       "context"
       "fmt"
       "log"

       commMetadata "github.com/LukmanulHakim18/core/metadata"
       "google.golang.org/grpc/metadata"
   )

   func main() {
       ctx := metadata.NewIncomingContext(context.Background(), metadata.MD{
           "Accept-Language": {"EN"},
           "App-Version":     {"v1.0.0"},
           "Device-Type":     {"android"},
           "Device-Uuid":     {"60e02f4d-3c9e-4a61-b265-7eac459e6fa4"},
           "Request-Id":      {"12345"},
           "User-Id":         {"1234"},
           "Token":           {"token-example"},
       })

       userId, err := commMetadata.GetUserId(ctx)
       if err != nil {
           log.Fatalf(err.Error())
       }
       fmt.Printf("user id: %s", userId)
       // output
       // user id: 1234
   }
   ```

9. Function `GetTokenFromCtx(ctx context.Context) (string, error)`

   ```go
   import (
       "context"
       "fmt"
       "log"

       commMetadata "github.com/LukmanulHakim18/core/metadata"
       "google.golang.org/grpc/metadata"
   )

   func main() {
       ctx := metadata.NewIncomingContext(context.Background(), metadata.MD{
           "Accept-Language": {"EN"},
           "App-Version":     {"v1.0.0"},
           "Device-Type":     {"android"},
           "Device-Uuid":     {"60e02f4d-3c9e-4a61-b265-7eac459e6fa4"},
           "Request-Id":      {"12345"},
           "User-Id":         {"1234"},
           "Token":           {"token-example"},
       })

       token, err := commMetadata.GetTokenFromCtx(ctx)
       if err != nil {
           log.Fatalf(err.Error())
       }
       fmt.Printf("token: %s", token)
       // output
       // token: token-example
   }
   ```
