$ErrorActionPreference = "Stop"

function Test-DockerInstalled {
    return [bool](Get-Command docker -ErrorAction SilentlyContinue)
}

if (Test-DockerInstalled) {
    Write-Host "OK: Docker already installed: $((Get-Command docker).Source)"
    & docker --version
    exit 0
}

$installer = Join-Path $env:TEMP "DockerDesktopInstaller.exe"
$url = "https://desktop.docker.com/win/main/amd64/Docker Desktop Installer.exe"

Write-Host "Docker not found. Downloading Docker Desktop installer..."
[Net.ServicePointManager]::SecurityProtocol = [Net.SecurityProtocolType]::Tls12
Invoke-WebRequest -Uri $url -OutFile $installer -UseBasicParsing

Write-Host "Installing Docker Desktop (silent)..."
$proc = Start-Process $installer -ArgumentList "install", "--quiet", "--accept-license", "--backend=wsl-2" -Wait -PassThru
if ($proc.ExitCode -ne 0) {
    throw "Docker Desktop installer failed with exit code $($proc.ExitCode)."
}

Write-Host ""
Write-Host "ACTION REQUIRED: Docker Desktop needs a restart to finish setup."
Write-Host "Please restart your PC, then re-run setup.ps1 to continue."
# 3010 is the standard Windows 'success, reboot required' code.
exit 3010
