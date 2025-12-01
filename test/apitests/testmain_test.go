package apitests

import (
    "log"
    "net/http"
    "os"
    "os/exec"
    "path/filepath"
    "testing"
    "time"
)

// TestMain automatically starts the server (go run .) before tests and stops it after.
// If a server is already running on localhost:8080, TestMain will reuse it.
func TestMain(m *testing.M) {
    // If there's already a server running, just run tests
    if isServerUp() {
        os.Exit(m.Run())
    }

    wd, err := os.Getwd()
    if err != nil {
        log.Fatalf("cannot get working dir: %v", err)
    }
    repoRoot := filepath.Join(wd, "..", "..")

    cmd := exec.Command("go", "run", ".")
    cmd.Dir = repoRoot
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    if err := cmd.Start(); err != nil {
        log.Fatalf("cannot start server: %v", err)
    }

    // Wait until server is ready (up to 5s)
    ready := waitForServerUp(5*time.Second, 100*time.Millisecond)
    if !ready {
        // attempt to kill the process and fail the tests
        _ = cmd.Process.Kill()
        log.Fatalf("server not ready after timeout")
    }

    // Run tests
    code := m.Run()

    // Stop server
    if err := cmd.Process.Kill(); err != nil {
        log.Printf("failed to kill server process: %v", err)
    }

    os.Exit(code)
}

func isServerUp() bool {
    client := http.Client{Timeout: 200 * time.Millisecond}
    resp, err := client.Get("http://localhost:8080/")
    if err != nil {
        return false
    }
    resp.Body.Close()
    return true
}

func waitForServerUp(timeout, tick time.Duration) bool {
    deadline := time.Now().Add(timeout)
    for time.Now().Before(deadline) {
        if isServerUp() {
            return true
        }
        time.Sleep(tick)
    }
    return false
}
