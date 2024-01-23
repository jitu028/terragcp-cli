package flags

const (
	ApiKey               = "api-key"
	Model                = "model"
	AutoSave             = "auto-save"
	Render               = "render"
	AllowHarmProbability = "allow-harm-probability"
	TopK                 = "top-k"
	TopP                 = "top-p"
	Temperature          = "temperature"
	CandidateCount       = "candidate-count"
	MaxOutputTokens      = "max-output-tokens"
	File                 = "file"
	Format               = "format"
)

const (
	RenderFormatHtml     = "html"
	RenderFormatMarkdown = "markdown"
	RenderFormatPretty   = "pretty"
)

const (
	MaxBlobBufferSizeBytes = 4194304
)

const (
	ApiKeyEnv = "GOOGLE_API_KEY"
)

const (
	ModelGeminiPro    = "models/gemini-pro"
	ModelEmbedding001 = "models/embedding-001"
	ModelGeminiProVision = "models/gemini-pro-vision"

)

const (
	HarmProbabilityUnspecified = "unspecified"
	HarmProbabilityNegligible  = "negligible"
	HarmProbabilityLow         = "low"
	HarmProbabilityMedium      = "medium"
	HarmProbabilityHigh        = "high"
)
