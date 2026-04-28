# Streaming API Contract Test Plan

This directory locks the streaming API contract before any shared streaming implementation is introduced.

At this stage, the artifacts here are planning inputs for test-driven development only:

- `contract_cases.md` defines the lifecycle and error cases that must become failing tests first.
- No production code changes are allowed until those failing tests exist.
- The initial failure check must be performed against the current Huffman and Range API surfaces, because they expose the most obvious contract gaps.
