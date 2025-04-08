# Cryptomus API Go Wrapper

This repository contains an **unofficial Go wrapper** for the Cryptomus API, a crypto payment gateway. This wrapper simplifies the process of integrating Cryptomus functionality into your Go projects.

## Features

- Easy-to-use Go interface for Cryptomus API
- Supports payment and payout operations
- Handles static wallet functionalities
- Supports refund operations
- Supports resending and testing webhook requests
- Supports verifying signature
- Provides strongly typed responses

## Installation

To install the Cryptomus API Go wrapper, use `go get`:

```
go get github.com/itsjoniur/go-cryptomus
```

## Usage

Here's a quick example of how to use the wrapper:

```go

import (
    "fmt"
    "github.com/itsjoniur/go-cryptomus"
)

func main() {
    httpClient := http.DefaultClient
    client := cryptomus.New(httpClient, "your-merchant-id", "your-payment-api-key", "your-payout-api-key")
    
    // Create an invoice
    invoiceReq := &cryptomus.InvoiceRequest{
        Amount: "10",
        Currency: "USD",
        OrderId: "your-order-id",
        InvoiceRequestOptions: &cryptomus.invoiceRequestOptions{
            Network: "tron",
            UrlCallback: "https://yourdomain.com/callback"
        },
    }
    invoice, err := cryptomus.CreateInvoice(invoiceReq)
    if err != nil {
        // Handle error
    }
	
    fmt.Printf("Invoice created: %+v\n", invoice)
	
    // Verify webhook
    var body []byte
	
    webhook, err := client.ParseWebhook(body, true)
    if err != nil { 
        switch {
        default:
            // Handle default error
        case errors.Is(err, cryptomus.ErrInvalidSign):
            // Handle invalid sign error
        case errors.Is(err, cryptomus.ErrMissingSign):
            // Handle missing sign error
        }
    }

    fmt.Printf("Verified webhook: %+v\n", webhook)
}
```

## API Coverage

This wrapper currently supports the following Cryptomus API functionalities:

- Payment operations
- Static wallet operations
- Refund operations
- Resending webhook requests

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
