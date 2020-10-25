// Code generated by piper's step-generator. DO NOT EDIT.

package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/SAP/jenkins-library/pkg/config"
	"github.com/SAP/jenkins-library/pkg/log"
	"github.com/SAP/jenkins-library/pkg/telemetry"
	"github.com/spf13/cobra"
)

type mavenExecuteStaticCodeChecksOptions struct {
	SpotBugs                     bool     `json:"spotBugs,omitempty"`
	Pmd                          bool     `json:"pmd,omitempty"`
	MavenModulesExcludes         []string `json:"mavenModulesExcludes,omitempty"`
	SpotBugsExcludeFilterFile    string   `json:"spotBugsExcludeFilterFile,omitempty"`
	SpotBugsIncludeFilterFile    string   `json:"spotBugsIncludeFilterFile,omitempty"`
	SpotBugsMaxAllowedViolations int      `json:"spotBugsMaxAllowedViolations,omitempty"`
	PmdFailurePriority           int      `json:"pmdFailurePriority,omitempty"`
	PmdMaxAllowedViolations      int      `json:"pmdMaxAllowedViolations,omitempty"`
	ProjectSettingsFile          string   `json:"projectSettingsFile,omitempty"`
	GlobalSettingsFile           string   `json:"globalSettingsFile,omitempty"`
	M2Path                       string   `json:"m2Path,omitempty"`
	LogSuccessfulMavenTransfers  bool     `json:"logSuccessfulMavenTransfers,omitempty"`
}

// MavenExecuteStaticCodeChecksCommand Execute static code checks for Maven based projects. The plugins SpotBugs and PMD are used.
func MavenExecuteStaticCodeChecksCommand() *cobra.Command {
	const STEP_NAME = "mavenExecuteStaticCodeChecks"

	metadata := mavenExecuteStaticCodeChecksMetadata()
	var stepConfig mavenExecuteStaticCodeChecksOptions
	var startTime time.Time

	var createMavenExecuteStaticCodeChecksCmd = &cobra.Command{
		Use:   STEP_NAME,
		Short: "Execute static code checks for Maven based projects. The plugins SpotBugs and PMD are used.",
		Long: `Executes Spotbugs Maven plugin as well as Pmd Maven plugin for static code checks.
SpotBugs is a program to find bugs in Java programs. It looks for instances of “bug patterns” — code instances that are likely to be errors.
For more information please visit https://spotbugs.readthedocs.io/en/latest/maven.html
PMD is a source code analyzer. It finds common programming flaws like unused variables, empty catch blocks, unnecessary object creation, and so forth. It supports Java, JavaScript, Salesforce.com Apex and Visualforce, PLSQL, Apache Velocity, XML, XSL.
For more information please visit https://pmd.github.io/.
The plugins should be configured in the respective pom.xml.
For SpotBugs include- and exclude filters as well as maximum allowed violations are conifgurable via .pipeline/config.yml.
For PMD the failure priority and the max allowed violations are configurable via .pipeline/config.yml.`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			startTime = time.Now()
			log.SetStepName(STEP_NAME)
			log.SetVerbose(GeneralConfig.Verbose)

			path, _ := os.Getwd()
			fatalHook := &log.FatalHook{CorrelationID: GeneralConfig.CorrelationID, Path: path}
			log.RegisterHook(fatalHook)

			err := PrepareConfig(cmd, &metadata, STEP_NAME, &stepConfig, config.OpenPiperFile)
			if err != nil {
				log.SetErrorCategory(log.ErrorConfiguration)
				return err
			}

			if len(GeneralConfig.HookConfig.SentryConfig.Dsn) > 0 {
				sentryHook := log.NewSentryHook(GeneralConfig.HookConfig.SentryConfig.Dsn, GeneralConfig.CorrelationID)
				log.RegisterHook(&sentryHook)
			}

			return nil
		},
		Run: func(_ *cobra.Command, _ []string) {
			telemetryData := telemetry.CustomData{}
			telemetryData.ErrorCode = "1"
			handler := func() {
				telemetryData.Duration = fmt.Sprintf("%v", time.Since(startTime).Milliseconds())
				telemetryData.ErrorCategory = log.GetErrorCategory().String()
				telemetry.Send(&telemetryData)
			}
			log.DeferExitHandler(handler)
			defer handler()
			telemetry.Initialize(GeneralConfig.NoTelemetry, STEP_NAME)
			mavenExecuteStaticCodeChecks(stepConfig, &telemetryData)
			telemetryData.ErrorCode = "0"
			log.Entry().Info("SUCCESS")
		},
	}

	addMavenExecuteStaticCodeChecksFlags(createMavenExecuteStaticCodeChecksCmd, &stepConfig)
	return createMavenExecuteStaticCodeChecksCmd
}

func addMavenExecuteStaticCodeChecksFlags(cmd *cobra.Command, stepConfig *mavenExecuteStaticCodeChecksOptions) {
	cmd.Flags().BoolVar(&stepConfig.SpotBugs, "spotBugs", true, "Parameter to turn off SpotBugs.")
	cmd.Flags().BoolVar(&stepConfig.Pmd, "pmd", true, "Parameter to turn off PMD.")
	cmd.Flags().StringSliceVar(&stepConfig.MavenModulesExcludes, "mavenModulesExcludes", []string{}, "Maven modules which should be excluded by the static code checks. By default the modules 'unit-tests' and 'integration-tests' will be excluded.")
	cmd.Flags().StringVar(&stepConfig.SpotBugsExcludeFilterFile, "spotBugsExcludeFilterFile", os.Getenv("PIPER_spotBugsExcludeFilterFile"), "Path to a filter file with bug definitions which should be excluded.")
	cmd.Flags().StringVar(&stepConfig.SpotBugsIncludeFilterFile, "spotBugsIncludeFilterFile", os.Getenv("PIPER_spotBugsIncludeFilterFile"), "Path to a filter file with bug definitions which should be included.")
	cmd.Flags().IntVar(&stepConfig.SpotBugsMaxAllowedViolations, "spotBugsMaxAllowedViolations", 0, "The maximum number of failures allowed before execution fails.")
	cmd.Flags().IntVar(&stepConfig.PmdFailurePriority, "pmdFailurePriority", 0, "What priority level to fail the build on. PMD violations are assigned a priority from 1 (most severe) to 5 (least severe) according the the rule's priority. Violations at or less than this priority level are considered failures and will fail the build if failOnViolation=true and the count exceeds maxAllowedViolations. The other violations will be regarded as warnings and will be displayed in the build output if verbose=true. Setting a value of 5 will treat all violations as failures, which may cause the build to fail. Setting a value of 1 will treat all violations as warnings. Only values from 1 to 5 are valid.")
	cmd.Flags().IntVar(&stepConfig.PmdMaxAllowedViolations, "pmdMaxAllowedViolations", 0, "The maximum number of failures allowed before execution fails. Used in conjunction with failOnViolation=true and utilizes failurePriority. This value has no meaning if failOnViolation=false. If the number of failures is greater than this number, the build will be failed. If the number of failures is less than or equal to this value, then the build will not be failed.")
	cmd.Flags().StringVar(&stepConfig.ProjectSettingsFile, "projectSettingsFile", os.Getenv("PIPER_projectSettingsFile"), "Path to the mvn settings file that should be used as project settings file.")
	cmd.Flags().StringVar(&stepConfig.GlobalSettingsFile, "globalSettingsFile", os.Getenv("PIPER_globalSettingsFile"), "Path to the mvn settings file that should be used as global settings file.")
	cmd.Flags().StringVar(&stepConfig.M2Path, "m2Path", os.Getenv("PIPER_m2Path"), "Path to the location of the local repository that should be used.")
	cmd.Flags().BoolVar(&stepConfig.LogSuccessfulMavenTransfers, "logSuccessfulMavenTransfers", false, "Configures maven to log successful downloads. This is set to `false` by default to reduce the noise in build logs.")

}

// retrieve step metadata
func mavenExecuteStaticCodeChecksMetadata() config.StepData {
	var theMetaData = config.StepData{
		Metadata: config.StepMetadata{
			Name:    "mavenExecuteStaticCodeChecks",
			Aliases: []config.Alias{{Name: "mavenExecute", Deprecated: false}},
		},
		Spec: config.StepSpec{
			Inputs: config.StepInputs{
				Parameters: []config.StepParameters{
					{
						Name:        "spotBugs",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "bool",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "pmd",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "bool",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "mavenModulesExcludes",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "[]string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "spotBugsExcludeFilterFile",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "spotBugs/excludeFilterFile"}},
					},
					{
						Name:        "spotBugsIncludeFilterFile",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "spotBugs/includeFilterFile"}},
					},
					{
						Name:        "spotBugsMaxAllowedViolations",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "int",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "spotBugs/maxAllowedViolations"}},
					},
					{
						Name:        "pmdFailurePriority",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "int",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "pmd/failurePriority"}},
					},
					{
						Name:        "pmdMaxAllowedViolations",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "int",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "pmd/maxAllowedViolations"}},
					},
					{
						Name:        "projectSettingsFile",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"GENERAL", "STEPS", "STAGES", "PARAMETERS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "maven/projectSettingsFile"}},
					},
					{
						Name:        "globalSettingsFile",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"GENERAL", "STEPS", "STAGES", "PARAMETERS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "maven/globalSettingsFile"}},
					},
					{
						Name:        "m2Path",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"GENERAL", "STEPS", "STAGES", "PARAMETERS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "maven/m2Path"}},
					},
					{
						Name:        "logSuccessfulMavenTransfers",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"GENERAL", "STEPS", "STAGES", "PARAMETERS"},
						Type:        "bool",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "maven/logSuccessfulMavenTransfers"}},
					},
				},
			},
		},
	}
	return theMetaData
}