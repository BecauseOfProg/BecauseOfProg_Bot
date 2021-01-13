package data

type Category struct {
	Name string
	ID   string
}

var Categories = map[string]Category{
	"software":    {ID: "software", Name: "🖥 Software"},
	"security":    {ID: "security", Name: "🔐 Sécurité"},
	"web":         {ID: "web", Name: "🌐 Web"},
	"hardware":    {ID: "hardware", Name: "💾 Hardware"},
	"programming": {ID: "programming", Name: "⌨ Programmation"},
	"android":     {ID: "android", Name: "📲 Android"},
	"linux":       {ID: "linux", Name: "🐧 Linux"},
	"windows":     {ID: "windows", Name: "💸 Windows"},
	"apple":       {ID: "apple", Name: "🍎 Apple"},
}
