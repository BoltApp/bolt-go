# GuestPaymentsInitializeResponseBody

The payment operation cannot complete


## Supported Types

### GenericError

```go
guestPaymentsInitializeResponseBody := sdkerrors.CreateGuestPaymentsInitializeResponseBodyGenericError(components.GenericError{/* values here */})
```

### FieldError

```go
guestPaymentsInitializeResponseBody := sdkerrors.CreateGuestPaymentsInitializeResponseBodyFieldError(components.FieldError{/* values here */})
```

### CartError

```go
guestPaymentsInitializeResponseBody := sdkerrors.CreateGuestPaymentsInitializeResponseBodyCartError(components.CartError{/* values here */})
```

### CreditCardError

```go
guestPaymentsInitializeResponseBody := sdkerrors.CreateGuestPaymentsInitializeResponseBodyCreditCardError(components.CreditCardError{/* values here */})
```

