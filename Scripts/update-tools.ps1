$ErrorActionPreference = "Stop"

function Assert-Choco {
    if (-not (Get-Command choco -ErrorAction SilentlyContinue)) {
        throw "Chocolatey is not installed or not on PATH. Run install-chocolatey.ps1 first."
    }
}

function Assert-Gup {
    if (-not (Get-Command gup -ErrorAction SilentlyContinue)) {
        throw "gup is not installed or not on PATH. Run install-gup.ps1 first."
    }
}

Assert-Choco
Assert-Gup

Write-Host "Upgrading all Chocolatey packages..."
choco upgrade all -y

Write-Host "Upgrading all go-installed binaries via gup..."
& gup update

Write-Host "OK: Tools upgraded successfully."
