🔐 token: Command-line tool to generate OAuth2 tokens for Google Workspace using a service account.

`google-oauth2-token` is a convenient and secure way to generate OAuth2 tokens for Google Workspace (formerly G Suite) by utilizing a service account key. It simplifies the process of acquiring and managing access tokens, making it easier for developers and system administrators to authenticate with Google APIs.

# 🚀 Features:
- Generate OAuth2 tokens for Google Workspace.
- Supports impersonation of superadmin users.
- Easy-to-use command-line interface.
- Cross-platform compatibility (Linux and macOS).

# 💻 Usage:
1. Clone the repository:
    ```bash
    git clone https://github.com/cldcvr/google-oauth2-token.git
    cd google-oauth2-token
    ```
2. Build the binary
    ```bash
    go build
    ```
3. Run the binary to generate the OAuth2 token, where
    - `xxx@yyy.com` is the user with appropriate permissions in the Google Workspace tenant
    - `credentials.json` is the service account key file
    - `scope1,scope2,scope3` are the scopes. For example, `https://www.googleapis.com/auth/admin.directory.user.readonly`
    ```bash
    ./token -email xxx@yyy.com -file credentials.json -scopes scope1,scope2,scope3
    ```
4. In the output you'll get the token in plaintext, which can be used in Postman or `cURL` commands. For example:
    ```bash
    curl -s -XGET \
      -H "Authorization: Bearer ya29.GOOGLE_OAUTH2_TOKEN" \
      -H 'Accept: application/json' \
      --compressed \
      "https://admin.googleapis.com/admin/directory/v1/users/aaa@yyy.com?projection=full"
    ```

# Google Workspace Service Account Setup Guide

## Create a Service Account

1. Go to the [Google Cloud Console](https://console.cloud.google.com/) and log in with your Google Workspace administrator account.
2. Click on the project dropdown and select a project or create a new one.
3. Navigate to "IAM & Admin" -> "Service Accounts" in the left-hand sidebar.
4. Click on the "Create Service Account" button.
5. Enter a name and optional description for the service account.
6. Choose the appropriate role(s) for the service account. For example, you might want to grant it the "Service Account Token Creator" role to generate OAuth2 tokens.
7. Click on the "Continue" button.
8. Optionally, add users who should have access to this service account. Typically, it's best to keep the list minimal.
9. Click on the "Done" button to create the service account.

## Generate a Key for the Service Account

1. In the list of service accounts, locate the service account you just created and click on the three dots (⋮) next to it.
2. Select "Manage keys" from the dropdown.
3. Click on the "Add Key" button.
4. Choose the key type as "JSON" and click on the "Create" button.
5. The key file will be downloaded to your local machine. Keep it secure, as it grants access to your Google Workspace resources.

## Assign scopes to the Service Account

1. Determine the necessary scopes (permissions) required for your service account to access Google Workspace APIs. Refer to the API documentation to identify the scopes needed for your specific use case.
2. Once you have identified the scopes, go to the [Google Admin console](https://admin.google.com/) and log in with your Google Workspace superadmin account.
3. Navigate to "Security" -> "Access and data control" -> "API Controls" -> "Manage Domain wide delegation" -> "Add new"
4. In the "Client ID" field, paste the Client ID (e.g., `118428936071234567890`) from the service account key file.
5. In the "OAuth Scopes" field, add the required scopes, separated by commas (`,`).
6. Click on the "Authorize" button to grant the specified scopes to the service account.

# License
This documentation is licensed under the Apache License. See the [LICENSE](LICENSE) file for details.

Contributions and feedback are welcome! Let's make OAuth2 token generation easier for everyone. 🌟

#oauth2 #googleworkspace #serviceaccount #golang #cli-tool #opensource
