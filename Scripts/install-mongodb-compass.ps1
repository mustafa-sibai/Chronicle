$ErrorActionPreference = "Stop"

function Assert-Choco {
    if (-not (Get-Command choco -ErrorAction SilentlyContinue)) {
        throw "Chocolatey is not installed or not on PATH. Run install-chocolatey.ps1 first."
    }
}

function Test-CompassInstalled {
    # Compass doesn't add a command to PATH, so probe the choco package and known install locations.
    if (choco list --local-only --exact mongodb-compass --limit-output) {
        return $true
    }
    $candidates = @(
        (Join-Path $env:LOCALAPPDATA "MongoDBCompass\MongoDBCompass.exe"),
        (Join-Path ${env:ProgramFiles} "MongoDB Compass\MongoDBCompass.exe")
    )
    return [bool]($candidates | Where-Object { Test-Path $_ })
}

Assert-Choco

if (Test-CompassInstalled) {
    Write-Host "OK: MongoDB Compass already installed."
    exit 0
}

Write-Host "MongoDB Compass not found. Installing latest via Chocolatey..."
choco install -y mongodb-compass --force

Write-Host "OK: MongoDB Compass installed successfully."
