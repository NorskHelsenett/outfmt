# outfmt

outfmt is a library which can take in any generic data and print it out. Either in a table view, json or yaml. Outfmt is very practical in CLI usecases.

## examples

some examples can be found under the _examples directory

```go
output, err := outfmt.Format(MY_GENERIC_DATA, &outfmt.Config{
    Format: outfmt.OutputFormatTable,
})
if err != nil {
    fmt.Fprintf(os.Stderr, "could not format data: %s", err.Error())
    return
}
```


## tags

the table view depends on tags to determine which fields to print. Example

```go
type Data struct {
    Id string
    Name string `outfmt:"NAME"`
}
```

in this example only the Name field will be printed when using a table view. Tagging is only neccesary for table view.

## formats

- JSON 
- YAML
- Table
- Field

## field

When wanting to only print one field pass the path of the field you want to print using the AdditionalField property on config.

```go
type Names struct {
    First string
    Last string
}

type Person struct {
    Names Names
    Age int
}

output, err := outfmt.Format(MY_GENERIC_DATA, &outfmt.Config{
    Format: outfmt.OutputFormatField,
    AdditionalField: "Names.First"
})
if err != nil {
    fmt.Fprintf(os.Stderr, "could not format data: %s", err.Error())
    return
}
```

here only the persons first name will be printed.
