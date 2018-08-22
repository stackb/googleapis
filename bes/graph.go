package main

import (
	"log"

	bes "go.stack.build/github.com/bazelbuild/bazel/build_event_stream/d6b40d94"
)

type BuildEventListener interface {
	HandleAborted(b *bes.BuildEvent, e *bes.Aborted)
	HandleProgress(b *bes.BuildEvent, e *bes.Progress)
	HandleConfiguration(b *bes.BuildEvent, id *bes.BuildEventId_ConfigurationId, e *bes.Configuration)
	HandleNamedSetOfFiles(b *bes.BuildEvent, id *bes.BuildEventId_NamedSetOfFilesId, e *bes.NamedSetOfFiles)
	HandleTargetComplete(b *bes.BuildEvent, id *bes.BuildEventId_TargetCompletedId, e *bes.TargetComplete)
	HandleBuildStarted(b *bes.BuildEvent, e *bes.BuildStarted)
	HandleBuildFinished(b *bes.BuildEvent, e *bes.BuildFinished)
}

func NotifyBuildEvents(events []*bes.BuildEvent, listeners []BuildEventListener) {
	for _, listener := range listeners {
		for _, event := range events {
			NotifyBuildEvent(event, listener)
		}
	}
}

func NotifyBuildEvent(event *bes.BuildEvent, listener BuildEventListener) {
	switch u := event.Payload.(type) {
	case *bes.BuildEvent_Progress:
		log.Printf("Got progress event: %s: %v\n", u, u.Progress)
		listener.HandleProgress(event, u.Progress)

	case *bes.BuildEvent_Aborted:
		log.Printf("Got aborted event: %s: %v\n", u, u.Aborted)
		listener.HandleAborted(event, u.Aborted)

	case *bes.BuildEvent_Started:
		log.Printf("Got started event: %s: %v\n", u, u.Started)
		listener.HandleBuildStarted(event, u.Started)

	case *bes.BuildEvent_UnstructuredCommandLine:
		log.Printf("Got unstructured command line event: %s: %v\n", u, u.UnstructuredCommandLine)

	case *bes.BuildEvent_StructuredCommandLine:
		log.Printf("Got structure command line event: %s: %v\n", u, u.StructuredCommandLine)

	case *bes.BuildEvent_OptionsParsed:
		log.Printf("Got optionsParsed event: %s: %v\n", u, u.OptionsParsed)

	case *bes.BuildEvent_WorkspaceStatus:
		log.Printf("Got workspaceStatus event: %s: %v\n", u, u.WorkspaceStatus)

	case *bes.BuildEvent_Fetch:
		log.Printf("Got fetch event: %s: %v\n", u, u.Fetch)

	case *bes.BuildEvent_Configuration:
		log.Printf("Got configuration event: %s: %v\n", u, u.Configuration)
		buildId := event.GetId()
		switch id := buildId.Id.(type) {
		case *bes.BuildEventId_Configuration:
			listener.HandleConfiguration(event, id.Configuration, u.Configuration)
		}

	case *bes.BuildEvent_Expanded:
		log.Printf("Got expanded event: %s: %v\n", u, u.Expanded)

	case *bes.BuildEvent_Configured:
		log.Printf("Got configured event: %s: %v\n", u, u.Configured)

	case *bes.BuildEvent_Action:
		log.Printf("Got action event: %s: %v\n", u, u.Action)

	case *bes.BuildEvent_NamedSetOfFiles:
		log.Printf("Got namedSetOfFiles event: %s: %v\n", u, u.NamedSetOfFiles)
		buildId := event.GetId()
		switch id := buildId.Id.(type) {
		case *bes.BuildEventId_NamedSet:
			listener.HandleNamedSetOfFiles(event, id.NamedSet, u.NamedSetOfFiles)
		}

	case *bes.BuildEvent_Completed:
		log.Printf("Got completed event: %s: %v\n", u, u.Completed)
		// for _, g := range u.Completed.OutputGroup {
		// 	log.Printf(" --> OutputGroup: %s\n", g.Name)
		// 	for _, s := range g.FileSets {
		// 		log.Printf("     --+ %s\n", s.Id)
		// 	}
		// }

		buildId := event.GetId()
		switch id := buildId.Id.(type) {
		case *bes.BuildEventId_TargetCompleted:
			listener.HandleTargetComplete(event, id.TargetCompleted, u.Completed)
		}

	case *bes.BuildEvent_TestResult:
		log.Printf("Got testResult event: %s: %v\n", u, u.TestResult)

	case *bes.BuildEvent_TestSummary:
		log.Printf("Got testSummary event: %s: %v\n", u, u.TestSummary)

	case *bes.BuildEvent_Finished:
		log.Printf("Got finished event: %s: %v\n", u, u.Finished)
		listener.HandleBuildFinished(event, u.Finished)

	case *bes.BuildEvent_BuildMetrics:
		log.Printf("Got build metrics event: %s: %v\n", u, u.BuildMetrics)

	case *bes.BuildEvent_BuildToolLogs:
		log.Printf("Got build tool logs event: %s: %v\n", u, u.BuildToolLogs)

	default:
		log.Fatalf("Unknown event type: %v", u)
	}
}

// Callback for agents that want to be notified when build successful and
// default output group.
type DefaultTargetCompleteListener interface {
	HandleDefaultTargetComplete(g *BuildEventGraph, id *bes.BuildEventId_TargetCompletedId, e *bes.TargetComplete)
}

// Callback for agents that want to be notified when build successful
// and default output group.
type BuildOutcomeListener interface {
	HandleBuildStarted(e *bes.BuildStarted)
	HandleBuildSuccess(e *bes.BuildFinished)
	HandleBuildFailure(e *bes.BuildFinished)
}

func NewBuildEventGraph() *BuildEventGraph {
	return &BuildEventGraph{
		NamedFileSets:    make(map[string]*bes.NamedSetOfFiles),
		NamedFileSetIds:  make(map[string]*bes.BuildEventId_NamedSetOfFilesId),
		Configurations:   make(map[string]*bes.Configuration),
		ConfigurationIds: make(map[string]*bes.BuildEventId_ConfigurationId),
		Completions:      make(map[string][]*bes.TargetComplete),
		CompletionIds:    make(map[string][]*bes.BuildEventId_TargetCompletedId),
	}
}

type BuildEventGraph struct {
	Aborted  *bes.Aborted
	Started  *bes.BuildStarted
	Finished *bes.BuildFinished
	// Mapping from Id to fileset
	NamedFileSets   map[string]*bes.NamedSetOfFiles
	NamedFileSetIds map[string]*bes.BuildEventId_NamedSetOfFilesId
	// Mapping from configration id (a hash) to Configuration object.
	Configurations   map[string]*bes.Configuration
	ConfigurationIds map[string]*bes.BuildEventId_ConfigurationId
	// Mapping from output group name ('default') to the list of Targets that completed within that group.
	Completions   map[string][]*bes.TargetComplete
	CompletionIds map[string][]*bes.BuildEventId_TargetCompletedId
}

func (l *BuildEventGraph) HandleProgress(b *bes.BuildEvent, e *bes.Progress) {
}

func (l *BuildEventGraph) HandleAborted(b *bes.BuildEvent, e *bes.Aborted) {
	l.Aborted = e
}

func (l *BuildEventGraph) HandleConfiguration(b *bes.BuildEvent, id *bes.BuildEventId_ConfigurationId, e *bes.Configuration) {
	//log.Printf("Configuration %s '%s' platform=%s, cpu=%s ----------------", id.GetId(), e.GetMnemonic(), e.GetPlatformName(), e.GetCpu())
	//for k, v := range e.GetMakeVariable() {
	//	log.Printf(" ---| %s=%s\n", k, v)
	//}
	l.Configurations[id.GetId()] = e
	l.ConfigurationIds[id.GetId()] = id
}

func (l *BuildEventGraph) HandleNamedSetOfFiles(b *bes.BuildEvent, id *bes.BuildEventId_NamedSetOfFilesId, e *bes.NamedSetOfFiles) {
	//log.Printf("NamedSetOfFiles %s ----------------", id.GetId())
	// for _, f := range e.GetFiles() {
	// 	log.Printf(" --- fs %s contains %s ----------------", id.Id, f.Name)
	// }
	// for _, fsId := range e.GetFileSets() {
	// 	log.Printf(" --- fs %s depends on %s ----------------", id.Id, fsId.Id)
	// }

	l.NamedFileSets[id.GetId()] = e
	l.NamedFileSetIds[id.GetId()] = id
}

func (l *BuildEventGraph) HandleTargetComplete(b *bes.BuildEvent, id *bes.BuildEventId_TargetCompletedId, e *bes.TargetComplete) {
	//log.Printf("TargetComplete %s %s (%s) via configuration %s", e.GetSuccess(), id.GetLabel(), e.GetTargetKind(), id.GetConfiguration().GetId())
	for _, g := range e.GetOutputGroup() {
		name := g.GetName()
		l.Completions[name] = append(l.Completions[name], e)
		l.CompletionIds[name] = append(l.CompletionIds[name], id)
	}
}

func (l *BuildEventGraph) HandleBuildStarted(b *bes.BuildEvent, e *bes.BuildStarted) {
	l.Started = e
}

func (l *BuildEventGraph) HandleBuildFinished(b *bes.BuildEvent, e *bes.BuildFinished) {
	l.Finished = e
}

func (l *BuildEventGraph) GetDefaultCompletions() []*bes.TargetComplete {
	return l.Completions["default"]
}

func (l *BuildEventGraph) GetDefaultCompletionIds() []*bes.BuildEventId_TargetCompletedId {
	return l.CompletionIds["default"]
}

func (l *BuildEventGraph) VisitAllFiles(fileSet *bes.NamedSetOfFiles, visitor func(parent *bes.NamedSetOfFiles, file *bes.File)) {
	for _, id := range fileSet.GetFileSets() {
		s := l.NamedFileSets[id.GetId()]
		l.VisitAllFiles(s, visitor)
	}
	for _, file := range fileSet.GetFiles() {
		visitor(fileSet, file)
	}
}

func (l *BuildEventGraph) GetFirstDefaultOutputOfTarget(e *bes.TargetComplete) *bes.File {
	for _, g := range e.GetOutputGroup() {
		if g.GetName() == "default" {
			for _, id := range g.GetFileSets() {
				outSet := l.NamedFileSets[id.GetId()]
				if outSet == nil {
					return nil
				}
				outFiles := outSet.GetFiles()
				return outFiles[0]
			}
		}
	}
	return nil
}

// ****************************************************************
// DefaultCompletionNotifier
// ****************************************************************

func NewDefaultCompletionNotifier(graph *BuildEventGraph, listeners []DefaultTargetCompleteListener) *DefaultCompletionNotifier {
	return &DefaultCompletionNotifier{
		graph:     graph,
		listeners: listeners,
	}
}

type DefaultCompletionNotifier struct {
	graph     *BuildEventGraph
	listeners []DefaultTargetCompleteListener
}

func (l *DefaultCompletionNotifier) HandleProgress(b *bes.BuildEvent, e *bes.Progress) {
}

func (l *DefaultCompletionNotifier) HandleConfiguration(b *bes.BuildEvent, id *bes.BuildEventId_ConfigurationId, e *bes.Configuration) {
	l.graph.HandleConfiguration(b, id, e)
}

func (l *DefaultCompletionNotifier) HandleNamedSetOfFiles(b *bes.BuildEvent, id *bes.BuildEventId_NamedSetOfFilesId, e *bes.NamedSetOfFiles) {
	l.graph.HandleNamedSetOfFiles(b, id, e)
}

func (l *DefaultCompletionNotifier) HandleTargetComplete(b *bes.BuildEvent, id *bes.BuildEventId_TargetCompletedId, e *bes.TargetComplete) {
	l.graph.HandleTargetComplete(b, id, e)
}

func (l *DefaultCompletionNotifier) HandleBuildStarted(b *bes.BuildEvent, e *bes.BuildStarted) {
	l.graph.HandleBuildStarted(b, e)
}

func (l *DefaultCompletionNotifier) HandleBuildFinished(b *bes.BuildEvent, e *bes.BuildFinished) {
	l.graph.HandleBuildFinished(b, e)

	if e.GetOverallSuccess() {
		l.handleBuildSuccess()
	}
}

func (l *DefaultCompletionNotifier) handleBuildSuccess() {
	targets := l.graph.GetDefaultCompletions()
	ids := l.graph.GetDefaultCompletionIds()
	n := len(ids)
	for i := 0; i < n; i++ {
		l.handleDefaultCompletion(ids[i], targets[i])
	}
}

func (l *DefaultCompletionNotifier) handleDefaultCompletion(id *bes.BuildEventId_TargetCompletedId, e *bes.TargetComplete) {
	for _, c := range l.listeners {
		c.HandleDefaultTargetComplete(l.graph, id, e)
	}
}

// ****************************************************************
// BuildOutcomeNotifier
// ****************************************************************

func NewBuildOutcomeNotifier(listeners []BuildOutcomeListener) *BuildOutcomeNotifier {
	return &BuildOutcomeNotifier{
		listeners: listeners,
	}
}

type BuildOutcomeNotifier struct {
	listeners []BuildOutcomeListener
}

func (l *BuildOutcomeNotifier) HandleProgress(b *bes.BuildEvent, e *bes.Progress) {
}

func (l *BuildOutcomeNotifier) HandleConfiguration(b *bes.BuildEvent, id *bes.BuildEventId_ConfigurationId, e *bes.Configuration) {
}

func (l *BuildOutcomeNotifier) HandleNamedSetOfFiles(b *bes.BuildEvent, id *bes.BuildEventId_NamedSetOfFilesId, e *bes.NamedSetOfFiles) {
}

func (l *BuildOutcomeNotifier) HandleTargetComplete(b *bes.BuildEvent, id *bes.BuildEventId_TargetCompletedId, e *bes.TargetComplete) {
}

func (l *BuildOutcomeNotifier) HandleBuildStarted(b *bes.BuildEvent, e *bes.BuildStarted) {
	for _, c := range l.listeners {
		c.HandleBuildStarted(e)
	}
}

func (l *BuildOutcomeNotifier) HandleBuildFinished(b *bes.BuildEvent, e *bes.BuildFinished) {
	for _, c := range l.listeners {
		if e.GetOverallSuccess() {
			c.HandleBuildSuccess(e)
		} else {
			c.HandleBuildFailure(e)
		}
	}
}
