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

Assert-DockerRunning

Write-Host "Creating and starting Valkey Admin..."
& docker run -d --name valkey-admin --network "container:valkey" -e DEPLOYMENT_MODE=Web --restart unless-stopped valkey/valkey-admin:latest

Write-Host "OK: Valkey Admin is running at http://localhost:8080"
Write-Host "    Connect to host 'localhost' on port 6379 (TLS off)."