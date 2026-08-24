$ErrorActionPreference = "Stop"

function Assert-Go {
    if (-not (Get-Command go -ErrorAction SilentlyContinue)) {
        throw "Go is not installed or not on PATH. Run install-go.ps1 first."
    }
}

function Test-AirInstalled {
    return [bool](Get-Command air -ErrorAction SilentlyContinue)
}

Assert-Go

if (Test-AirInstalled) {
    Write-Host "OK: Air already installed: $((Get-Command air).Source)"
    & air -v
    exit 0
}

# Resolve the Go bin directory (where 'go install' places binaries)
$goBin = & go env GOBIN
if ([string]::IsNullOrWhiteSpace($goBin)) {
    $goBin = Join-Path (& go env GOPATH) "bin"
}

Write-Host "Air not found. Installing via 'go install'..."
& go install github.com/air-verse/air@latest

# Add to PATH + make usable immediately
& "$PSScriptRoot\add-application-to-system-path.ps1" $goBin

# Verify install
if (-not (Get-Command air -ErrorAction SilentlyContinue)) {
    throw "Air installed but still not accessible from PATH. Try opening a new PowerShell window."
}

Write-Host "OK: Air installed successfully: $((Get-Command air).Source)"
& air -v
