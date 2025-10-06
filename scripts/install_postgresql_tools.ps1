# PostgreSQL Tools Installation Script for Windows
# This script helps install PostgreSQL client tools for the backup API

Write-Host "PostgreSQL Tools Installation Script" -ForegroundColor Green
Write-Host "=====================================" -ForegroundColor Green

# Check if pg_dump is already available
$pgDumpPath = Get-Command pg_dump -ErrorAction SilentlyContinue
if ($pgDumpPath) {
    Write-Host "✅ pg_dump is already available at: $($pgDumpPath.Source)" -ForegroundColor Green
    exit 0
}

Write-Host "❌ pg_dump not found. Installing PostgreSQL client tools..." -ForegroundColor Yellow

# Check if Chocolatey is available
$choco = Get-Command choco -ErrorAction SilentlyContinue
if ($choco) {
    Write-Host "Installing PostgreSQL using Chocolatey..." -ForegroundColor Blue
    choco install postgresql --version=15.4 -y
    if ($LASTEXITCODE -eq 0) {
        Write-Host "✅ PostgreSQL installed successfully!" -ForegroundColor Green
        Write-Host "Please restart your terminal and try the backup API again." -ForegroundColor Yellow
        exit 0
    }
}

# Check if winget is available
$winget = Get-Command winget -ErrorAction SilentlyContinue
if ($winget) {
    Write-Host "Installing PostgreSQL using winget..." -ForegroundColor Blue
    winget install PostgreSQL.PostgreSQL
    if ($LASTEXITCODE -eq 0) {
        Write-Host "✅ PostgreSQL installed successfully!" -ForegroundColor Green
        Write-Host "Please restart your terminal and try the backup API again." -ForegroundColor Yellow
        exit 0
    }
}

# Manual installation instructions
Write-Host "`nManual Installation Required:" -ForegroundColor Red
Write-Host "=============================" -ForegroundColor Red
Write-Host "1. Download PostgreSQL from: https://www.postgresql.org/download/windows/" -ForegroundColor White
Write-Host "2. Run the installer and select 'Command Line Tools' during installation" -ForegroundColor White
Write-Host "3. Add PostgreSQL bin directory to your PATH environment variable:" -ForegroundColor White
Write-Host "   - Open System Properties > Environment Variables" -ForegroundColor White
Write-Host "   - Add to PATH: C:\Program Files\PostgreSQL\15\bin" -ForegroundColor White
Write-Host "   - (Adjust version number as needed)" -ForegroundColor White
Write-Host "4. Restart your terminal and try the backup API again" -ForegroundColor White

Write-Host "`nAlternative: Download only client tools:" -ForegroundColor Yellow
Write-Host "1. Go to: https://www.postgresql.org/download/windows/" -ForegroundColor White
Write-Host "2. Download 'Command Line Tools' (not the full installer)" -ForegroundColor White
Write-Host "3. Extract to a folder (e.g., C:\PostgreSQL\bin)" -ForegroundColor White
Write-Host "4. Add that folder to your PATH environment variable" -ForegroundColor White

Write-Host "`nAfter installation, verify with: pg_dump --version" -ForegroundColor Cyan
