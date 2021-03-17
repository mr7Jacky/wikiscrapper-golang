# Wikipedia Quick View

#### Introduction

This project focuses on create a quick view of terminology
by searching in wikipedia. The GUI is a web-based html file 
that support searching the top relevant contents from wikipedia

#### Project structure
1. All the GoLang file ending with `.go` is the source code
2. [go.mod](go.mod) contains the all the required package needs to be installed
3. [Environment file](assets/.env) contains necessary configuration information that will be used by program
4. [HTML file](assets/index.html) the web-based GUI page
5. [CSS file](assets/style.css) includes the aesthetic design for GUI

#### How to use
1. Compile the source code
    ```bash
    go build -o Server *.go
    ```
2. Run the server
    ```bash
    ./Server
    ```
3. Open your browser and entering 
    ```
    http://localhost:<port number>/
    ```
    Default port number is set to 3000, so by default using the url:
    ```
    http://localhost:3000/
    ```
4. In the text box, entering anything you want to search and press `ENTER`
5. The default number of articles is set to 5, you can change [here](wiki.go)
6. Client the title of any article for more details (jump to wiki page)

#### Reference
1. [Colly - Scraping Framework for Gophers](http://go-colly.org/)
2. [Go Doc - Goquery](https://pkg.go.dev/github.com/PuerkitoBio/goquery)
3. [Go Doc - Colly](https://pkg.go.dev/github.com/gocolly/colly)
4. [Freshman - How to build your first web application with Go](https://freshman.tech/web-development-with-go/)