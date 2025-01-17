package shared

import (
	"fmt"
	"strings"

	"github.com/cli/cli/pkg/iostreams"
)

func RenderRunHeader(cs *iostreams.ColorScheme, run Run, ago, prNumber string) string {
	title := fmt.Sprintf("%s %s%s",
		cs.Bold(run.HeadBranch), run.Name, prNumber)
	symbol, symbolColor := Symbol(cs, run.Status, run.Conclusion)
	id := cs.Cyanf("%d", run.ID)

	header := ""
	header += fmt.Sprintf("%s %s · %s\n", symbolColor(symbol), title, id)
	header += fmt.Sprintf("Triggered via %s %s", run.Event, ago)

	return header
}

func RenderJobs(cs *iostreams.ColorScheme, jobs []Job, verbose bool) string {
	lines := []string{}
	for _, job := range jobs {
		symbol, symbolColor := Symbol(cs, job.Status, job.Conclusion)
		id := cs.Cyanf("%d", job.ID)
		lines = append(lines, fmt.Sprintf("%s %s (ID %s)", symbolColor(symbol), job.Name, id))
		if verbose || IsFailureState(job.Conclusion) {
			for _, step := range job.Steps {
				stepSymbol, stepSymColor := Symbol(cs, step.Status, step.Conclusion)
				lines = append(lines, fmt.Sprintf("  %s %s", stepSymColor(stepSymbol), step.Name))
			}
		}
	}

	return strings.Join(lines, "\n")
}

func RenderAnnotations(cs *iostreams.ColorScheme, annotations []Annotation) string {
	lines := []string{}

	for _, a := range annotations {
		lines = append(lines, fmt.Sprintf("%s %s", AnnotationSymbol(cs, a), a.Message))
		lines = append(lines, cs.Grayf("%s: %s#%d\n", a.JobName, a.Path, a.StartLine))
	}

	return strings.Join(lines, "\n")
}
