# Change Proposal: add-shared-frame-format

## Summary

Define a unified binary frame envelope for all CompressKit compressed files, adding a versioned header with magic bytes, algorithm ID, flags, content size, checksum, and optional extension fields (dictionary ID, skippable metadata blocks).

## Motivation

Currently each algorithm writes its own bespoke header (or none). The existing `REQ-ARCH-003` frequency table format is algorithm-specific. There is no:
- common magic to identify CompressKit files,
- version field to enable forward-compatibility,
- checksum to detect corruption,
- `content_size` hint to allow decoders to pre-allocate,
- extension point for future features (dictionaries, metadata).

Without a shared envelope, interoperability testing (`add-interoperability-conformance`) cannot define a canonical parse path, and the streaming API (`add-streaming-api-foundation`) has nowhere to write the end-of-stream marker.

## Scope

### In scope
- Byte-level frame format: layout, endianness, field widths
- Magic bytes value and rationale
- Version field semantics (current = 1)
- Algorithm ID registry
- Flags bitmap definitions
- `content_size` field (0 = unknown)
- `checksum` field (xxHash-64 of uncompressed content)
- Optional `dictionary_id` extension block
- Optional `skippable metadata` block mechanism
- Format versioning / backward-compatibility rules

### Out of scope
- Streaming API lifecycle (→ `add-streaming-api-foundation`)
- Conformance test matrix (→ `add-interoperability-conformance`)
- Benchmark infrastructure (→ `add-benchmark-governance`)
- The existing per-algorithm frequency table format (REQ-ARCH-003 unchanged)

## Impact

| Spec | Change type |
|------|-------------|
| `core-architecture/spec.md` | ADDED requirements REQ-ARCH-012, REQ-ARCH-013, REQ-ARCH-014 |
| `cross-language-testing/spec.md` | ADDED requirement REQ-TEST-007 |

## Dependencies

None. This change is self-contained; `add-streaming-api-foundation` references the frame format but does not define it.

## Risks

- Adding a new header breaks existing encoded files. Mitigation: version field + migration tooling task in tasks.md.
- xxHash-64 dependency may be unavailable in some embedded targets — flagged as a future concern.
- Skippable block mechanism must use a well-known marker to avoid false-positive detection.
