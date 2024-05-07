# Nuclei: Automating Web Application and Network Service Testing [Cheat Sheet]

![img](https://miro.medium.com/v2/resize:fit:498/0*mhpaOsayH5Uqntqn.png)

Nuclei is an open-source framework designed for automating the detection and exploitation of vulnerabilities in web applications and other network services. It allows penetration testers and security researchers to define custom templates that specify a set of HTTP requests to send to a target, along with corresponding matching rules that can be used to identify vulnerabilities or misconfigurations.

## Installation

To install Nuclei, you can download the latest version from the official [GitHub](https://github.com/projectdiscovery/nuclei) repository and follow the instructions

## Cheat Sheet

here is a comprehensive cheat sheet with some commonly used Nuclei commands for bug bounty hunting:

```
# Display help information
nuclei -h

# Display the current version of Nuclei
nuclei -version

# Load a list of targets from a file
nuclei -l targets.txt -t ~/nuclei-templates/

# Specify a single target to test
nuclei -t https://example.com -t ~/nuclei-templates/

# Specify a URL to test
nuclei -u https://example.com -t ~/nuclei-templates/

# Run Nuclei in silent mode (suppress output)
nuclei -silent -t https://example.com -t ~/nuclei-templates/

# Specify the number of concurrent threads to use
nuclei -c 10 -t https://example.com -t ~/nuclei-templates/

# Skip templates that require authentication
nuclei -no-verify -t https://example.com -t ~/nuclei-templates/

# Customize the output format of the Nuclei report
nuclei -o output.txt -t https://example.com -t ~/nuclei-templates/

# Ignore SSL certificate errors
nuclei -insecure -t https://example.com -t ~/nuclei-templates/

# Specify a custom HTTP header to include in requests
nuclei -headers "Authorization: Bearer TOKEN" -t https://example.com -t ~/nuclei-templates/

# Specify a custom user agent string to include in requests
nuclei -user-agent "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3" -t https://example.com -t ~/nuclei-templates/
```

These commands should help you get started with running Nuclei and customizing your tests according to your specific needs. It’s important to note that Nuclei supports many more commands and options beyond these, so be sure to consult the official documentation for more information.

## Templates

To create a template for Nuclei, follow these steps:

1. Choose a target to test, such as a web application or network service.
2. Identify a vulnerability or misconfiguration that you want to test for.
3. Use the Nuclei template language to define the test. The template language is based on YAML syntax and supports a wide range of features, such as HTTP requests, response matching, and variables. Refer to the Nuclei documentation for more information on the template language and its features.
4. Save the template to a file with a `.yaml` extension. It's recommended to save templates in a folder within the Nuclei templates directory for easier management and organization.
5. Test the template by running Nuclei with the `-t` flag and specifying the path to the template file. For example:

```
nuclei -t https://example.com -t ~/nuclei-templates/my-template.yaml
```

Analyze the output to determine if the test was successful and if any vulnerabilities or misconfigurations were identified. Make changes to the template as necessary and repeat the testing process until you achieve the desired results.

**Remember to test your templates carefully and responsibly, and always obtain permission from the target before conducting any vulnerability testing.*

**Custom Nuclei template SQL Injection script with explanations**:

```
id: example-template   # A unique identifier for the template
info:                   # Information about the template
  name: Example Template
  author: Your Name
  severity: low

requests:               # List of HTTP requests to send
  - method: GET         # HTTP method (GET, POST, etc.)
    path: /             # Request path
    headers:            # Optional request headers
      User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64)
    matchers:           # List of matchers to apply to the response
      - type: word       # Matcher type (word, status, regex, etc.)
        words:
          - "Example"   # Keyword to search for in the response
    attacks:            # Optional list of attack payloads to send
      - payload: "' OR true--"
```

Explanation:

- `id`: A unique identifier for the template that can be used to reference it in command-line arguments and other templates.
- `info`: Metadata about the template, such as the name, author, and severity of the vulnerabilities or misconfigurations it tests for.
- `requests`: A list of HTTP requests to send to the target. This section defines the HTTP method, path, headers, and other attributes of the requests. The requests can be customized further with variables, authentication credentials, and other features.
- `matchers`: A list of matchers to apply to the response of the requests. The matchers define how to search for specific patterns or keywords in the response body or headers. Nuclei supports various types of matchers, such as word, regex, status, and binary.
- `attacks`: An optional list of attack payloads to send in the requests. These payloads can be used to test for vulnerabilities such as SQL injection, XSS, and command injection. Nuclei supports various types of payloads, including file-based, parameter-based, and JSON-based.

## Conclusion

Nuclei is a powerful open-source tool that automates web application and network service vulnerability scanning. Its flexible and customizable approach allows users to create and use templates to identify vulnerabilities and misconfigurations. By using Nuclei with other tools, bug bounty hunters can streamline their testing workflows and improve security. Nuclei’s development and community support make it a valuable tool for security testing.