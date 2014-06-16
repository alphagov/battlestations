
A Golang web app to manage deployment of applications into the GOV.UK infrastructure.

# Why?

We are unable to give permission to our deployment Jenkins instances to a significant subset
of the developers working on GOV.UK due to the level of access to our infrastructure it
implies. In order for these developers to deploy they currently need to find an authorised
member of GOV.UK to assume the risk and execute the deploy on their behalf. This is a time
consuming activity for the authorised developers so this is an experiment to see whether
we can replace the current system with a more automated one.

# How?

A user wants to request a release:
  - Authenticates with an OAuth provider
  - Lists available applications for that user
  - Picks an application
  - Chooses a release tag to deploy
  - Submits request to deploy

The server then processes the request to release:
  - Get the currently deployed tag of the application
  - Confirm the requested tag is newer than deployed tag
  - Get the diff between the two patches
  - Store against a HMACd random identifier:
    + User requesting
    + Application
    + Requested tag
    + SHA of patch
  - Email review list with patch and link to accept or reject the deployment

Someone accepts or rejects the deployment:
  - Receives an email with a link to accept or reject deployment
  - Authenticates with an OAuth provider
  - If has permission to accept or rejected deploy, presented with a page allowing them to

If the deployment was rejected
  - Remove request to deploy
  - Email original user saying it was denied and by whom

If the deployment was accepted
  - Store the accepting user against the original deployment request
  - Generate an expiring one time identifier which references deployment
  - Email original user with a link to start the deployment and when there slot will expire

User wants to start the deployment
  - Clicks the link in their email
  - Deployment ident is checked for validity (existence and expiry)
  - Destroy the one time ident
  - A session key is generated that is stored as a cookie and references the deployment struct
  - Redirect to deployment page

The deployment page and onwards are still to spec out.
