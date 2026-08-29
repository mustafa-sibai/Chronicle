$ErrorActionPreference = "Stop"

$ReplicaSetName = "rs0"

function Assert-Choco {
    if (-not (Get-Command choco -ErrorAction SilentlyContinue)) {
        throw "Chocolatey is not installed or not on PATH. Run install-chocolatey.ps1 first."
    }
}

function Test-MongoInstalled {
    return [bool](Get-Command mongod -ErrorAction SilentlyContinue)
}

function Get-MongoBinDir {
    return Get-ChildItem "C:\Program Files\MongoDB\Server" -Directory -ErrorAction SilentlyContinue |
        Sort-Object Name -Descending |
        Select-Object -First 1 |
        ForEach-Object { Join-Path $_.FullName "bin" }
}

function Install-MongoShell {
    if (Get-Command mongosh -ErrorAction SilentlyContinue) {
        return
    }

    Write-Host "mongosh not found. Installing via Chocolatey..."
    choco install -y mongodb-shell --force

    # Refresh current session PATH from machine PATH so mongosh is usable immediately
    $machinePath = [Environment]::GetEnvironmentVariable("Path", "Machine")
    $env:Path = "$machinePath;$env:Path"

    if (-not (Get-Command mongosh -ErrorAction SilentlyContinue)) {
        throw "mongosh installed but still not accessible from PATH. Try opening a new PowerShell window."
    }

    Write-Host "OK: mongosh installed successfully."
}

# Multi-document transactions require MongoDB's oplog, which only exists once
# replication is enabled - even a single-member replica set unlocks it.
function Enable-MongoReplicaSet {
    $mongoBin = Get-MongoBinDir
    $configPath = if ($mongoBin) { Join-Path $mongoBin "mongod.cfg" } else { $null }
    if (-not $configPath -or -not (Test-Path $configPath)) {
        Write-Warning "Could not locate mongod.cfg; skipping replica set setup."
        return
    }

    $config = Get-Content $configPath -Raw
    $escapedName = [regex]::Escape($ReplicaSetName)
    if ($config -match "(?m)^\s*replSetName:\s*$escapedName\s*$") {
        Write-Host "OK: MongoDB replica set '$ReplicaSetName' already configured."
    } else {
        Write-Host "Enabling replica set '$ReplicaSetName' in mongod.cfg (required for multi-document transactions)..."
        Add-Content -Path $configPath -Value "`r`nreplication:`r`n  replSetName: $ReplicaSetName`r`n"

        Write-Host "Restarting MongoDB service to apply config..."
        Restart-Service -Name MongoDB
        Start-Sleep -Seconds 3
    }

    Install-MongoShell

    $isInitiated = $false
    try {
        $result = & mongosh --quiet --eval "rs.status().ok" 2>$null
        if ($result -match "1") { $isInitiated = $true }
    } catch {}

    if ($isInitiated) {
        Write-Host "OK: Replica set already initiated."
    } else {
        Write-Host "Initiating replica set '$ReplicaSetName'..."
        & mongosh --quiet --eval "rs.initiate()"
        Start-Sleep -Seconds 2
        Write-Host "OK: Replica set initiated."
    }
}

Assert-Choco

if (Test-MongoInstalled) {
    Write-Host "OK: MongoDB already installed: $((Get-Command mongod).Source)"
} else {
    Write-Host "MongoDB not found. Installing latest MongoDB Community Server via Chocolatey..."
    choco install -y mongodb --force

    $mongoBin = Get-MongoBinDir
    if ($mongoBin) {
        & "$PSScriptRoot\add-application-to-system-path.ps1" $mongoBin
    }

    if (-not (Get-Command mongod -ErrorAction SilentlyContinue)) {
        throw "MongoDB installed but still not accessible from PATH. Try opening a new PowerShell window."
    }

    Write-Host "OK: MongoDB installed successfully: $((Get-Command mongod).Source)"
}

& mongod --version

Enable-MongoReplicaSet
