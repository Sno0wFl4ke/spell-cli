# spell-cli ✨
A command line tool for small dev oprations.

### Command list
- `spell` - Show the list of commands

#### API module
- `spell api GET <url>` - GET request to the specified URL
- `spell api POST <url> <json>` - POST request to the specified URL with a json body
- `spell api PATCH <url> <json>` - PATCH request to the specified URL with a json body
- `spell api DELETE <url>` - DELETE request to the specified URL

#### JSON module
- `spell json read/r <file>` - Read a json file in your terminal
- `spell json write/w <file> <json>` - Write a json file in your terminal
  - Example: `spell json write users.json users.elliot.job "Engineer"`

#### ENV module
- `spell env get <file> <key>` - Read a key from a .env file
- `spell env set <file> <key> <value>` - Write a key to a .env file
- `spell env list <file>` - List all keys in a .env file

more to come soon...


