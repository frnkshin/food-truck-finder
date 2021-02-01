## Getting Started

To get this running, follow the following steps

### Prerequisites

* go (Please follow the instructions from https://golang.org/doc/install to get latest go)

### Installation

* Clone this repo
    ```sh
    git clone https://github.com/frnkshin/food-truck-finder
    ```
* cd to cmd/food-truck-finder
    ```sh
    cd food-truck-finder/cmd/food-truck-finder
    ```
* Build the console application
    ```sh
    go build
    ```
* Run the application
    ```sh
    # To get help messages
    ./food-truck-finder

    # To get open food trucks (by default paginated at 10 listings)
    ./food-truck-finder find now

    # To see all open food trucks
    ./food-truck-finder find all
    ```

### Tips
* To sort desc
    ```sh
    ./food-truck-finder --sort=false find now
    ```   
* To change api endpoint
    ```sh
    ./food-truck-finder --url <URL-VALUE> find now
    ```
* To change number of listings
    ```sh
    ./food-truck-finder --limit <LIMIT> find now
    ```
