
<div align="center">
  <img src="https://github.com/user-attachments/assets/6b5aae46-2ac0-4c2f-a98c-675f2bf02350"  alt="Cloud Cup">
     <h3>Cloud Cup Cup CLI Tool </h3>
  <p><strong>Cup CLI tool to manage the <a href="https://github.com/aliamerj/cloud-cup">Cloud-Cup</a> reverse proxy server.</strong></p>
</div>


**Cup** is a command-line interface (CLI) tool designed to manage and interact with Cloud-Cup, providing an easy way to monitor, configure, and maintain your Cloud-Cup reverse proxy instance.

The CLI includes commands to check the status, view configurations, test backends and routes, and apply new configurations.

## Usage

Here are the commands you can use with cup-cli:
1. Shows whether the Cloud-Cup application is running or not.
```bash
cup show status
```
2. Displays the current configuration the application is running with.
```bash
cup show config
```
4. Checks the health of all backend servers.
```bash
cup check
```
6. Checks the health of a specific backend server by its address.
```bash
cup check -b <backend address>
```
8. Checks whether the specified route is functioning properly.
```bash
cup check -r <route name>
```
10. Applies a new configuration file and performs a hot reload of the application with the changes.
```bash
cup apply -f <file path>
```

## Contributing

We welcome contributions to enhance cup-cli. Feel free to open issues, suggest new features, or submit pull requests!
