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

$existing = & docker ps -a --filter "name=^valkey$" --format "{{.Names}}"
if ($existing -eq "valkey") {
    Write-Host "Valkey container already exists. Ensuring it is running..."
    & docker start valkey *> $null
    if ($LASTEXITCODE -ne 0) {
        throw "Failed to start existing Valkey container."
    }
}
else {
    Write-Host "Creating and starting Valkey..."
    & docker run -d --name valkey -p 6379:6379 -p 8080:8080 --restart unless-stopped valkey/valkey
    if ($LASTEXITCODE -ne 0) {
        throw "Failed to create Valkey container."
    }
}

Write-Host "OK: Valkey is running on localhost:6379"