# Exercise for Week 3

## Export our Social Media feeds to JSON and XML file types

Your task is to extend the [social media program](https://github.com/timolinn/go_training/blob/fc60d4d9a0ae33abbfe0d4d0c96cc23489da809f/week3/cmd/main.go#L31) that we wrote in the video lesson.

In the video, I wrote an `export` function that fetches the social media feeds and writes them to a `txt` file. The function signature looked like this:

```go
    func export(u week3.SocialMedia, filename string) error
```

Your job is to extend this function by making it possible to export to multiple file types depending on what the developer wants. Here's is a step by step guide:

+ Create a new package call it `exporter`. Ensure that this package contains only features that perform tasks related to exporting social media data to a file
+ Write a function or functions that are exported in the `exporter` package, which our `main` program can use to export social media feeds from Twitter, Facebook and Linkedin to the following file types:

  + JSON
  + XML
  + _YAML (Optional)_

  Note that they must contain valid data structures of the file type eg. `fbdata.json` must contain VALID json data.
  
  ```json
    { "1": "Facebook feeds", "2": "Hey, here's my cool new selfie" }
  ```

+ Write a program in your `main.go` that use this package to export data from all the existing social media platforms (mentioned above).
+ Upload your solution to github INCLUDING the exported data and submit the project link on HNG Board.

> I know I have not specifically treated how to handle JSON or XML data in Go, so I don't expect you to know this 100%. However, I do expect you to make good efforts to find the solution to this problem online or through peer help. My goal is that you search for the solution yourself and learn from it, afterall most of your software engineering career will revolve around finding solutions online or from your team mates and peer.

## Resources

+ <https://blog.golang.org/json>
+ <https://gobyexample.com/json>
+ <https://gobyexample.com/xml>
