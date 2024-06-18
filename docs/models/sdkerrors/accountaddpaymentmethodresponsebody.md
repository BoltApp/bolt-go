# AccountAddPaymentMethodResponseBody

The payment method is invalid and cannot be added, or some other error has occurred


## Supported Types

### GenericError

```go
accountAddPaymentMethodResponseBody := sdkerrors.CreateAccountAddPaymentMethodResponseBodyGenericError(components.GenericError{/* values here */})
```

### FieldError

```go
accountAddPaymentMethodResponseBody := sdkerrors.CreateAccountAddPaymentMethodResponseBodyFieldError(components.FieldError{/* values here */})
```

### CreditCardError

```go
accountAddPaymentMethodResponseBody := sdkerrors.CreateAccountAddPaymentMethodResponseBodyCreditCardError(components.CreditCardError{/* values here */})
```

