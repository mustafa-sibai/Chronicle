$ErrorActionPreference = "Stop"

function Assert-Choco {
    if (-not (Get-Command choco -ErrorAction SilentlyContinue)) {
        throw "Chocolatey is not installed or not on PATH. Run install-chocolatey.ps1 first."
    }
}

function Test-MongoInstalled {
    return [bool](Get-Command mongod -ErrorAction SilentlyContinue)
}

Assert-Choco

if (Test-MongoInstalled) {
    Write-Host "OK: MongoDB already installed: $((Get-Command mongod).Source)"
    & mongod --version
    exit 0
}

Write-Host "MongoDB not found. Installing latest MongoDB Community Server via Chocolatey..."
choco install -y mongodb --force

# Resolve the highest-versioned Server bin dir and add to PATH + make usable immediately
$mongoBin = Get-ChildItem "C:\Program Files\MongoDB\Server" -Directory -ErrorAction SilentlyContinue |
    Sort-Object Name -Descending |
    Select-Object -First 1 |
    ForEach-Object { Join-Path $_.FullName "bin" }
if ($mongoBin) {
    & "$PSScriptRoot\add-application-to-system-path.ps1" $mongoBin
}

# Verify install
if (-not (Get-Command mongod -ErrorAction SilentlyContinue)) {
    throw "MongoDB installed but still not accessible from PATH. Try opening a new PowerShell window."
}

Write-Host "OK: MongoDB installed successfully: $((Get-Command mongod).Source)"
& mongod --version
