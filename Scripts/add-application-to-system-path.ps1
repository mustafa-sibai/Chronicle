param(
    [Parameter(Mandatory = $true)]
    [string]$PathToAdd
)

$ErrorActionPreference = "Stop"

# Add to MACHINE path
$machinePath = [Environment]::GetEnvironmentVariable("Path", "Machine")
[Environment]::SetEnvironmentVariable("Path", "$machinePath;$PathToAdd", "Machine")

# Add to CURRENT session so no terminal restart needed
$env:Path = "$env:Path;$PathToAdd"

Write-Host "Added to PATH: $PathToAdd"
