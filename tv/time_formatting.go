package tv
import "fmt"

// Parses hh:mm:ss format to number of seconds.
func parseTime(str string) int {
  var h, m, s int
  fmt.Sscanf(str, "%d:%d:%d", &h, &m, &s)
  // Deliberately ignore errors, and just accept 0s
  return h * 60 * 60 + m * 60 + s
}

// Format seconds into hh:mm:ss
func formatTime(secs int) string {
  s := secs % 60
  m := (secs / 60) % 60
  h := (secs / (60*60))
  return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
}
