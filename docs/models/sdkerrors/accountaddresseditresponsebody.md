# AccountAddressEditResponseBody

The address is invalid and cannot be added, or some other error has occurred


## Supported Types

### GenericError

```go
accountAddressEditResponseBody := sdkerrors.CreateAccountAddressEditResponseBodyGenericError(components.GenericError{/* values here */})
```

### FieldError

```go
accountAddressEditResponseBody := sdkerrors.CreateAccountAddressEditResponseBodyFieldError(components.FieldError{/* values here */})
```

