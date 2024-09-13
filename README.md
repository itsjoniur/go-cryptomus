# Cryptomus API Go Wrapper

This repository contains a Go wrapper for the Cryptomus API, a crypto payment gateway. This wrapper simplifies the process of integrating Cryptomus functionality into your Go projects.

## Features

- Easy-to-use Go interface for Cryptomus API
- Supports payment operations
- Handles static wallet functionalities
- Supports refund operations
- Supports resending webhook requests
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
    httpClient := http.Client{}
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
}
```

## API Coverage

This wrapper currently supports the following Cryptomus API functionalities:

- Payment operations
- Static wallet operations
- Refund operations
- Resending webhook requests

## Configuration

The wrapper can be configured with the following options:

- Merchant ID (required)
- Payment API Key (required)
- Payout API Key (optional)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Disclaimer

This is an unofficial wrapper for the Cryptomus API. It is not affiliated with or endorsed by Cryptomus.

## Support

If you encounter any problems or have any questions, please open an issue in this repository.