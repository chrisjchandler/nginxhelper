package main

import (
    "fmt"
    "os"
    "text/template"
)

type BackendServer struct {
    Host string
    Port int
}

type NginxConfig struct {
    ServerName    string
    ListenPort    int
    BackendServers []BackendServer
    SSLCertPath   string
    SSLKeyPath    string
}

func main() {
    // User prompts for common junk
    var config NginxConfig
    fmt.Print("What's the darn Server Name? ")
    fmt.Scanln(&config.ServerName)
    fmt.Print("And the Listen Port, because why not? ")
    fmt.Scanln(&config.ListenPort)
    fmt.Print("SSL Certificate Path - because security, duh! ")
    fmt.Scanln(&config.SSLCertPath)
    fmt.Print("SSL Key Path - who even cares? ")
    fmt.Scanln(&config.SSLKeyPath)

    // Prompt for the number of backend servers
    var numServers int
    fmt.Print("How many of those pesky backend servers do you have? ")
    fmt.Scanln(&numServers)

    // Gather configuration for each backend server
    for i := 1; i <= numServers; i++ {
        var backendServer BackendServer
        fmt.Printf("Backend Server %d Host (or just press Enter to finish): ", i)
        _, err := fmt.Scanln(&backendServer.Host)
        if err != nil || backendServer.Host == "" {
            break
        }
        fmt.Printf("Backend Server %d Port - if you must: ", i)
        fmt.Scanln(&backendServer.Port)
        config.BackendServers = append(config.BackendServers, backendServer)
    }

    // Define the Nginx configuration template
    nginxTemplate := `
server {
    listen {{.ListenPort}};
    server_name {{.ServerName}};

    ssl_certificate {{.SSLCertPath}};
    ssl_certificate_key {{.SSLKeyPath}};

    {{range $index, $server := .BackendServers}}
    location /server{{$index}} {
        proxy_pass http://{{$server.Host}}:{{$server.Port}};
    }
    {{end}}
}

http {
    upstream backend {
        {{range $server := .BackendServers}}
        server {{$server.Host}}:{{$server.Port}};
        {{end}}
    }
}
`

    // Parse the template
    tmpl, err := template.New("nginx").Parse(nginxTemplate)
    if err != nil {
        fmt.Println("Ugh! Error parsing the template:", err)
        return
    }

    // Create the output file
    outputFile, err := os.Create("nginx.conf")
    if err != nil {
        fmt.Println("Error creating output file. Seriously? Come on:", err)
        return
    }
    defer outputFile.Close()

    // Execute the template and write the output to the file
    err = tmpl.Execute(outputFile, config)
    if err != nil {
        fmt.Println("Error executing the template. This is just great:", err)
        return
    }

    fmt.Println("Nginx configuration file 'nginx.conf' generated. Finally!")
}
