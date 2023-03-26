// This is a more verbose version of the openapi.yaml example file - with comments.

{
    "schema_version": "v1",
    "name_for_human": "Fun facts", // this will be displayed in the ChatGPT Plugins UI
    "name_for_model": "Hi_and_pi", // can have digits, letters and under_score
    "description_for_human": "Different fun facts, like how to say hi in different languages, or the first 20 digits of Pi",
    "description_for_model": "Fun facts like how to say hi in different langauges: thai, english, chinese, hebrew and arabic; and the first 20 digits of Pi",
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
        "response_format": {
            … // response format and examples go here
        }
      }
    ]
  }
  
  
  