$EXECUTABLE_NAME = "crump8term"
$EXECUTABLE_PATH = ".\$EXECUTABLE_NAME"
$GoPath = ((go env | Select-String -Pattern "GOPATH=" | Out-String) -split "=")[1].TrimEnd()
$GoPath += "\bin"
Set-Location $EXECUTABLE_PATH
go build -gcflags "-N -l"
if($LastExitCode -eq 0) {
    Write-Output 'build success'
} else {
    Write-Error 'go build failed'
    Exit 1
}
Start-Process ".\$EXECUTABLE_NAME.exe"
$timeOut = 20
$started = $false
# wait for process to start
Do {
    Start-Sleep -Milliseconds 250
    $timeOut--
    $Proc = Get-Process $EXECUTABLE_NAME -ErrorAction SilentlyContinue
    If ($Proc) {
        $started = $true
    }
}
Until ($started -or $timeOut -eq 0)
If (!($started)) {
    Write-Error 'Process did not start'
    Exit 1
}
$ProcId = ($Proc | Select-Object -expand Id)
Start-Process -FilePath "$GoPath\dlv.exe" -ArgumentList "attach $ProcId --headless --listen=:2345 --log" -WindowStyle Hidden