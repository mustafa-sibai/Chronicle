$ErrorActionPreference = "Stop"

function Assert-Go {
    if (-not (Get-Command go -ErrorAction SilentlyContinue)) {
        throw "Go is not installed or not on PATH. Run install-go.ps1 first."
    }
}

function Test-GupInstalled {
    return [bool](Get-Command gup -ErrorAction SilentlyContinue)
}

Assert-Go

if (Test-GupInstalled) {
    Write-Host "OK: gup already installed: $((Get-Command gup).Source)"
    & gup version
    exit 0
}

Write-Host "gup not found. Installing latest via go install..."
& go install github.com/nao1215/gup@latest

# Resolve GOBIN (falls back to GOPATH\bin) and add to PATH + make usable immediately
$goBin = & go env GOBIN
if ([string]::IsNullOrWhiteSpace($goBin)) {
    $goBin = Join-Path (& go env GOPATH) "bin"
}
& "$PSScriptRoot\add-application-to-system-path.ps1" $goBin

# Verify install
if (-not (Get-Command gup -ErrorAction SilentlyContinue)) {
    throw "gup installed but still not accessible from PATH. Try opening a new PowerShell window."
}

Write-Host "OK: gup installed successfully: $((Get-Command gup).Source)"
& gup version
