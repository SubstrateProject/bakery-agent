# bakery-agent

A small golang helper to deal with  discrete tasks related to ami "baking"

- grab a tarball from s3 and unpack it to a directory
- send a message denoting success or failure of a provisioning run to an sqs queue

Later it may grow the ability to "run" the provisioning script and ship it's output.
