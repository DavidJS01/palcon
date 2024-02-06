<a name="readme-top"></a>

<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/othneildrew/Best-README-Template">
    <img src="https://cdn2.steamgriddb.com/icon/9ce60c64ac4510df68537de96631261f.ico" alt="Logo" width="80" height="80">
  </a>

  <h3 align="center">Palcon</h3>

  <p align="center">
    A Palworld RCON Server Management CLI Tool
    <br />
    <a href="https://github.com/DavidJS01/Palworld/issues">Report Bug</a>
    ·
    <a href="https://github.com/DavidJS01/Palworld/issues">Request Feature</a>
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-palcon">About Palcon</a>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About Palcon
Palcon is a CLI tool for interacting with dedicated Palworld servers using the RCON protocol. It currently supports the below features:

| Command   | Implemented? | Description                                    |
|-----------|--------------|------------------------------------------------|
| ban       |       ✅      | Ban a player given a Palworld UID or Steam ID  |
| broadcast |       ✅      | Broadcast a message to the server              |
| exit      |       ✅      | Save the server, and then shut down the server |
| info      |       ✅      | Get info about the server                      |
| kick      |       ✅      | Kick a player given a Palworld UID or Steam ID |
| save      |       ✅      | Save the server                                |
| start     |       ✅      | Start a server on the **local** computer         |


<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- GETTING STARTED -->
## Getting Started

### Building From Source
```sh
# install Go https://go.dev/doc/install

# clone repo
git clone git@github.com:DavidJS01/palcon.git

# build executable
go build -o palcon.exe
```


### Installation with Go Install

#### Install From Github (Recommended)
```sh
go install github.com/DavidJS01/palcon@latest

> palcon
Palcon is a CLI tool that lets you interact with dedicated palworld servers.
```

#### Install Local Project
After cloning the project locally and changing your terminal's directory into the project's root:
```sh
# clone project
git clone git@github.com:DavidJS01/palcon.git

# change working directory
cd palcon

# install project
go install

# run palcon
> palcon
Palcon is a CLI tool that lets you interact with dedicated palworld servers.
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- USAGE EXAMPLES -->
## Usage
```sh
# CLI format:
# palcon --host <SERVER IP> --port <SERVER PORT> --password <ADMIN PASSWORD> COMMAND

# get info on server (good for checking if server is online)
palcon --host 173.194.0.0 --port 25575 --password PASSWORD info

# start a local Palworld server
palcon start

# save a server
palcon --host 173.194.0.0 --port 25575 --password PASSWORD save

# get help message with ALL commands
palcon
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- ROADMAP -->
## Roadmap

- [] Add file based configuration for default host, port, and password

See the [open issues](https://github.com/DavidJS01/palcon/issues) for a full list of proposed features and known issues.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- CONTRIBUTING -->
## Contributing
Contributions are appreciated.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/featureName`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/featureName`)
5. Open a Pull Request

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- LICENSE -->
## License

Distributed under the GPL V3 License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- CONTACT -->
## Contact

David Shipman - chemdev.me@protonmail.com

<p align="right">(<a href="#readme-top">back to top</a>)</p>