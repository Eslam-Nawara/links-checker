# Links checker
Links checker is a software that crawl all the websites defined in a config file, validates the links in that file and every internal link that has the same hostname as its parent and report all the broken links to stdout.

### Note
-  Any link that require authentication like `LinkedIn` profiles would be recorded as broken links.
 - Link is considered healthy if it responds to head/get requests with 20X or 30X

## Config file
A configuration file is a `TOML` file that contains a list of all sites need to be checked, categorized in a key-value format.
```toml
[sites]
    [sites.threefold]
        url = 'threefold.io'
    [sites.codescalers]
        url = 'codescalers.com'
```

## How to use 
- Get the package
```sh 
 go get github.com/Eslam-Nawara/linkschecker
```

**Package API:**
|Function | Description |
| :--- | :--- |
| `CheckLinksInFile(configFile string) error` | Crawls a `TOML` file, extract all Links in the file and check if healthy |
| `LinksFromConfig(configFile string) ([]string, error)` | Parses a `TOML` file and convert it into array of strings |

## Usage and Installation:
- Install linkschecker
```sh 
 go install github.com/Eslam-Nawara/linkschecker/cmd/linkschecker@latest
```
- Create a `TOML` file with links to be checked.
- Run the checker	
```sh
 linkschecker --config=<config file>
```
