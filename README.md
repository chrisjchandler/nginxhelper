# Nginx Config Generator

## Overview

tool designed to make your life a bit easier when it comes to creating Nginx configuration files. No more manual editing and typos!

## What Does It Do?

This simple Go-based tool prompts you for the necessary information and generates an Nginx configuration file for you. It can handle basic server configurations, SSL certificates, and even multiple backend servers.

## How to Use

1. Clone this repository to your local machine.
2. Open your terminal and navigate to the project directory.
3. Run `go run nginxhelper.go` to start the configuration process.
4. Follow the prompts, providing the requested information.
5. The tool will generate an Nginx configuration file named `nginx.conf` in the same directory.

## Multiple Backend Servers

If you have multiple backend servers to configure, the tool will keep asking for server details until you're done. Simply press Enter to move on to the next section.

## License

This tool is provided under the MIT License, so you can use it freely in your projects. Refer to the [LICENSE](LICENSE) file for more details.

## Feedback and Contributions

Feel free to provide feedback or contribute to this project. We're open to suggestions and improvements to make this tool even more useful.

