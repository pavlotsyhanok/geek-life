# Geek-Life v2 Roadmap

This file is now the working roadmap for the `geek-life` v2 release.

## Product Direction

`geek-life` v2 is a rebuild of the app around a modern Go TUI architecture.

Primary goals:

- Ship a maintainable, testable TUI built around explicit app state
- Replace widget-global coupling with a model-driven architecture
- Replace legacy persistence abstractions with a clearer storage layer
- Preserve the core strengths of the app: keyboard-first flow, fast startup, local-first storage, markdown-friendly task notes
- Reach a release-quality `v2.0.0` with CI, tests, packaging, and migration story

## Target Stack

Architecture decisions for v2:

- TUI runtime: `Bubble Tea v2`
- TUI components: `Bubbles`
- Styling and layout: `Lip Gloss v2`
- Markdown rendering: `Glamour`
- Storage: `SQLite`
- Query layer: `sqlc`
- Migrations: SQL migration files, with `Atlas` acceptable if it helps workflow
- Logging: `log/slog`
- Packaging and releases: `goreleaser`

## Release Strategy

Release approach:

- Stabilize the current codebase enough to keep it buildable during the rewrite
- Build v2 in parallel rather than incrementally mutating the old UI architecture
- Preserve existing user workflows where they are good, but do not preserve old implementation choices for compatibility alone
- Introduce explicit migration/import from the legacy local database if feasible

## Current Status

Completed before the v2 rebuild:

- [x] Raised the repo to a modern Go toolchain
- [x] Updated maintained direct dependencies in the current app
- [x] Fixed vet/build blockers in the legacy codebase
- [x] Added repository and utility tests around the current persistence layer
- [x] Added CI, linting, vuln scanning, and reproducible local build commands

These changes keep the repository healthy while v2 is developed.

## v2 Milestones

### Milestone 0: Foundation

- [ ] Create a v2 application layout in the repository
- [ ] Decide whether v2 lives under `cmd/geek-life`, `internal/`, and `pkg/`, or another clean structure
- [ ] Define versioning, branch strategy, and release criteria for `v2.0.0`
- [ ] Add architecture notes for the event model, storage layer, and screen model

### Milestone 1: Domain and Persistence

- [ ] Define the v2 domain model for projects, tasks, due dates, completion state, and note content
- [ ] Design the SQLite schema
- [ ] Add SQL migrations
- [ ] Generate typed queries with `sqlc`
- [ ] Implement repository/storage services on top of SQLite
- [ ] Add tests for CRUD, filtering, due-date queries, and migration behavior
- [ ] Decide whether notes stay inline in SQLite or move to separate markdown files

### Milestone 2: Application Core

- [ ] Implement the root Bubble Tea model
- [ ] Define message types, commands, and async boundaries
- [ ] Remove package-global UI state in v2
- [ ] Add a central app state container for selected project, selected task, filters, edit mode, and status messages
- [ ] Add a command/service layer for storage, clipboard, config, and external editor integration
- [ ] Add unit tests for model transitions and update logic

### Milestone 3: Core TUI Screens

- [ ] Build the projects pane/list
- [ ] Build the tasks pane/list
- [ ] Build the task detail view
- [ ] Build keyboard navigation between views
- [ ] Build status and help surfaces
- [ ] Preserve fast keyboard-first flows from v1
- [ ] Verify desktop terminal behavior on Linux, macOS, and Windows terminals

### Milestone 4: Task Editing Experience

- [ ] Implement task creation and rename flows
- [ ] Implement due-date editing flow
- [ ] Implement completion toggle flow
- [ ] Implement markdown note editing flow
- [ ] Decide whether in-app note editing uses `textarea`, split edit/preview, or external-editor-first mode
- [ ] Implement markdown preview with `Glamour`
- [ ] Re-add clipboard export

### Milestone 5: Dynamic Views and Filtering

- [ ] Rebuild Today view
- [ ] Rebuild Tomorrow view
- [ ] Rebuild Upcoming view
- [ ] Rebuild Unscheduled view
- [ ] Add search/filter support
- [ ] Consider tags, pinning, and saved filters if they fit the release scope

### Milestone 6: Migration and Compatibility

- [ ] Define a migration path from the legacy Storm/Bolt database
- [ ] Build a one-time import command or migration utility
- [ ] Test migration on real sample data
- [ ] Document backward compatibility limits clearly

### Milestone 7: Config, Packaging, and Release

- [ ] Add config file support
- [ ] Use XDG-compliant default paths
- [ ] Add build-time version injection for releases
- [ ] Add `goreleaser` configuration
- [ ] Produce release artifacts for Linux, macOS, and Windows
- [ ] Document installation and upgrade instructions
- [ ] Prepare release notes for `v2.0.0`

### Milestone 8: Quality Gate

- [ ] Add end-to-end interaction coverage where practical
- [ ] Add snapshot/render tests for critical screens where helpful
- [ ] Run `go test`, `go vet`, `staticcheck`, and `govulncheck` in CI for v2
- [ ] Add manual QA checklist for terminal rendering and keyboard flows
- [ ] Freeze release scope
- [ ] Cut `v2.0.0-rc1`
- [ ] Fix release-candidate issues
- [ ] Ship `v2.0.0`

## Proposed Repository Shape

Proposed direction, subject to refinement:

- `cmd/geek-life/` for the executable entrypoint
- `internal/app/` for the Bubble Tea root model and screen orchestration
- `internal/domain/` for core entities and business rules
- `internal/storage/` for SQLite and generated queries
- `internal/clipboard/` for clipboard integration
- `internal/config/` for config loading and path resolution
- `internal/migrate/` for schema and legacy-import logic

## Non-Goals For Initial v2 Release

Do not expand scope unnecessarily before `v2.0.0`:

- Cloud sync
- Multi-user collaboration
- Plugin systems
- Full embedded IDE/editor behavior for notes
- Complex calendar views

## First Build Slice

The first meaningful vertical slice for v2 should be:

- [ ] Open app
- [ ] Load projects from SQLite
- [ ] Select a project
- [ ] View its tasks
- [ ] Create a task
- [ ] Toggle task completion
- [ ] Quit cleanly

Once that slice exists, build outward from it.

## Open Decisions

- [ ] Should note editing be inline-first or external-editor-first?
- [ ] Should we keep the three-pane mental model exactly, or modernize the layout while preserving shortcuts?
- [ ] Should we preserve the current on-disk DB path for import discovery only, or also for continued storage?
- [ ] Should `v2` coexist in the same binary initially, or ship as a clearly separate release candidate path until stable?

## Immediate Next Tasks

The next implementation tasks should be:

- [ ] Create the v2 directory structure
- [ ] Add a minimal Bubble Tea app that boots and renders a placeholder shell
- [ ] Add the SQLite schema and first migration
- [ ] Add `sqlc` configuration and generated query scaffolding
- [ ] Wire the root model to load projects from the new storage layer
