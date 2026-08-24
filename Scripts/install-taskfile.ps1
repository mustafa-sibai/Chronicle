$ErrorActionPreference = "Stop"

function Assert-Go {
    if (-not (Get-Command go -ErrorAction SilentlyContinue)) {
        throw "Go is not installed or not on PATH. Run install-go.ps1 first."
    }
}

function Test-TaskInstalled {
    return [bool](Get-Command task -ErrorAction SilentlyContinue)
}

Assert-Go

if (Test-TaskInstalled) {
    Write-Host "OK: Task already installed: $((Get-Command task).Source)"
    & task --version
    exit 0
}

Write-Host "Task not found. Installing latest via go install..."
& go install github.com/go-task/task/v3/cmd/task@latest

# Resolve GOBIN (falls back to GOPATH\bin) and add to PATH + make usable immediately
$goBin = & go env GOBIN
if ([string]::IsNullOrWhiteSpace($goBin)) {
    $goBin = Join-Path (& go env GOPATH) "bin"
}
& "$PSScriptRoot\add-application-to-system-path.ps1" $goBin

# Verify install
if (-not (Get-Command task -ErrorAction SilentlyContinue)) {
    throw "Task installed but still not accessible from PATH. Try opening a new PowerShell window."
}

Write-Host "OK: Task installed successfully: $((Get-Command task).Source)"
& task --version
