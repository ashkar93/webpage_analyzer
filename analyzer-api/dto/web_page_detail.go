package dto

type WebPageDetail struct {
	Version              string
	Title                string
	H1                   int
	H2                   int
	H3                   int
	H4                   int
	H5                   int
	H6                   int
	ExternalLink         int
	InternalLink         int
	InternalDeadIdLink   int
	InternalDeadPathLink int
	ExternalDeadLink     int
	IsWithLogin          bool
}
