#### Canonical Terraform Provider

I wanted to explore the JSON schema from Terraform providers and built a canonical example provider. It doesn't do anything, so don't expect too much, but it does enough to explore the JSON schema aspects.

This is a canonical example of a Terraform provider that exercises different schema types including:

```yaml
- set
- map
- list
- string
- int
- float
```

It also exercises nested schemas like:

```yaml
- list
- map
- set
```

The reason for this is to dump the JSON schema in v0.12.12 of Terraform.

#### Instructions

1. Grab this repository

```bash
git clone https://github.com/davedotdev/canonicaltfprovider.git
cd canonicaltfprovider
go mod download
go build
```

2. You'll need Terraform (version 0.12.12 ideally).

```bash
terraform providers schema -json 
```

The terraform command dumps out the JSON schema for the canonical provider which looks like below (prettified). 

```json
{
  "format_version": "0.1",
  "provider_schemas": {
    "canonical": {
      "provider": {
        "version": 0,
        "block": {
          "attributes": {
            "host": {
              "type": "string",
              "required": true
            },
            "password": {
              "type": "string",
              "required": true
            },
            "port": {
              "type": "number",
              "required": true
            },
            "sshkey": {
              "type": "string",
              "required": true
            },
            "username": {
              "type": "string",
              "required": true
            }
          }
        }
      },
      "resource_schemas": {
        "canonical": {
          "version": 0,
          "block": {
            "attributes": {
              "booltest": {
                "type": "bool",
                "required": true
              },
              "floattest": {
                "type": "number",
                "required": true
              },
              "id": {
                "type": "string",
                "optional": true,
                "computed": true
              },
              "inttest": {
                "type": "number",
                "required": true
              },
              "liststringtest": {
                "type": [
                  "list",
                  "string"
                ],
                "required": true
              },
              "maptest": {
                "type": [
                  "map",
                  "string"
                ]
              },
              "stringtest": {
                "type": "string",
                "required": true
              }
            },
            "block_types": {
              "listsettest": {
                "nesting_mode": "list",
                "block": {
                  "attributes": {
                    "setbooltest": {
                      "type": "bool",
                      "required": true
                    },
                    "setfloattest": {
                      "type": "number",
                      "required": true
                    },
                    "setinttest": {
                      "type": "number",
                      "required": true
                    },
                    "setstringtest": {
                      "type": "string",
                      "required": true
                    }
                  }
                },
                "min_items": 1,
                "max_items": 1
              },
              "settest": {
                "nesting_mode": "set",
                "block": {
                  "attributes": {
                    "setbooltest": {
                      "type": "bool",
                      "required": true
                    },
                    "setfloattest": {
                      "type": "number",
                      "required": true
                    },
                    "setinttest": {
                      "type": "number",
                      "required": true
                    },
                    "setstringtest": {
                      "type": "string",
                      "required": true
                    }
                  }
                },
                "min_items": 1,
                "max_items": 1
              }
            }
          }
        }
      }
    }
  }
}
```


