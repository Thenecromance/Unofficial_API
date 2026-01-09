# API Auto generated tool

This is a very interesting little program. With the powerful Go-template, it can help you automatically generate Blizzard API client code.

Theoretically, this program can generate the latest client code at any time based on the latest definition of the Blizzard API.

The structure of the return value will also be automatically generated through github.com/twpayne/go-jsonstruct. (You need to have the client ID and secret of the Blizzard API.)

This tool will automatically generate the following contents:
- API call functions
- Request parameter structures
- Return value structures - Routers(for gins so far)


You can also choose to generate APIs of different categories [main.go](./main.go)

## Extra files 
tools will auto generate ./params.json, which you can modify different parameters for requesting response to generate to models 

## Want to support other languages? 
1. just rewrite all the api.tmpl, model.tmpl
2. import json converter to target language (ex:quicktype)
3. run this tool, done 