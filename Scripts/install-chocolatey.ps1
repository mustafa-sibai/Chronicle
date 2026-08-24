$ErrorActionPreference = "Stop"

function Test-ChocolateyInstalled {
    return [bool](Get-Command choco -ErrorAction SilentlyContinue)
}

if (Test-ChocolateyInstalled) {
    Write-Host "OK: Chocolatey is already installed: $((Get-Command choco).Source)"
    exit 0
}

Write-Host "Chocolatey not found. Installing Chocolatey..."

# Chocolatey install requires TLS 1.2+ and running as Admin.
if (-not ([Security.Principal.WindowsPrincipal][Security.Principal.WindowsIdentity]::GetCurrent()
    ).IsInRole([Security.Principal.WindowsBuiltInRole]::Administrator)) {
    throw "Please run this script in an elevated PowerShell (Run as Administrator)."
}

# Ensure TLS 1.2
[Net.ServicePointManager]::SecurityProtocol = [Net.SecurityProtocolType]::Tls12

Set-ExecutionPolicy Bypass -Scope Process -Force

$installScript = (New-Object System.Net.WebClient).DownloadString("https://community.chocolatey.org/install.ps1")
Invoke-Expression $installScript

if (-not (Get-Command choco -ErrorAction SilentlyContinue)) {
    throw "Chocolatey install ran but 'choco' is still not available in PATH. Try opening a new PowerShell window."
}

Write-Host "OK: Chocolatey installed successfully: $((Get-Command choco).Source)"
