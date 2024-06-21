# PaymentsActionResponseBody

An error has occurred, and further details are contained in the response


## Supported Types

### GenericError

```go
paymentsActionResponseBody := sdkerrors.CreatePaymentsActionResponseBodyGenericError(components.GenericError{/* values here */})
```

### FieldError

```go
paymentsActionResponseBody := sdkerrors.CreatePaymentsActionResponseBodyFieldError(components.FieldError{/* values here */})
```

