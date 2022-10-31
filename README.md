# Links checker
Links checker is a software that crawl all the websites defined in a config file, validates the links in that file and every internal link that has the same hostname as its parent and report all the broken links to stdout.

**Note**
 Any link that require authentication as `LinkedIn` profiles would be recorded as broken links.
 Make sure to edit the config file to add your website link.

## How to use 
- Clone the repo
```sh
$ git clone https://github.com/Eslam-Nawara/links-checker
```
- Run the program
```sh
$ go run *.go --config=<config file>
```
