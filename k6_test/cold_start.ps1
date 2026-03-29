param(
    [string]$ContainerName,
    [string]$Port
)

Write-Host "Preparing Cold Start for $ContainerName..." -ForegroundColor Cyan
docker stop $ContainerName | Out-Null
Start-Sleep -Seconds 2

Write-Host "Starting container and starting the stopwatch..." -ForegroundColor Yellow
$stopwatch = [System.Diagnostics.Stopwatch]::StartNew()
docker start $ContainerName | Out-Null

# Changed localhost to 127.0.0.1 to avoid Windows IPv6 routing issues
$url = "http://127.0.0.1:$Port/api/user/test3"
$isUp = $false
$attempts = 0
$lastError = ""

while (-not $isUp -and $attempts -lt 200) {
    try {
        $response = Invoke-WebRequest -Uri $url -UseBasicParsing -TimeoutSec 1 -ErrorAction Stop
        if ($response.StatusCode -eq 200) {
            $isUp = $true
        }
    } catch {
        # Capture the actual error message so we know what went wrong
        $lastError = $_.Exception.Message
        $attempts++
        Start-Sleep -Milliseconds 100 
    }
}

$stopwatch.Stop()

if ($isUp) {
    Write-Host "SUCCESS: $ContainerName woke up and served traffic in: $($stopwatch.Elapsed.TotalSeconds) seconds" -ForegroundColor Green
} else {
    Write-Host "FAILED: $ContainerName failed to start within 20 seconds." -ForegroundColor Red
    Write-Host "Last Error was: $lastError" -ForegroundColor Red
}