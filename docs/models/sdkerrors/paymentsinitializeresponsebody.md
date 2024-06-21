# PaymentsInitializeResponseBody

The payment operation cannot complete


## Supported Types

### GenericError

```go
paymentsInitializeResponseBody := sdkerrors.CreatePaymentsInitializeResponseBodyGenericError(components.GenericError{/* values here */})
```

### FieldError

```go
paymentsInitializeResponseBody := sdkerrors.CreatePaymentsInitializeResponseBodyFieldError(components.FieldError{/* values here */})
```

### CartError

```go
paymentsInitializeResponseBody := sdkerrors.CreatePaymentsInitializeResponseBodyCartError(components.CartError{/* values here */})
```

### CreditCardError

```go
paymentsInitializeResponseBody := sdkerrors.CreatePaymentsInitializeResponseBodyCreditCardError(components.CreditCardError{/* values here */})
```

