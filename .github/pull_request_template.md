## Description
<!---
  Describe your changes in detail.
-->
## Checklist
- [ ] 1.
- [ ] 2.
- [ ] 3.

```changes
section: <kebab-case of a module name> | <1st level dir in the repo>
type: fix | feature | chore
summary: <ONE-LINE of what effectively changes for a user>
impact: <what to expect for users, possibly MULTI-LINE>, required if impact_level is high ↓
impact_level: default | high | low
```

<!---
`impact_level: default` adds to changelog as usual, this is the default that can be omitted
`impact_level: high`    something important for users, the impact will be copied to "Know Before Update" section
`impact_level: low`     omitted in changelog YAML; note there is `type:chore` for chores

Tip for the section field:

  - <kebab-case of a module>, e.g. "cloud-provider-aws", "node-manager"
  - "ci", has forced low impact
  - "docs", includes website changes, should have low impact
