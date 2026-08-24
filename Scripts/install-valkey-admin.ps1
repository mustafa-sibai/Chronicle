$ErrorActionPreference = "Stop"

function Assert-DockerRunning {
    if (-not (Get-Command docker -ErrorAction SilentlyContinue)) {
        throw "Docker is not installed or not on PATH. Run install-docker.ps1 first."
    }
    & docker info *> $null
    if ($LASTEXITCODE -ne 0) {
        throw "Docker engine is not running. Start Docker Desktop (or re-run install-docker.ps1) and try again."
    }
}

function Test-ContainerExists {
    param([string]$Name)
    $existing = & docker ps -a --filter "name=^/$Name$" --format "{{.Names}}"
    return ($existing -eq $Name)
}

Assert-DockerRunning

$containerName = "valkey-admin"

if (Test-ContainerExists $containerName) {
    Write-Host "OK: Valkey Admin container already exists. Ensuring it is running..."
    & docker start $containerName | Out-Null
    exit 0
}

Write-Host "Valkey Admin container not found. Creating and starting it..."
& docker run -d --name $containerName -p 8080:8080 -e DEPLOYMENT_MODE=Web --restart unless-stopped valkey/valkey-admin:latest

if (-not (Test-ContainerExists $containerName)) {
    throw "Failed to create the Valkey Admin container."
}

# The valkey container is reachable from here at host.docker.internal:6379
Write-Host "OK: Valkey Admin is running at http://localhost:8080"
