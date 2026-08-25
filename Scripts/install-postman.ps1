$ErrorActionPreference = "Stop"

function Assert-Choco {
    if (-not (Get-Command choco -ErrorAction SilentlyContinue)) {
        throw "Chocolatey is not installed or not on PATH. Run install-chocolatey.ps1 first."
    }
}

function Test-PostmanInstalled {
    # Postman doesn't add a command to PATH, so probe the choco package and known install locations.
    if (choco list --local-only --exact postman --limit-output) {
        return $true
    }
    $candidates = @(
        (Join-Path $env:LOCALAPPDATA "Postman\Postman.exe"),
        (Join-Path ${env:ProgramFiles} "Postman\Postman.exe")
    )
    return [bool]($candidates | Where-Object { Test-Path $_ })
}

Assert-Choco

if (Test-PostmanInstalled) {
    Write-Host "OK: Postman already installed."
    exit 0
}

Write-Host "Postman not found. Installing latest via Chocolatey..."
choco install -y postman --force

Write-Host "OK: Postman installed successfully."
