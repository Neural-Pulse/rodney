# SSH Sniffer in Go

This project is a simple SSH sniffer written in Go, capable of detecting SSH connection attempts on all ports, not just the standard port 22. It uses raw sockets to capture TCP packets and identifies SSH connections by the protocol's signature in the payload.

## Prerequisites

To run this sniffer, you will need:

- Go installed on your system (version 1.13 or higher recommended).
- Superuser (root) permissions to run the sniffer, as it requires access to raw sockets.

## Installation

Clone this repository to your local environment using:

`bash git clone https://your-repository/ssh-sniffer.git cd ssh-sniffer`

## Usage

To start the sniffer, navigate to the project directory and execute:

`bash sudo go run sshSniffer.go`

The sniffer will begin monitoring all network interfaces for SSH connection attempts. When an SSH connection is detected, a message will be printed to the console.

## Testing the Sniffer

To test the sniffer, you can initiate an SSH session from another machine or from the same system to any destination. If you do not have a remote machine available, set up an SSH server on your system or use a Docker container for this purpose.

Example SSH command:

`bash ssh user@localhost -p 2222`

Replace `user` with your username and `2222` with the port your SSH server is listening on if you are using a non-standard port.

## Contributions

Contributions are welcome! To contribute, please fork the repository, make your changes, and submit a Pull Request.

## License

This project is distributed under the MIT License. See the `LICENSE` file for more details.
