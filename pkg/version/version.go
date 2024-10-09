package version

var (
	Version      = "v1.0.0"
	BuildDateUTC = "2024-10-09T10:20:55Z"
	GoVersion    = "go1.22.1"
	Platform     = "linux/amd64"
	GitRepo      = "http://github.com/yecaowulei/sre-tool-kit.git"
)

// PrintVersion
// @Description: 打印构建版本信息
func PrintVersion() {
	println("Version:", Version)
	println("BuildDateUTC:", BuildDateUTC)
	println("GoVersion:", GoVersion)
	println("Platform:", Platform)
	println("GitRepo:", GitRepo)
	println()
}
