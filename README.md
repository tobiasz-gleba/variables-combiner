# Variables Combiner CLI

## This tool will render your configuration from scoped variable groups that you can decide on and provide from various sources.

In pipeline often you have multiple scopes of variables. This tool will render you final configuration file.

### Concepts

- **scope** - Probably you have multiple scopes of configuration `variables: values`. 
Example: `global` < `environment` < `application specyfic values`
- **variable source** - Every scope have to read variables from some source/s.
- **configuration output** - The desired output for your configuration
- **configuration template** - Often you don't want to have all variables combined from sources. You can template your configuration, to get exact output, without all variables, that will not be not consumed by final application.

### Supported variables sources

1. Yaml

### Supported variables outputs

1. Yaml

### Contributing