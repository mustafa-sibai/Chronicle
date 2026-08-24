$ErrorActionPreference = "Stop"

function Assert-Choco {
    if (-not (Get-Command choco -ErrorAction SilentlyContinue)) {
        throw "Chocolatey is not installed or not on PATH. Run install-chocolatey.ps1 first."
    }
}

function Test-PowerShellCoreInstalled {
    return [bool](Get-Command pwsh -ErrorAction SilentlyContinue)
}

Assert-Choco

if (Test-PowerShellCoreInstalled) {
    Write-Host "OK: PowerShell Core already installed: $((Get-Command pwsh).Source)"
    & pwsh --version
    exit 0
}

Write-Host "PowerShell Core not found. Installing via Chocolatey..."
choco install -y powershell-core --force

# Add to PATH + make usable immediately
& "$PSScriptRoot\add-application-to-system-path.ps1" "C:\Program Files\PowerShell\7"

# Verify install
if (-not (Get-Command pwsh -ErrorAction SilentlyContinue)) {
    throw "PowerShell Core installed but still not accessible from PATH. Try opening a new PowerShell window."
}

Write-Host "OK: PowerShell Core installed successfully: $((Get-Command pwsh).Source)"
& pwsh --version
