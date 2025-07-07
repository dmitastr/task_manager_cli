<a id="readme-top"></a>


<!-- PROJECT LOGO -->
<br />
<div align="center">

  <h3 align="center">Task manager CLI</h3>

</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

Simple CLI tool for task management written in pure GO.

<p align="right">(<a href="#readme-top">back to top</a>)</p>



### Built With

* [![go][golang]][golang-url]

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- GETTING STARTED -->
## Getting Started

This is an example of how you may give instructions on setting up your project locally.
To get a local copy up and running follow these simple example steps.

### Prerequisites

Install go from official website.


### Installation


2. Clone the repo
   ```sh
   git clone https://github.com/dmastr/task_manager_cli.git
   ```
3. Create file for the database
   ```sh
   echo "{}" > ./internal/database/jsonstorage/data/db.json
   ```
3. Build
   ```sh
   go build -o task_manager.exe ./cmd/task_manager/main.go
   ```
4. Run it
   ```sh
   task_manager.exe
   ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- USAGE EXAMPLES -->
## Usage

Accepted commands:
* `list` list all tasks
* `add -title <task title>` add new task
* `complete -id <task_id>` mark task as completed
* `complete -id <task_id> -u`  unmark task as completed

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- LICENSE -->
## License

Distributed under the Unlicense License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>





<!-- MARKDOWN LINKS & IMAGES -->
[golang]: https://img.shields.io/badge/GOLANG-grey?style=for-the-badge&logo=go
[golang-url]: https://go.dev/