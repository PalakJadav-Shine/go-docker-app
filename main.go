package main

import (
	"archive/tar"
	"bytes"
	"cmp"
	"compress/gzip"
	"context"
	"crypto/sha256"
	"embed"
	"encoding/json"
	"fmt"
	"hash/crc32"
	"html"
	"io"
	"log"
	"log/slog" // The modern structured logger
	"maps"
	"math"
	"net/http"
	"os"
	"path"
	"reflect"
	"regexp"
	"runtime"
	"slices"
	"sort"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"
	"unicode"
	"unicode/utf8" // Fixes string length problems
	"unique"
	"unsafe"

	"github.com/spf13/cobra"
)

//go:embed pod_log.txt
var embeddedFile embed.FS

func main() {
	// 1. Initialize slog to output JSON (Perfect for K8s/GCP logs)
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
        Level: slog.LevelInfo,
    }))
    slog.SetDefault(logger)

	var rootCmd = &cobra.Command{Use: "superapp"}

	var demoCmd = &cobra.Command{
		Use:   "run-all",
		Short: "Demonstrate All Standard Libraries",
		Run: func(cmd *cobra.Command, args []string) {
			// --- 1. Strings, Unicode (Length Fix), & Regexp ---
			raw := "  K8s-Demo-2026! 🚀 "
			trimmed := strings.TrimSpace(raw)
			
			// Solving the "length" problem: Bytes vs Runes
			byteLen := len(trimmed)
			charLen := utf8.RuneCountInString(trimmed) 
			
			isUpper := unicode.IsUpper(rune(trimmed[0]))
			re := regexp.MustCompile(`[0-9]+`)
			year := re.FindString(trimmed)

			// --- 2. Sorting & Comparison ---
			nums := []int{math.MaxInt8, 10, 5, 100}
			slices.Sort(nums)
			sort.Ints(nums)
			maxVal := cmp.Or(nums[len(nums)-1], 0)
			
			m1 := map[string]int{"cpu": 1}
			m2 := map[string]int{"ram": 2}
			maps.Copy(m1, m2)

			// --- 3. Concurrency & Context ---
			var wg sync.WaitGroup
			wg.Add(1)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			go func(c context.Context) {
				defer wg.Done()
				slog.Info("System Health", "os", runtime.GOOS, "cpus", runtime.NumCPU())
			}(ctx)

			// --- 4. Crypto, Hash, & Encoding ---
			h := sha256.Sum256([]byte("secret"))
			checksum := crc32.ChecksumIEEE([]byte("data"))
			
			statusMap := map[string]string{
				"app":  "healthy", 
				"hash": strconv.FormatUint(uint64(checksum), 16),
			}
			jsonBody, _ := json.Marshal(statusMap)

			// --- 5. Archive & Compression ---
			var buf bytes.Buffer
			gw := gzip.NewWriter(&buf)
			tw := tar.NewWriter(gw)
			hdr := &tar.Header{Name: "demo.txt", Size: int64(len("hello"))}
			_ = tw.WriteHeader(hdr)
			_, _ = io.WriteString(tw, "hello")
			tw.Close(); gw.Close()

			// --- 6. Filesystem & Low Level ---
			cwd, _ := os.Getwd()
			logPath := path.Join(cwd, "logs", "pod.log")
			uid := syscall.Getuid()
			uID := unique.Make("session-123")

			// --- 7. Reflection & Unsafe ---
			typ := reflect.TypeOf(nums)
			size := unsafe.Sizeof(nums)

			// --- 8. FINAL OUTPUT USING SLOG (No more Fprintln) ---
			// We consume all variables here to avoid build errors.
			// --- Category 1: String & Encoding Analysis ---
			slog.Info("TEXT_AUDIT", 
					"content", trimmed, 
					"year", year, 
					"is_upper", isUpper,
			)

			// --- Category 2: Unicode Diagnostics (The length fix) ---
			slog.Info("UNICODE_METRICS", 
					"byte_count", byteLen, 
					"char_count", charLen, 
					"safe_html", html.EscapeString("<b>Safe</b>"),
			)

			// --- Category 3: Memory & Math ---
			slog.Info("COMPUTE_STATS", 
					"max_val", maxVal, 
					"reflect_type", typ.String(), 
					"memory_size_bytes", size,
			)

			// --- Category 4: Security & System ---
			slog.Info("SYSTEM_SECURITY", 
					"sha256", fmt.Sprintf("%x", h[0:8]), // Shorter for readability
					"path", logPath, 
					"uid", uid, 
					"unique_id", uID.Value(),
			)

			// --- Category 5: Data Serialization ---
			slog.Info("ENCODING_TEST", "json_payload", string(jsonBody))
			

			wg.Wait()
		},
	}

	var serveCmd = &cobra.Command{
		Use: "serve",
		Run: func(cmd *cobra.Command, args []string) {
			http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				slog.Info("Inbound Request", "method", r.Method, "path", r.URL.Path)
				fmt.Fprint(w, "Standard Library Demo is Running")
			})
			slog.Info("Server starting", "port", 8080)
			log.Fatal(http.ListenAndServe(":8080", nil))
		},
	}

	rootCmd.AddCommand(demoCmd, serveCmd)
	rootCmd.Execute()
}