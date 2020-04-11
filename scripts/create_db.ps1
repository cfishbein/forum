if(Test-Path forum.db) {
    Write-Host "Database already exists"
    exit 0
}

Get-Content $PSScriptRoot/create_tables.sql -Raw | sqlite3 forum.db
