package ecspresso

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
)

var (
	SortTaskDefinitionForDiff    = sortTaskDefinitionForDiff
	SortServiceDefinitionForDiff = sortServiceDefinitionForDiff
	EqualString                  = equalString
	ToNumberCPU                  = toNumberCPU
	ToNumberMemory               = toNumberMemory
	CalcDesiredCount             = calcDesiredCount
	ParseTags                    = parseTags
	ExtractRoleName              = extractRoleName
	IsLongArnFormat              = isLongArnFormat
	ECRImageURLRegex             = ecrImageURLRegex
	NewLogger                    = newLogger
	NewLogFilter                 = newLogFilter
	NewConfigLoader              = newConfigLoader
	ArnToName                    = arnToName
	InitVerifyState              = initVerifyState
	VerifyResource               = verifyResource
)

func (d *App) SetLogger(logger *log.Logger) {
	d.logger = logger
}

func SetLogger(logger *log.Logger) {
	commonLogger = logger
}

func SetAWSV2ConfigLoadOptionsFunc(f []func(*config.LoadOptions) error) {
	awsv2ConfigLoadOptionsFunc = f
}

func ResetAWSV2ConfigLoadOptionsFunc() {
	awsv2ConfigLoadOptionsFunc = nil
}

func (d *App) TaskDefinitionArnForRun(ctx context.Context, opt RunOption) (string, error) {
	return d.taskDefinitionArnForRun(ctx, opt)
}
