$elapsed = Measure-Command{
  ./scurl
}
Write-Host "Elapsed time: $($elapsed.TotalMilliseconds) ms"
