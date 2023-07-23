// Bytes are single quoted and base64 encoded when output:

b: '\x03abc\U0001F604'

// cue eval bytes.cue
// cue export bytes.cue --out json