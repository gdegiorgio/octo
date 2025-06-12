$OCTO_HOME = "$env:USERPROFILE\.octo"

function Install-Octo {
    $octoBinary = Join-Path $OCTO_HOME "bin\octo.exe"

    if (Test-Path $octoBinary) {
        Write-Host "Octo binary already exists at $octoBinary"
        return
    }

    # Detect architecture
    $arch = if ([System.Environment]::Is64BitOperatingSystem) {
        if ($env:PROCESSOR_ARCHITECTURE -match "ARM64") { "arm64" } else { "amd64" }
    } else {
        Write-Host "Unsupported architecture"
        return
    }

    $os = "windows"
    Write-Host "Installing Octo for $os $arch"

    $binDir = Join-Path $OCTO_HOME "bin"
    New-Item -ItemType Directory -Force -Path $binDir | Out-Null

    $zipFile = Join-Path $OCTO_HOME "octo_${os}_${arch}.zip"
    $downloadUrl = "https://github.com/gdegiorgio/octo/releases/latest/download/octo_${os}_${arch}.zip"

    Invoke-WebRequest -Uri $downloadUrl -OutFile $zipFile
    Expand-Archive -Path $zipFile -DestinationPath $OCTO_HOME -Force
    Remove-Item $zipFile

    $extractedPath = Get-ChildItem -Path $OCTO_HOME -Recurse -Filter "octo.exe" | Select-Object -First 1

    if (-not $extractedPath) {
        Write-Host "Failed to find octo.exe after extraction."
        return
    }

    Move-Item -Path $extractedPath.FullName -Destination $octoBinary -Force

    # Add to PATH if not already there
    $envPath = [System.Environment]::GetEnvironmentVariable("Path", "User")
    if ($envPath -notlike "*$($binDir)*") {
        [System.Environment]::SetEnvironmentVariable("OCTO_HOME", $OCTO_HOME, "User")
        [System.Environment]::SetEnvironmentVariable("Path", "$envPath;$binDir", "User")
        Write-Host "Added $binDir to PATH"
    }

    Write-Host "Octo installed successfully. Restart your terminal to use it now."
}

Install-Octo
