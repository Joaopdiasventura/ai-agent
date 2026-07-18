# Business Rules

## Question Must Produce Searchable Tokens

- Condition: the question is empty after tokenization or contains no usable token.
- Processing: `search.Engine.Search` returns `Found=false`.
- Result: CLI or HTTP returns a fallback path.
- Component: `internal/search` and `internal/tokenizer`.

## HTTP Body Must Be A Single Valid JSON Object

- Condition: the API receives invalid JSON, an empty body, more than one JSON object, an unknown field, or an invalid field type.
- Processing: `internal/handlers/ask` rejects the request before calling `app.AgentResponse`.
- Result: the handler returns `400` with a JSON `response` field.
- Component: `internal/handlers/ask`.

## HTTP Content Is Required

- Condition: `content` is missing, empty, or whitespace after trimming.
- Processing: the handler rejects the request.
- Result: `400` with `O campo content é obrigatório.`
- Component: `internal/handlers/ask`.

## Request Body Size Is Limited

- Condition: request body exceeds `maxRequestBodyBytes`, currently `2048`.
- Processing: `http.MaxBytesReader` causes decoding to fail with `http.MaxBytesError`.
- Result: `413` with a JSON error response.
- Component: `internal/handlers/ask`.

## Language Controls Document Selection

- Condition: a question is detected as Portuguese or English.
- Processing: `FilterDocumentsByIntent` ignores documents whose `Document.Language` does not match `QueryAnalysis.Language`.
- Result: Portuguese questions return `pt` documents; English questions return `en` documents.
- Component: `internal/nlp` and `internal/search`.

## Similarity Threshold Controls Answerability

- Condition: ranked results have similarity lower than `minimumSimilarity`, currently `0.1`.
- Processing: `FilterRelevantResults` removes them.
- Result: if no result remains, the app returns a localized fallback.
- Component: `internal/search`.

## Entity-Specific Results Must Mention The Entity

- Condition: an entity is detected in the question.
- Processing: `FilterRelevantResults` requires `result.Document.Content` to contain `analysis.Entity.Value`.
- Result: results that do not mention the resolved entity value are excluded.
- Component: `internal/search`.

## Project And Technology Intents Can Return Multiple Documents

- Condition: intent is `IntentTechnologies` or `IntentProject`.
- Processing: `ShouldSearchMultipleDocuments` allows using the configured result limit.
- Result: up to `maximumSearchResults`, currently `5`, can be used for the response.
- Component: `internal/search` and `internal/app`.

## Localized Fallbacks Are Returned Instead Of Fabricated Answers

- Condition: search returns `Found=false`.
- Processing: `app.NotFoundMessage` selects a Portuguese or English fallback.
- Result: the system does not generate an answer outside the retrieved facts.
- Component: `internal/app`.
