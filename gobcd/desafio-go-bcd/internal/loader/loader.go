package loader

import "app/internal"

type LoaderCustomer interface {
	Load() (c []internal.Customer, err error)
}

type LoaderInvoice interface {
	Load() (i []internal.Invoice, err error)
}

type LoaderProduct interface {
	Load() (p []internal.Product, err error)
}

type LoaderSale interface {
	Load() (s []internal.Sale, err error)
}
