This is a the same version of the `ai-plugin.json` example file, but with comments"

```json
{
    "schema_version": "v1",
    "name_for_human": "Fun facts", // this will be displayed in the ChatGPT Plugins UI, 50 character max
    "name_for_model": "Hi_and_pi", // can have digits, letters and under_score, 50 character max
    "description_for_human": "Different fun facts, like how to say hi in different languages, or the first 20 digits of Pi", // 120 character max
    "description_for_model": "Fun facts like how to say hi in different langauges: thai, english, chinese, hebrew and arabic; and the first 20 digits of Pi", // currently 8000 chars max
    "auth": {
      "type": "none" 
    },
    "api": {
      "type": "openapi",
      "url": "http://localhost:8081/openapi.yaml", // link to the other file you’re making
      "is_user_authenticated": false
    },
    "logo_url": "http://localhost:8081/image.jpg",
    "contact_email": "hi@no-real-email-here.com",
    "endpoints": [
      {
        "name": "getPi",
        "description": "Endpoint for saying hi in different languages",
        "url": "http://localhost:8081/hi", // or your Tunneling alternative, like ngrok
        "request_type": "GET", // or any other HTTP method
        "response_type": "JSON", 
        "response_format": { // 100k character limit for the API response
            … // response format and examples go here
        }
      }
    ]
  }
  ```
  
All of the possible options for the plugin fields:
 ```md
| Field                      | Type                  | Description / Options                                                   |
| -------------------------- | --------------------- | ----------------------------------------------------------------------- |
| schema_version             | String                | Manifest schema version                                                 | 
| name_for_model             | String                | Name the model will used to target the plugin                           |
| name_for_human             | String                | Human-readable name, such as the full company name                      |
| description_for_model      | String                | Description better tailored to the model, such as token context length considerations or keyword usage for improved plugin prompting. |
| description_for_human      | String                | Human-readable description of the plugin                                |
| auth                       | ManifestAuth          | Authentication scheme                                                   |
| api                        | Object                | API specification                                                       |
| logo_url                   | String                | URL used to fetch the plugin's logo                                     |
| contact_email              | String                | Email contact for safety/moderation reachout, support, and deactivation |
| legal_info_url             | String                | Redirect URL for users to view plugin information                       |
| HttpAuthorizationType      | HttpAuthorizationType | "bearer" or "basic"                                                     |
| ManifestAuthType           | ManifestAuthType      | "none",  "user_http", "service_http", or "oauth"                        |
| interface BaseManifestAuth | BaseManifestAuth      | type: ManifestAuthType; instructions: string;                           |
| ManifestNoAuth             | ManifestNoAuth        | No authentication required: BaseManifestAuth & { type: 'none', }        |
| ManifestAuth               | ManifestAuth          | ManifestNoAuth, ManifestServiceHttpAuth, ManifestUserHttpAuth, ManifestOAuthAuth |
```

The following are examples with different authentication methods: 
```yaml
// App-level API keys
type ManifestServiceHttpAuth  = BaseManifestAuth & {
  type: 'service_http',
  authorization_type: HttpAuthorizationType
  verification_tokens: {
    [service: string]?: string
  }
}

// User-level HTTP authentication
type ManifestUserHttpAuth  = BaseManifestAuth & {
  type: 'user_http'
  authorization_type: HttpAuthorizationType
}

type ManifestOAuthAuth  = BaseManifestAuth & {
  type: 'oauth',

  // OAuth URL where a user is directed to for the OAuth authentication flow to begin.
  client_url: string

  // OAuth scopes required to accomplish operations on the user's behalf.
  scope: string

  // Endpoint used to exchange OAuth code with access token.
  authorization_url: string

  // When exchanging OAuth code with access token, the expected header 'content-type'. For example: - "application/json": 'content-type: application/json'
 authorization_content_type: string;

  // When registering the OAuth client ID and secrets, the plugin service will surface a unique token. 
verification_tokens: {
[service: string]?: string
}
}
```
