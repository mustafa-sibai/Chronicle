$ErrorActionPreference = "Stop"

$SkipDbs = $args -contains "--skip-dbs"

# Only Windows is supported for now (macOS/Linux setup.sh may come later).
if ($env:OS -ne "Windows_NT") {
    Write-Host "This setup script only supports Windows. Exiting."
    exit 1
}

# Check if the script is running with administrative privileges
$principal = New-Object Security.Principal.WindowsPrincipal([Security.Principal.WindowsIdentity]::GetCurrent())
if (-not $principal.IsInRole([Security.Principal.WindowsBuiltInRole]::Administrator)) {
    Write-Host "This script must be run as an administrator. Please restart PowerShell with 'Run as Administrator' and try again."
    exit 1
}

Write-Host "The following applications will be installed on your system:"
Write-Host "  - Chocolatey (package manager)"
Write-Host "  - PowerShell Core (pwsh)"
Write-Host "  - Go (programming language)"
Write-Host "  - gup (Go binary updater)"
Write-Host "  - Air (Go live-reload tool)"
Write-Host "  - Task (task runner)"
Write-Host "  - MongoDB (Community Server)"
Write-Host "  - MongoDB Compass (GUI)"
Write-Host "  - Docker Desktop (container runtime)"
Write-Host "  - Valkey (via Docker)"
Write-Host "  - Valkey Admin (via Docker)"
Write-Host ""
Write-Host "Note: You can skip installing MongoDB, MongoDB Compass, Valkey server, Valkey Admin, and Docker by passing the argument --skip-dbs when running this script" -ForegroundColor Yellow

$continue = Read-Host "Do you want to continue? (Y/N)"
if ($continue -ne "Y" -and $continue -ne "y") {
    Write-Host "Setup aborted by user."
    exit 1
}

$ScriptsDir = Join-Path $PSScriptRoot "Scripts"

& "$ScriptsDir\install-chocolatey.ps1"
& "$ScriptsDir\install-powershell-core.ps1"
& "$ScriptsDir\install-go.ps1"
& "$ScriptsDir\install-gup.ps1"
& "$ScriptsDir\install-air.ps1"
& "$ScriptsDir\install-taskfile.ps1"

if ($SkipDbs) {
    Write-Host "Skipping MongoDB, MongoDB Compass, Docker, Valkey, and Valkey Admin (--skip-dbs)."
    Write-Host "OK: Setup complete."
    exit 0
}

& "$ScriptsDir\install-mongodb.ps1"
& "$ScriptsDir\install-mongodb-compass.ps1"

& "$ScriptsDir\install-docker.ps1"
# A fresh Docker install needs a restart. Stop here so the user can reboot, then
# re-run setup.ps1 to install the Docker-based tools.
if ($LASTEXITCODE -eq 3010) {
    Write-Host "Setup paused: restart your PC, then run setup.ps1 again to finish."
    exit 0
}
& "$ScriptsDir\install-valkey.ps1"
& "$ScriptsDir\install-valkey-admin.ps1"

Write-Host "OK: Setup complete."
