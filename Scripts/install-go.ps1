$ErrorActionPreference = "Stop"

function Assert-Choco {
    if (-not (Get-Command choco -ErrorAction SilentlyContinue)) {
        throw "Chocolatey is not installed or not on PATH. Run install-chocolatey.ps1 first."
    }
}

function Test-GoInstalled {
    return [bool](Get-Command go -ErrorAction SilentlyContinue)
}

Assert-Choco

if (Test-GoInstalled) {
    Write-Host "OK: Go already installed: $((Get-Command go).Source)"
    & go version
    exit 0
}

Write-Host "Go not found. Installing via Chocolatey..."
choco install -y golang --force

# Add to PATH + make usable immediately
& "$PSScriptRoot\add-application-to-system-path.ps1" "C:\Program Files\Go\bin"

# Verify install
if (-not (Get-Command go -ErrorAction SilentlyContinue)) {
    throw "Go installed but still not accessible from PATH. Try opening a new PowerShell window."
}

Write-Host "OK: Go installed successfully: $((Get-Command go).Source)"
& go version
