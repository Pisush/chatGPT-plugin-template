# chatGPT-plugin-template
For Gophers!


-------


Here are some motivational screenshots:

<img width="685" alt="Screenshot 2023-03-26 at 20 35 16" src="https://user-images.githubusercontent.com/7592392/227796730-1bbaef6e-b8b7-487c-8b1b-6273e75e7af8.png">

<img width="1286" alt="Screenshot 2023-03-26 at 20 34 21" src="https://user-images.githubusercontent.com/7592392/227796699-8253da12-4b73-4e32-b22e-fe5a3bd5d289.png">


-------


## What are ChatGPT plugins?
In 2 sentences: It's a way to enable ChatGPT to use your app's API. You share the specs, and it can translate a query from the end user to an API call.
The full answer is available [here](https://openai.com/blog/chatgpt-plugins). ATM it's in limited alpha.


-------


## Steps for Building a ChatGPT Plugin Based on Your Go Web Server
You need to add at least 2 files from root:
1. Add and edit the file `openapi.yaml`
2. Add and edit the file `.well-known/ai-plugin.json` (you might need to sudo this)
As comments are not allowed in a JSON file, check out `ai-plugin_example.md` for some more details on structuring this file
3. As a sanity check, double check the paths from file 1 match the endpoints from path 2, and that you changed all the ngrok URLs to your URLs.
4. execute: `$ go run main.go`
5. Add the plugin via the ChatGPT UI by going to the Plugin Store -> Install An Unverified Plugin, and insert the top website domain

### FYI
- If you are running in localhost, use a tunneling tool like ngrok to expose the files. Change in `openapi.yaml` and in `.well-known/ai-plugin.json`. And your `main.go` can remain on localhost.
- As long as the plugin is unverified, max 15 people can install the plugin

-------


## Debug
**Problem:**
You can't access the `.well-known/ai-plugin.json` file, it might be a permissions issues.

**Solution:**
Go to the root of the repo (e.g. `~/Projects/chatGPT-plugin-template`) and run `ls -la`. If you see that the permissions for the `.well-known` folder look different - make it the same.

-------

**Problem:**
You are running on localhost and the file `.well-known/ai-plugin.json` exists in the right place, but the ChatGPT Plugin UI can't find a manifest.

**Solution:**
Use a tunneling tool like ngrok to expose the files. Change to the ngrok url in `openapi.yaml` and in `.well-known/ai-plugin.json`. And your `main.go` can remain on localhost.

-------

**Problem:**
You have successfully added the plugin but it shows an empty field.

**Solution:**
Get your hands dirty to make the prompt engineering easy: change the descriptions and use more keywords for the different actions, do this in the files: `openapi.yaml` and  `.well-known/ai-plugin.json`.

-------

**Problem:**
The plugin was added successfully, but returns the error `ResponseTooLargeError`

**Solution:**
The response is limited at 100k characters. If your response is designed to be smaller, it can be that there's an error in the API and it's just returning the entire website. This is mostly common at private projects where security measurements are not in place.
