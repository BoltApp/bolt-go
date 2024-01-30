# Amount

A monetary amount, i.e. a base unit amount and a supported currency.


## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            | Example                                                                |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `Currency`                                                             | [components.Currency](../../models/components/currency.md)             | :heavy_check_mark:                                                     | A supported currency.                                                  | USD                                                                    |
| `Units`                                                                | *int64*                                                                | :heavy_check_mark:                                                     | A monetary amount, represented in its base units (e.g. USD/EUR cents). | 900                                                                    |