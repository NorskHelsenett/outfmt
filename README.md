# outfmt

outfmt is a library which can take in any generic data and print it out. Either in a table view, json or yaml. Outfmt is very practical in CLI usecases.

## examples

some examples can be found under the _examples directory

```go
output, err := outfmt.Format($MY_GENERIC_OBJECT, &outfmt.Config{
    Format: outfmt.OutputFormatTable,
})
if err != nil {
    fmt.Fprintf(os.Stderr, "could not format data: %s", err.Error())
    return
}
```

## registering types.

when wanting to output structs as a table the struct firstly needs to be registed. The registration process also defined
the fields which will be printed and the names they will get. Printing happens on a whitelist and not a blacklist like the json
and yaml printers.

```go

type Data struct {
    Id     string
    Name   string
    Active bool
}

func init() {
    outfmt.Register(Data{}, &outfmt.Spec{
        "default": {
            {"ID", "Id"},
            {"NAME", "Name"},
        },
        "wide": {
            {"ID", "Id"},
            {"NAME", "Name"},
            {"ACTIVE", "Active"},
        },
    })
}

// output
-------------------------------------
ID      NAME  
12323   cool  
534324  cool  
1gerfs  cool 
```

Now, when formatting the struct as a table outfmt will only print the Id and Names field of the struct. The `outfmt.Spec` consists of a 
map, this allows us to create certain conditions for what fields outfmt will print. If not being told otherwise outfmt will choose the 
default condition.

```go
// output
-------------------------------------
ID      NAME  ACTIVE  
12323   cool  true    
534324  cool  false   
1gerfs  cool  true 
```

### changing the condition 

when specifying the Condition output format we must pass the name of the condition in the AdditionalField property of the `outfmt.Config`

## formats

- JSON 
- YAML
- Table
- Field
- Condition

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

output, err := outfmt.Format($MY_GENERIC_OBJECT, &outfmt.Config{
    Format: outfmt.OutputFormatField,
    AdditionalField: "Names.First"
})
if err != nil {
    fmt.Fprintf(os.Stderr, "could not format data: %s", err.Error())
    return
}

// output
-------------------------------------
John
Dick
```

here only the persons first name will be printed. AddtionalField can also be a comma seperated list of fields. When using a comma seperated
list this is the output we get.

```go
// output
-------------------------------------
John Deere
Dick McDonald
```