# docker_container.web:
resource "docker_container" "web" {
    attach            = false
    command           = [
        "nginx",
        "-g",
        "daemon off;",
    ]
    cpu_shares        = 0
    dns               = []
    dns_opts          = []
    entrypoint        = [
        "/docker-entrypoint.sh",
    ]
    //gateway           = "172.17.0.1"
    hostname          = "f5c8cf143f7d"
   // id                = "f5c8cf143f7dd10abd4fa694a6a46de60b2ae51ac618c42fb2d56e9156cc4153"
    image             = "sha256:519e12e2a84a9eb18094635ae1edfd97b26f95dbc66e317eefb657a1cb08c8dc"
    //ip_address        = "172.17.0.2"
    //ip_prefix_length  = 16
    ipc_mode          = "private"
    log_driver        = "json-file"
    log_opts          = {}
    logs              = false
    max_retry_count   = 0
    memory            = 0
    memory_swap       = 0
    must_run          = true
    name              = "hashicorp-learn"
    
    network_mode      = "default"
    privileged        = false
    publish_all_ports = false
    read_only         = false
    restart           = "no"
    rm                = false
    shm_size          = 64
    start             = true

    ports {
        external = 8081
        internal = 80
        ip       = "0.0.0.0"
        protocol = "tcp"
    }
}

# docker_image.nginx:
resource "docker_image" "nginx" {
    id     = "sha256:519e12e2a84a9eb18094635ae1edfd97b26f95dbc66e317eefb657a1cb08c8dcnginx:latest"
    latest = "sha256:519e12e2a84a9eb18094635ae1edfd97b26f95dbc66e317eefb657a1cb08c8dc"
    name   = "nginx:latest"
}
