# Provider Evaluation Report

This example shows how to compare LLM providers/models for Smara workflows.

## Goal

Choose a provider/model based on:

- task quality
- latency
- cost
- tool-call reliability
- long-context behavior
- Indonesian/English response quality

## 1. List providers

```bash
smara provider list
```

## 2. Configure candidates

Example candidates:

```text
openai/gpt-4.1
openrouter/deepseek/deepseek-chat
anthropic/claude-3-5-sonnet
custom/local-model
ollama/qwen2.5-coder
```

## 3. Create an eval suite

Create `eval-suite.json`:

```json
{
  "name": "smara-docs-and-agent-eval",
  "cases": [
    {
      "name": "summarize-cli-reference",
      "prompt": "Summarize the Smara CLI command groups from cmd/smara in a docs-friendly outline."
    },
    {
      "name": "debug-web-session",
      "prompt": "Given a multi-session chat bug, propose a safe backend/frontend fix plan."
    },
    {
      "name": "write-vitepress-page",
      "prompt": "Write a VitePress guide for using Smara Graphify to analyze a Go codebase."
    }
  ]
}
```

## 4. Run evaluation

```bash
smara eval run --file eval-suite.json --json > eval-results.json
```

If you want to compare one provider at a time:

```bash
smara provider set openai
smara provider set-model gpt-4.1
smara eval run --file eval-suite.json --json > eval-openai.json

smara provider set openrouter
smara provider set-model deepseek/deepseek-chat
smara eval run --file eval-suite.json --json > eval-openrouter.json
```

## 5. Evaluate qualitative output

Score each case from 1-5:

| Criterion | Meaning |
|---|---|
| Accuracy | Correct about Smara/source behavior. |
| Completeness | Covers important edge cases. |
| Actionability | Produces commands/files/checklists. |
| Safety | Avoids risky destructive actions. |
| Conciseness | Useful without too much noise. |

## 6. Report template

```md
# Provider Evaluation Report

Date: YYYY-MM-DD
Smara version: vX.Y.Z
Eval suite: eval-suite.json

## Candidates

| Provider | Model | Notes |
|---|---|---|
| openai | gpt-4.1 | baseline |
| openrouter | deepseek/deepseek-chat | cost candidate |

## Results

| Case | Provider/model | Quality | Latency | Cost | Notes |
|---|---|---:|---:|---:|---|
| summarize-cli-reference | openai/gpt-4.1 | 5 | 8.2s | $... | strong structure |

## Recommendation

Use `provider/model` for default agent mode because ...
Use `provider/model` for docs generation because ...
```

## 7. Apply winner

```bash
smara provider set openai
smara provider set-model gpt-4.1
smara provider test
```

## 8. Track usage after change

```bash
smara analytics --days 7
```

Look for:

- request count
- token usage
- estimated cost
- repeated failures
- model distribution

## Notes

- Use the same prompts across providers.
- Run evals on the same network/environment when possible.
- Include tool-heavy cases if the provider will be used in agent mode.
- Keep a cheaper fallback model for routine docs cleanup.
